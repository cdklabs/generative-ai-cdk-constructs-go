//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSourceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DataSourceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DataSourceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDataSourceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDataSourceBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDataSourceBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDataSourceBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

