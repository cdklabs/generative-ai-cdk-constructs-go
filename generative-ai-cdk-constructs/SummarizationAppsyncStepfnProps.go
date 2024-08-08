package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type SummarizationAppsyncStepfnProps struct {
	// Required.
	//
	// Cognito user pool used for authentication.
	// Default: - None.
	//
	// Experimental.
	CognitoUserPool awscognito.IUserPool `field:"required" json:"cognitoUserPool" yaml:"cognitoUserPool"`
	// Optional.
	//
	// User provided props to override the default props for the S3 Bucket.
	// Providing both this and `existingInputAssetsBucketObj` will cause an error.
	// Default: - Default props are used.
	//
	// Experimental.
	BucketInputsAssetsProps *awss3.BucketProps `field:"optional" json:"bucketInputsAssetsProps" yaml:"bucketInputsAssetsProps"`
	// Optional.
	//
	// User provided props to override the default props for the S3 Bucket.
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
	CustomDocumentReaderDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customDocumentReaderDockerLambdaProps" yaml:"customDocumentReaderDockerLambdaProps"`
	// Optional.
	//
	// Allows to provide Input Validation custom lambda code
	// and settings instead of the existing.
	// Experimental.
	CustomInputValidationDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customInputValidationDockerLambdaProps" yaml:"customInputValidationDockerLambdaProps"`
	// Optional.
	//
	// Allows to provide File Transformer custom lambda code
	// and settings instead of the existing.
	// Experimental.
	CustomSummaryGeneratorDockerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customSummaryGeneratorDockerLambdaProps" yaml:"customSummaryGeneratorDockerLambdaProps"`
	// Optional.
	//
	// A new custom EventBus is created with provided props.
	// Providing existingEventBusInterface and eventBusProps both will result in validation error.
	// Default: - None.
	//
	// Experimental.
	EventBusProps *awsevents.EventBusProps `field:"optional" json:"eventBusProps" yaml:"eventBusProps"`
	// Optional.
	//
	// Existing instance of EventBus. The summary construct integrate appsync with event bridge'
	// to route the request to step functions.
	// Default: - None.
	//
	// Experimental.
	ExistingBusInterface awsevents.IEventBus `field:"optional" json:"existingBusInterface" yaml:"existingBusInterface"`
	// Optional.
	//
	// Existing s3 Bucket to store the input document which needs to be summarized.
	// pdf is the supported input document format. If transformed (txt format) file is
	// available then this bucket is optional.
	// Default: - None.
	//
	// Experimental.
	ExistingInputAssetsBucketObj awss3.IBucket `field:"optional" json:"existingInputAssetsBucketObj" yaml:"existingInputAssetsBucketObj"`
	// Optional - Existing merged Appsync GraphQL api.
	// Default: - None.
	//
	// Experimental.
	ExistingMergedApi awsappsync.CfnGraphQLApi `field:"optional" json:"existingMergedApi" yaml:"existingMergedApi"`
	// Optional.
	//
	// This bucket stores the transformed (txt) assets for generating summary.
	// If None is provided then this contruct will create one.
	// Default: - None.
	//
	// Experimental.
	ExistingProcessedAssetsBucketObj awss3.IBucket `field:"optional" json:"existingProcessedAssetsBucketObj" yaml:"existingProcessedAssetsBucketObj"`
	// Optional.
	//
	// Security group for the lambda function which this construct will use.
	// If no exisiting security group is provided it will create one from the vpc.
	// Default: - none.
	//
	// Experimental.
	ExistingSecurityGroup awsec2.ISecurityGroup `field:"optional" json:"existingSecurityGroup" yaml:"existingSecurityGroup"`
	// Optional.
	//
	// An existing VPC can be used to deploy the construct.
	// Providing both this and vpcProps is an error.
	// Default: - none.
	//
	// Experimental.
	ExistingVpc awsec2.IVpc `field:"optional" json:"existingVpc" yaml:"existingVpc"`
	// Optional.
	//
	// The summary construct transform the input document into txt format. If the
	// transformation is not required then this flag can be set to false. If set to true
	// then a transformed asset bucket is created which transform the input document from
	// input asset bucket to txt format.
	// Default: - False.
	//
	// Experimental.
	IsFileTransformationRequired *string `field:"optional" json:"isFileTransformationRequired" yaml:"isFileTransformationRequired"`
	// Enable observability.
	//
	// Warning: associated cost with the services
	// used. Best practice to enable by default.
	// Default: - true.
	//
	// Experimental.
	Observability *bool `field:"optional" json:"observability" yaml:"observability"`
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Optional.
	//
	// User provided Name for summary api on appsync.
	// A graphql api will be created by this construct with this name.
	// Default: 'summaryApi'.
	//
	// Experimental.
	SummaryApiName *string `field:"optional" json:"summaryApiName" yaml:"summaryApiName"`
	// Optional.
	//
	// Chain type defines how to pass the document to LLM.
	// there are three types of chain types.
	// Stuff: Simply "stuff" all your documents into a single prompt.
	// Map-reduce: Summarize each document on it's own in a "map" step and then "reduce" the summaries into a final summary
	// Refine :  This constructs a response by looping over the input documents and iteratively updating its answer.
	// Default: - Stuff.
	//
	// Experimental.
	SummaryChainType *string `field:"optional" json:"summaryChainType" yaml:"summaryChainType"`
	// Optional.
	//
	// The construct creates a custom VPC based on vpcProps.
	// Providing both this and existingVpc is an error.
	// Default: - none.
	//
	// Experimental.
	VpcProps *awsec2.VpcProps `field:"optional" json:"vpcProps" yaml:"vpcProps"`
}

