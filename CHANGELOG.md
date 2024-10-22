## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#74](https://github.com/turbot/steampipe-plugin-whois/pull/74))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#74](https://github.com/turbot/steampipe-plugin-whois/pull/74))

## v0.11.0 [2024-01-22]

_What's new?_

- New tables added
  - [rdap_domain](https://hub.steampipe.io/plugins/turbot/whois/tables/rdap_domain) ([#46](https://github.com/turbot/steampipe-plugin-whois/pull/46))

## v0.10.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#51](https://github.com/turbot/steampipe-plugin-whois/pull/51))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#51](https://github.com/turbot/steampipe-plugin-whois/pull/51))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-whois/blob/main/docs/LICENSE). ([#51](https://github.com/turbot/steampipe-plugin-whois/pull/51))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#50](https://github.com/turbot/steampipe-plugin-whois/pull/50))

## v0.8.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#42](https://github.com/turbot/steampipe-plugin-whois/pull/42))

## v0.8.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#40](https://github.com/turbot/steampipe-plugin-whois/pull/40))
- Recompiled plugin with Go version `1.21`. ([#40](https://github.com/turbot/steampipe-plugin-whois/pull/40))

## v0.7.0 [2023-03-22]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#32](https://github.com/turbot/steampipe-plugin-whois/pull/32))

## v0.6.0 [2022-09-09]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.6](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v416-2022-09-02) which includes several caching and memory management improvements. ([#30](https://github.com/turbot/steampipe-plugin-whois/pull/30))
- Recompiled plugin with Go version `1.19`. ([#30](https://github.com/turbot/steampipe-plugin-whois/pull/30))

## v0.5.0 [2022-06-05]

_Enhancements_

- Recompiled plugin with [whois-parser v1.24.0](https://github.com/likexian/whois-parser/releases/tag/v1.24.0). ([#28](https://github.com/turbot/steampipe-plugin-whois/pull/28))
- Recompiled plugin with [steampipe-plugin-sdk v3.3.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v331--2022-06-30), which includes several caching fixes. ([#27](https://github.com/turbot/steampipe-plugin-whois/pull/27))

_Bug fixes_

- Fixed `.dk` domains returning incorrect nameservers. ([#28](https://github.com/turbot/steampipe-plugin-whois/pull/28))

## v0.4.1 [2022-05-23]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#25](https://github.com/turbot/steampipe-plugin-whois/pull/25))

## v0.4.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#22](https://github.com/turbot/steampipe-plugin-whois/pull/22))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#21](https://github.com/turbot/steampipe-plugin-whois/pull/21))

## v0.3.0 [2021-12-16]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#18](https://github.com/turbot/steampipe-plugin-whois/pull/18))
- Recompiled plugin with Go version 1.17 ([#18](https://github.com/turbot/steampipe-plugin-whois/pull/18))

## v0.2.0 [2021-08-18]

_Enhancements_

- Accept more date formats and don't error if unparseable (e.g. `tesco.co.uk` created date is `before Aug-1996`).
- Updated Steampipe SDK version for better caching, etc.

## v0.1.3 [2021-04-02]

_What's new?_

- Updated README and docs home page ([#9](https://github.com/turbot/steampipe-plugin-whois/pull/9))

## v0.1.2 [2021-03-18]

_Enhancements_

- Updated examples for `whois_domain` table ([#7](https://github.com/turbot/steampipe-plugin-whois/pull/7))
- Recompiled plugin with [steampipe-plugin-sdk v0.2.4](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v024-2021-03-16)

## v0.1.1 [2021-02-25]

_Bug fixes_

- Update to display the version of the plugin.
- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)
