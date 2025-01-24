//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IKnowledgeBase) validateAddConfluenceDataSourceParameters(props *ConfluenceDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateAddS3DataSourceParameters(props *S3DataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateAddSalesforceDataSourceParameters(props *SalesforceDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateAddSharePointDataSourceParameters(props *SharePointDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateAddWebCrawlerDataSourceParameters(props *WebCrawlerDataSourceAssociationProps) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateGrantRetrieveParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKnowledgeBase) validateGrantRetrieveAndGenerateParameters(grantee awsiam.IGrantable) error {
	return nil
}

