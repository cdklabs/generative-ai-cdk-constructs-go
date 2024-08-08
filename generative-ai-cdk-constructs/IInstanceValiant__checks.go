//go:build !no_runtime_type_checking

package generative-ai-cdk-constructs

import (
	"fmt"
)

func (j *jsiiProxy_IInstanceValiant) validateSetInstanceTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

