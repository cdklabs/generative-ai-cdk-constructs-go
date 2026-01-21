package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// ****************************************************************************                        DEF - Action Group Class ***************************************************************************.
// Experimental.
type AgentActionGroup interface {
	// The api schema for this action group (if defined).
	// Experimental.
	ApiSchema() ApiSchema
	// A description of the action group.
	// Experimental.
	Description() *string
	// Whether this action group is available for the agent to invoke or not.
	// Experimental.
	Enabled() *bool
	// The action group executor for this action group (if defined).
	// Experimental.
	Executor() ActionGroupExecutor
	// Whether to delete the resource even if it's in use.
	// Experimental.
	ForceDelete() *bool
	// The function schema for this action group (if defined).
	// Experimental.
	FunctionSchema() *awsbedrock.CfnAgent_FunctionSchemaProperty
	// The name of the action group.
	// Experimental.
	Name() *string
	// The AWS Defined signature (if defined).
	// Experimental.
	ParentActionGroupSignature() ParentActionGroupSignature
}

// The jsii proxy struct for AgentActionGroup
type jsiiProxy_AgentActionGroup struct {
	_ byte // padding
}

func (j *jsiiProxy_AgentActionGroup) ApiSchema() ApiSchema {
	var returns ApiSchema
	_jsii_.Get(
		j,
		"apiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) Executor() ActionGroupExecutor {
	var returns ActionGroupExecutor
	_jsii_.Get(
		j,
		"executor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ForceDelete() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) FunctionSchema() *awsbedrock.CfnAgent_FunctionSchemaProperty {
	var returns *awsbedrock.CfnAgent_FunctionSchemaProperty
	_jsii_.Get(
		j,
		"functionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ParentActionGroupSignature() ParentActionGroupSignature {
	var returns ParentActionGroupSignature
	_jsii_.Get(
		j,
		"parentActionGroupSignature",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentActionGroup(props *AgentActionGroupProps) AgentActionGroup {
	_init_.Initialize()

	if err := validateNewAgentActionGroupParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentActionGroup{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentActionGroup_Override(a AgentActionGroup, props *AgentActionGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		[]interface{}{props},
		a,
	)
}

// Defines an action group that allows your agent to request the user for additional information when trying to complete a task.
// Experimental.
func AgentActionGroup_CodeInterpreter(enabled *bool) AgentActionGroup {
	_init_.Initialize()

	if err := validateAgentActionGroup_CodeInterpreterParameters(enabled); err != nil {
		panic(err)
	}
	var returns AgentActionGroup

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		"codeInterpreter",
		[]interface{}{enabled},
		&returns,
	)

	return returns
}

// Defines an action group that allows your agent to request the user for additional information when trying to complete a task.
// Experimental.
func AgentActionGroup_UserInput(enabled *bool) AgentActionGroup {
	_init_.Initialize()

	if err := validateAgentActionGroup_UserInputParameters(enabled); err != nil {
		panic(err)
	}
	var returns AgentActionGroup

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		"userInput",
		[]interface{}{enabled},
		&returns,
	)

	return returns
}

