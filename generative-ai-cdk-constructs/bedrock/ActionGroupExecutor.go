package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type ActionGroupExecutor struct {
	// To return the action group invocation results directly in the InvokeAgent response, specify RETURN_CONTROL .
	// Experimental.
	CustomControl *string `field:"optional" json:"customControl" yaml:"customControl"`
	// The Lambda function containing the business logic that is carried out upon invoking the action.
	// Experimental.
	Lambda awslambda.IFunction `field:"optional" json:"lambda" yaml:"lambda"`
}

