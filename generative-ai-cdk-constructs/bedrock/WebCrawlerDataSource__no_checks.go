//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebCrawlerDataSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_WebCrawlerDataSource) validateFormatAsCfnPropsParameters(props *DataSourceAssociationProps, dataSourceConfiguration *awsbedrock.CfnDataSource_DataSourceConfigurationProperty) error {
	return nil
}

func (w *jsiiProxy_WebCrawlerDataSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_WebCrawlerDataSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_WebCrawlerDataSource) validateHandleCommonPermissionsParameters(props *DataSourceAssociationProps) error {
	return nil
}

func validateWebCrawlerDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWebCrawlerDataSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWebCrawlerDataSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewWebCrawlerDataSourceParameters(scope constructs.Construct, id *string, props *WebCrawlerDataSourceProps) error {
	return nil
}

