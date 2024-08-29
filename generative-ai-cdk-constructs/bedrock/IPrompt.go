package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Prompt, either created with CDK or imported.
// Experimental.
type IPrompt interface {
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
}

// The jsii proxy for IPrompt
type jsiiProxy_IPrompt struct {
	_ byte // padding
}

func (j *jsiiProxy_IPrompt) PromptArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrompt) PromptId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptId",
		&returns,
	)
	return returns
}

