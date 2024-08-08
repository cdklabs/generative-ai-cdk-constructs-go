package bedrock


// Properties to add an Alias to an Agent.
// Experimental.
type AddAgentAliasProps struct {
	// The name for the agent alias.
	// Experimental.
	AliasName *string `field:"required" json:"aliasName" yaml:"aliasName"`
	// The version of the agent to associate with the agent alias.
	// Default: - Creates a new version of the agent.
	//
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Description for the agent alias.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

