package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// The properties for the RagAppsyncStepfnOpensearchProps class.
// Experimental.
type RagAppsyncStepfnOpensearchProps struct {
	// Cognito user pool used for authentication.
	// Default: - None.
	//
	// Experimental.
	CognitoUserPool awscognito.IUserPool `field:"required" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// Index name for the OpenSearch Service.
	// Default: - None.
	//
	// Experimental.
	OpenSearchIndexName *string `field:"required" json:"openSearchIndexName" yaml:"openSearchIndexName"`
	// Optional user provided props to override the default props for the S3 Bucket.
	//
	// Providing both this and `existingInputAssetsBucketObj` will cause an error.
	// Default: - Default props are used.
	//
	// Experimental.
	BucketInputsAssetsProps *awss3.BucketProps `field:"optional" json:"bucketInputsAssetsProps" yaml:"bucketInputsAssetsProps"`
	// Optional user provided props to override the default props for the S3 Bucket.
	//
	// Providing both this and `existingProcessedAssetsBucketObj` will cause an error.
	// Default: - Default props are used.
	//
	// Experimental.
	BucketProcessedAssetsProps *awss3.BucketProps `field:"optional" json:"bucketProcessedAssetsProps" yaml:"bucketProcessedAssetsProps"`
	// Optional.
	//
	// Allows to provide Embeddings custom lambda code
	// and settings instead of the existing.
	// Experimental.
	CustomEmbeddingsDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customEmbeddingsDockerLambdaProps" yaml:"customEmbeddingsDockerLambdaProps"`
	// Optional.
	//
	// Allows to provide File Transformer custom lambda code
	// and settings instead of the existing.
	// Experimental.
	CustomFileTransformerDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customFileTransformerDockerLambdaProps" yaml:"customFileTransformerDockerLambdaProps"`
	// Optional.
	//
	// Allows to provide Input Validation custom lambda code
	// and settings instead of the existing.
	// Experimental.
	CustomInputValidationDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customInputValidationDockerLambdaProps" yaml:"customInputValidationDockerLambdaProps"`
	// Optional Existing instance of an EventBridge bus.
	//
	// If not provided, the construct will create one.
	// Default: - None.
	//
	// Experimental.
	ExistingBusInterface awsevents.IEventBus `field:"optional" json:"existingBusInterface" yaml:"existingBusInterface"`
	// Existing instance of S3 Bucket object, providing both this and `bucketInputsAssetsProps` will cause an error.
	// Default: - None.
	//
	// Experimental.
	ExistingInputAssetsBucketObj awss3.IBucket `field:"optional" json:"existingInputAssetsBucketObj" yaml:"existingInputAssetsBucketObj"`
	// Existing merged Appsync GraphQL api.
	// Default: - None.
	//
	// Experimental.
	ExistingMergedApi awsappsync.CfnGraphQLApi `field:"optional" json:"existingMergedApi" yaml:"existingMergedApi"`
	// Optional existing Amazon OpenSearch Service domain.
	// Default: - None.
	//
	// Experimental.
	ExistingOpensearchDomain awsopensearchservice.IDomain `field:"optional" json:"existingOpensearchDomain" yaml:"existingOpensearchDomain"`
	// Optional existing Amazon Amazon OpenSearch Serverless collection.
	// Default: - None.
	//
	// Experimental.
	ExistingOpensearchServerlessCollection awsopensearchserverless.CfnCollection `field:"optional" json:"existingOpensearchServerlessCollection" yaml:"existingOpensearchServerlessCollection"`
	// Existing instance of S3 Bucket object, providing both this and `bucketProcessedAssetsProps` will cause an error.
	// Default: - None.
	//
	// Experimental.
	ExistingProcessedAssetsBucketObj awss3.IBucket `field:"optional" json:"existingProcessedAssetsBucketObj" yaml:"existingProcessedAssetsBucketObj"`
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
	// Enable observability.
	//
	// Warning: associated cost with the services
	// used. Best practice to enable by default.
	// Default: - true.
	//
	// Experimental.
	Observability *bool `field:"optional" json:"observability" yaml:"observability"`
	// Optional.
	//
	// SecretsManager secret to authenticate against the OpenSearch Service domain if
	// domain is configured with Username/Password.
	// Default: - None.
	//
	// Experimental.
	OpenSearchSecret awssecretsmanager.ISecret `field:"optional" json:"openSearchSecret" yaml:"openSearchSecret"`
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

