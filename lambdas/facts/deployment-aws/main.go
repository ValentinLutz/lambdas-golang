package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	config := NewStageConfig()

	app := awscdk.NewApp(nil)
	tags := awscdk.Tags_Of(app)
	tags.Add(
		jsii.String("project"), jsii.String("facts"), &awscdk.TagProps{},
	)
	tags.Add(
		jsii.String("environment"), jsii.String("dev"), &awscdk.TagProps{},
	)

	NewStack(app, "stack", config)

	app.Synth(nil)
}
