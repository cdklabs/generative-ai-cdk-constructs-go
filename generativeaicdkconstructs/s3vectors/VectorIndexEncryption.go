package s3vectors


// What kind of encryption to apply to this index.
//
// By default, if you don't specify, all new vectors in Amazon S3 vector indexes
// use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically AES256.
// Experimental.
type VectorIndexEncryption string

const (
	// Encryption with a master key managed by S3.
	// Experimental.
	VectorIndexEncryption_S3_MANAGED VectorIndexEncryption = "S3_MANAGED"
	// Encryption with a KMS key managed by the user.
	// Experimental.
	VectorIndexEncryption_KMS VectorIndexEncryption = "KMS"
)

