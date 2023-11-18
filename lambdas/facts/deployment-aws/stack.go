package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewGetFunction(stack awscdk.Stack) awslambda.Function {
	dockerImageCode := awslambda.DockerImageCode_FromImageAsset(
		jsii.String("../../../"), &awslambda.AssetImageCodeProps{
			File: jsii.String("./lambdas/facts/app-v1-get/Dockerfile"),
		},
	)
	return awslambda.NewDockerImageFunction(
		stack, jsii.String("GetFactsFunctionV1"), &awslambda.DockerImageFunctionProps{
			Code: dockerImageCode,
			Environment: &map[string]*string{
				"DB_HOST": jsii.String("database-postgres"),
				"DB_PORT": jsii.String("5432"),
				"DB_NAME": jsii.String("test"),
				"DB_USER": jsii.String("test"),
				"DB_PASS": jsii.String("test"),
			},
		},
	)
}

func NewPostFunction(stack awscdk.Stack) awslambda.Function {
	dockerImageCode := awslambda.DockerImageCode_FromImageAsset(
		jsii.String("../../../"), &awslambda.AssetImageCodeProps{
			File: jsii.String("./lambdas/facts/app-v1-post/Dockerfile"),
		},
	)
	return awslambda.NewDockerImageFunction(
		stack, jsii.String("PostFactsFunctionV1"), &awslambda.DockerImageFunctionProps{
			Code: dockerImageCode,
			Environment: &map[string]*string{
				"DB_HOST": jsii.String("database-postgres"),
				"DB_PORT": jsii.String("5432"),
				"DB_NAME": jsii.String("test"),
				"DB_USER": jsii.String("test"),
				"DB_PASS": jsii.String("test"),
			},
		},
	)
}

func NewRestApi(stack awscdk.Stack) awscdk.Stack {
	getFunction := NewGetFunction(stack)

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

	postFunction := NewPostFunction(stack)
	factsResource.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(postFunction, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{},
	)
	return stack
}

func NewStack(scope constructs.Construct, id string, props *StageProps) awscdk.Stack {
	stack := awscdk.NewStack(
		scope, &id, &awscdk.StackProps{Env: &awscdk.Environment{Account: &props.account, Region: &props.region}},
	)

	return NewRestApi(stack)
}
