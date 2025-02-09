package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/better/terraform-provider-snowflake/pkg/provider"
	"github.com/chanzuckerberg/go-misc/ver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	version := flag.Bool("version", false, "spit out version for resources here")
	flag.Parse()

	if *version {
		verString, err := ver.VersionStr()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(verString)
		return
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
