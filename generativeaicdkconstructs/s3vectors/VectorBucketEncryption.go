package s3vectors


// What kind of server-side encryption to apply to this bucket.
// Experimental.
type VectorBucketEncryption string

const (
	// Server-side encryption with a master key managed by S3.
	// Experimental.
	VectorBucketEncryption_S3_MANAGED VectorBucketEncryption = "S3_MANAGED"
	// Server-side encryption with a KMS key managed by the user.
	//
	// If `encryptionKey` is specified, this key will be used, otherwise, one will be defined.
	// Experimental.
	VectorBucketEncryption_KMS VectorBucketEncryption = "KMS"
)

