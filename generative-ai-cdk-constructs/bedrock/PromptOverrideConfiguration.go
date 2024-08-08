package bedrock


// Contains configurations to override prompts in different parts of an agent sequence.
// Experimental.
type PromptOverrideConfiguration struct {
	// Contains configurations to override a prompt template in one part of an agent sequence.
	// Experimental.
	PromptConfigurations *[]*PromptConfiguration `field:"required" json:"promptConfigurations" yaml:"promptConfigurations"`
	// The ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence.
	//
	// If you specify this field,
	// at least one of the promptConfigurations must contain a parserMode value
	// that is set to OVERRIDDEN.
	// Experimental.
	OverrideLambda *string `field:"optional" json:"overrideLambda" yaml:"overrideLambda"`
}

