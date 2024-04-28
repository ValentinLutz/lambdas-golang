package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dynamodbtable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamaccesskey"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamuser"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucket"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketlifecycleconfiguration"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketpublicaccessblock"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketserversideencryptionconfiguration"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketversioning"
	"github.com/cdktf/cdktf-provider-random-go/random/v11/id"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"root/libraries/cdkutil"
	"root/libraries/cdkutil/provider"
)

func NewStack(scope constructs.Construct, region string, env string) {
	stageConfig, err := NewStageConfig(region, env)
	if err != nil {
		panic(err)
	}

	resource := "bootstrap"
	stack := cdktf.NewTerraformStack(scope, jsii.String(cdkutil.StackName(stageConfig.Region, stageConfig.Environment, resource)))

	awsProviderConfig := provider.AwsProviderConfig{
		Region:      stageConfig.Region,
		Environment: stageConfig.Environment,
		Profile:     stageConfig.Profile,
		Resource:    resource,
		Commit:      cdkutil.GitCommit(),
		Bucket:      stageConfig.Bucket,
	}

	if stageConfig.Bucket != nil {
		provider.NewS3Backend(stack, awsProviderConfig)
	}

	provider.NewAwsProvider(stack, awsProviderConfig)
	provider.NewRandomProvider(stack)

	bootstrap(stack)
}

func bootstrap(stack cdktf.TerraformStack) {
	iamUser := iamuser.NewIamUser(stack, jsii.String("terraform-user"), &iamuser.IamUserConfig{
		Name: jsii.String("terraform"),
	})

	iamAccessKey := iamaccesskey.NewIamAccessKey(stack, jsii.String("terraform_access_key"), &iamaccesskey.IamAccessKeyConfig{
		User: iamUser.Id(),
	})

	terraformStateSuffix := id.NewId(stack, jsii.String("terraform-state-suffix"), &id.IdConfig{
		ByteLength: jsii.Number(8),
	})

	bucket := s3bucket.NewS3Bucket(stack, jsii.String("terraform-state"), &s3bucket.S3BucketConfig{
		Bucket: jsii.String("terraform-state-" + *terraformStateSuffix.Dec()),
		Lifecycle: &cdktf.TerraformResourceLifecycle{
			PreventDestroy: jsii.Bool(true),
		},
	})

	s3bucketversioning.NewS3BucketVersioningA(stack, jsii.String("terraform-state-versioning"), &s3bucketversioning.S3BucketVersioningAConfig{
		Bucket: bucket.Id(),
		VersioningConfiguration: &s3bucketversioning.S3BucketVersioningVersioningConfiguration{
			Status: jsii.String("Enabled"),
		},
	})

	s3bucketserversideencryptionconfiguration.NewS3BucketServerSideEncryptionConfigurationA(
		stack, jsii.String("terraform-state-encryption"), &s3bucketserversideencryptionconfiguration.S3BucketServerSideEncryptionConfigurationAConfig{
			Bucket: bucket.Id(),
			Rule: &[]s3bucketserversideencryptionconfiguration.S3BucketServerSideEncryptionConfigurationRuleA{
				{
					ApplyServerSideEncryptionByDefault: &s3bucketserversideencryptionconfiguration.S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultA{
						SseAlgorithm: jsii.String("AES256"),
					},
					BucketKeyEnabled: jsii.Bool(true),
				},
			},
		})

	s3bucketlifecycleconfiguration.NewS3BucketLifecycleConfiguration(
		stack, jsii.String("terraform-state-lifecycle"), &s3bucketlifecycleconfiguration.S3BucketLifecycleConfigurationConfig{
			Bucket: bucket.Id(),
			Rule: &[]s3bucketlifecycleconfiguration.S3BucketLifecycleConfigurationRule{
				{
					Id:     jsii.String("terraform-state-lifecycle-rule-1"),
					Status: jsii.String("Enabled"),
					NoncurrentVersionExpiration: &s3bucketlifecycleconfiguration.S3BucketLifecycleConfigurationRuleNoncurrentVersionExpiration{
						NewerNoncurrentVersions: jsii.String("10"),
						NoncurrentDays:          jsii.Number(1),
					},
				},
			},
		})

	s3bucketpublicaccessblock.NewS3BucketPublicAccessBlock(stack, jsii.String("terraform-state-block-access"), &s3bucketpublicaccessblock.S3BucketPublicAccessBlockConfig{
		Bucket:                bucket.Id(),
		BlockPublicAcls:       jsii.Bool(true),
		BlockPublicPolicy:     jsii.Bool(true),
		IgnorePublicAcls:      jsii.Bool(true),
		RestrictPublicBuckets: jsii.Bool(true),
	})

	dynamodbtable.NewDynamodbTable(stack, jsii.String("terraform-lock"), &dynamodbtable.DynamodbTableConfig{
		Name:        jsii.String("terraform-lock"),
		BillingMode: jsii.String("PAY_PER_REQUEST"),
		HashKey:     jsii.String("LockID"),
		Attribute: &[]dynamodbtable.DynamodbTableAttribute{
			{Name: jsii.String("LockID"), Type: jsii.String("S")},
		},
	})

	cdktf.NewTerraformOutput(stack, jsii.String("access_key_id"), &cdktf.TerraformOutputConfig{
		Value: iamAccessKey.Id(),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("secret_access_key"), &cdktf.TerraformOutputConfig{
		Value:     iamAccessKey.Secret(),
		Sensitive: jsii.Bool(true),
	})
	cdktf.NewTerraformOutput(stack, jsii.String("bucket"), &cdktf.TerraformOutputConfig{
		Value: bucket.Bucket(),
	})
}
