//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSourceNew) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DataSourceNew) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (d *jsiiProxy_DataSourceNew) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DataSourceNew) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DataSourceNew) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateDataSourceNew_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataSourceNew_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDataSourceNew_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDataSourceNewParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

