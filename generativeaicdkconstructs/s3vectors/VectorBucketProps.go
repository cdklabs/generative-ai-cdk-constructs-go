package s3vectors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a VectorBucket resource.
// Experimental.
type VectorBucketProps struct {
	// Whether all objects should be automatically deleted when the bucket is removed from the stack or when the stack is deleted.
	//
	// Requires the `removalPolicy` to be set to `RemovalPolicy.DESTROY`.
	//
	// Setting `autoDeleteObjects` to true on a bucket will add `s3:PutBucketPolicy` to the
	// bucket policy. This is because during bucket deletion, the custom resource provider
	// needs to update the bucket policy by adding a deny policy for `s3:PutObject` to
	// prevent race conditions with external bucket writers.
	// Default: false.
	//
	// Experimental.
	AutoDeleteObjects *bool `field:"optional" json:"autoDeleteObjects" yaml:"autoDeleteObjects"`
	// The kind of server-side encryption to apply to this bucket.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	// Default: - `KMS` if `encryptionKey` is specified, or `S3_MANAGED` otherwise.
	//
	// Experimental.
	Encryption VectorBucketEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for bucket encryption.
	//
	// The `encryption` property must be either not specified or set to `KMS`.
	// An error will be emitted if `encryption` is set to `S3_MANAGED`.
	// Default: - If `encryption` is set to `KMS` and this property is undefined,
	// a new KMS key will be created and associated with this bucket.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Policy to apply when the bucket is removed from this stack.
	// Default: - - The bucket will be orphaned.
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Physical name of this bucket.
	// Default: - Assigned by CloudFormation (recommended).
	//
	// Experimental.
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

