package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// ****************************************************************************                        DEF - Agent Collaborator Class ***************************************************************************.
// Experimental.
type AgentCollaborator interface {
	// Experimental.
	AgentAlias() IAgentAlias
	// Instructions on how this agent should collaborate with the main agent.
	// Experimental.
	CollaborationInstruction() *string
	// A friendly name for the collaborator.
	// Experimental.
	CollaboratorName() *string
	// Whether to relay conversation history to this collaborator.
	// Default: - undefined (uses service default).
	//
	// Experimental.
	RelayConversationHistory() *bool
	// Grants the specified principal permissions to get the agent alias and invoke the agent from this collaborator.
	//
	// Returns: The Grant object.
	// Experimental.
	Grant(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for AgentCollaborator
type jsiiProxy_AgentCollaborator struct {
	_ byte // padding
}

func (j *jsiiProxy_AgentCollaborator) AgentAlias() IAgentAlias {
	var returns IAgentAlias
	_jsii_.Get(
		j,
		"agentAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaborator) CollaborationInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaborator) CollaboratorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentCollaborator) RelayConversationHistory() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"relayConversationHistory",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentCollaborator(props *AgentCollaboratorProps) AgentCollaborator {
	_init_.Initialize()

	if err := validateNewAgentCollaboratorParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentCollaborator{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentCollaborator",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentCollaborator_Override(a AgentCollaborator, props *AgentCollaboratorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentCollaborator",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AgentCollaborator) Grant(grantee awsiam.IGrantable) awsiam.Grant {
	if err := a.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grant",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

