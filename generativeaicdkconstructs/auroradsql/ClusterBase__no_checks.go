//go:build no_runtime_type_checking

package auroradsql

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGrantConnectAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateClusterBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateClusterBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClusterBaseParameters(scope constructs.Construct, id *string) error {
	return nil
}

