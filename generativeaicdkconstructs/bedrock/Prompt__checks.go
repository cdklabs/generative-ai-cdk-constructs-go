//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_Prompt) validateAddVariantParameters(variant PromptVariant) error {
	if variant == nil {
		return fmt.Errorf("parameter variant is required, but nil was provided")
	}

	return nil
}

func validatePrompt_FromPromptAttributesParameters(scope constructs.Construct, id *string, attrs *PromptAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validatePrompt_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Prompt) validateSetPromptVersionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPromptParameters(scope constructs.Construct, id *string, props *PromptProps) error {
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

