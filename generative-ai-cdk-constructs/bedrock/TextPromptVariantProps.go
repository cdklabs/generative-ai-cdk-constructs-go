package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Experimental.
type TextPromptVariantProps struct {
	// The model which is used to run the prompt.
	//
	// The model could be a foundation
	// model, a custom model, or a provisioned model.
	// Experimental.
	Model IInvokable `field:"required" json:"model" yaml:"model"`
	// The name of the prompt variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// The variables in the prompt template that can be filled in at runtime.
	// Experimental.
	PromptVariables *[]*string `field:"optional" json:"promptVariables" yaml:"promptVariables"`
	// The text prompt.
	//
	// Variables are used by enclosing its name with double curly braces
	// as in `{{variable_name}}`.
	// Experimental.
	PromptText *string `field:"required" json:"promptText" yaml:"promptText"`
	// Inference configuration for the Text Prompt.
	// Experimental.
	InferenceConfiguration *awsbedrock.CfnPrompt_PromptModelInferenceConfigurationProperty `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
}

