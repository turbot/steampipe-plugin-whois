---
Title: WHOIS
Organization: Turbot HQ, Inc
iconURL: https://turbot.com/images/turbot-icon-original.png

---

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
