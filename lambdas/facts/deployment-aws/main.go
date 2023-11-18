package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

type TestStackProps struct {
	awscdk.StackProps
}

func NewStack(scope constructs.Construct, id string, props *TestStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	function := awslambda.NewDockerImageFunction(
		stack, jsii.String("FactsFunction"), &awslambda.DockerImageFunctionProps{
			Code: awslambda.DockerImageCode_FromImageAsset(
				jsii.String("../"), &awslambda.AssetImageCodeProps{
					File: jsii.String("Dockerfile"),
				},
			),
			Environment: &map[string]*string{
				"DB_HOST": jsii.String("database-postgres"),
				"DB_PORT": jsii.String("5432"),
				"DB_NAME": jsii.String("test"),
				"DB_USER": jsii.String("test"),
				"DB_PASS": jsii.String("test"),
			},
		},
	)

	lambdaRestApi := awsapigateway.NewLambdaRestApi(
		stack, jsii.String("FactsRestApi"), &awsapigateway.LambdaRestApiProps{
			Handler: function,
			Proxy:   jsii.Bool(false),
		},
	)

	factsResource := lambdaRestApi.Root().AddResource(
		jsii.String("facts"), &awsapigateway.ResourceOptions{},
	)
	factsResource.AddMethod(jsii.String("GET"), nil, &awsapigateway.MethodOptions{})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewStack(
		app, "stack", &TestStackProps{
			awscdk.StackProps{
				Env: env(),
			},
		},
	)

	app.Synth(nil)
}
