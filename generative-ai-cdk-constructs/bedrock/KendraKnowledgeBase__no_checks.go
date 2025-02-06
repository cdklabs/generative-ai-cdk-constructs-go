//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraKnowledgeBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBase) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBase) validateGrantRetrieveParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBase) validateGrantRetrieveAndGenerateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateKendraKnowledgeBase_FromKnowledgeBaseAttributesParameters(scope constructs.Construct, id *string, attrs *KendraKnowledgeBaseAttributes) error {
	return nil
}

func validateKendraKnowledgeBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKendraKnowledgeBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKendraKnowledgeBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKendraKnowledgeBaseParameters(scope constructs.Construct, id *string, props *KendraKnowledgeBaseProps) error {
	return nil
}

