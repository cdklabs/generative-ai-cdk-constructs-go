package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Specifies interface for resources created with CDK or imported into CDK.
// Experimental.
type IDataSource interface {
	awscdk.IResource
	// The unique identifier of the data source.
	//
	// Example:
	//   'JHUEVXUZMU'
	//
	// Experimental.
	DataSourceId() *string
}

// The jsii proxy for IDataSource
type jsiiProxy_IDataSource struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IDataSource) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

