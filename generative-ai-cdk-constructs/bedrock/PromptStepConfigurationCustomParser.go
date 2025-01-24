package bedrock


// Experimental.
type PromptStepConfigurationCustomParser struct {
	// The step in the agent sequence where to set a specific prompt configuration.
	// Experimental.
	StepType AgentStepType `field:"required" json:"stepType" yaml:"stepType"`
	// The custom prompt template to be used.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-placeholders.html
	//
	// Default: - The default prompt template will be used.
	//
	// Experimental.
	CustomPromptTemplate *string `field:"optional" json:"customPromptTemplate" yaml:"customPromptTemplate"`
	// The inference configuration parameters to use.
	// Experimental.
	InferenceConfig *InferenceConfiguration `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
	// Whether to enable or skip this step in the agent sequence.
	// Default: - The default state for each step type is as follows.
	//
	// PRE_PROCESSING – ENABLED
	// ORCHESTRATION – ENABLED
	// KNOWLEDGE_BASE_RESPONSE_GENERATION – ENABLED
	// POST_PROCESSING – DISABLED.
	//
	// Experimental.
	StepEnabled *bool `field:"optional" json:"stepEnabled" yaml:"stepEnabled"`
	// Whether to use the custom Lambda parser defined for the sequence.
	// Default: - false.
	//
	// Experimental.
	UseCustomParser *bool `field:"optional" json:"useCustomParser" yaml:"useCustomParser"`
}

