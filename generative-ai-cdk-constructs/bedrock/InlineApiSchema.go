package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// API Schema from a string value.
// Experimental.
type InlineApiSchema interface {
	ApiSchema
	// Called when the action group is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(_scope constructs.Construct) *ApiSchemaConfig
}

// The jsii proxy struct for InlineApiSchema
type jsiiProxy_InlineApiSchema struct {
	jsiiProxy_ApiSchema
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

// Loads the API Schema from a local disk path.
//
// Returns: `InlineApiSchema` with the contents of `path`.
// Experimental.
func InlineApiSchema_FromAsset(path *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromAssetParameters(path); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		"fromAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// API Schema as an S3 object.
//
// Returns: `S3ApiSchema` with the S3 bucket and key.
// Experimental.
func InlineApiSchema_FromBucket(bucket awss3.IBucket, key *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		"fromBucket",
		[]interface{}{bucket, key},
		&returns,
	)

	return returns
}

// Inline code for API Schema.
//
// Returns: `InlineApiSchema` with inline schema.
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

func (i *jsiiProxy_InlineApiSchema) Bind(_scope constructs.Construct) *ApiSchemaConfig {
	if err := i.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ApiSchemaConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

