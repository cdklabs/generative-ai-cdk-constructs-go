//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateToolChoice_SpecificToolParameters(toolName *string) error {
	return nil
}

func validateNewToolChoiceParameters(any interface{}, auto interface{}) error {
	return nil
}

