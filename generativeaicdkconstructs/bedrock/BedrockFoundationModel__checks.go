//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BedrockFoundationModel) validateAsArnParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockFoundationModel) validateAsIModelParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockFoundationModel) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockFoundationModel) validateGrantInvokeAllRegionsParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func validateBedrockFoundationModel_FromCdkFoundationModelParameters(modelId awsbedrock.FoundationModel, props *BedrockFoundationModelProps) error {
	if modelId == nil {
		return fmt.Errorf("parameter modelId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateBedrockFoundationModel_FromCdkFoundationModelIdParameters(modelId awsbedrock.FoundationModelIdentifier, props *BedrockFoundationModelProps) error {
	if modelId == nil {
		return fmt.Errorf("parameter modelId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewBedrockFoundationModelParameters(value *string, props *BedrockFoundationModelProps) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

