package bedrock


// Experimental.
type CommonPromptVariantProps struct {
	// The model which is used to run the prompt.
	//
	// The model could be a foundation
	// model, a custom model, or a provisioned model.
	// Experimental.
	Model IInvokable `field:"required" json:"model" yaml:"model"`
	// The name of the prompt variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
}

