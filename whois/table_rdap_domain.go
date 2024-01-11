package whois

import (
	"context"
	"strings"

	"github.com/openrdap/rdap"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableRdapDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "rdap_domain",
		Description: "RDAP domain information including expiration, DNS servers and contact details.",
		List: &plugin.ListConfig{
			KeyColumns: plugin.SingleColumn("domain"),
			Hydrate:    getRdapDomain,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromQual("domain"), Description: "Domain name the RDAP information relates to."},

			{Name: "handle", Type: proto.ColumnType_STRING, Transform: transform.FromField("Handle").NullIfZero(), Description: "Registry-unique string identifier referencing the domain (domain name)."},
			{Name: "object_class_name", Type: proto.ColumnType_STRING, Description: "The data structure, a member named 'objectClassName', gives the object class name of a particular object as a string. This identifies the type of object being processed.  An objectClassName is REQUIRED in all RDAP response objects so that the type of the object can be interpreted."},
			{Name: "ldh_name", Type: proto.ColumnType_STRING, Description: "Textual representation of DNS names.", Transform: transform.FromField("LDHName")},

			// JSON fields
			{Name: "conformance", Type: proto.ColumnType_JSON, Description: "An array of strings, each providing a hint on the used specification."},
			{Name: "entities", Type: proto.ColumnType_JSON, Description: "An array of entities (linked contacts and the designated registrar)."},
			{Name: "events", Type: proto.ColumnType_JSON, Description: "An array of events that have occurred on the domain."},
			{Name: "links", Type: proto.ColumnType_JSON, Description: "Navigation to related on-line resources."},
			{Name: "nameservers", Type: proto.ColumnType_JSON, Description: "An array of nameserver objects."},
			{Name: "network", Type: proto.ColumnType_JSON, Description: "Information about IP address blocks and network allocations."},
			{Name: "notices", Type: proto.ColumnType_JSON, Description: "Information about the service."},
			{Name: "secure_dns", Type: proto.ColumnType_JSON, Description: "Secure DNS information.", Transform: transform.FromField("SecureDNS")},
			{Name: "status", Type: proto.ColumnType_JSON, Description: "An array of status flags describing the object state."},
			{Name: "variants", Type: proto.ColumnType_JSON, Description: "The internationalized domain name (IDN) variants."},
		},
	}
}

func getRdapDomain(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.EqualsQuals
	domain := quals["domain"].GetStringValue()

	client := &rdap.Client{}
	rdapResult, err := client.QueryDomain(domain)

	// Handle not found error
	if err != nil {
		if strings.Contains(err.Error(), "No RDAP servers found") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("rdap_domain.getRdapDomain", "api_error", err)
		return nil, err
	}

	d.StreamListItem(ctx, rdapResult)

	return nil, nil
}
