package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Represents a Prompt, either created with CDK or imported.
// Experimental.
type IPrompt interface {
	// Optional KMS encryption key associated with this prompt.
	// Experimental.
	KmsKey() awskms.IKey
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
	// The version of the prompt.
	// Default: - "DRAFT".
	//
	// Experimental.
	PromptVersion() *string
	// Experimental.
	SetPromptVersion(p *string)
}

// The jsii proxy for IPrompt
type jsiiProxy_IPrompt struct {
	_ byte // padding
}

func (j *jsiiProxy_IPrompt) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_IPrompt) PromptVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrompt)SetPromptVersion(val *string) {
	if err := j.validateSetPromptVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promptVersion",
		val,
	)
}

