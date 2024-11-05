//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IPrompt) validateSetPromptVersionParameters(val *string) error {
	return nil
}

