package bedrock


// Specifies whether to allow the agent to carry out the step specified in the promptType.
//
// If you set this value to DISABLED, the agent skips that step.
// The default state for each promptType is as follows.
//
//     PRE_PROCESSING – ENABLED
//     ORCHESTRATION – ENABLED
//     KNOWLEDGE_BASE_RESPONSE_GENERATION – ENABLED
// POST_PROCESSING – DISABLED.
// Experimental.
type PromptState string

const (
	// Experimental.
	PromptState_ENABLED PromptState = "ENABLED"
	// Experimental.
	PromptState_DISABLED PromptState = "DISABLED"
)

