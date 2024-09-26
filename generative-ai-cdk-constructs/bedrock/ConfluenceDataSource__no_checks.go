//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfluenceDataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConfluenceDataSource) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (c *jsiiProxy_ConfluenceDataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConfluenceDataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ConfluenceDataSource) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateConfluenceDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfluenceDataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConfluenceDataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConfluenceDataSourceParameters(scope constructs.Construct, id *string, props *ConfluenceDataSourceProps) error {
	return nil
}

