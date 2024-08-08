package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Experimental.
type Topic interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	CreateTopic(props *TopicProps)
	// Experimental.
	FinancialAdviceTopic(props *TopicProps)
	// Experimental.
	InappropriateContent(props *TopicProps)
	// Experimental.
	LegalAdvice(props *TopicProps)
	// Experimental.
	MedicalAdvice(props *TopicProps)
	// Experimental.
	PoliticalAdviceTopic(props *TopicProps)
	// Experimental.
	TopicConfigPropertyList() *[]*awsbedrock.CfnGuardrail_TopicConfigProperty
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Topic
type jsiiProxy_Topic struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Topic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopic(scope constructs.Construct, id *string) Topic {
	_init_.Initialize()

	if err := validateNewTopicParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_Topic{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewTopic_Override(t Topic, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		[]interface{}{scope, id},
		t,
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
func Topic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) CreateTopic(props *TopicProps) {
	if err := t.validateCreateTopicParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"createTopic",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_Topic) FinancialAdviceTopic(props *TopicProps) {
	if err := t.validateFinancialAdviceTopicParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"financialAdviceTopic",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_Topic) InappropriateContent(props *TopicProps) {
	if err := t.validateInappropriateContentParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"inappropriateContent",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_Topic) LegalAdvice(props *TopicProps) {
	if err := t.validateLegalAdviceParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"legalAdvice",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_Topic) MedicalAdvice(props *TopicProps) {
	if err := t.validateMedicalAdviceParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"medicalAdvice",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_Topic) PoliticalAdviceTopic(props *TopicProps) {
	if err := t.validatePoliticalAdviceTopicParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"politicalAdviceTopic",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_Topic) TopicConfigPropertyList() *[]*awsbedrock.CfnGuardrail_TopicConfigProperty {
	var returns *[]*awsbedrock.CfnGuardrail_TopicConfigProperty

	_jsii_.Invoke(
		t,
		"topicConfigPropertyList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Topic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

