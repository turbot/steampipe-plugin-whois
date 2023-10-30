# Table: rdap_domain

An RDAP domain refers to a specific domain name for which you can retrieve registration information using the RDAP (Registration Data Access Protocol) system. RDAP allows users to query domain registration data in a structured and standardized way. The term "RDAP domain" is used to indicate that you are accessing or querying information for a particular domain name through the RDAP protocol.

**Note:** It's not practical to list all domains in the world, so this table requires a
`domain` qualifier to be passed in the `where` or `join` clause for all queries.


## Examples

### Basic whois info

```sql
select
  domain,
  domain_id,
  ldh_name,
  object_class_name
from
  rdap_domain
where
  domain = 'steampipe.io';
```

### Get name server information for a domain

```sql
select
  domain,
  n ->> 'Handle' as nameserver_handle,
  n ->> 'LDHName' as nameserver_ldh_name,
  n ->> 'UnicodeName' as nameserver_unicode_name,
  n ->> 'Port43' as nameserver_port43,
  n -> 'Conformance' as nameserver_conformance,
  n -> 'Events' as nameserver_events,
  n -> 'Status' as nameserver_status,
  n -> 'Entities' as nameserver_entities
from
  rdap_domain,
  jsonb_array_elements(nameservers) as n
where
  domain = 'steampipe.io';
```

### Check domain status codes

Commonly used protections:

```sql
select
  domain,
  s as atatus_code
from
  rdap_domain,
  jsonb_array_elements_text(status) as s
where
  domain = 'steampipe.io';
```

### Get domain variants

```sql
select
  domain,
  v ->> 'IDNTable' as idn_table,
  v ->> 'Relation' as relation,
  v ->> 'VariantNames' as variant_names
from
  rdap_domain,
  jsonb_array_elements(variants) as v
where
  domain = 'steampipe.io';
```

### Get event details of a domain

```sql
select
  domain,
  e ->> 'Action' as action,
  e ->> 'Actor' as Actor,
  e ->> 'Date' as date,
  e -> 'Links' as links
from
  rdap_domain,
  jsonb_array_elements(events) as e
where
  domain = 'steampipe.io';
```

### Get entity details of a domain

```sql
select
  domain,
  domain_id,
  e ->> 'Handle' as handle,
  e ->> 'Port43' as port43,
  e -> 'AsEventActor' as as_event_actor,
  e -> 'VCard' as v_card,
  e -> 'Autnums' as autnums,
  e -> 'Roles' as roles,
  e -> 'Notices' as notices,
  e -> 'Remarks' as Remarks,
  e -> 'Networks' as Networks
from
  rdap_domain,
  jsonb_array_elements(entities) as e
where
  domain = 'steampipe.io';
```

### Get public IP details of a domain

```sql
select
  domain,
  p ->> 'Type' as public_ip_type,
  p ->> 'Identifier' as public_ip_identifier
from
  rdap_domain,
  jsonb_array_elements(public_ids) as p
where
  domain = 'steampipe.io';
```

### Working with multiple domains

```sql
select
  domain,
  status
from
  rdap_domain
where
  domain in (
    'github.com',
    'google.com',
    'steampipe.io',
    'yahoo.com'
  );
```