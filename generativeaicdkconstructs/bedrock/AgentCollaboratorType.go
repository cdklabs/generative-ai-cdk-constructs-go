package bedrock


// Enum for collaborator's relay conversation history types.
// Experimental.
type AgentCollaboratorType string

const (
	// Supervisor agent.
	// Experimental.
	AgentCollaboratorType_SUPERVISOR AgentCollaboratorType = "SUPERVISOR"
	// Disabling collaboration.
	// Experimental.
	AgentCollaboratorType_DISABLED AgentCollaboratorType = "DISABLED"
	// Supervisor router.
	// Experimental.
	AgentCollaboratorType_SUPERVISOR_ROUTER AgentCollaboratorType = "SUPERVISOR_ROUTER"
)

