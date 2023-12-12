---
organization: Turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/whois.svg"
brand_color: "#005A9C"
display_name: WHOIS
name: whois
description: Steampipe plugin for querying domains, name servers and contact information from WHOIS.
og_description: Query WHOIS with SQL! Open source CLI. No DB required. 
og_image: "/images/plugins/turbot/whois-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# WHOIS + Steampipe

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

[WHOIS](https://whois.icann.org/en/about-whois) is a widely used Internet record listing that identifies who owns a domain and how to get in contact with them. The Internet Corporation for Assigned Names and Numbers (ICANN) regulates domain name registration and ownership.

For example:

```sql
select
  domain,
  expiration_date
from
  whois_domain
where
  domain = 'steampipe.io';
```

```
+--------------+---------------------+
| domain       | expiration_date     |
+--------------+---------------------+
| steampipe.io | 2021-10-13 19:28:29 |
+--------------+---------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/whois/tables)**

## Get started

### Install

Download and install the latest WHOIS plugin:

```bash
steampipe plugin install whois
```

### Credentials

| Item | Description |
| - | - |
| Credentials | No creds required |
| Permissions | n/a |
| Radius | Steampipe connects to the correct WHOIS server based on the TLD |
| Resolution | n/a |

### Configuration

No configuration is needed. Installing the latest whois plugin will create a config file (`~/.steampipe/config/whois.spc`) with a single connection named `whois`:

```hcl
connection "whois" {
  plugin = "whois"
}
```


