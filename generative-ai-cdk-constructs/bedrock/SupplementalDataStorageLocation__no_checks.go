//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateSupplementalDataStorageLocation_S3Parameters(config *SupplementalDataStorageS3Config) error {
	return nil
}

func validateNewSupplementalDataStorageLocationParameters(type_ SupplementalDataStorageLocationType, locationConfig *SupplementalDataStorageS3Config) error {
	return nil
}

