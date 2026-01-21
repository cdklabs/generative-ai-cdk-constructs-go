package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"
)

// Represents a supplemental data storage location for images extracted from multimodal documents in your data source.
// Experimental.
type SupplementalDataStorageLocation interface {
	// The configuration for the storage location.
	// Experimental.
	LocationConfig() *SupplementalDataStorageS3Config
	// The type of the storage location.
	// Experimental.
	Type() SupplementalDataStorageLocationType
}

// The jsii proxy struct for SupplementalDataStorageLocation
type jsiiProxy_SupplementalDataStorageLocation struct {
	_ byte // padding
}

func (j *jsiiProxy_SupplementalDataStorageLocation) LocationConfig() *SupplementalDataStorageS3Config {
	var returns *SupplementalDataStorageS3Config
	_jsii_.Get(
		j,
		"locationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SupplementalDataStorageLocation) Type() SupplementalDataStorageLocationType {
	var returns SupplementalDataStorageLocationType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates a new SupplementalDataStorageLocation.
// Experimental.
func NewSupplementalDataStorageLocation(type_ SupplementalDataStorageLocationType, locationConfig *SupplementalDataStorageS3Config) SupplementalDataStorageLocation {
	_init_.Initialize()

	if err := validateNewSupplementalDataStorageLocationParameters(type_, locationConfig); err != nil {
		panic(err)
	}
	j := jsiiProxy_SupplementalDataStorageLocation{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SupplementalDataStorageLocation",
		[]interface{}{type_, locationConfig},
		&j,
	)

	return &j
}

// Creates a new SupplementalDataStorageLocation.
// Experimental.
func NewSupplementalDataStorageLocation_Override(s SupplementalDataStorageLocation, type_ SupplementalDataStorageLocationType, locationConfig *SupplementalDataStorageS3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SupplementalDataStorageLocation",
		[]interface{}{type_, locationConfig},
		s,
	)
}

// Creates a new S3 supplemental data storage location.
//
// Returns: A new SupplementalDataStorageLocation instance.
// Experimental.
func SupplementalDataStorageLocation_S3(config *SupplementalDataStorageS3Config) SupplementalDataStorageLocation {
	_init_.Initialize()

	if err := validateSupplementalDataStorageLocation_S3Parameters(config); err != nil {
		panic(err)
	}
	var returns SupplementalDataStorageLocation

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SupplementalDataStorageLocation",
		"s3",
		[]interface{}{config},
		&returns,
	)

	return returns
}

