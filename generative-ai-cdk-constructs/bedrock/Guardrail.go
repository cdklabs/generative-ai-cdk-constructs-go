package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Deploy bedrock guardrail .
// Experimental.
type Guardrail interface {
	constructs.Construct
	// guardrail Id.
	// Experimental.
	GuardrailId() *string
	// Instance of guardrail.
	// Experimental.
	GuardrailInstance() awsbedrock.CfnGuardrail
	// guardrail version.
	// Experimental.
	GuardrailVersion() *string
	// The ARN of the AWS KMS key used to encrypt the guardrail.
	// Experimental.
	KmsKeyArn() *string
	// The name of the guardrail.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	AddSensitiveInformationPolicyConfig(props *[]*SensitiveInformationPolicyConfigProps, guardrailRegexesConfig *awsbedrock.CfnGuardrail_RegexConfigProperty)
	// Experimental.
	AddTags(props *GuardrailProps)
	// Experimental.
	AddTopicPolicyConfig(topic Topic)
	// Creates a version of the guardrail.
	// Experimental.
	AddVersion(id *string, description *string) GuardrailVersion
	// Experimental.
	AddWordPolicyConfig(wordsFilter *[]*awsbedrock.CfnGuardrail_WordConfigProperty)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	UploadWordPolicyFromFile(filePath *string)
}

// The jsii proxy struct for Guardrail
type jsiiProxy_Guardrail struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Guardrail) GuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Guardrail) GuardrailInstance() awsbedrock.CfnGuardrail {
	var returns awsbedrock.CfnGuardrail
	_jsii_.Get(
		j,
		"guardrailInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Guardrail) GuardrailVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Guardrail) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Guardrail) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Guardrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewGuardrail(scope constructs.Construct, id *string, props *GuardrailProps) Guardrail {
	_init_.Initialize()

	if err := validateNewGuardrailParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Guardrail{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Guardrail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGuardrail_Override(g Guardrail, scope constructs.Construct, id *string, props *GuardrailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Guardrail",
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
func Guardrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGuardrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Guardrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Guardrail) AddSensitiveInformationPolicyConfig(props *[]*SensitiveInformationPolicyConfigProps, guardrailRegexesConfig *awsbedrock.CfnGuardrail_RegexConfigProperty) {
	if err := g.validateAddSensitiveInformationPolicyConfigParameters(props, guardrailRegexesConfig); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addSensitiveInformationPolicyConfig",
		[]interface{}{props, guardrailRegexesConfig},
	)
}

func (g *jsiiProxy_Guardrail) AddTags(props *GuardrailProps) {
	if err := g.validateAddTagsParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addTags",
		[]interface{}{props},
	)
}

func (g *jsiiProxy_Guardrail) AddTopicPolicyConfig(topic Topic) {
	if err := g.validateAddTopicPolicyConfigParameters(topic); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addTopicPolicyConfig",
		[]interface{}{topic},
	)
}

func (g *jsiiProxy_Guardrail) AddVersion(id *string, description *string) GuardrailVersion {
	if err := g.validateAddVersionParameters(id); err != nil {
		panic(err)
	}
	var returns GuardrailVersion

	_jsii_.Invoke(
		g,
		"addVersion",
		[]interface{}{id, description},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Guardrail) AddWordPolicyConfig(wordsFilter *[]*awsbedrock.CfnGuardrail_WordConfigProperty) {
	if err := g.validateAddWordPolicyConfigParameters(wordsFilter); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addWordPolicyConfig",
		[]interface{}{wordsFilter},
	)
}

func (g *jsiiProxy_Guardrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Guardrail) UploadWordPolicyFromFile(filePath *string) {
	if err := g.validateUploadWordPolicyFromFileParameters(filePath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"uploadWordPolicyFromFile",
		[]interface{}{filePath},
	)
}

