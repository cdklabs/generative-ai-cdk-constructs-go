package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Defines how fulfillment of the action group is handled after the necessary information has been elicited from the user.
//
// Valid executors are:
// - Lambda function
// - Return Control.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/action-handle.html
//
// Experimental.
type ActionGroupExecutor interface {
	// Experimental.
	CustomControl() *string
	// Experimental.
	LambdaFunction() awslambda.IFunction
}

// The jsii proxy struct for ActionGroupExecutor
type jsiiProxy_ActionGroupExecutor struct {
	_ byte // padding
}

func (j *jsiiProxy_ActionGroupExecutor) CustomControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionGroupExecutor) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}


// Defines an action group with a Lambda function containing the business logic.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-lambda.html
//
// Experimental.
func ActionGroupExecutor_FromlambdaFunction(lambdaFunction awslambda.IFunction) ActionGroupExecutor {
	_init_.Initialize()

	if err := validateActionGroupExecutor_FromlambdaFunctionParameters(lambdaFunction); err != nil {
		panic(err)
	}
	var returns ActionGroupExecutor

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ActionGroupExecutor",
		"fromlambdaFunction",
		[]interface{}{lambdaFunction},
		&returns,
	)

	return returns
}

func ActionGroupExecutor_RETURN_CONTROL() ActionGroupExecutor {
	_init_.Initialize()
	var returns ActionGroupExecutor
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ActionGroupExecutor",
		"RETURN_CONTROL",
		&returns,
	)
	return returns
}

