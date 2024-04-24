package main

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	NewStack(app, "eu-central-1", "test")

	app.Synth()
}
