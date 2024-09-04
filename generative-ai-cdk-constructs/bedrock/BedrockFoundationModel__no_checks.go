//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockFoundationModel) validateAsArnParameters(construct constructs.IConstruct) error {
	return nil
}

func (b *jsiiProxy_BedrockFoundationModel) validateAsIModelParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBedrockFoundationModelParameters(value *string, props *BedrockFoundationModelProps) error {
	return nil
}

