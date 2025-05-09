package bedrock


// Contains configurations to override a prompt template in one part of an agent sequence.
// Experimental.
type PromptStepConfiguration struct {
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
	// The foundation model to use for this specific prompt step.
	//
	// This allows using different models for different steps in the agent sequence.
	// Default: - The agent's default foundation model will be used.
	//
	// Experimental.
	FoundationModel IInvokable `field:"optional" json:"foundationModel" yaml:"foundationModel"`
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
}

