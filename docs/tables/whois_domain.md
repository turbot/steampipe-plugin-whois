# Table: whois_domain

WHOIS domain information including expiration, DNS servers and contact details.

Note: It's not practical to list all domains in the world, so this table requires a
`domain` qualifier to be passed in the `where` or `join` clause for all queries.


## Examples

### Basic whois info

```sql
select
  domain,
  expiration_date
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Days until expiration

```sql
select
  domain,
  expiration_date,
  date_part('day', expiration_date - current_date) as days_until_expiration
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Get name server information

```sql
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

```sql
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

```sql
select
  domain,
  status,
  status ? 'pendingtransfer' as pending_transfer
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Contact information

```sql
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

### Registrar managing the domain

```sql
select
  domain,
  registrar->>'name' as registrar
from
  whois_domain
where
  domain = 'steampipe.io';
```

### Working with multiple domains

```sql
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
