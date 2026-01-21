//go:build !no_runtime_type_checking

package sagemakerdeployment

import (
	"fmt"
)

func validateJumpStartModel_OfParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewJumpStartModelParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

