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
type IVectorBucket interface {
	awscdk.IResource
	// Adds a statement to the resource policy for a principal (i.e. account/role/service) to perform actions on this bucket and/or its contents. Use `bucketArn` to obtain ARNs for this bucket.
	//
	// Note that the policy statement may or may not be added to the policy.
	// For example, when an `IBucket` is created from an existing bucket,
	// it's not possible to tell whether the bucket already has a policy
	// attached, let alone to re-use that policy to add more statements to it.
	// So it's safest to do nothing in these cases.
	//
	// Returns: metadata about the execution of this method. If the policy
	// was not added, the value of `statementAdded` will be `false`. You
	// should always check this value to make sure that the operation was
	// actually carried out. Otherwise, synthesis and deploy will terminate
	// silently, which may be confusing.
	// Experimental.
	AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Grants IAM actions to the IAM Principal to delete the vector bucket and indexes.
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantDelete(grantee awsiam.IGrantable, indexIds interface{}) awsiam.Grant
	// Grants IAM actions to the IAM Principal.
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantRead(grantee awsiam.IGrantable, indexIds interface{}) awsiam.Grant
	// Grants IAM actions to the IAM Principal.
	//
	// Returns: An IAM Grant object representing the granted permissions.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable, indexIds interface{}) awsiam.Grant
	// The timestamp when the vector bucket was created, in ISO 8601 format.
	// Experimental.
	CreationTime() *string
	// Optional KMS encryption key associated with this vector bucket.
	// Experimental.
	EncryptionKey() awskms.IKey
	// The resource policy associated with this bucket.
	//
	// If `autoCreatePolicy` is true, a `BucketPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	// Experimental.
	Policy() VectorBucketPolicy
	// Experimental.
	SetPolicy(p VectorBucketPolicy)
	// The ARN of the vector bucket.
	// Experimental.
	VectorBucketArn() *string
	// The name of the vector bucket.
	// Experimental.
	VectorBucketName() *string
}

// The jsii proxy for IVectorBucket
type jsiiProxy_IVectorBucket struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVectorBucket) AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(permission); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{permission},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorBucket) GrantDelete(grantee awsiam.IGrantable, indexIds interface{}) awsiam.Grant {
	if err := i.validateGrantDeleteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantDelete",
		[]interface{}{grantee, indexIds},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorBucket) GrantRead(grantee awsiam.IGrantable, indexIds interface{}) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee, indexIds},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorBucket) GrantWrite(grantee awsiam.IGrantable, indexIds interface{}) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee, indexIds},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVectorBucket) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucket) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucket) Policy() VectorBucketPolicy {
	var returns VectorBucketPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucket)SetPolicy(val VectorBucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_IVectorBucket) VectorBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorBucket) VectorBucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorBucketName",
		&returns,
	)
	return returns
}

