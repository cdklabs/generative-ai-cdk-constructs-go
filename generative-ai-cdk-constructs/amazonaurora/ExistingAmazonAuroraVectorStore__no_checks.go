//go:build no_runtime_type_checking

package amazonaurora

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateAddIngressRuleToAuroraSecurityGroupParameters(lambdaSecurityGroup awsec2.ISecurityGroup, auroraSecurityGroup awsec2.ISecurityGroup) error {
	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateCreateAuroraPgCRPolicyParameters(clusterIdentifier *string) error {
	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateCreateLambdaSecurityGroupParameters(vpc awsec2.IVpc) error {
	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateGenerateResourceArnParameters(clusterIdentifier *string) error {
	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateSetupCustomResourceParameters(databaseClusterResources *DatabaseClusterResources, lambdaSecurityGroup awsec2.SecurityGroup, auroraPgCRPolicy awsiam.ManagedPolicy) error {
	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateSetupDatabaseClusterResourcesParameters(vpc awsec2.IVpc, secret awssecretsmanager.ISecret, clusterIdentifier *string, auroraSecurityGroup awsec2.ISecurityGroup) error {
	return nil
}

func validateExistingAmazonAuroraVectorStore_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewExistingAmazonAuroraVectorStoreParameters(scope constructs.Construct, id *string, props *ExistingAmazonAuroraVectorStoreProps) error {
	return nil
}

