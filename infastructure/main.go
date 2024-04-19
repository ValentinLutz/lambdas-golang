package main

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"root/infastructure/bootstrap"
)

func main() {
	app := cdktf.NewApp(nil)

	bootstrap.NewStack(app, "eu-central-1", "test")

	app.Synth()
}
