package auroradsql

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/auroradsql/internal"
)

// Interface for Aurora DSQL cluster resources.
// Experimental.
type ICluster interface {
	awscdk.IResource
	// Grants IAM actions to the IAM Principal.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// grants connection authorization for a custom database role to the IAM Principal.
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// Grants connection authorization for the admin role to the IAM Principal.
	// Experimental.
	GrantConnectAdmin(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the cluster.
	// Experimental.
	ClusterArn() *string
	// The id of the cluster.
	// Experimental.
	ClusterId() *string
	// The timestamp when the cluster was created, in ISO 8601 format.
	// Experimental.
	CreationTime() *string
	// Optional KMS encryption key associated with this bucket.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The status of the cluster.
	// Experimental.
	Status() *string
	// Experimental.
	VpcEndpointServiceName() *string
}

// The jsii proxy for ICluster
type jsiiProxy_ICluster struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ICluster) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_ICluster) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ICluster) GrantConnectAdmin(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConnectAdminParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnectAdmin",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ICluster) ClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICluster) VpcEndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointServiceName",
		&returns,
	)
	return returns
}

