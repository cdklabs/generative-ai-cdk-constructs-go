//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSupplementalDataStorageLocation_S3Parameters(config *SupplementalDataStorageS3Config) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

func validateNewSupplementalDataStorageLocationParameters(type_ SupplementalDataStorageLocationType, locationConfig *SupplementalDataStorageS3Config) error {
	if type_ == "" {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	if locationConfig == nil {
		return fmt.Errorf("parameter locationConfig is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(locationConfig, func() string { return "parameter locationConfig" }); err != nil {
		return err
	}

	return nil
}

