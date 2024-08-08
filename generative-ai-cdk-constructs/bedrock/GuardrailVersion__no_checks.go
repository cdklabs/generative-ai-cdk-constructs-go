//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateGuardrailVersion_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGuardrailVersionParameters(scope constructs.Construct, id *string, props *awsbedrock.CfnGuardrailVersionProps) error {
	return nil
}

