//go:build no_runtime_type_checking

package amazonaurora

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AmazonAuroraVectorStore) validateAddIngressRuleToAuroraSecurityGroupParameters(lambdaSecurityGroup awsec2.ISecurityGroup, auroraSecurityGroup awsec2.ISecurityGroup) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraVectorStore) validateCreateAuroraPgCRPolicyParameters(clusterIdentifier *string) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraVectorStore) validateCreateLambdaSecurityGroupParameters(vpc awsec2.IVpc) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraVectorStore) validateGenerateResourceArnParameters(clusterIdentifier *string) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraVectorStore) validateSetupCustomResourceParameters(databaseClusterResources *DatabaseClusterResources, lambdaSecurityGroup awsec2.SecurityGroup, auroraPgCRPolicy awsiam.ManagedPolicy) error {
	return nil
}

func (a *jsiiProxy_AmazonAuroraVectorStore) validateSetupDatabaseClusterResourcesParameters(vpc awsec2.IVpc, secret awssecretsmanager.ISecret, clusterIdentifier *string, auroraSecurityGroup awsec2.ISecurityGroup) error {
	return nil
}

func validateAmazonAuroraVectorStore_FromExistingAuroraVectorStoreParameters(scope constructs.Construct, id *string, props *ExistingAmazonAuroraVectorStoreProps) error {
	return nil
}

func validateAmazonAuroraVectorStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAmazonAuroraVectorStoreParameters(scope constructs.Construct, id *string, props *AmazonAuroraVectorStoreProps) error {
	return nil
}

