//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SharePointDataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SharePointDataSource) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (s *jsiiProxy_SharePointDataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SharePointDataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SharePointDataSource) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateSharePointDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSharePointDataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSharePointDataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSharePointDataSourceParameters(scope constructs.Construct, id *string, props *SharePointDataSourceProps) error {
	return nil
}

