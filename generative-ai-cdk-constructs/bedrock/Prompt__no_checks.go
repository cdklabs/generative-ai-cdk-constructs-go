//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Prompt) validateAddVariantParameters(variant PromptVariant) error {
	return nil
}

func validatePrompt_FromPromptAttributesParameters(scope constructs.Construct, id *string, attrs *PromptAttributes) error {
	return nil
}

func validatePrompt_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Prompt) validateSetPromptVersionParameters(val *string) error {
	return nil
}

func validateNewPromptParameters(scope constructs.Construct, id *string, props *PromptProps) error {
	return nil
}

