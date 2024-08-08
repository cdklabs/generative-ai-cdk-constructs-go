//go:build !no_runtime_type_checking

package generative-ai-cdk-constructs

import (
	"fmt"
)

func validateSageMakerInstanceType_OfParameters(instanceType *string) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

func validateNewSageMakerInstanceTypeParameters(instanceType *string) error {
	if instanceType == nil {
		return fmt.Errorf("parameter instanceType is required, but nil was provided")
	}

	return nil
}

