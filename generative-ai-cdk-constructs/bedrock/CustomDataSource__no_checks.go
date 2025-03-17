//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomDataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CustomDataSource) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (c *jsiiProxy_CustomDataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CustomDataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_CustomDataSource) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateCustomDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCustomDataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCustomDataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCustomDataSourceParameters(scope constructs.Construct, id *string, props *CustomDataSourceProps) error {
	return nil
}

