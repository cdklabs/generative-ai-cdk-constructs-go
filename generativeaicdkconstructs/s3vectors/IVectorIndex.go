package s3vectors

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/s3vectors/internal"
)

// Interface for S3 vector bucket resources.
// Experimental.
type IVectorIndex interface {
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// The timestamp when the vector index was created, in ISO 8601 format.
	// Experimental.
	CreationTime() *string
	// Optional KMS encryption key associated with this vector index.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The ARN of the vector index.
	// Experimental.
	VectorIndexArn() *string
	// The name of the vector index.
	// Experimental.
	VectorIndexName() *string
}

// The jsii proxy for IVectorIndex
type jsiiProxy_IVectorIndex struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVectorIndex) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVectorIndex) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorIndex) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorIndex) VectorIndexArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorIndex) VectorIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexName",
		&returns,
	)
	return returns
}

