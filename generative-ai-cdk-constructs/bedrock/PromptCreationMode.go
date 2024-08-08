package bedrock


// Specifies whether to override the default prompt template for this promptType.
//
// Set this value to OVERRIDDEN to use the prompt that you
// provide in the basePromptTemplate. If you leave it as DEFAULT, the agent
// uses a default prompt template.
// Experimental.
type PromptCreationMode string

const (
	// Experimental.
	PromptCreationMode_DEFAULT PromptCreationMode = "DEFAULT"
	// Experimental.
	PromptCreationMode_OVERRIDDEN PromptCreationMode = "OVERRIDDEN"
)

