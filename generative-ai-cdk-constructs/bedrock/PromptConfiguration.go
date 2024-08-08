package bedrock


// Contains configurations to override a prompt template in one part of an agent sequence.
// Experimental.
type PromptConfiguration struct {
	// Defines the prompt template with which to replace the default prompt template.
	//
	// length 0-100000.
	// Experimental.
	BasePromptTemplate *string `field:"required" json:"basePromptTemplate" yaml:"basePromptTemplate"`
	// Contains inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the promptType.
	// Experimental.
	InferenceConfiguration *InferenceConfiguration `field:"required" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Specifies whether to override the default prompt template for this promptType.
	//
	// Set this value to OVERRIDDEN to use the prompt that you
	// provide in the basePromptTemplate. If you leave it as DEFAULT, the agent
	// uses a default prompt template.
	// Experimental.
	PromptCreationMode PromptCreationMode `field:"required" json:"promptCreationMode" yaml:"promptCreationMode"`
	// Specifies whether to allow the agent to carry out the step specified in the promptType.
	//
	// If you set this value to DISABLED, the agent skips that
	// step. The default state for each promptType is as follows.
	//
	//     PRE_PROCESSING – ENABLED
	//     ORCHESTRATION – ENABLED
	//     KNOWLEDGE_BASE_RESPONSE_GENERATION – ENABLED
	// POST_PROCESSING – DISABLED.
	// Experimental.
	PromptState PromptState `field:"required" json:"promptState" yaml:"promptState"`
	// The step in the agent sequence that this prompt configuration applies to.
	// Experimental.
	PromptType PromptType `field:"required" json:"promptType" yaml:"promptType"`
	// Specifies whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the promptType.
	//
	// If you set the field as OVERRIDEN, the
	// overrideLambda field in the PromptOverrideConfiguration must be specified
	// with the ARN of a Lambda function.
	// Experimental.
	ParserMode ParserMode `field:"optional" json:"parserMode" yaml:"parserMode"`
}

