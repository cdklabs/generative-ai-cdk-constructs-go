package bedrock


// Properties for creating a CDK-Managed Agent Alias.
// Experimental.
type AgentAliasProps struct {
	// The agent associated to this alias.
	// Experimental.
	Agent IAgent `field:"required" json:"agent" yaml:"agent"`
	// The version of the agent to associate with the agent alias.
	// Default: - Creates a new version of the agent.
	//
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// The name for the agent alias.
	// Default: - "latest".
	//
	// Experimental.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// Description for the agent alias.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

