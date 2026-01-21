package s3vectors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes for specifying an imported S3 vector bucket.
// Experimental.
type VectorBucketAttributes struct {
	// The ARN of the vector bucket.
	// Experimental.
	VectorBucketArn *string `field:"required" json:"vectorBucketArn" yaml:"vectorBucketArn"`
	// The account this existing bucket belongs to.
	// Default: - it's assumed the bucket belongs to the same account as the scope it's being imported into.
	//
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The timestamp when the cluster was created, in ISO 8601 format.
	// Default: undefined - No creation time is provided.
	//
	// Experimental.
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// The encryption key associated with this bucket.
	// Default: - No encryption key.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The region this existing bucket is in.
	//
	// Features that require the region (e.g. `bucketWebsiteUrl`) won't fully work
	// if the region cannot be correctly inferred.
	// Default: - it's assumed the bucket is in the same region as the scope it's being imported into.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

