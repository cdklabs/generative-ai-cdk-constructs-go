//go:build !no_runtime_type_checking

package sagemakerdeployment

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs"
)

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateAddObservabilityToConstructParameters(props *generativeaicdkconstructs.BaseClassProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	if statement == nil {
		return fmt.Errorf("parameter statement is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateUpdateConstructUsageMetricCodeParameters(props *generativeaicdkconstructs.BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if lambdaFunctions == nil {
		return fmt.Errorf("parameter lambdaFunctions is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateUpdateEnvSuffixParameters(props *generativeaicdkconstructs.BaseClassProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateJumpStartSageMakerEndpoint_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateSetEnablexrayParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateSetFieldLogLevelParameters(val awsappsync.FieldLogLevel) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateSetLambdaTracingParameters(val awslambda.Tracing) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateSetRetentionParameters(val awslogs.RetentionDays) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) validateSetStageParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateJumpStartSageMakerEndpoint_SetUsageMetricMapParameters(val *map[string]*float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewJumpStartSageMakerEndpointParameters(scope constructs.Construct, id *string, props *JumpStartSageMakerEndpointProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

