//go:build no_runtime_type_checking

package opensearchvectorindex

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorIndex) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VectorIndex) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VectorIndex) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVectorIndex_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVectorIndex_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVectorIndex_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVectorIndexParameters(scope constructs.Construct, id *string, props *VectorIndexProps) error {
	return nil
}

