package whois

import (
	"context"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/likexian/whois"
	whoisparser "github.com/likexian/whois-parser"
	"github.com/openrdap/rdap"
	"github.com/sethvargo/go-retry"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableWhoisDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "whois_domain",
		Description: "WHOIS domain information including expiration, DNS servers and contact details.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    listDomain,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromField("Domain.Domain").NullIfZero(), Description: "Domain name the WHOIS information relates to."},

			// Domain information
			{Name: "domain_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Domain.ID").NullIfZero(), Description: "Unique identifier for the domain."},
			{Name: "domain_punycode", Type: proto.ColumnType_STRING, Transform: transform.FromField("Domain.Punycode").NullIfZero(), Description: "Punycode ASCII variation of the Unicode domain name."},
			{Name: "domain_extension", Type: proto.ColumnType_STRING, Transform: transform.FromField("Domain.Extension").NullIfZero(), Description: "Extension of the domain."},
			{Name: "whois_server", Type: proto.ColumnType_STRING, Transform: transform.FromField("Domain.WhoisServer").NullIfZero(), Description: "WHOIS server that manages the domain."},
			{Name: "status", Type: proto.ColumnType_JSON, Transform: transform.FromField("Domain.Status").NullIfZero(), Description: "Extensible Provisioning Protocol (EPP) status codes set on the domain. Common status codes (e.g. client_transfer_prohibited) are also elevated to column level. A full list is available at https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en"},
			{Name: "name_servers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Domain.NameServers").NullIfZero(), Description: "List of name servers for the domain."},
			{Name: "dns_sec", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Domain.DnsSec").NullIfZero(), Description: "True if the domain has enabled DNSSEC."},
			{Name: "created_date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Domain.CreatedDate").Transform(whoisDateToTimestamp).NullIfZero(), Description: "Date when the domain was first registered."},
			{Name: "updated_date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Domain.UpdatedDate").Transform(whoisDateToTimestamp).NullIfZero(), Description: "Last date when the domain record was updated."},
			{Name: "expiration_date", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Domain.ExpirationDate").Transform(whoisDateToTimestamp).NullIfZero(), Description: "Expiration date for the domain."},

			// Commonly used EPP status codes
			{Name: "client_delete_prohibited", Type: proto.ColumnType_BOOL, Transform: transform.FromP(statusToBool, "clientdeleteprohibited"), Description: "This status code tells your domain's registry to reject requests to delete the domain."},
			{Name: "client_transfer_prohibited", Type: proto.ColumnType_BOOL, Transform: transform.FromP(statusToBool, "clienttransferprohibited"), Description: "This status code tells your domain's registry to reject requests to transfer the domain from your current registrar to another."},
			{Name: "client_update_prohibited", Type: proto.ColumnType_BOOL, Transform: transform.FromP(statusToBool, "clientupdateprohibited"), Description: "This status code tells your domain's registry to reject requests to update the domain."},
			{Name: "server_delete_prohibited", Type: proto.ColumnType_BOOL, Transform: transform.FromP(statusToBool, "serverdeleteprohibited"), Description: "This status code prevents your domain from being deleted. clientdeleteprohibited is more commonly used."},
			{Name: "server_transfer_prohibited", Type: proto.ColumnType_BOOL, Transform: transform.FromP(statusToBool, "servertransferprohibited"), Description: "This status code prevents your domain from being transferred from your current registrar to another. clienttransferprohibited is more commonly used."},
			{Name: "server_update_prohibited", Type: proto.ColumnType_BOOL, Transform: transform.FromP(statusToBool, "serverupdateprohibited"), Description: "This status code locks your domain preventing it from being updated. clientupdateprohibited is more commonly used."},

			// Contact information
			{Name: "registrar", Type: proto.ColumnType_JSON, Description: "Registrar contact information."},
			{Name: "registrant", Type: proto.ColumnType_JSON, Description: "Registrant contact information."},
			{Name: "admin", Type: proto.ColumnType_JSON, Description: "Administrative contact information.", Transform: transform.FromField("Administrative")},
			{Name: "technical", Type: proto.ColumnType_JSON, Description: "Technical contact information."},
			{Name: "billing", Type: proto.ColumnType_JSON, Description: "Billing contact information."},
		},
	}
}

func listDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	domain := quals["domain"].GetStringValue()

	// Attempt rdap lookup first.
	client := &rdap.Client{}
	rdapResult, err := client.QueryDomain(domain)

	if err == nil {
		plugin.Logger(ctx).Debug("whois_domain.listDomain", "rdapResult", rdapResult)
		mapped := rdapToWhoisDomain(domain, rdapResult)
		plugin.Logger(ctx).Debug("whois_domain.listDomain", "mapped.Domain", mapped.Domain)
		d.StreamListItem(ctx, mapped)
		return nil, nil
	}

	// Drop to whois.
	var whoisRaw string

	// WHOIS servers are fussy about load, so retry with backoff
	b := retry.NewFibonacci(100 * time.Millisecond)

	err = retry.Do(ctx, retry.WithMaxRetries(10, b), func(ctx context.Context) error {
		var err error
		whoisRaw, err = whois.Whois(domain)
		if err != nil {
			plugin.Logger(ctx).Warn("whois_domain.getDomain", "lookup_error", err)
			if strings.Contains(err.Error(), "connection reset by peer") || strings.Contains(err.Error(), "domain is empty") {
				return retry.RetryableError(err)
			}
		}
		return nil
	})

	if err != nil {
		plugin.Logger(ctx).Error("whois_domain.getDomain", "lookup_error_final", err)
		return nil, err
	}

	// Not found errors come back as success with a not found string to match
	if strings.Index(whoisRaw, "No match for") == 0 {
		plugin.Logger(ctx).Warn("whois_domain.getDomain", "not_found_error", domain)
		return nil, nil
	}

	result, err := whoisparser.Parse(whoisRaw)
	switch err {
	case nil:
		break
	case whoisparser.ErrNotFoundDomain:
		plugin.Logger(ctx).Warn("whois_domain.getDomain", "ErrDomainNotFound", err)
		plugin.Logger(ctx).Debug("whois_domain.getDomain", "domain", domain)
		return nil, nil
	case whoisparser.ErrDomainDataInvalid:
		plugin.Logger(ctx).Warn("whois_domain.getDomain", "ErrDomainInvalidData", err)
		plugin.Logger(ctx).Debug("whois_domain.getDomain", "whoisRaw", whoisRaw)
		return nil, nil
	default:
		plugin.Logger(ctx).Error("whois_domain.getDomain", "parse_error", err)
		plugin.Logger(ctx).Debug("whois_domain.getDomain", "whoisRaw", whoisRaw)
		return nil, err
	}

	d.StreamListItem(ctx, result)
	return nil, nil
}

// WHOIS dates can be pretty crazy, so be forgiving in our parsing and fail to null
// tesco.co.uk -> "before Aug-1996"
// stripe.co.uk -> "14-Jul-2011"
func whoisDateToTimestamp(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	ds := d.Value.(string)
	t, err := dateparse.ParseAny(ds)
	if err != nil {
		return nil, nil
	}
	return t, nil
}

func statusToBool(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	wd := d.HydrateItem.(whoisparser.WhoisInfo)
	param := d.Param.(string)
	if wd.Domain == nil || wd.Domain.Status == nil {
		return false, nil
	}
	return containsString(wd.Domain.Status, param), nil
}

func containsString(items []string, target string) bool {
	for _, i := range items {
		if i == target {
			return true
		}
	}
	return false
}

func rdapToWhoisDomain(passedDomain string, r *rdap.Domain) whoisparser.WhoisInfo {
	var nameservers []string
	for _, ns := range r.Nameservers {
		nameservers = append(nameservers, ns.LDHName)
	}
	dnssec := false
	if r.SecureDNS != nil && r.SecureDNS.DelegationSigned != nil && *r.SecureDNS.DelegationSigned {
		dnssec = true
	}
	createdDate := ""
	updatedDate := ""
	expireDate := ""
	for _, event := range r.Events {
		switch event.Action {
		case "registration":
			createdDate = event.Date
		case "expiration":
			expireDate = event.Date
		case "last changed":
			updatedDate = event.Date
		}
	}

	domain := &whoisparser.Domain{
		ID:             r.Handle,
		Domain:         passedDomain,
		Punycode:       r.LDHName,
		Extension:      r.LDHName[strings.LastIndex(r.LDHName, ".")+1:],
		WhoisServer:    r.Port43,
		Status:         spacesSpaced(r.Status),
		NameServers:    nameservers,
		DNSSec:         dnssec,
		CreatedDate:    createdDate,
		UpdatedDate:    updatedDate,
		ExpirationDate: expireDate,
	}

	registrar := entityToContact("registrar", r.Entities)
	registrant := entityToContact("registrant", r.Entities)
	admin := entityToContact("admin", r.Entities)
	technical := entityToContact("technical", r.Entities)
	billing := entityToContact("billing", r.Entities)

	return whoisparser.WhoisInfo{
		Domain:         domain,
		Registrar:      registrar,
		Registrant:     registrant,
		Administrative: admin,
		Technical:      technical,
		Billing:        billing,
	}
}

func spacesSpaced(sa []string) (spacelessSA []string) {
	for _, s := range sa {
		spacelessSA = append(spacelessSA, strings.ReplaceAll(s, " ", ""))
	}
	return
}

func assignIfEmpty(target string, newValue string) string {
	if target == "" {
		return newValue
	}
	return target
}

func vcardToContact(contact *whoisparser.Contact, entity rdap.Entity) {
	if entity.VCard != nil && entity.VCard.Properties != nil {
		for _, property := range entity.VCard.Properties {
			switch property.Name {
			case "fn":
				contact.Name = assignIfEmpty(contact.Name, strings.Join(property.Values(), " "))
			case "org":
				contact.Organization = strings.Join(property.Values(), " ")
			case "adr":
				// TODO: https://datatracker.ietf.org/doc/html/rfc6350#section-6.3.1
				contact.Street = ""
				contact.City = ""
				contact.Province = ""
				contact.PostalCode = ""
				contact.Country = ""
			case "tel":
				// TODO: parse these more to detrimine if it's fax number.
				contact.Phone = strings.Join(property.Values(), " ")
				contact.Fax = ""
			case "email":
				contact.Email = strings.Join(property.Values(), " ")
			case "url":
				contact.ReferralURL = strings.Join(property.Values(), " ")
			}
		}

	}
}

func entityToContact(role string, entities []rdap.Entity) *whoisparser.Contact {
	contact := &whoisparser.Contact{}
	for _, entity := range entities {
		if containsString(entity.Roles, role) {
			contact.ID = entity.Handle
			vcardToContact(contact, entity)
			for _, inner := range entity.Entities {
				if containsString(inner.Roles, "abuse") {
					vcardToContact(contact, inner)
				}
			}
		}
	}
	return contact
}
