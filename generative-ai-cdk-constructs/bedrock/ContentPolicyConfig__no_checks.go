//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateContentPolicyConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewContentPolicyConfigParameters(scope constructs.Construct, id *string, props *[]*ContentPolicyConfigProps) error {
	return nil
}

