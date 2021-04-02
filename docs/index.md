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
---

# WHOIS + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

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

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-whois
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)













# WHOIS

WHOIS record information including expiration, DNS servers and contact details.

Data returned in WHOIS records [varies by TLD](https://tools.ietf.org/html/rfc7485). This plugin targets the most common fields.


## Installation

Download and install the latest whois plugin:

```bash
$ steampipe plugin install whois
Installing plugin whois...
$
```

## Run a query


```bash
$ steampipe query
Welcome to Steampipe v0.0.12
Type ".inspect" for more information.
> select domain, domain_id, status, expiration_date from whois_record where domain = 'steampipe.io';
+--------------+--------------------------+---------------------------------------------------------+---------------------+
|    domain    |        domain_id         |                         status                          |   expiration_date   |
+--------------+--------------------------+---------------------------------------------------------+---------------------+
| steampipe.io | D503300001187474055-LRMS | ["clienttransferprohibited","servertransferprohibited"] | 2021-10-13 19:28:29 |
+--------------+--------------------------+---------------------------------------------------------+---------------------+
```
