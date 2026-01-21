//go:build no_runtime_type_checking

package kendra

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraGenAiIndex) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KendraGenAiIndex) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KendraGenAiIndex) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKendraGenAiIndex_FromAttrsParameters(scope constructs.Construct, id *string, attrs *KendraGenAiIndexAttributes) error {
	return nil
}

func validateKendraGenAiIndex_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKendraGenAiIndex_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKendraGenAiIndex_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKendraGenAiIndexParameters(scope constructs.Construct, id *string, props *KendraGenAiIndexProps) error {
	return nil
}

