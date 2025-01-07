//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validatePromptVariant_AgentParameters(props *AgentPromptVariantProps) error {
	return nil
}

func validatePromptVariant_ChatParameters(props *ChatPromptVariantProps) error {
	return nil
}

func validatePromptVariant_TextParameters(props *TextPromptVariantProps) error {
	return nil
}

func (j *jsiiProxy_PromptVariant) validateSetGenAiResourceParameters(val *awsbedrock.CfnPrompt_PromptGenAiResourceProperty) error {
	return nil
}

func (j *jsiiProxy_PromptVariant) validateSetInferenceConfigurationParameters(val *awsbedrock.CfnPrompt_PromptInferenceConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_PromptVariant) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PromptVariant) validateSetTemplateConfigurationParameters(val *awsbedrock.CfnPrompt_PromptTemplateConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_PromptVariant) validateSetTemplateTypeParameters(val PromptTemplateType) error {
	return nil
}

