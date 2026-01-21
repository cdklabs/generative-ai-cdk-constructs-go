package bedrock


// Experimental.
type PromptVersionProps struct {
	// The prompt to use for this version.
	// Experimental.
	Prompt Prompt `field:"required" json:"prompt" yaml:"prompt"`
	// The description of the prompt version.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

