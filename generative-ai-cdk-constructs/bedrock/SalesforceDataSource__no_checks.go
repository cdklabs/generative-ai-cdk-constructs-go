//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SalesforceDataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SalesforceDataSource) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (s *jsiiProxy_SalesforceDataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SalesforceDataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SalesforceDataSource) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateSalesforceDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSalesforceDataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSalesforceDataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSalesforceDataSourceParameters(scope constructs.Construct, id *string, props *SalesforceDataSourceProps) error {
	return nil
}

