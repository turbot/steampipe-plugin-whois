---
title: "Steampipe Table: rdap_domain - Query RDAP Domains using SQL"
description: "Allows users to query RDAP domain registration data in a structured and standardized way."
---

# Table: rdap_domain - Query RDAP information for Domains using SQL

An RDAP domain refers to a specific domain name for which you can retrieve registration information using the RDAP (Registration Data Access Protocol) system. RDAP allows users to query domain registration data in a structured and standardized way. The term "RDAP domain" is used to indicate that you are accessing or querying information for a particular domain name through the RDAP protocol.

## Table Usage Guide

The `rdap_domain` table provides insights about an RDAP domain query specifically refers to querying information about a domain name. This can include details such as the domain name's registration status, the registrar information, the domain's creation and expiration dates, and contact information associated with the domain. RDAP provides a more secure and standardized way to access this information compared to WHOIS, and it is becoming the preferred method for domain-related queries in the internet infrastructure community.

**Important Notes**
It's not practical to list all domains in the world, so this table requires a
`domain` qualifier to be passed in the `where` or `join` clause for all queries.

## Examples

### Basic RDAP info

```sql+postgres
select
  domain,
  handle,
  ldh_name,
  object_class_name
from
  rdap_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  domain,
  handle,
  ldh_name,
  object_class_name
from
  rdap_domain
where
  domain = 'steampipe.io';
```

### Get nameserver information for a domain

```sql+postgres
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

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(n.value, '$.Handle') as nameserver_handle,
  json_extract(n.value, '$.LDHName') as nameserver_ldh_name,
  json_extract(n.value, '$.UnicodeName') as nameserver_unicode_name,
  json_extract(n.value, '$.Port43') as nameserver_port43,
  json_extract(n.value, '$.Conformance') as nameserver_conformance,
  json_extract(n.value, '$.Events') as nameserver_events,
  json_extract(n.value, '$.Status') as nameserver_status,
  json_extract(n.value, '$.Entities') as nameserver_entities
from
  rdap_domain,
  json_each(rdap_domain.nameservers) as n
where
  rdap_domain.domain = 'steampipe.io';

```

### Check domain status codes

Commonly used protections:

```sql+postgres
select
  domain,
  s as atatus_code
from
  rdap_domain,
  jsonb_array_elements_text(status) as s
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(s.value, '$') as status_code
from
  rdap_domain,
  json_each(rdap_domain.status) as s
where
  rdap_domain.domain = 'steampipe.io';

```

### Get domain variants

```sql+postgres
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

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(v.value, '$.IDNTable') as idn_table,
  json_extract(v.value, '$.Relation') as relation,
  json_extract(v.value, '$.VariantNames') as variant_names
from
  rdap_domain,
  json_each(rdap_domain.variants) as v
where
  rdap_domain.domain = 'steampipe.io';
```

### Get event details of a domain

```sql+postgres
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

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(e.value, '$.Action') as action,
  json_extract(e.value, '$.Actor') as actor,
  json_extract(e.value, '$.Date') as date,
  json_extract(e.value, '$.Links') as links
from
  rdap_domain,
  json_each(rdap_domain.events) as e
where
  rdap_domain.domain = 'steampipe.io';

```

### Get entity details of a domain

```sql+postgres
select
  domain,
  handle,
  e ->> 'Handle' as handle,
  e ->> 'Port43' as port43,
  e -> 'AsEventActor' as as_event_actor,
  e -> 'VCard' as v_card,
  e -> 'Autnums' as autnums,
  e -> 'Roles' as roles,
  e -> 'Notices' as notices,
  e -> 'Remarks' as remarks,
  e -> 'Networks' as networks
from
  rdap_domain,
  jsonb_array_elements(entities) as e
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  rdap_domain.domain,
  rdap_domain.handle,
  json_extract(e.value, '$.Handle') as handle,
  json_extract(e.value, '$.Port43') as port43,
  json_extract(e.value, '$.AsEventActor') as as_event_actor,
  json_extract(e.value, '$.VCard') as v_card,
  json_extract(e.value, '$.Autnums') as autnums,
  json_extract(e.value, '$.Roles') as roles,
  json_extract(e.value, '$.Notices') as notices,
  json_extract(e.value, '$.Remarks') as remarks,
  json_extract(e.value, '$.Networks') as networks
from
  rdap_domain,
  json_each(rdap_domain.entities) as e
where
  rdap_domain.domain = 'steampipe.io';

```

### Get public IP details of a domain

```sql+postgres
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

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(p.value, '$.Type') as public_ip_type,
  json_extract(p.value, '$.Identifier') as public_ip_identifier
from
  rdap_domain,
  json_each(rdap_domain.public_ids) as p
where
  rdap_domain.domain = 'steampipe.io';

```

### Get network information of a domain

```sql+postgres
select
  domain,
  network ->> 'Handle' as network_handle,
  network ->> 'ObjectClassName' as network_object_class_name,
  network ->> 'StartAddress' as network_start_address,
  network ->> 'EndAddress' as network_end_address,
  network ->> 'IPVersion' as network_ip_version,
  network ->> 'Name' as network_name,
  network ->> 'Type' as network_type,
  network ->> 'Country' as network_country,
  network ->> 'ParentHandle' as network_parent_handle,
  network -> 'Status' as network_status
from
  rdap_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(rdap_domain.network, '$.Handle') as network_handle,
  json_extract(rdap_domain.network, '$.ObjectClassName') as network_object_class_name,
  json_extract(rdap_domain.network, '$.StartAddress') as network_start_address,
  json_extract(rdap_domain.network, '$.EndAddress') as network_end_address,
  json_extract(rdap_domain.network, '$.IPVersion') as network_ip_version,
  json_extract(rdap_domain.network, '$.Name') as network_name,
  json_extract(rdap_domain.network, '$.Type') as network_type,
  json_extract(rdap_domain.network, '$.Country') as network_country,
  json_extract(rdap_domain.network, '$.ParentHandle') as network_parent_handle,
  json_extract(rdap_domain.network, '$.Status') as network_status
from
  rdap_domain
where
  rdap_domain.domain = 'steampipe.io';

```

### Get secure DNS details of a domain

```sql+postgres
select
  domain,
  secure_dns ->> 'ZoneSigned' as zone_signed,
  secure_dns ->> 'DelegationSigned' as delegation_signed,
  secure_dns ->> 'MaxSigLife' as max_sig_life,
  secure_dns ->> 'Ds' as data_structure,
  secure_dns ->> 'Keys' as keys
from
  rdap_domain
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(rdap_domain.secure_dns, '$.ZoneSigned') as zone_signed,
  json_extract(rdap_domain.secure_dns, '$.DelegationSigned') as delegation_signed,
  json_extract(rdap_domain.secure_dns, '$.MaxSigLife') as max_sig_life,
  json_extract(rdap_domain.secure_dns, '$.Ds') as data_structure,
  json_extract(rdap_domain.secure_dns, '$.Keys') as keys
from
  rdap_domain
where
  rdap_domain.domain = 'steampipe.io';

```

### Get a domain remarks

```sql+postgres
select
  domain,
  r ->> 'Title' as title,
  r ->> 'Type' as type,
  r ->> 'Description' as description,
  r ->> 'Links' as links
from
  rdap_domain,
  jsonb_array_elements(remarks) as r
where
  domain = 'steampipe.io';
```

```sql+sqlite
select
  rdap_domain.domain,
  json_extract(r.value, '$.Title') as title,
  json_extract(r.value, '$.Type') as type,
  json_extract(r.value, '$.Description') as description,
  json_extract(r.value, '$.Links') as links
from
  rdap_domain,
  json_each(rdap_domain.remarks) as r
where
  rdap_domain.domain = 'steampipe.io';

```

### Working with multiple domains

```sql+postgres
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

```sql+sqlite
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