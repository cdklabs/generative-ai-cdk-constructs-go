package bedrock


// Experimental.
type AgentAliasProps struct {
	// The unique identifier of the agent.
	// Experimental.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// The version of the agent to associate with the agent alias.
	// Default: - Creates a new version of the agent.
	//
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// The name for the agent alias.
	// Default: - 'latest'.
	//
	// Experimental.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// Description for the agent alias.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of resource update timestamps to let CloudFormation determine when to update the alias.
	// Experimental.
	ResourceUpdates *[]*string `field:"optional" json:"resourceUpdates" yaml:"resourceUpdates"`
	// OPTIONAL: Tag (KEY-VALUE) bedrock agent resource.
	// Default: - false.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

