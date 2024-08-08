package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Creates a version of the guardrail.
//
// Use this API to create a snapshot of the guardrail when you are satisfied with
// a configuration, or to compare the configuration with another version.
// Experimental.
type GuardrailVersion interface {
	constructs.Construct
	// Instance of guardrail version.
	// Experimental.
	GuardrailVersionInstance() awsbedrock.CfnGuardrailVersion
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuardrailVersion
type jsiiProxy_GuardrailVersion struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_GuardrailVersion) GuardrailVersionInstance() awsbedrock.CfnGuardrailVersion {
	var returns awsbedrock.CfnGuardrailVersion
	_jsii_.Get(
		j,
		"guardrailVersionInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewGuardrailVersion(scope constructs.Construct, id *string, props *awsbedrock.CfnGuardrailVersionProps) GuardrailVersion {
	_init_.Initialize()

	if err := validateNewGuardrailVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GuardrailVersion{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGuardrailVersion_Override(g GuardrailVersion, scope constructs.Construct, id *string, props *awsbedrock.CfnGuardrailVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailVersion",
		[]interface{}{scope, id, props},
		g,
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
func GuardrailVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGuardrailVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

