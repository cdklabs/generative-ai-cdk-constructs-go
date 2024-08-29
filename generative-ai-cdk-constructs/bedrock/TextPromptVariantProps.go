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
	Model awsbedrock.IModel `field:"required" json:"model" yaml:"model"`
	// The name of the prompt variant.
	// Experimental.
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
	// Inference configuration for the Text Prompt.
	// Experimental.
	InferenceConfiguration *awsbedrock.CfnPrompt_PromptModelInferenceConfigurationProperty `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Template Configuration for the text prompt.
	// Experimental.
	TemplateConfiguration *awsbedrock.CfnPrompt_TextPromptTemplateConfigurationProperty `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
}

