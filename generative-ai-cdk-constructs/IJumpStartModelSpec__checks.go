//go:build !no_runtime_type_checking

package generative-ai-cdk-constructs

import (
	"fmt"
)

func (j *jsiiProxy_IJumpStartModelSpec) validateSetDefaultInstanceTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IJumpStartModelSpec) validateSetEnvironmentParameters(val *map[string]interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		switch v.(type) {
		case *string:
			// ok
		case string:
			// ok
		case *float64:
			// ok
		case float64:
			// ok
		case *int:
			// ok
		case int:
			// ok
		case *uint:
			// ok
		case uint:
			// ok
		case *int8:
			// ok
		case int8:
			// ok
		case *int16:
			// ok
		case int16:
			// ok
		case *int32:
			// ok
		case int32:
			// ok
		case *int64:
			// ok
		case int64:
			// ok
		case *uint8:
			// ok
		case uint8:
			// ok
		case *uint16:
			// ok
		case uint16:
			// ok
		case *uint32:
			// ok
		case uint32:
			// ok
		case *uint64:
			// ok
		case uint64:
			// ok
		case *bool:
			// ok
		case bool:
			// ok
		default:
			return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *string, *float64, *bool; received %#v (a %T)", idx_97dfc6, v, v)
		}
	}

	return nil
}

func (j *jsiiProxy_IJumpStartModelSpec) validateSetGatedBucketParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IJumpStartModelSpec) validateSetInstanceTypesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IJumpStartModelSpec) validateSetModelIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IJumpStartModelSpec) validateSetRequiresEulaParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IJumpStartModelSpec) validateSetVersionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

