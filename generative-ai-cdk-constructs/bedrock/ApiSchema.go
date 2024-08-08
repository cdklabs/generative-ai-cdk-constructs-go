package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Bedrock Agents Action Group API Schema definition.
// Experimental.
type ApiSchema interface {
	// Called when the action group is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(scope constructs.Construct) *ApiSchemaConfig
}

// The jsii proxy struct for ApiSchema
type jsiiProxy_ApiSchema struct {
	_ byte // padding
}

// Experimental.
func NewApiSchema_Override(a ApiSchema) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		nil, // no parameters
		a,
	)
}

// Loads the API Schema from a local disk path.
//
// Returns: `InlineApiSchema` with the contents of `path`.
// Experimental.
func ApiSchema_FromAsset(path *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromAssetParameters(path); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
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
func ApiSchema_FromBucket(bucket awss3.IBucket, key *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
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

func (a *jsiiProxy_ApiSchema) Bind(scope constructs.Construct) *ApiSchemaConfig {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *ApiSchemaConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

