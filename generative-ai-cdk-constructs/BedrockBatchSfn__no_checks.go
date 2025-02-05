//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockBatchSfn) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (b *jsiiProxy_BedrockBatchSfn) validateToSingleStateParameters(options *awsstepfunctions.SingleStateOptions) error {
	return nil
}

func validateBedrockBatchSfn_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBedrockBatchSfnParameters(parent constructs.Construct, id *string, props *BedrockBatchSfnProps) error {
	return nil
}

