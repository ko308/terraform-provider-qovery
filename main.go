package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"

	"github.com/qovery/terraform-provider-qovery/qovery"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var version = "dev"

func main() {
	opts := tfsdk.ServeOpts{
		Name: "qovery",
	}

	if err := tfsdk.Serve(context.Background(), qovery.New(version), opts); err != nil {
		log.Fatal(err)
	}
}
