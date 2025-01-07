package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// Experimental.
type ToolChoice interface {
	// Experimental.
	Any() interface{}
	// Experimental.
	Auto() interface{}
	// Experimental.
	Tool() *string
}

// The jsii proxy struct for ToolChoice
type jsiiProxy_ToolChoice struct {
	_ byte // padding
}

func (j *jsiiProxy_ToolChoice) Any() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"any",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolChoice) Auto() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolChoice) Tool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}


// Experimental.
func NewToolChoice(any interface{}, auto interface{}, tool *string) ToolChoice {
	_init_.Initialize()

	if err := validateNewToolChoiceParameters(any, auto); err != nil {
		panic(err)
	}
	j := jsiiProxy_ToolChoice{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolChoice",
		[]interface{}{any, auto, tool},
		&j,
	)

	return &j
}

// Experimental.
func NewToolChoice_Override(t ToolChoice, any interface{}, auto interface{}, tool *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolChoice",
		[]interface{}{any, auto, tool},
		t,
	)
}

// The Model must request the specified tool.
//
// Only supported by some models like Anthropic Claude 3 models.
// Experimental.
func ToolChoice_SpecificTool(toolName *string) ToolChoice {
	_init_.Initialize()

	if err := validateToolChoice_SpecificToolParameters(toolName); err != nil {
		panic(err)
	}
	var returns ToolChoice

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolChoice",
		"specificTool",
		[]interface{}{toolName},
		&returns,
	)

	return returns
}

func ToolChoice_ANY() ToolChoice {
	_init_.Initialize()
	var returns ToolChoice
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolChoice",
		"ANY",
		&returns,
	)
	return returns
}

func ToolChoice_AUTO() ToolChoice {
	_init_.Initialize()
	var returns ToolChoice
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolChoice",
		"AUTO",
		&returns,
	)
	return returns
}

