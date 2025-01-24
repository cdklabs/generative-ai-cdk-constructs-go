package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Represents the concept of an API Schema for a Bedrock Agent Action Group.
// Experimental.
type ApiSchema interface {
	// Experimental.
	InlineSchema() *string
	// Experimental.
	S3File() *awss3.Location
}

// The jsii proxy struct for ApiSchema
type jsiiProxy_ApiSchema struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Constructor accessible only to extending classes.
// Experimental.
func NewApiSchema_Override(a ApiSchema, s3File *awss3.Location, inlineSchema *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		[]interface{}{s3File, inlineSchema},
		a,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func ApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func ApiSchema_FromLocalAsset(path *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func ApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}

