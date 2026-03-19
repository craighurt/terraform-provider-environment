package main

import (
	"context"
	"flag"
	"log"

	"github.com/craighurt/terraform-provider-environment/internal/provider"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/craighurt/environment",
	}

	if debug {
		opts.Debug = true
	}

	err := providerserver.Serve(context.Background(), func() fwprovider.Provider {
		return &provider.EnvironmentProvider{}
	}, opts)

	if err != nil {
		log.Fatal(err)
	}
}
