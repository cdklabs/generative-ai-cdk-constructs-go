//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateParsingStrategy_FoundationModelParameters(props *FoundationModelParsingStrategyProps) error {
	return nil
}

func (j *jsiiProxy_ParsingStrategy) validateSetConfigurationParameters(val *awsbedrock.CfnDataSource_ParsingConfigurationProperty) error {
	return nil
}

