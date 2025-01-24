package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Class to define an API Schema from an S3 object.
// Experimental.
type S3ApiSchema interface {
	ApiSchema
	// Experimental.
	InlineSchema() *string
	// Experimental.
	S3File() *awss3.Location
}

// The jsii proxy struct for S3ApiSchema
type jsiiProxy_S3ApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_S3ApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3ApiSchema(location *awss3.Location) S3ApiSchema {
	_init_.Initialize()

	if err := validateNewS3ApiSchemaParameters(location); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ApiSchema{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		[]interface{}{location},
		&j,
	)

	return &j
}

// Experimental.
func NewS3ApiSchema_Override(s S3ApiSchema, location *awss3.Location) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		[]interface{}{location},
		s,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func S3ApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func S3ApiSchema_FromLocalAsset(path *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func S3ApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}

