package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewGetFunction(stack awscdk.Stack, config *StageConfig) awslambda.Function {
	dockerImageCode := awslambda.DockerImageCode_FromImageAsset(
		jsii.String("../../../"), &awslambda.AssetImageCodeProps{
			File: jsii.String("./lambdas/facts/app-v1-get/Dockerfile"),
		},
	)
	env := map[string]*string{
		"DB_HOST":      &config.databaseProps.host,
		"DB_PORT":      &config.databaseProps.port,
		"DB_NAME":      &config.databaseProps.name,
		"DB_SECRET_ID": &config.databaseProps.secret,
	}
	if config.endpointUrl != nil {
		env["AWS_ENDPOINT_URL"] = config.endpointUrl
	}

	return awslambda.NewDockerImageFunction(
		stack, jsii.String("GetFactsFunctionV1"), &awslambda.DockerImageFunctionProps{
			Code:        dockerImageCode,
			Environment: &env,
		},
	)
}

func NewPostFunction(stack awscdk.Stack, config *StageConfig) awslambda.Function {
	dockerImageCode := awslambda.DockerImageCode_FromImageAsset(
		jsii.String("../../../"), &awslambda.AssetImageCodeProps{
			File: jsii.String("./lambdas/facts/app-v1-post/Dockerfile"),
		},
	)

	env := map[string]*string{
		"DB_HOST":      &config.databaseProps.host,
		"DB_PORT":      &config.databaseProps.port,
		"DB_NAME":      &config.databaseProps.name,
		"DB_SECRET_ID": &config.databaseProps.secret,
	}
	if config.endpointUrl != nil {
		env["AWS_ENDPOINT_URL"] = config.endpointUrl
	}
	return awslambda.NewDockerImageFunction(
		stack, jsii.String("PostFactsFunctionV1"), &awslambda.DockerImageFunctionProps{
			Code:        dockerImageCode,
			Environment: &env,
		},
	)
}

func NewRestApi(stack awscdk.Stack, config *StageConfig) awscdk.Stack {
	getFunction := NewGetFunction(stack, config)

	lambdaRestApi := awsapigateway.NewLambdaRestApi(
		stack, jsii.String("FactsRestApi"), &awsapigateway.LambdaRestApiProps{
			Proxy:   jsii.Bool(false),
			Handler: getFunction,
		},
	)

	factsResource := lambdaRestApi.Root().AddResource(
		jsii.String("facts"), &awsapigateway.ResourceOptions{},
	)

	factsResource.AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getFunction, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
	)

	postFunction := NewPostFunction(stack, config)
	factsResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(postFunction, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
	)
	return stack
}

func NewStack(scope constructs.Construct, id string, config *StageConfig) awscdk.Stack {
	stack := awscdk.NewStack(
		scope, &id, &awscdk.StackProps{Env: &awscdk.Environment{Account: &config.account, Region: &config.region}},
	)

	return NewRestApi(stack, config)
}
