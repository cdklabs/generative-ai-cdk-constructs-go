package bedrock


// Attributes for specifying an imported Bedrock Agent.
// Experimental.
type AgentAttributes struct {
	// The ARN of the agent.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:agent/OKDSJOGKMO"@attributeundefined
	//
	// Experimental.
	AgentArn *string `field:"required" json:"agentArn" yaml:"agentArn"`
	// The ARN of the IAM role associated to the agent.
	//
	// Example:
	//   "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"@attributeundefined
	//
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The agent version.
	//
	// If no explicit versions have been created,
	// leave this  empty to use the DRAFT version. Otherwise, use the
	// version number (e.g. 1).
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Optional KMS encryption key associated with this agent.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// When this agent was last updated.
	// Experimental.
	LastUpdated *string `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
}

