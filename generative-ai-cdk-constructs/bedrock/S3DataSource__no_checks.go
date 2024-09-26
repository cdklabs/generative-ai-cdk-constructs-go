//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3DataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_S3DataSource) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (s *jsiiProxy_S3DataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_S3DataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_S3DataSource) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateS3DataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3DataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateS3DataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewS3DataSourceParameters(scope constructs.Construct, id *string, props *S3DataSourceProps) error {
	return nil
}

