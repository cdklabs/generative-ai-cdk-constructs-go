package s3vectors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes for specifying an imported S3 vector bucket.
// Experimental.
type VectorIndexAttributes struct {
	// The ARN of the vector index.
	// Experimental.
	VectorIndexArn *string `field:"required" json:"vectorIndexArn" yaml:"vectorIndexArn"`
	// The timestamp when the vector index was created, in ISO 8601 format.
	// Default: undefined - No creation time is provided.
	//
	// Experimental.
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Optional KMS encryption key associated with this vector index.
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
}

