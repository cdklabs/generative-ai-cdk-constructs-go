//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomTransformation) validateGeneratePolicyStatementsParameters(scope constructs.Construct) error {
	return nil
}

func validateCustomTransformation_LambdaParameters(props *LambdaCustomTransformationProps) error {
	return nil
}

func (j *jsiiProxy_CustomTransformation) validateSetConfigurationParameters(val *awsbedrock.CfnDataSource_CustomTransformationConfigurationProperty) error {
	return nil
}

