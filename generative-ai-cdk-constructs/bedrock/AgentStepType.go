package bedrock


// The step in the agent sequence that this prompt configuration applies to.
// Experimental.
type AgentStepType string

const (
	// Experimental.
	AgentStepType_PRE_PROCESSING AgentStepType = "PRE_PROCESSING"
	// Experimental.
	AgentStepType_ORCHESTRATION AgentStepType = "ORCHESTRATION"
	// Experimental.
	AgentStepType_POST_PROCESSING AgentStepType = "POST_PROCESSING"
	// Experimental.
	AgentStepType_ROUTING_CLASSIFIER AgentStepType = "ROUTING_CLASSIFIER"
	// Experimental.
	AgentStepType_MEMORY_SUMMARIZATION AgentStepType = "MEMORY_SUMMARIZATION"
	// Experimental.
	AgentStepType_KNOWLEDGE_BASE_RESPONSE_GENERATION AgentStepType = "KNOWLEDGE_BASE_RESPONSE_GENERATION"
)

