package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Prompts are a specific set of inputs that guide FMs on Amazon Bedrock to generate an appropriate response or output for a given task or instruction.
//
// You can optimize the prompt for specific use cases and models.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management.html
//
// Experimental.
type Prompt interface {
	constructs.Construct
	IPrompt
	// The KMS key that the prompt is encrypted with.
	// Experimental.
	KmsKey() awskms.IKey
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The ARN of the prompt.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345"
	//
	// Experimental.
	PromptArn() *string
	// The ID of the prompt.
	//
	// Example:
	//   "PROMPT12345"
	//
	// Experimental.
	PromptId() *string
	// The name of the prompt.
	// Experimental.
	PromptName() *string
	// The version of the prompt.
	// Experimental.
	PromptVersion() *string
	// Experimental.
	SetPromptVersion(val *string)
	// The variants of the prompt.
	// Experimental.
	Variants() *[]PromptVariant
	// Adds a prompt variant.
	// Experimental.
	AddVariant(variant PromptVariant)
	// Creates a prompt version, a static snapshot of your prompt that can be deployed to production.
	// Experimental.
	CreateVersion(description *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Prompt
type jsiiProxy_Prompt struct {
	internal.Type__constructsConstruct
	jsiiProxy_IPrompt
}

func (j *jsiiProxy_Prompt) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prompt) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prompt) PromptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prompt) PromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prompt) PromptName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prompt) PromptVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prompt) Variants() *[]PromptVariant {
	var returns *[]PromptVariant
	_jsii_.Get(
		j,
		"variants",
		&returns,
	)
	return returns
}


// Experimental.
func NewPrompt(scope constructs.Construct, id *string, props *PromptProps) Prompt {
	_init_.Initialize()

	if err := validateNewPromptParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Prompt{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Prompt",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPrompt_Override(p Prompt, scope constructs.Construct, id *string, props *PromptProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Prompt",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_Prompt)SetPromptVersion(val *string) {
	if err := j.validateSetPromptVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptVersion",
		val,
	)
}

// Experimental.
func Prompt_FromPromptAttributes(scope constructs.Construct, id *string, attrs *PromptAttributes) IPrompt {
	_init_.Initialize()

	if err := validatePrompt_FromPromptAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IPrompt

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Prompt",
		"fromPromptAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
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
func Prompt_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrompt_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Prompt",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prompt) AddVariant(variant PromptVariant) {
	if err := p.validateAddVariantParameters(variant); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addVariant",
		[]interface{}{variant},
	)
}

func (p *jsiiProxy_Prompt) CreateVersion(description *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"createVersion",
		[]interface{}{description},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prompt) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

