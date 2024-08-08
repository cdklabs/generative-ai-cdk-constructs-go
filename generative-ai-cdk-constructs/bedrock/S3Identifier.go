package bedrock


// Result of the bind when `S3ApiSchema` is used.
// Experimental.
type S3Identifier struct {
	// The name of the S3 bucket.
	// Experimental.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The S3 object key containing the resource.
	// Experimental.
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

