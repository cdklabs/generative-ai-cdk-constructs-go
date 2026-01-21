package kendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/kendra/internal"
)

// Represents a Kendra Index, either created with CDK or imported.
// Experimental.
type IKendraGenAiIndex interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the index.
	// Experimental.
	IndexArn() *string
	// The identifier of the index.
	// Experimental.
	IndexId() *string
	// An IAM role that gives Amazon Kendra permissions to access your Amazon CloudWatch logs and metrics.
	//
	// This is also the
	// role used when you use the BatchPutDocument operation to index
	// documents from an Amazon S3 bucket.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IKendraGenAiIndex
type jsiiProxy_IKendraGenAiIndex struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IKendraGenAiIndex) IndexArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKendraGenAiIndex) IndexId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKendraGenAiIndex) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

