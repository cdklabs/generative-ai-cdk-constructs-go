package s3vectors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Experimental.
type VectorBucketPolicyProps struct {
	// The S3 vector bucket that the policy applies to.
	// Experimental.
	Bucket IVectorBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Policy document to apply to the bucket.
	// Default: - A new empty PolicyDocument will be created.
	//
	// Experimental.
	Document awsiam.PolicyDocument `field:"optional" json:"document" yaml:"document"`
	// Policy to apply when the policy is removed from this stack.
	// Default: - RemovalPolicy.DESTROY.
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

