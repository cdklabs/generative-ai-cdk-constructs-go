//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func validateLangchainCommonDepsLayer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewLangchainCommonDepsLayerParameters(scope constructs.Construct, id *string, props *LangchainLayerProps) error {
	return nil
}

