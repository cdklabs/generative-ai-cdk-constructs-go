package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type PromptOverrideConfiguration interface {
	// The custom Lambda parser function to use.
	//
	// The Lambda parser processes and interprets the raw foundation model output.
	// It receives an input event with:
	// - messageVersion: Version of message format (1.0)
	// - agent: Info about the agent (name, id, alias, version)
	// - invokeModelRawResponse: Raw model output to parse
	// - promptType: Type of prompt being parsed
	// - overrideType: Type of override (OUTPUT_PARSER)
	//
	// The Lambda must return a response that the agent uses for next actions.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/lambda-parser.html
	//
	// Experimental.
	Parser() awslambda.IFunction
	// The prompt configurations to override the prompt templates in the agent sequence.
	// Default: - No prompt configuration will be overridden.
	//
	// Experimental.
	Steps() *[]*PromptStepConfigurationCustomParser
}

// The jsii proxy struct for PromptOverrideConfiguration
type jsiiProxy_PromptOverrideConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_PromptOverrideConfiguration) Parser() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"parser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) Steps() *[]*PromptStepConfigurationCustomParser {
	var returns *[]*PromptStepConfigurationCustomParser
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}


// Experimental.
func PromptOverrideConfiguration_FromSteps(steps *[]*PromptStepConfiguration) PromptOverrideConfiguration {
	_init_.Initialize()

	if err := validatePromptOverrideConfiguration_FromStepsParameters(steps); err != nil {
		panic(err)
	}
	var returns PromptOverrideConfiguration

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptOverrideConfiguration",
		"fromSteps",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Creates a PromptOverrideConfiguration with a custom Lambda parser function.
// Experimental.
func PromptOverrideConfiguration_WithCustomParser(props *CustomParserProps) PromptOverrideConfiguration {
	_init_.Initialize()

	if err := validatePromptOverrideConfiguration_WithCustomParserParameters(props); err != nil {
		panic(err)
	}
	var returns PromptOverrideConfiguration

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptOverrideConfiguration",
		"withCustomParser",
		[]interface{}{props},
		&returns,
	)

	return returns
}

