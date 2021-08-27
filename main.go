package main

import (
	"context"
	"flag"
	"log"

	"github.com/SecurityGeekIO/terraform-provider-zpa/zscaler"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", true, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "github.com/SecurityGeekIO/terraform-provider-zpa",
			&plugin.ServeOpts{
				ProviderFunc: zscaler.Provider,
			})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: zscaler.Provider})
	}
}

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: zscaler.Provider,
// 	})
// }
