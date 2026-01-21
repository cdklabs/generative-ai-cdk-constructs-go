package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Contains details about the Lambda function containing the orchestration logic carried out upon invoking the custom orchestration.
// Experimental.
type OrchestrationExecutor interface {
	// Experimental.
	LambdaFunction() awslambda.IFunction
}

// The jsii proxy struct for OrchestrationExecutor
type jsiiProxy_OrchestrationExecutor struct {
	_ byte // padding
}

func (j *jsiiProxy_OrchestrationExecutor) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}


// Defines an orchestration executor with a Lambda function containing the business logic.
// Experimental.
func OrchestrationExecutor_FromlambdaFunction(lambdaFunction awslambda.IFunction) OrchestrationExecutor {
	_init_.Initialize()

	if err := validateOrchestrationExecutor_FromlambdaFunctionParameters(lambdaFunction); err != nil {
		panic(err)
	}
	var returns OrchestrationExecutor

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.OrchestrationExecutor",
		"fromlambdaFunction",
		[]interface{}{lambdaFunction},
		&returns,
	)

	return returns
}

