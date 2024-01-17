package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	defer jsii.Close()

	config, err := NewStageConfig()
	if err != nil {
		panic(err)
	}

	app := awscdk.NewApp(nil)
	tags := awscdk.Tags_Of(app)
	tags.Add(jsii.String("region"), &config.region, &awscdk.TagProps{})
	tags.Add(jsii.String("environment"), &config.environment, &awscdk.TagProps{})

	NewStack(app, NewIdWithStage(config, "Infrastructure"), config)

	app.Synth(nil)
}
