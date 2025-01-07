package bedrock


// Experimental.
type AgentPromptVariantProps struct {
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
	// An alias pointing to the agent version to be used.
	// Experimental.
	AgentAlias IAgentAlias `field:"required" json:"agentAlias" yaml:"agentAlias"`
	// The text prompt.
	//
	// Variables are used by enclosing its name with double curly braces
	// as in `{{variable_name}}`.
	// Experimental.
	PromptText *string `field:"required" json:"promptText" yaml:"promptText"`
}

