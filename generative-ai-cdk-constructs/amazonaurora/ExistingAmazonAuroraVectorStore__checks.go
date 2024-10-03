//go:build !no_runtime_type_checking

package amazonaurora

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
)

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateAddIngressRuleToAuroraSecurityGroupParameters(lambdaSecurityGroup awsec2.SecurityGroup, auroraSecurityGroup awsec2.SecurityGroup) error {
	if lambdaSecurityGroup == nil {
		return fmt.Errorf("parameter lambdaSecurityGroup is required, but nil was provided")
	}

	if auroraSecurityGroup == nil {
		return fmt.Errorf("parameter auroraSecurityGroup is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateCreateAuroraPgCRPolicyParameters(clusterIdentifier *string) error {
	if clusterIdentifier == nil {
		return fmt.Errorf("parameter clusterIdentifier is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateCreateLambdaSecurityGroupParameters(vpc awsec2.IVpc) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateGenerateResourceArnParameters(clusterIdentifier *string) error {
	if clusterIdentifier == nil {
		return fmt.Errorf("parameter clusterIdentifier is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateSetupCustomResourceParameters(databaseClusterResources *DatabaseClusterResources, lambdaSecurityGroup awsec2.SecurityGroup, auroraPgCRPolicy awsiam.ManagedPolicy) error {
	if databaseClusterResources == nil {
		return fmt.Errorf("parameter databaseClusterResources is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(databaseClusterResources, func() string { return "parameter databaseClusterResources" }); err != nil {
		return err
	}

	if lambdaSecurityGroup == nil {
		return fmt.Errorf("parameter lambdaSecurityGroup is required, but nil was provided")
	}

	if auroraPgCRPolicy == nil {
		return fmt.Errorf("parameter auroraPgCRPolicy is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExistingAmazonAuroraVectorStore) validateSetupDatabaseClusterResourcesParameters(vpc awsec2.IVpc, secret awssecretsmanager.ISecret, clusterIdentifier *string, auroraSecurityGroupId *string) error {
	if vpc == nil {
		return fmt.Errorf("parameter vpc is required, but nil was provided")
	}

	if secret == nil {
		return fmt.Errorf("parameter secret is required, but nil was provided")
	}

	if clusterIdentifier == nil {
		return fmt.Errorf("parameter clusterIdentifier is required, but nil was provided")
	}

	if auroraSecurityGroupId == nil {
		return fmt.Errorf("parameter auroraSecurityGroupId is required, but nil was provided")
	}

	return nil
}

func validateExistingAmazonAuroraVectorStore_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewExistingAmazonAuroraVectorStoreParameters(scope constructs.Construct, id *string, props *ExistingAmazonAuroraVectorStoreProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

