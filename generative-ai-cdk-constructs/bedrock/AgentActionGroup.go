package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Experimental.
type AgentActionGroup interface {
	constructs.Construct
	// The Lambda function containing the business logic that is carried out upon invoking the action.
	// Experimental.
	ActionGroupExecutor() *ActionGroupExecutor
	// The unique identifier of the action group.
	// Experimental.
	ActionGroupName() *string
	// The action group.
	// Experimental.
	ActionGroupProperty() *awsbedrock.CfnAgent_AgentActionGroupProperty
	// The action group state.
	// Experimental.
	ActionGroupState() *string
	// The API schema.
	// Experimental.
	ApiSchema() *ApiSchemaConfig
	// The description.
	// Experimental.
	Description() *string
	// A list of action groups associated with the agent.
	// Experimental.
	FunctionSchema() *awsbedrock.CfnAgent_FunctionSchemaProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The parent action group signature.
	// Experimental.
	ParentActionGroupSignature() *string
	// The skip resource in use check on delete.
	// Default: - false.
	//
	// Experimental.
	SkipResourceInUseCheckOnDelete() *bool
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentActionGroup
type jsiiProxy_AgentActionGroup struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AgentActionGroup) ActionGroupExecutor() *ActionGroupExecutor {
	var returns *ActionGroupExecutor
	_jsii_.Get(
		j,
		"actionGroupExecutor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ActionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ActionGroupProperty() *awsbedrock.CfnAgent_AgentActionGroupProperty {
	var returns *awsbedrock.CfnAgent_AgentActionGroupProperty
	_jsii_.Get(
		j,
		"actionGroupProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ActionGroupState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ApiSchema() *ApiSchemaConfig {
	var returns *ApiSchemaConfig
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

func (j *jsiiProxy_AgentActionGroup) FunctionSchema() *awsbedrock.CfnAgent_FunctionSchemaProperty {
	var returns *awsbedrock.CfnAgent_FunctionSchemaProperty
	_jsii_.Get(
		j,
		"functionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) ParentActionGroupSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentActionGroup) SkipResourceInUseCheckOnDelete() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipResourceInUseCheckOnDelete",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentActionGroup(scope constructs.Construct, id *string, props *AgentActionGroupProps) AgentActionGroup {
	_init_.Initialize()

	if err := validateNewAgentActionGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentActionGroup{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentActionGroup_Override(a AgentActionGroup, scope constructs.Construct, id *string, props *AgentActionGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AgentActionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgentActionGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentActionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

