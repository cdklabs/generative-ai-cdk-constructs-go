package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// API Schema in an S3 object.
// Experimental.
type S3ApiSchema interface {
	ApiSchema
	// Called when the action group is initialized to allow this object to bind to the stack, add resources and have fun.
	// Experimental.
	Bind(_scope constructs.Construct) *ApiSchemaConfig
}

// The jsii proxy struct for S3ApiSchema
type jsiiProxy_S3ApiSchema struct {
	jsiiProxy_ApiSchema
}

// Experimental.
func NewS3ApiSchema(bucket awss3.IBucket, key *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateNewS3ApiSchemaParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ApiSchema{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		[]interface{}{bucket, key},
		&j,
	)

	return &j
}

// Experimental.
func NewS3ApiSchema_Override(s S3ApiSchema, bucket awss3.IBucket, key *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		[]interface{}{bucket, key},
		s,
	)
}

// Loads the API Schema from a local disk path.
//
// Returns: `InlineApiSchema` with the contents of `path`.
// Experimental.
func S3ApiSchema_FromAsset(path *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromAssetParameters(path); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
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
func S3ApiSchema_FromBucket(bucket awss3.IBucket, key *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
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

func (s *jsiiProxy_S3ApiSchema) Bind(_scope constructs.Construct) *ApiSchemaConfig {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ApiSchemaConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

