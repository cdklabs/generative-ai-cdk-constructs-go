//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KnowledgeBase) validateAddConfluenceDataSourceParameters(props *ConfluenceDataSourceAssociationProps) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateAddS3DataSourceParameters(props *S3DataSourceAssociationProps) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateAddSalesforceDataSourceParameters(props *SalesforceDataSourceAssociationProps) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateAddSharePointDataSourceParameters(props *SharePointDataSourceAssociationProps) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateAddWebCrawlerDataSourceParameters(props *WebCrawlerDataSourceAssociationProps) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateAssociateToAgentParameters(agent Agent) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (k *jsiiProxy_KnowledgeBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateKnowledgeBase_FromKnowledgeBaseAttributesParameters(scope constructs.Construct, id *string, attrs *KnowledgeBaseAttributes) error {
	return nil
}

func validateKnowledgeBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKnowledgeBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateKnowledgeBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewKnowledgeBaseParameters(scope constructs.Construct, id *string, props *KnowledgeBaseProps) error {
	return nil
}

