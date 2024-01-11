---
title: "Steampipe Table: whois_domain - Query Whois Domains using SQL"
description: "Allows users to query Whois Domains, providing specific details about domain names, their registration, and ownership information."
---

# Table: whois_domain - Query Whois Domains using SQL

Whois is a protocol that is used to query databases that store the registered users or assignees of an Internet resource, such as a domain name or an IP address block. It provides information related to the registration and ownership of a domain name. This includes details about the registrant, administrative, billing and technical contacts.

## Table Usage Guide

The `whois_domain` table provides insights into domain names within the Whois protocol. As a security analyst, explore domain-specific details through this table, including registration, ownership, and associated metadata. Utilize it to uncover information about domains, such as their registrant details, administrative contacts, and the status of the domain.

**Important Notes**
- It's not practical to list all domains in the world, so this table requires a
`domain` qualifier to be passed in the `where` or `join` clause for all queries.

## Examples

### Basic whois info

```sql+postgres
select
  domain,
  expiration_date
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  expiration_date
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Days until expiration

```sql+postgres
select
  domain,
  expiration_date,
  date_part('day', expiration_date - current_date) as days_until_expiration
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  expiration_date,
  julianday(expiration_date) - julianday(date('now')) as days_until_expiration
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Get name server information

```sql+postgres
select
  domain,
  name_servers
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  name_servers
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Check domain status codes

Commonly used protections:

```sql+postgres
select
  domain,
  client_delete_prohibited,
  client_transfer_prohibited,
  client_update_prohibited,
  server_delete_prohibited,
  server_transfer_prohibited,
  server_update_prohibited
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  client_delete_prohibited,
  client_transfer_prohibited,
  client_update_prohibited,
  server_delete_prohibited,
  server_transfer_prohibited,
  server_update_prohibited
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Check for any EPP status code:

```sql+postgres
select
  domain,
  status,
  status ? 'pendingtransfer' as pending_transfer
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  status,
  json_extract(status, '$.pendingtransfer') as pending_transfer
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Contact information

```sql+postgres
select
  domain,
  jsonb_pretty(admin) as admin,
  jsonb_pretty(billing) as billing,
  jsonb_pretty(registrant) as registrant,
  jsonb_pretty(technical) as technical
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  admin,
  billing,
  registrant,
  technical
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Registrar managing the domain

```sql+postgres
select
  domain,
  registrar->>'name' as registrar
from
  whois_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  json_extract(registrar, '$.name') as registrar
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Working with multiple domains

```sql+postgres
select
  domain,
  expiration_date
from
  whois_domain
where
  domain in (
    'github.com',
    'google.com',
    'steampipe.io',
    'yahoo.com'
  );
```

```sql+sqlite
select
  domain,
  expiration_date
from
  whois_domain
where
  domain in (
    'github.com',
    'google.com',
    'steampipe.io',
    'yahoo.com'
  );
```

## Implementation notes

* Automatically retries with backoff. WHOIS servers are fussy with throttling.
* May return partial results, some WHOIS servers return domain info but throttle / skip contact information.
