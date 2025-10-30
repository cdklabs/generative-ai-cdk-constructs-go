//go:build !no_runtime_type_checking

package aosscwdashboard

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AossCwDashboard) validateAddCollectionMonitoringbyAttributesParameters(collectionName *string, collectionId *string, props *CollectionMonitoringProps) error {
	if collectionName == nil {
		return fmt.Errorf("parameter collectionName is required, but nil was provided")
	}

	if collectionId == nil {
		return fmt.Errorf("parameter collectionId is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AossCwDashboard) validateAddCollectionMonitoringByCollectionParameters(collection awsopensearchserverless.CfnCollection, props *CollectionMonitoringProps) error {
	if collection == nil {
		return fmt.Errorf("parameter collection is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AossCwDashboard) validateAddIndexMonitoringByAtributesParameters(collectionName *string, collectionId *string, IndexName *string, IndexId *string, props *IndexMonitoringProps) error {
	if collectionName == nil {
		return fmt.Errorf("parameter collectionName is required, but nil was provided")
	}

	if collectionId == nil {
		return fmt.Errorf("parameter collectionId is required, but nil was provided")
	}

	if IndexName == nil {
		return fmt.Errorf("parameter IndexName is required, but nil was provided")
	}

	if IndexId == nil {
		return fmt.Errorf("parameter IndexId is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAossCwDashboard_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewAossCwDashboardParameters(scope constructs.Construct, id *string, props *AossCwDashboardProps) error {
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

