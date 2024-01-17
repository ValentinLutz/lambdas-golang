package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewVpc(stack awscdk.Stack, config *StageConfig) awsec2.Vpc {
	vpc := awsec2.NewVpc(
		stack, jsii.String("LoadTestVpc"), &awsec2.VpcProps{
			MaxAzs: jsii.Number(1),
		},
	)

	return vpc
}

func NewTest(stack awscdk.Stack, config *StageConfig) awscdk.Stack {
	vpc := NewVpc(stack, config)

	//vpc := awsec2.Vpc_FromVpcAttributes(
	//	stack, jsii.String("DefaultVpc"), &awsec2.VpcAttributes{
	//		VpcId: jsii.String("vpc-804e10eb"),
	//		AvailabilityZones: &[]*string{
	//			jsii.String("eu-central-1a"),
	//			jsii.String("eu-central-1b"),
	//			jsii.String("eu-central-1c"),
	//		},
	//		PublicSubnetIds: &[]*string{
	//			jsii.String("subnet-17b3b07c"),
	//			jsii.String("subnet-d0433aad"),
	//			jsii.String("subnet-bebcf0f3"),
	//		},
	//		VpcCidrBlock: jsii.String("172.31.0.0/16"),
	//	},
	//)

	// port 443 is in ingress rule by default
	securityGroup := awsec2.NewSecurityGroup(
		stack, jsii.String("LoadTestSecurityGroup"), &awsec2.SecurityGroupProps{
			Vpc:              vpc,
			AllowAllOutbound: jsii.Bool(true),
		},
	)
	//securityGroup.AddIngressRule(
	//	awsec2.Peer_AnyIpv4(), awsec2.Port_Tcp(jsii.Number(22)), jsii.String("SSH"), jsii.Bool(false),
	//)

	awsec2.NewInterfaceVpcEndpoint(
		stack, jsii.String("LoadTestVpce"), &awsec2.InterfaceVpcEndpointProps{
			Vpc:     vpc,
			Service: awsec2.InterfaceVpcEndpointAwsService_APIGATEWAY(),
			SecurityGroups: &[]awsec2.ISecurityGroup{
				securityGroup,
			},
			Subnets: &awsec2.SubnetSelection{
				SubnetType: awsec2.SubnetType_PUBLIC,
			},
			PrivateDnsEnabled: jsii.Bool(true),
		},
	)

	awscodebuild.NewProject(
		stack, jsii.String("LoadTestProject"), &awscodebuild.ProjectProps{
			Vpc:       vpc,
			BuildSpec: awscodebuild.BuildSpec_FromAsset(jsii.String("./test.yaml")),
			Environment: &awscodebuild.BuildEnvironment{
				BuildImage: awscodebuild.LinuxBuildImage_AMAZON_LINUX_2_5(),
			},
			//SecurityGroups: nil,
		},
	)

	//sshKeyPair := awsec2.NewKeyPair(
	//	stack, jsii.String("LoadTestKeyPair"), &awsec2.KeyPairProps{},
	//)

	//instance := awsec2.NewInstance(
	//	stack, jsii.String("LoadTestInstance"), &awsec2.InstanceProps{
	//		InstanceType: awsec2.InstanceType_Of(awsec2.InstanceClass_T4G, awsec2.InstanceSize_NANO),
	//		KeyPair:      sshKeyPair,
	//		MachineImage: awsec2.MachineImage_LatestAmazonLinux2023(
	//			&awsec2.AmazonLinux2023ImageSsmParameterProps{
	//				CpuType: awsec2.AmazonLinuxCpuType_ARM_64,
	//			},
	//		),
	//		Vpc:           vpc,
	//		SecurityGroup: securityGroup,
	//	},
	//)

	return stack
}

//func NewDatabase(stack awscdk.Stack, config *StageConfig, vpc awsec2.Vpc) awsrds.DatabaseInstance {
//	database := awsrds.NewDatabaseInstance(
//		stack, jsii.String("SmallDatabase"), &awsrds.DatabaseInstanceProps{
//			Vpc:                     vpc,
//			BackupRetention:         awscdk.Duration_Days(jsii.Number(0)),
//			CloudwatchLogsRetention: awslogs.RetentionDays_ONE_MONTH,
//			Engine: awsrds.DatabaseInstanceEngine_Postgres(
//				&awsrds.PostgresInstanceEngineProps{
//					Version: awsrds.PostgresEngineVersion_VER_16_1(),
//				},
//			),
//			InstanceType: awsec2.InstanceType_Of(awsec2.InstanceClass_T4G, awsec2.InstanceSize_MICRO),
//			MultiAz:      jsii.Bool(false),
//		},
//	)
//
//	return database
//}

func NewStack(scope constructs.Construct, id *string, config *StageConfig) awscdk.Stack {
	stack := awscdk.NewStack(
		scope, id, &awscdk.StackProps{Env: &awscdk.Environment{Account: &config.account, Region: &config.region}},
	)

	//vpc := NewVpc(stack, config)
	NewTest(stack, config)

	return stack
}
