package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Represents an Agent, either created with CDK or imported.
// Experimental.
type IAgent interface {
	awscdk.IResource
	// The ARN of the agent.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:agent/OKDSJOGKMO"@attributeundefined
	//
	// Experimental.
	AgentArn() *string
	// The ID of the Agent.
	//
	// Example:
	//   "OKDSJOGKMO"@attributeundefined
	//
	// Experimental.
	AgentId() *string
	// Optional KMS encryption key associated with this agent.
	// Experimental.
	KmsKey() awskms.IKey
	// When this agent was last updated.
	// Experimental.
	LastUpdated() *string
	// The IAM role associated to the agent.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IAgent
type jsiiProxy_IAgent struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAgent) AgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgent) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgent) LastUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgent) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

