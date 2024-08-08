package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Bedrock models.
//
// If you need to use a model name that doesn't exist as a static member, you
// can instantiate a `BedrockFoundationModel` object, e.g: `new BedrockFoundationModel('my-model')`.
// Experimental.
type BedrockFoundationModel interface {
	// Experimental.
	ModelId() *string
	// Experimental.
	SupportsAgents() *bool
	// Experimental.
	SupportsKnowledgeBase() *bool
	// Experimental.
	VectorDimensions() *float64
	// Experimental.
	AsArn(construct constructs.IConstruct) *string
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockFoundationModel
type jsiiProxy_BedrockFoundationModel struct {
	_ byte // padding
}

func (j *jsiiProxy_BedrockFoundationModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) SupportsAgents() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) SupportsKnowledgeBase() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) VectorDimensions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vectorDimensions",
		&returns,
	)
	return returns
}


// Experimental.
func NewBedrockFoundationModel(value *string, props *BedrockFoundationModelProps) BedrockFoundationModel {
	_init_.Initialize()

	if err := validateNewBedrockFoundationModelParameters(value, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockFoundationModel{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		[]interface{}{value, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBedrockFoundationModel_Override(b BedrockFoundationModel, value *string, props *BedrockFoundationModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		[]interface{}{value, props},
		b,
	)
}

func BedrockFoundationModel_AMAZON_TITAN_PREMIER_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"AMAZON_TITAN_PREMIER_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_TITAN_TEXT_EXPRESS_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"AMAZON_TITAN_TEXT_EXPRESS_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_HAIKU_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_INSTANT_V1_2() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_INSTANT_V1_2",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_SONNET_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_V2() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_V2",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_V2_1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_V2_1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_COHERE_EMBED_ENGLISH_V3() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"COHERE_EMBED_ENGLISH_V3",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_COHERE_EMBED_MULTILINGUAL_V3() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"COHERE_EMBED_MULTILINGUAL_V3",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V2_1024() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V2_1024",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) AsArn(construct constructs.IConstruct) *string {
	if err := b.validateAsArnParameters(construct); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"asArn",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

