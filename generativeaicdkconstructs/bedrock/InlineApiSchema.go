package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type InlineApiSchema interface {
	ApiSchema
	// Experimental.
	InlineSchema() *string
	// Experimental.
	S3File() *awss3.Location
}

// The jsii proxy struct for InlineApiSchema
type jsiiProxy_InlineApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_InlineApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InlineApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Experimental.
func NewInlineApiSchema(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateNewInlineApiSchemaParameters(schema); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineApiSchema{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		[]interface{}{schema},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineApiSchema_Override(i InlineApiSchema, schema *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		[]interface{}{schema},
		i,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func InlineApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func InlineApiSchema_FromLocalAsset(path *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func InlineApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}

