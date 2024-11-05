package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// ****************************************************************************                       PROPS FOR NEW CONSTRUCT ***************************************************************************.
// Experimental.
type PromptProps struct {
	// The name of the prompt.
	// Experimental.
	PromptName *string `field:"required" json:"promptName" yaml:"promptName"`
	// The Prompt Variant that will be used by default.
	// Default: - No default variant provided.
	//
	// Experimental.
	DefaultVariant PromptVariant `field:"optional" json:"defaultVariant" yaml:"defaultVariant"`
	// A description of what the prompt does.
	// Default: - No description provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key that the prompt is encrypted with.
	// Default: - AWS owned and managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The variants of your prompt.
	//
	// Variants can use different messages, models,
	// or configurations so that you can compare their outputs to decide the best
	// variant for your use case. Maximum of 3 variants.
	// Experimental.
	Variants *[]PromptVariant `field:"optional" json:"variants" yaml:"variants"`
}

