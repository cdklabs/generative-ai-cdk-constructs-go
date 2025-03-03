package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Memory class for managing Bedrock Agent memory configurations.
//
// Enables conversational context retention
// across multiple sessions through session identifiers. Memory context is stored with unique
// memory IDs per user, allowing access to conversation history and summaries. Supports viewing
// stored sessions and clearing memory.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-memory.html
//
// Experimental.
type Memory interface {
}

// The jsii proxy struct for Memory
type jsiiProxy_Memory struct {
	_ byte // padding
}

// Experimental.
func NewMemory() Memory {
	_init_.Initialize()

	j := jsiiProxy_Memory{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Memory",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewMemory_Override(m Memory) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Memory",
		nil, // no parameters
		m,
	)
}

// Creates a session summary memory with custom configuration.
//
// Returns: Memory configuration object.
// Experimental.
func Memory_SessionSummary(props *SessionSummaryMemoryProps) *awsbedrock.CfnAgent_MemoryConfigurationProperty {
	_init_.Initialize()

	if err := validateMemory_SessionSummaryParameters(props); err != nil {
		panic(err)
	}
	var returns *awsbedrock.CfnAgent_MemoryConfigurationProperty

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Memory",
		"sessionSummary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func Memory_SESSION_SUMMARY() *awsbedrock.CfnAgent_MemoryConfigurationProperty {
	_init_.Initialize()
	var returns *awsbedrock.CfnAgent_MemoryConfigurationProperty
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Memory",
		"SESSION_SUMMARY",
		&returns,
	)
	return returns
}

