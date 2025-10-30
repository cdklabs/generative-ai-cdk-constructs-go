//go:build !no_runtime_type_checking

package bedrockcwdashboard

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock"
)

func (b *jsiiProxy_BedrockCwDashboard) validateAddAllModelsMonitoringParameters(props *ModelMonitoringProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockCwDashboard) validateAddGuardrailMonitoringParameters(guardrail bedrock.IGuardrail) error {
	if guardrail == nil {
		return fmt.Errorf("parameter guardrail is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockCwDashboard) validateAddModelMonitoringParameters(modelName *string, modelId *string, props *ModelMonitoringProps) error {
	if modelName == nil {
		return fmt.Errorf("parameter modelName is required, but nil was provided")
	}

	if modelId == nil {
		return fmt.Errorf("parameter modelId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateBedrockCwDashboard_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewBedrockCwDashboardParameters(scope constructs.Construct, id *string, props *BedrockCwDashboardProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

