//go:build no_runtime_type_checking

package kendra

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraGenAiIndexBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KendraGenAiIndexBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KendraGenAiIndexBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKendraGenAiIndexBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKendraGenAiIndexBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKendraGenAiIndexBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKendraGenAiIndexBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

