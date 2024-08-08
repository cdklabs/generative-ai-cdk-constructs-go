package bedrock


// Result of binding `ApiSchema` into an `ActionGroup`.
// Experimental.
type ApiSchemaConfig struct {
	// The JSON or YAML-formatted payload defining the OpenAPI schema for the action group.
	//
	// (mutually exclusive with `s3`).
	// Experimental.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// Contains details about the S3 object containing the OpenAPI schema for the action group.
	//
	// (mutually exclusive with `payload`).
	// Experimental.
	S3 *S3Identifier `field:"optional" json:"s3" yaml:"s3"`
}

