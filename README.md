![image](https://hub.steampipe.io/images/plugins/turbot/whois-social-graphic.png)

# WHOIS Plugin for Steampipe

Use SQL to query domain records, name servers and contact information from WHOIS.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/whois)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/whois/tables)
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-whois/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install whois
```

Run a query:
```sql
select * from whois_domain where domain = 'steampipe.io';
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-whois.git
cd steampipe-plugin-whois
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
```

Try it!
```
steampipe query
> .inspect whois
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-whois/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [WHOIS Plugin](https://github.com/turbot/steampipe-plugin-whois/labels/help%20wanted)
