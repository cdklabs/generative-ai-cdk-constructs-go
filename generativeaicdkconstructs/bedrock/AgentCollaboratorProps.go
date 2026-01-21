package bedrock


// ****************************************************************************                   PROPS - Agent Collaborator Class ***************************************************************************.
// Experimental.
type AgentCollaboratorProps struct {
	// Descriptor for the collaborating agent.
	//
	// This cannot be the TSTALIASID (`agent.testAlias`).
	// Experimental.
	AgentAlias IAgentAlias `field:"required" json:"agentAlias" yaml:"agentAlias"`
	// Instructions on how this agent should collaborate with the main agent.
	// Experimental.
	CollaborationInstruction *string `field:"required" json:"collaborationInstruction" yaml:"collaborationInstruction"`
	// A friendly name for the collaborator.
	// Experimental.
	CollaboratorName *string `field:"required" json:"collaboratorName" yaml:"collaboratorName"`
	// Whether to relay conversation history to this collaborator.
	// Default: - undefined (uses service default).
	//
	// Experimental.
	RelayConversationHistory *bool `field:"optional" json:"relayConversationHistory" yaml:"relayConversationHistory"`
}

