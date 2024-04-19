package bootstrap

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dynamodbtable"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamaccesskey"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamuser"
	awsprov "github.com/cdktf/cdktf-provider-aws-go/aws/v19/provider"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucket"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/s3bucketpublicaccessblock"
	"github.com/cdktf/cdktf-provider-random-go/random/v11/id"
	randprov "github.com/cdktf/cdktf-provider-random-go/random/v11/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

const ResourceName = "bootstrap"

func NewStack(scope constructs.Construct, region string, env string) {
	stageConfig, err := NewStageConfig(region, env)
	if err != nil {
		panic(err)
	}

	stackName := ResourceName + "-" + region + "-" + env
	stack := cdktf.NewTerraformStack(scope, jsii.String(stackName))

	commitHash := cdktf.NewTerraformVariable(stack, jsii.String("commit_hash"), &cdktf.TerraformVariableConfig{
		Type: jsii.String("string"),
	})

	if stageConfig.Bucket != nil {
		cdktf.NewS3Backend(stack, &cdktf.S3BackendConfig{
			Region:        jsii.String(stageConfig.Region),
			Bucket:        stageConfig.Bucket,
			Key:           jsii.String(stackName + ".terraform.tfstate"),
			DynamodbTable: jsii.String("terraform-lock"),
			Encrypt:       jsii.Bool(true),
		})
	}

	awsprov.NewAwsProvider(stack, jsii.String("aws-provider"), &awsprov.AwsProviderConfig{
		Region:  jsii.String(stageConfig.Region),
		Profile: jsii.String(stageConfig.Profile),
		DefaultTags: &[]awsprov.AwsProviderDefaultTags{{&map[string]*string{
			"custom:environment": jsii.String(stageConfig.Environment),
			"custom:region":      jsii.String(stageConfig.Region),
			"custom:resource":    jsii.String(ResourceName),
			"custom:stack":       jsii.String(stackName),
			"custom:iac":         jsii.String("cdktf"),
			"custom:commit":      commitHash.StringValue(),
		}}},
	})

	randprov.NewRandomProvider(stack, jsii.String("random-provider"), &randprov.RandomProviderConfig{})

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
