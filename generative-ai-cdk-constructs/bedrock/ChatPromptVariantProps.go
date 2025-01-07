package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Experimental.
type ChatPromptVariantProps struct {
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
	// Inference configuration for the Chat Prompt.
	//
	// Must include at least one User Message.
	// The messages should alternate between User and Assistant.
	// Experimental.
	Messages *[]ChatMessage `field:"required" json:"messages" yaml:"messages"`
	// Inference configuration for the Text Prompt.
	// Experimental.
	InferenceConfiguration *awsbedrock.CfnPrompt_PromptModelInferenceConfigurationProperty `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Context or instructions for the model to consider before generating a response.
	// Experimental.
	System *string `field:"optional" json:"system" yaml:"system"`
	// The configuration with available tools to the model and how it must use them.
	// Experimental.
	ToolConfiguration *ToolConfiguration `field:"optional" json:"toolConfiguration" yaml:"toolConfiguration"`
}

