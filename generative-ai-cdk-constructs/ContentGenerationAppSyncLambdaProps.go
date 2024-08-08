package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// The properties for the ContentGenerationAppSyncLambdaProps class.
// Experimental.
type ContentGenerationAppSyncLambdaProps struct {
	// Cognito user pool used for authentication.
	// Default: - None.
	//
	// Experimental.
	CognitoUserPool awscognito.IUserPool `field:"required" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// Optional.
	//
	// Allows to provide custom lambda code
	// and settings instead of the existing.
	// Experimental.
	CustomDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customDockerLambdaProps" yaml:"customDockerLambdaProps"`
	// Optional Existing instance of an EventBridge bus.
	//
	// If not provided, the construct will create one.
	// Default: - None.
	//
	// Experimental.
	ExistingBusInterface awsevents.IEventBus `field:"optional" json:"existingBusInterface" yaml:"existingBusInterface"`
	// Existing instance of S3 Bucket object, providing both this and `generatedAssetsBucketProps` will cause an error.
	// Default: - None.
	//
	// Experimental.
	ExistingGeneratedAssetsBucketObj awss3.IBucket `field:"optional" json:"existingGeneratedAssetsBucketObj" yaml:"existingGeneratedAssetsBucketObj"`
	// Existing merged Appsync GraphQL api.
	// Default: - None.
	//
	// Experimental.
	ExistingMergedApi awsappsync.CfnGraphQLApi `field:"optional" json:"existingMergedApi" yaml:"existingMergedApi"`
	// Optional existing security group allowing access to opensearch.
	//
	// Used by the lambda functions
	// built by this construct. If not provided, the construct will create one.
	// Default: - none.
	//
	// Experimental.
	ExistingSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"existingSecurityGroup" yaml:"existingSecurityGroup"`
	// Optional An existing VPC in which to deploy the construct.
	//
	// Providing both this and
	// vpcProps is an error.
	// Default: - none.
	//
	// Experimental.
	ExistingVpc awsec2.IVpc `field:"optional" json:"existingVpc" yaml:"existingVpc"`
	// Optional user provided props to override the default props for the S3 Bucket.
	//
	// Providing both this and `existingGeneratedAssetsBucketObj` will cause an error.
	// Default: - Default props are used.
	//
	// Experimental.
	GeneratedAssetsBucketProps *awss3.BucketProps `field:"optional" json:"generatedAssetsBucketProps" yaml:"generatedAssetsBucketProps"`
	// Optional.
	//
	// Allows a user to configure
	// Lambda provisioned concurrency for consistent performance.
	// Experimental.
	LambdaProvisionedConcurrency *float64 `field:"optional" json:"lambdaProvisionedConcurrency" yaml:"lambdaProvisionedConcurrency"`
	// Enable observability.
	//
	// Warning: associated cost with the services
	// used. Best practive to enable by default.
	// Default: - true.
	//
	// Experimental.
	Observability *bool `field:"optional" json:"observability" yaml:"observability"`
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Optional custom properties for a VPC the construct will create.
	//
	// This VPC will
	// be used by the Lambda functions the construct creates. Providing
	// both this and existingVpc is an error.
	// Default: - none.
	//
	// Experimental.
	VpcProps *awsec2.VpcProps `field:"optional" json:"vpcProps" yaml:"vpcProps"`
}

