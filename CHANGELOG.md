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
