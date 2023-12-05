package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewGetFunction(stack awscdk.Stack, config *StageConfig) awslambda.Function {
	env := map[string]*string{
		"DB_HOST":      &config.databaseProps.host,
		"DB_PORT":      &config.databaseProps.port,
		"DB_NAME":      &config.databaseProps.name,
		"DB_SECRET_ID": &config.databaseProps.secret,
	}
	if config.endpointUrl != nil {
		env["AWS_ENDPOINT_URL"] = config.endpointUrl
	}

	return awslambda.NewFunction(
		stack, NewIdWithStage(config, "facts-function-v1-get"), &awslambda.FunctionProps{
			Code: awslambda.Code_FromAsset(
				jsii.String("../lambda-v1-get"),
				&awss3assets.AssetOptions{},
			),
			Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
			Handler:      jsii.String("bootstrap"),
			Architecture: config.lambdaConfig.architecture,
			Environment:  &env,
			Tracing:      awslambda.Tracing_ACTIVE,
		},
	)
}

func NewPostFunction(stack awscdk.Stack, config *StageConfig) awslambda.Function {
	env := map[string]*string{
		"DB_HOST":      &config.databaseProps.host,
		"DB_PORT":      &config.databaseProps.port,
		"DB_NAME":      &config.databaseProps.name,
		"DB_SECRET_ID": &config.databaseProps.secret,
	}
	if config.endpointUrl != nil {
		env["AWS_ENDPOINT_URL"] = config.endpointUrl
	}

	return awslambda.NewFunction(
		stack, NewIdWithStage(config, "facts-function-v1-post"), &awslambda.FunctionProps{
			Code: awslambda.Code_FromAsset(
				jsii.String("../lambda-v1-post"),
				&awss3assets.AssetOptions{},
			),
			Runtime:      awslambda.Runtime_PROVIDED_AL2(),
			Handler:      jsii.String("bootstrap"),
			Architecture: config.lambdaConfig.architecture,
			Environment:  &env,
			Tracing:      awslambda.Tracing_ACTIVE,
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
