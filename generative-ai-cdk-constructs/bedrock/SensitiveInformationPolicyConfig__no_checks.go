//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateSensitiveInformationPolicyConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSensitiveInformationPolicyConfigParameters(scope constructs.Construct, id *string, props *[]*SensitiveInformationPolicyConfigProps) error {
	return nil
}

