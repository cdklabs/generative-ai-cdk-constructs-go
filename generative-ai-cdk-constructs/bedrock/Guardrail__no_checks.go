//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Guardrail) validateAddSensitiveInformationPolicyConfigParameters(props *[]*SensitiveInformationPolicyConfigProps, guardrailRegexesConfig *awsbedrock.CfnGuardrail_RegexConfigProperty) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddTagsParameters(props *GuardrailProps) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddTopicPolicyConfigParameters(topic Topic) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddVersionParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddWordPolicyConfigParameters(wordsFilter *[]*awsbedrock.CfnGuardrail_WordConfigProperty) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateUploadWordPolicyFromFileParameters(filePath *string) error {
	return nil
}

func validateGuardrail_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewGuardrailParameters(scope constructs.Construct, id *string, props *GuardrailProps) error {
	return nil
}

