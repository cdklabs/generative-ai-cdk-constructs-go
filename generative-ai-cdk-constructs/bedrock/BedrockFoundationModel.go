package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Bedrock models.
//
// If you need to use a model name that doesn't exist as a static member, you
// can instantiate a `BedrockFoundationModel` object, e.g: `new BedrockFoundationModel('my-model')`.
// Experimental.
type BedrockFoundationModel interface {
	IInvokable
	// The ARN of the Bedrock invokable abstraction.
	// Experimental.
	InvokableArn() *string
	// Experimental.
	ModelArn() *string
	// **************************************************************************                           Constructor *************************************************************************.
	// Experimental.
	ModelId() *string
	// Experimental.
	SupportsAgents() *bool
	// Experimental.
	SupportsCrossRegion() *bool
	// Experimental.
	SupportsKnowledgeBase() *bool
	// Experimental.
	VectorDimensions() *float64
	// Returns the ARN of the foundation model in the following format: `arn:${Partition}:bedrock:${Region}::foundation-model/${ResourceId}`.
	// Experimental.
	AsArn(construct constructs.IConstruct) *string
	// Experimental.
	AsIModel(construct constructs.IConstruct) awsbedrock.IModel
	// Gives the appropriate policies to invoke and use the Foundation Model in the stack region.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Gives the appropriate policies to invoke and use the Foundation Model in all regions.
	// Experimental.
	GrantInvokeAllRegions(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of an object.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockFoundationModel
type jsiiProxy_BedrockFoundationModel struct {
	jsiiProxy_IInvokable
}

func (j *jsiiProxy_BedrockFoundationModel) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockFoundationModel) ModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArn",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_BedrockFoundationModel) SupportsCrossRegion() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"supportsCrossRegion",
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

// Experimental.
func BedrockFoundationModel_FromCdkFoundationModel(modelId awsbedrock.FoundationModel, props *BedrockFoundationModelProps) BedrockFoundationModel {
	_init_.Initialize()

	if err := validateBedrockFoundationModel_FromCdkFoundationModelParameters(modelId, props); err != nil {
		panic(err)
	}
	var returns BedrockFoundationModel

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"fromCdkFoundationModel",
		[]interface{}{modelId, props},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockFoundationModel_FromCdkFoundationModelId(modelId awsbedrock.FoundationModelIdentifier, props *BedrockFoundationModelProps) BedrockFoundationModel {
	_init_.Initialize()

	if err := validateBedrockFoundationModel_FromCdkFoundationModelIdParameters(modelId, props); err != nil {
		panic(err)
	}
	var returns BedrockFoundationModel

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"fromCdkFoundationModelId",
		[]interface{}{modelId, props},
		&returns,
	)

	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_LITE_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"AMAZON_NOVA_LITE_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_MICRO_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"AMAZON_NOVA_MICRO_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_AMAZON_NOVA_PRO_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"AMAZON_NOVA_PRO_V1",
		&returns,
	)
	return returns
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

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_HAIKU_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_5_HAIKU_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_5_SONNET_V1_0",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V2_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_3_5_SONNET_V2_0",
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

func BedrockFoundationModel_ANTHROPIC_CLAUDE_OPUS_V1_0() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"ANTHROPIC_CLAUDE_OPUS_V1_0",
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

func BedrockFoundationModel_META_LLAMA_3_2_11B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"META_LLAMA_3_2_11B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_2_1B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"META_LLAMA_3_2_1B_INSTRUCT_V1",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_META_LLAMA_3_2_3B_INSTRUCT_V1() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"META_LLAMA_3_2_3B_INSTRUCT_V1",
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

func BedrockFoundationModel_TITAN_EMBED_TEXT_V2_256() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V2_256",
		&returns,
	)
	return returns
}

func BedrockFoundationModel_TITAN_EMBED_TEXT_V2_512() BedrockFoundationModel {
	_init_.Initialize()
	var returns BedrockFoundationModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		"TITAN_EMBED_TEXT_V2_512",
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

func (b *jsiiProxy_BedrockFoundationModel) AsIModel(construct constructs.IConstruct) awsbedrock.IModel {
	if err := b.validateAsIModelParameters(construct); err != nil {
		panic(err)
	}
	var returns awsbedrock.IModel

	_jsii_.Invoke(
		b,
		"asIModel",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := b.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockFoundationModel) GrantInvokeAllRegions(grantee awsiam.IGrantable) awsiam.Grant {
	if err := b.validateGrantInvokeAllRegionsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		b,
		"grantInvokeAllRegions",
		[]interface{}{grantee},
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

