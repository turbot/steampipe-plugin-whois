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
