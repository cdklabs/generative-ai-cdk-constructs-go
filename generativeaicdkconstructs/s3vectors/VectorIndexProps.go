package s3vectors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a VectorIndex resource.
// Experimental.
type VectorIndexProps struct {
	// A dimension is the number of values in a vector. A larger dimension needs more storage.
	//
	// All vectors added to the index must have exactly this number of values.
	// Must be an integer between 1 and 4096.
	// Experimental.
	Dimension *float64 `field:"required" json:"dimension" yaml:"dimension"`
	// The vector bucket to use for the vector index.
	// Experimental.
	VectorBucket IVectorBucket `field:"required" json:"vectorBucket" yaml:"vectorBucket"`
	// The data type of the vectors to be inserted into the vector index.
	// Default: - FLOAT_32.
	//
	// Experimental.
	DataType VectorIndexDataType `field:"optional" json:"dataType" yaml:"dataType"`
	// The distance metric to be used for similarity search.
	// Default: - COSINE.
	//
	// Experimental.
	DistanceMetric VectorIndexDistanceMetric `field:"optional" json:"distanceMetric" yaml:"distanceMetric"`
	// The kind of server-side encryption to apply to this index.
	//
	// If you choose KMS, you can specify a KMS key via `encryptionKey`. If
	// encryption key is not specified, a key will automatically be created.
	// Default: - `KMS` if `encryptionKey` is specified, or `S3_MANAGED` otherwise.
	//
	// Experimental.
	Encryption VectorIndexEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// External KMS key to use for index encryption.
	//
	// The `encryption` property must be either not specified or set to `KMS`.
	// An error will be emitted if `encryption` is set to `S3_MANAGED`.
	// Default: - If `encryption` is set to `KMS` and this property is undefined,
	// a new KMS key will be created and associated with this index.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Non-filterable metadata keys allow you to enrich vectors with additional context during storage and retrieval.
	//
	// Unlike default metadata keys, these keys can't be used as query filters.
	// Non-filterable metadata keys can be retrieved but can't be searched, queried, or filtered.
	// You can access non-filterable metadata keys of your vectors after finding the vectors.
	// Default: - All metadata attached to vectors is filterable and can be used as filters in a similarity query.
	//
	// Experimental.
	NonFilterableMetadataKeys *[]*string `field:"optional" json:"nonFilterableMetadataKeys" yaml:"nonFilterableMetadataKeys"`
	// The name of the vector index.
	// Default: - Assigned by CloudFormation (recommended).
	//
	// Experimental.
	VectorIndexName *string `field:"optional" json:"vectorIndexName" yaml:"vectorIndexName"`
}

