package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Experimental.
type AgentAlias interface {
	constructs.Construct
	// The ARN of the agent alias.
	// Experimental.
	AliasArn() *string
	// The unique identifier of the agent alias.
	// Experimental.
	AliasId() *string
	// The name for the agent alias.
	// Experimental.
	AliasName() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentAlias
type jsiiProxy_AgentAlias struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AgentAlias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentAlias(scope constructs.Construct, id *string, props *AgentAliasProps) AgentAlias {
	_init_.Initialize()

	if err := validateNewAgentAliasParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentAlias{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentAlias_Override(a AgentAlias, scope constructs.Construct, id *string, props *AgentAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
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
func AgentAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgentAlias_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

