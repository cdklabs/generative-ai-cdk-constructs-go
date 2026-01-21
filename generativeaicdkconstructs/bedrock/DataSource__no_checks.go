//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDataSource_FromDataSourceIdParameters(scope constructs.Construct, id *string, dataSourceId *string) error {
	return nil
}

func validateDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

