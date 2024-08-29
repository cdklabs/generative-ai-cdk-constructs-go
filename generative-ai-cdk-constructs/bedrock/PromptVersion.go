package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Creates a version of the prompt.
//
// Use this to create a static snapshot of your prompt that can be deployed
// to production. Versions allow you to easily switch between different
// configurations for your prompt and update your application with the most
// appropriate version for your use-case.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-deploy.html
//
// Experimental.
type PromptVersion interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The prompt used by this version.
	// Experimental.
	Prompt() Prompt
	// The version of the prompt that was created.
	// Experimental.
	Version() *string
	// The Amazon Resource Name (ARN) of the prompt version.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345:1"
	//
	// Experimental.
	VersionArn() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PromptVersion
type jsiiProxy_PromptVersion struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PromptVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) Prompt() Prompt {
	var returns Prompt
	_jsii_.Get(
		j,
		"prompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVersion) VersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewPromptVersion(scope constructs.Construct, id *string, props *PromptVersionProps) PromptVersion {
	_init_.Initialize()

	if err := validateNewPromptVersionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PromptVersion{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVersion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPromptVersion_Override(p PromptVersion, scope constructs.Construct, id *string, props *PromptVersionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVersion",
		[]interface{}{scope, id, props},
		p,
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
func PromptVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePromptVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PromptVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
