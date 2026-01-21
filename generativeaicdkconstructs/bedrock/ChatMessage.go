package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"
)

// Experimental.
type ChatMessage interface {
	// Experimental.
	Role() ChatMessageRole
	// Experimental.
	Text() *string
}

// The jsii proxy struct for ChatMessage
type jsiiProxy_ChatMessage struct {
	_ byte // padding
}

func (j *jsiiProxy_ChatMessage) Role() ChatMessageRole {
	var returns ChatMessageRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChatMessage) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}


// Experimental.
func NewChatMessage(role ChatMessageRole, text *string) ChatMessage {
	_init_.Initialize()

	if err := validateNewChatMessageParameters(role, text); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChatMessage{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatMessage",
		[]interface{}{role, text},
		&j,
	)

	return &j
}

// Experimental.
func NewChatMessage_Override(c ChatMessage, role ChatMessageRole, text *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatMessage",
		[]interface{}{role, text},
		c,
	)
}

// Experimental.
func ChatMessage_Assistant(text *string) ChatMessage {
	_init_.Initialize()

	if err := validateChatMessage_AssistantParameters(text); err != nil {
		panic(err)
	}
	var returns ChatMessage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatMessage",
		"assistant",
		[]interface{}{text},
		&returns,
	)

	return returns
}

// Experimental.
func ChatMessage_User(text *string) ChatMessage {
	_init_.Initialize()

	if err := validateChatMessage_UserParameters(text); err != nil {
		panic(err)
	}
	var returns ChatMessage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatMessage",
		"user",
		[]interface{}{text},
		&returns,
	)

	return returns
}

