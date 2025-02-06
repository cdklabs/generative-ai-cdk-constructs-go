//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateGrantRetrieveParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) validateGrantRetrieveAndGenerateParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateKendraKnowledgeBaseBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKendraKnowledgeBaseBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKendraKnowledgeBaseBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKendraKnowledgeBaseBaseParameters(scope constructs.Construct, id *string) error {
	return nil
}

