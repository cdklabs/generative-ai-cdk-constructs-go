//go:build !no_runtime_type_checking

package generative-ai-cdk-constructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

func (q *jsiiProxy_QaAppsyncOpensearch) validateAddObservabilityToConstructParameters(props *BaseClassProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (q *jsiiProxy_QaAppsyncOpensearch) validateUpdateConstructUsageMetricCodeParameters(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) error {
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

func (q *jsiiProxy_QaAppsyncOpensearch) validateUpdateEnvSuffixParameters(props *BaseClassProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateQaAppsyncOpensearch_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QaAppsyncOpensearch) validateSetEnablexrayParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QaAppsyncOpensearch) validateSetFieldLogLevelParameters(val awsappsync.FieldLogLevel) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QaAppsyncOpensearch) validateSetLambdaTracingParameters(val awslambda.Tracing) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QaAppsyncOpensearch) validateSetRetentionParameters(val awslogs.RetentionDays) error {
	if val == "" {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_QaAppsyncOpensearch) validateSetStageParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateQaAppsyncOpensearch_SetUsageMetricMapParameters(val *map[string]*float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewQaAppsyncOpensearchParameters(scope constructs.Construct, id *string, props *QaAppsyncOpensearchProps) error {
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

