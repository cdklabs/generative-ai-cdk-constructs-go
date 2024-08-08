//go:build no_runtime_type_checking

package amazonaurora

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmazonAuroraDefaultVectorStore) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraDefaultVectorStore) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraDefaultVectorStore) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAmazonAuroraDefaultVectorStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAmazonAuroraDefaultVectorStore_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAmazonAuroraDefaultVectorStore_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAmazonAuroraDefaultVectorStoreParameters(scope constructs.Construct, id *string, props *AmazonAuroraDefaultVectorStoreProps) error {
	return nil
}

