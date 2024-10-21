package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// Defines a topic to deny.
// Experimental.
type Topic interface {
	// Definition of the topic.
	// Experimental.
	Definition() *string
	// Representative phrases that refer to the topic.
	// Experimental.
	Examples() *[]*string
	// The name of the topic to deny.
	// Experimental.
	Name() *string
}

// The jsii proxy struct for Topic
type jsiiProxy_Topic struct {
	_ byte // padding
}

func (j *jsiiProxy_Topic) Definition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Examples() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"examples",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Topic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopic(props *CustomTopicProps) Topic {
	_init_.Initialize()

	if err := validateNewTopicParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Topic{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewTopic_Override(t Topic, props *CustomTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		[]interface{}{props},
		t,
	)
}

// Experimental.
func Topic_Custom(props *CustomTopicProps) Topic {
	_init_.Initialize()

	if err := validateTopic_CustomParameters(props); err != nil {
		panic(err)
	}
	var returns Topic

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"custom",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Topic_FINANCIAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"FINANCIAL_ADVICE",
		&returns,
	)
	return returns
}

func Topic_INNOPROPRIATE_CONTENT() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"INNOPROPRIATE_CONTENT",
		&returns,
	)
	return returns
}

func Topic_LEGAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"LEGAL_ADVICE",
		&returns,
	)
	return returns
}

func Topic_MEDICAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"MEDICAL_ADVICE",
		&returns,
	)
	return returns
}

func Topic_POLITICAL_ADVICE() Topic {
	_init_.Initialize()
	var returns Topic
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		"POLITICAL_ADVICE",
		&returns,
	)
	return returns
}

