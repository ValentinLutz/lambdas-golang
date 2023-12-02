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
			File: jsii.String("./resources/facts/lambda-v1-get/Dockerfile"),
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
		stack, NewIdWithStage(config, "facts-function-v1-get"), &awslambda.DockerImageFunctionProps{
			Code:         dockerImageCode,
			Environment:  &env,
			Architecture: config.lambdaConfig.architecture,
		},
	)
}

func NewPostFunction(stack awscdk.Stack, config *StageConfig) awslambda.Function {
	dockerImageCode := awslambda.DockerImageCode_FromImageAsset(
		jsii.String("../../../"), &awslambda.AssetImageCodeProps{
			File: jsii.String("./resources/facts/lambda-v1-post/Dockerfile"),
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
		stack, NewIdWithStage(config, "facts-function-v1-post"), &awslambda.DockerImageFunctionProps{
			Code:         dockerImageCode,
			Environment:  &env,
			Architecture: config.lambdaConfig.architecture,
		},
	)
}

func NewRestApi(stack awscdk.Stack, config *StageConfig) awscdk.Stack {
	getFunction := NewGetFunction(stack, config)

	lambdaRestApi := awsapigateway.NewLambdaRestApi(
		stack, NewIdWithStage(config, "facts-rest-api"), &awsapigateway.LambdaRestApiProps{
			Proxy:   jsii.Bool(false),
			Handler: getFunction,
			DeployOptions: &awsapigateway.StageOptions{
				StageName: jsii.String(config.environment),
			},
		},
	)

	v1 := lambdaRestApi.Root().AddResource(
		jsii.String("v1"), &awsapigateway.ResourceOptions{},
	)

	factsResourceV1 := v1.AddResource(
		jsii.String("facts"), &awsapigateway.ResourceOptions{},
	)

	factsResourceV1.AddMethod(
		jsii.String("GET"),
		awsapigateway.NewLambdaIntegration(getFunction, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
		},
	)

	postFunction := NewPostFunction(stack, config)
	factsResourceV1.AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(postFunction, &awsapigateway.LambdaIntegrationOptions{}),
		&awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_IAM,
		},
	)
	return stack
}

func NewStack(scope constructs.Construct, id *string, config *StageConfig) awscdk.Stack {
	stack := awscdk.NewStack(
		scope, id, &awscdk.StackProps{Env: &awscdk.Environment{Account: &config.account, Region: &config.region}},
	)

	return NewRestApi(stack, config)
}
