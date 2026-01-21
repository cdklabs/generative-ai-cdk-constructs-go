package bedrock


// ****************************************************************************                     ATTRS FOR IMPORTED CONSTRUCT ***************************************************************************.
// Experimental.
type ApplicationInferenceProfileAttributes struct {
	// The ARN of the application inference profile.
	// Experimental.
	InferenceProfileArn *string `field:"required" json:"inferenceProfileArn" yaml:"inferenceProfileArn"`
	// The ID or Amazon Resource Name (ARN) of the inference profile.
	// Experimental.
	InferenceProfileIdentifier *string `field:"required" json:"inferenceProfileIdentifier" yaml:"inferenceProfileIdentifier"`
}

