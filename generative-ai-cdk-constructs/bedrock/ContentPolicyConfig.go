package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Experimental.
type ContentPolicyConfig interface {
	constructs.Construct
	// Experimental.
	ContentPolicyConfigList() *[]*awsbedrock.CfnGuardrail_ContentFilterConfigProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContentPolicyConfig
type jsiiProxy_ContentPolicyConfig struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ContentPolicyConfig) ContentPolicyConfigList() *[]*awsbedrock.CfnGuardrail_ContentFilterConfigProperty {
	var returns *[]*awsbedrock.CfnGuardrail_ContentFilterConfigProperty
	_jsii_.Get(
		j,
		"contentPolicyConfigList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentPolicyConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewContentPolicyConfig(scope constructs.Construct, id *string, props *[]*ContentPolicyConfigProps) ContentPolicyConfig {
	_init_.Initialize()

	if err := validateNewContentPolicyConfigParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContentPolicyConfig{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentPolicyConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewContentPolicyConfig_Override(c ContentPolicyConfig, scope constructs.Construct, id *string, props *[]*ContentPolicyConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentPolicyConfig",
		[]interface{}{scope, id, props},
		c,
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
func ContentPolicyConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContentPolicyConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentPolicyConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContentPolicyConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

