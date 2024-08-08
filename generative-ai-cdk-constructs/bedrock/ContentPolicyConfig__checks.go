//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateContentPolicyConfig_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewContentPolicyConfigParameters(scope constructs.Construct, id *string, props *[]*ContentPolicyConfigProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	for idx_b26480, v := range *props {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter props[%#v]", idx_b26480) }); err != nil {
			return err
		}
	}

	return nil
}

