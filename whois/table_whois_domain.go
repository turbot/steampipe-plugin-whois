package whois

import (
	"context"
	"strings"
	"time"

	"github.com/likexian/whois-go"
	"github.com/likexian/whois-parser-go"
	"github.com/sethvargo/go-retry"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableWhoisDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "whois_domain",
		Description: "WHOIS domain information including expiration, DNS servers and contact details.",
		List: &plugin.ListConfig{
			Hydrate: listDomain,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    getDomain,
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
			{Name: "created_date", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Domain.CreatedDate").NullIfZero(), Description: "Date when the domain was first registered."},
			{Name: "updated_date", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Domain.UpdatedDate").NullIfZero(), Description: "Last date when the domain record was updated."},
			{Name: "expiration_date", Type: proto.ColumnType_DATETIME, Transform: transform.FromField("Domain.ExpirationDate").NullIfZero(), Description: "Expiration date for the domain."},

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
	// If the user requests a list of domains with no quals, then just return empty.
	return nil, nil
}

func getDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	domain := quals["domain"].GetStringValue()
	var whoisRaw string

	// WHOIS servers are fussy about load, so retry with backoff
	b, err := retry.NewFibonacci(100 * time.Millisecond)
	if err != nil {
		plugin.Logger(ctx).Error("whois_domain.getDomain", "retry_init_error", err)
		return nil, err
	}

	err = retry.Do(ctx, retry.WithMaxRetries(10, b), func(ctx context.Context) error {
		var err error
		whoisRaw, err = whois.Whois(domain)
		if err != nil {
			plugin.Logger(ctx).Warn("whois_domain.getDomain", "lookup_error", err)
			if strings.Contains(err.Error(), "connection reset by peer") {
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
		plugin.Logger(ctx).Warn("whois_domain.getDomain", "not_found_error", err)
		return nil, nil
	}

	result, err := whoisparser.Parse(whoisRaw)
	if err != nil {
		plugin.Logger(ctx).Error("whois_domain.getDomain", "parse_error", err)
		return nil, err
	}

	return result, nil
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
