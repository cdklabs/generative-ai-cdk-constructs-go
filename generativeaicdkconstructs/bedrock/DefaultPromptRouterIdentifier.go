package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"
)

// Represents identifiers for default prompt routers in Bedrock.
// Experimental.
type DefaultPromptRouterIdentifier interface {
	// Experimental.
	PromptRouterId() *string
	// Experimental.
	RoutingModels() *[]BedrockFoundationModel
}

// The jsii proxy struct for DefaultPromptRouterIdentifier
type jsiiProxy_DefaultPromptRouterIdentifier struct {
	_ byte // padding
}

func (j *jsiiProxy_DefaultPromptRouterIdentifier) PromptRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultPromptRouterIdentifier) RoutingModels() *[]BedrockFoundationModel {
	var returns *[]BedrockFoundationModel
	_jsii_.Get(
		j,
		"routingModels",
		&returns,
	)
	return returns
}


func DefaultPromptRouterIdentifier_ANTHROPIC_CLAUDE_V1() DefaultPromptRouterIdentifier {
	_init_.Initialize()
	var returns DefaultPromptRouterIdentifier
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DefaultPromptRouterIdentifier",
		"ANTHROPIC_CLAUDE_V1",
		&returns,
	)
	return returns
}

func DefaultPromptRouterIdentifier_META_LLAMA_3_1() DefaultPromptRouterIdentifier {
	_init_.Initialize()
	var returns DefaultPromptRouterIdentifier
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DefaultPromptRouterIdentifier",
		"META_LLAMA_3_1",
		&returns,
	)
	return returns
}

