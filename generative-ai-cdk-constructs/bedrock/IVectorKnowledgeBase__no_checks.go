//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVectorKnowledgeBase) validateAddConfluenceDataSourceParameters(props *ConfluenceDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IVectorKnowledgeBase) validateAddS3DataSourceParameters(props *S3DataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IVectorKnowledgeBase) validateAddSalesforceDataSourceParameters(props *SalesforceDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IVectorKnowledgeBase) validateAddSharePointDataSourceParameters(props *SharePointDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IVectorKnowledgeBase) validateAddWebCrawlerDataSourceParameters(props *WebCrawlerDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IVectorKnowledgeBase) validateGrantRetrieveParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVectorKnowledgeBase) validateGrantRetrieveAndGenerateParameters(grantee awsiam.IGrantable) error {
	return nil
}

