package bedrock


// The step in the agent sequence that this prompt configuration applies to.
// Experimental.
type PromptType string

const (
	// Experimental.
	PromptType_PRE_PROCESSING PromptType = "PRE_PROCESSING"
	// Experimental.
	PromptType_ORCHESTRATION PromptType = "ORCHESTRATION"
	// Experimental.
	PromptType_POST_PROCESSING PromptType = "POST_PROCESSING"
	// Experimental.
	PromptType_KNOWLEDGE_BASE_RESPONSE_GENERATION PromptType = "KNOWLEDGE_BASE_RESPONSE_GENERATION"
)

