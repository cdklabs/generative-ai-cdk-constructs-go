package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for both Imported and CDK-created Agent Aliases.
// Experimental.
type IAgentAlias interface {
	// The unique identifier of the agent.
	//
	// Example:
	//   `DNCJJYQKSU`
	//
	// Experimental.
	AgentId() *string
	// The ARN of the agent alias.
	//
	// Example:
	//   `arn:aws:bedrock:us-east-1:123456789012:agent-alias/DNCJJYQKSU/TCLCITFZTN`
	//
	// Experimental.
	AliasArn() *string
	// The unique identifier of the agent alias.
	//
	// Example:
	//   `TCLCITFZTN`
	//
	// Experimental.
	AliasId() *string
}

// The jsii proxy for IAgentAlias
type jsiiProxy_IAgentAlias struct {
	_ byte // padding
}

func (j *jsiiProxy_IAgentAlias) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentAlias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentAlias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

