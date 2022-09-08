package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-whois/whois"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: whois.Plugin})
}
