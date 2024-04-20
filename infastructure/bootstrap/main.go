package bootstrap

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dynamodbtable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamaccesskey"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamuser"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucket"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketpublicaccessblock"
	"github.com/cdktf/cdktf-provider-random-go/random/v11/id"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"root/infastructure/provider"
	"root/infastructure/util"
)

func NewStack(scope constructs.Construct, region string, env string) {
	stageConfig, err := NewStageConfig(region, env)
	if err != nil {
		panic(err)
	}

	resource := "bootstrap"
	stack := cdktf.NewTerraformStack(scope, jsii.String(util.StackName(resource, region, env)))

	awsProviderConfig := provider.AwsProviderConfig{
		Region:      stageConfig.Region,
		Environment: stageConfig.Environment,
		Profile:     stageConfig.Profile,
		Resource:    resource,
		Commit:      util.GitCommit(),
		Bucket:      *stageConfig.Bucket,
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
		Versioning: &s3bucket.S3BucketVersioning{
			Enabled: jsii.Bool(true),
		},
		ServerSideEncryptionConfiguration: &s3bucket.S3BucketServerSideEncryptionConfiguration{
			Rule: &s3bucket.S3BucketServerSideEncryptionConfigurationRule{
				ApplyServerSideEncryptionByDefault: &s3bucket.S3BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault{
					SseAlgorithm: jsii.String("AES256"),
				},
			},
		},
	})

	s3bucketpublicaccessblock.NewS3BucketPublicAccessBlock(stack, jsii.String("terraform-block-access"), &s3bucketpublicaccessblock.S3BucketPublicAccessBlockConfig{
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
