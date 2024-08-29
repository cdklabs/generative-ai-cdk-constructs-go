package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Variants are specific sets of inputs that guide FMs on Amazon Bedrock to generate an appropriate response or output for a given task or instruction.
//
// You can optimize the prompt for specific use cases and models.
// Experimental.
type PromptVariant interface {
	// The inference configuration.
	// Experimental.
	InferenceConfiguration() *awsbedrock.CfnPrompt_PromptInferenceConfigurationProperty
	// Experimental.
	SetInferenceConfiguration(val *awsbedrock.CfnPrompt_PromptInferenceConfigurationProperty)
	// The unique identifier of the model with which to run inference on the prompt.
	// Experimental.
	ModelId() *string
	// Experimental.
	SetModelId(val *string)
	// The name of the prompt variant.
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// The template configuration.
	// Experimental.
	TemplateConfiguration() *awsbedrock.CfnPrompt_PromptTemplateConfigurationProperty
	// Experimental.
	SetTemplateConfiguration(val *awsbedrock.CfnPrompt_PromptTemplateConfigurationProperty)
	// The type of prompt template.
	// Experimental.
	TemplateType() PromptTemplateType
	// Experimental.
	SetTemplateType(val PromptTemplateType)
}

// The jsii proxy struct for PromptVariant
type jsiiProxy_PromptVariant struct {
	_ byte // padding
}

func (j *jsiiProxy_PromptVariant) InferenceConfiguration() *awsbedrock.CfnPrompt_PromptInferenceConfigurationProperty {
	var returns *awsbedrock.CfnPrompt_PromptInferenceConfigurationProperty
	_jsii_.Get(
		j,
		"inferenceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVariant) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVariant) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVariant) TemplateConfiguration() *awsbedrock.CfnPrompt_PromptTemplateConfigurationProperty {
	var returns *awsbedrock.CfnPrompt_PromptTemplateConfigurationProperty
	_jsii_.Get(
		j,
		"templateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptVariant) TemplateType() PromptTemplateType {
	var returns PromptTemplateType
	_jsii_.Get(
		j,
		"templateType",
		&returns,
	)
	return returns
}


// Experimental.
func NewPromptVariant_Override(p PromptVariant) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVariant",
		nil, // no parameters
		p,
	)
}

func (j *jsiiProxy_PromptVariant)SetInferenceConfiguration(val *awsbedrock.CfnPrompt_PromptInferenceConfigurationProperty) {
	if err := j.validateSetInferenceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceConfiguration",
		val,
	)
}

func (j *jsiiProxy_PromptVariant)SetModelId(val *string) {
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_PromptVariant)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PromptVariant)SetTemplateConfiguration(val *awsbedrock.CfnPrompt_PromptTemplateConfigurationProperty) {
	if err := j.validateSetTemplateConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateConfiguration",
		val,
	)
}

func (j *jsiiProxy_PromptVariant)SetTemplateType(val PromptTemplateType) {
	if err := j.validateSetTemplateTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateType",
		val,
	)
}

// Static method to create a text template.
// Experimental.
func PromptVariant_Text(props *TextPromptVariantProps) PromptVariant {
	_init_.Initialize()

	if err := validatePromptVariant_TextParameters(props); err != nil {
		panic(err)
	}
	var returns PromptVariant

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVariant",
		"text",
		[]interface{}{props},
		&returns,
	)

	return returns
}

