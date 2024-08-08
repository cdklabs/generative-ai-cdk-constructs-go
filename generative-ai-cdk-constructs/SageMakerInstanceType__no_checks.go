//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func validateSageMakerInstanceType_OfParameters(instanceType *string) error {
	return nil
}

func validateNewSageMakerInstanceTypeParameters(instanceType *string) error {
	return nil
}

