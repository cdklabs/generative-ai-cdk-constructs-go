//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"
)

func (j *jsiiProxy_IPrompt) validateSetPromptVersionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

