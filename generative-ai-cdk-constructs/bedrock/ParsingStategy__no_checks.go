//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateParsingStategy_FoundationModelParameters(props *FoundationModelParsingStategyProps) error {
	return nil
}

func (j *jsiiProxy_ParsingStategy) validateSetConfigurationParameters(val *awsbedrock.CfnDataSource_ParsingConfigurationProperty) error {
	return nil
}

