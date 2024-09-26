//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateChunkingStrategy_FixedSizeParameters(props *awsbedrock.CfnDataSource_FixedSizeChunkingConfigurationProperty) error {
	return nil
}

func validateChunkingStrategy_HierarchicalParameters(props *HierarchicalChunkingProps) error {
	return nil
}

func validateChunkingStrategy_SemanticParameters(props *awsbedrock.CfnDataSource_SemanticChunkingConfigurationProperty) error {
	return nil
}

func (j *jsiiProxy_ChunkingStrategy) validateSetConfigurationParameters(val *awsbedrock.CfnDataSource_ChunkingConfigurationProperty) error {
	return nil
}

