package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/internal"
)

// The LangchainCommonDepsLayer class.
// Experimental.
type LangchainCommonDepsLayer interface {
	constructs.Construct
	// Returns the instance of lambda.LayerVersion created by the construct.
	// Experimental.
	Layer() awslambda.LayerVersion
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LangchainCommonDepsLayer
type jsiiProxy_LangchainCommonDepsLayer struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_LangchainCommonDepsLayer) Layer() awslambda.LayerVersion {
	var returns awslambda.LayerVersion
	_jsii_.Get(
		j,
		"layer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LangchainCommonDepsLayer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// This construct creates a lambda layer loaded with relevant libraries to run genai applications.
//
// Libraries include boto3, botocore, requests, requests-aws4auth, langchain, opensearch-py and openai.
// Experimental.
func NewLangchainCommonDepsLayer(scope constructs.Construct, id *string, props *LangchainLayerProps) LangchainCommonDepsLayer {
	_init_.Initialize()

	if err := validateNewLangchainCommonDepsLayerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LangchainCommonDepsLayer{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.LangchainCommonDepsLayer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// This construct creates a lambda layer loaded with relevant libraries to run genai applications.
//
// Libraries include boto3, botocore, requests, requests-aws4auth, langchain, opensearch-py and openai.
// Experimental.
func NewLangchainCommonDepsLayer_Override(l LangchainCommonDepsLayer, scope constructs.Construct, id *string, props *LangchainLayerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.LangchainCommonDepsLayer",
		[]interface{}{scope, id, props},
		l,
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
func LangchainCommonDepsLayer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLangchainCommonDepsLayer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.LangchainCommonDepsLayer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LangchainCommonDepsLayer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

