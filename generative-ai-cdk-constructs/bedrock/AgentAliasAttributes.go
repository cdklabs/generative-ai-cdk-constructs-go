package bedrock


// Attributes needed to create an import.
// Experimental.
type AgentAliasAttributes struct {
	// The underlying agent for this alias.
	// Experimental.
	Agent IAgent `field:"required" json:"agent" yaml:"agent"`
	// The agent version for this alias.
	// Experimental.
	AgentVersion *string `field:"required" json:"agentVersion" yaml:"agentVersion"`
	// The Id of the agent alias.
	// Experimental.
	AliasId *string `field:"required" json:"aliasId" yaml:"aliasId"`
	// The name of the agent alias.
	// Experimental.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

