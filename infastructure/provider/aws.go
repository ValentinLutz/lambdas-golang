package provider

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"root/infastructure/util"
)

type AwsProviderConfig struct {
	Region      string
	Environment string
	Profile     string
	Resource    string
	Commit      string
	Bucket      string
}

func NewS3Backend(scope constructs.Construct, config AwsProviderConfig) cdktf.S3Backend {
	stackName := util.StackName(config.Resource, config.Region, config.Environment)

	return cdktf.NewS3Backend(scope, &cdktf.S3BackendConfig{
		Region:        jsii.String(config.Region),
		Profile:       jsii.String(config.Profile),
		Key:           jsii.String(stackName + ".terraform.tfstate"),
		DynamodbTable: jsii.String("terraform-lock"),
		Encrypt:       jsii.Bool(true),
		Bucket:        jsii.String(config.Bucket),
	})
}

func NewAwsProvider(scope constructs.Construct, config AwsProviderConfig) provider.AwsProvider {
	stackName := util.StackName(config.Resource, config.Region, config.Environment)

	awsProvider := provider.NewAwsProvider(scope, jsii.String("aws-provider"), &provider.AwsProviderConfig{
		Region:  jsii.String(config.Region),
		Profile: jsii.String(config.Profile),
		DefaultTags: &[]provider.AwsProviderDefaultTags{{&map[string]*string{
			"custom:environment": jsii.String(config.Environment),
			"custom:region":      jsii.String(config.Region),
			"custom:resource":    jsii.String(config.Resource),
			"custom:iac":         jsii.String("cdktf"),
			"custom:commit":      jsii.String(config.Commit),
			"custom:stack":       jsii.String(stackName),
		}}},
	})

	return awsProvider
}
