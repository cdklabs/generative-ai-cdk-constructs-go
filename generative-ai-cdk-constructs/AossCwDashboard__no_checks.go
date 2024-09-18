//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AossCwDashboard) validateAddCollectionMonitoringbyAttributesParameters(collectionName *string, collectionId *string, props *CollectionMonitoringProps) error {
	return nil
}

func (a *jsiiProxy_AossCwDashboard) validateAddCollectionMonitoringByCollectionParameters(collection awsopensearchserverless.CfnCollection, props *CollectionMonitoringProps) error {
	return nil
}

func (a *jsiiProxy_AossCwDashboard) validateAddIndexMonitoringByAtributesParameters(collectionName *string, collectionId *string, IndexName *string, IndexId *string, props *IndexMonitoringProps) error {
	return nil
}

func validateAossCwDashboard_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAossCwDashboardParameters(scope constructs.Construct, id *string, props *AossCwDashboardProps) error {
	return nil
}

