package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for configuring a Lambda-based custom transformation.
// Experimental.
type LambdaCustomTransformationProps struct {
	// The Lambda function to use for custom document processing.
	// Experimental.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// An S3 bucket URL/path to store input documents for Lambda processing and to store the output of the processed documents.
	//
	// Example:
	//   "s3://my-bucket/chunk-processor/"
	//
	// Experimental.
	S3BucketUri *string `field:"required" json:"s3BucketUri" yaml:"s3BucketUri"`
}

