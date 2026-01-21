//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateContextEnrichment_FoundationModelParameters(props *FoundationModelContextEnrichmentProps) error {
	return nil
}

func (j *jsiiProxy_ContextEnrichment) validateSetConfigurationParameters(val *awsbedrock.CfnDataSource_ContextEnrichmentConfigurationProperty) error {
	return nil
}

