//go:build !no_runtime_type_checking

package sagemakerdeployment

import (
	"fmt"
)

func (j *jsiiProxy_IInstanceAliase) validateSetAliasesParameters(val *map[string]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IInstanceAliase) validateSetRegionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

