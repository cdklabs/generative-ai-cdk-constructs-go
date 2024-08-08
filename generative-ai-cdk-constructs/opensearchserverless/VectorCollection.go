package opensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/opensearchserverless/internal"
)

// Deploys an OpenSearch Serverless Collection to be used as a vector store.
//
// It includes all policies.
// Experimental.
type VectorCollection interface {
	constructs.Construct
	// An IAM policy that allows API access to the collection.
	// Experimental.
	AossPolicy() awsiam.ManagedPolicy
	// Experimental.
	SetAossPolicy(val awsiam.ManagedPolicy)
	// The ARN of the collection.
	// Experimental.
	CollectionArn() *string
	// Experimental.
	SetCollectionArn(val *string)
	// The ID of the collection.
	// Experimental.
	CollectionId() *string
	// Experimental.
	SetCollectionId(val *string)
	// The name of the collection.
	// Experimental.
	CollectionName() *string
	// Experimental.
	SetCollectionName(val *string)
	// An OpenSearch Access Policy that allows access to the index.
	// Experimental.
	DataAccessPolicy() awsopensearchserverless.CfnAccessPolicy
	// Experimental.
	SetDataAccessPolicy(val awsopensearchserverless.CfnAccessPolicy)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Indicates whether to use standby replicas for the collection.
	// Experimental.
	StandbyReplicas() VectorCollectionStandbyReplicas
	// Experimental.
	SetStandbyReplicas(val VectorCollectionStandbyReplicas)
	// Grants the specified role access to data in the collection.
	// Experimental.
	GrantDataAccess(grantee awsiam.IRole)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VectorCollection
type jsiiProxy_VectorCollection struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_VectorCollection) AossPolicy() awsiam.ManagedPolicy {
	var returns awsiam.ManagedPolicy
	_jsii_.Get(
		j,
		"aossPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorCollection) CollectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorCollection) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorCollection) CollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorCollection) DataAccessPolicy() awsopensearchserverless.CfnAccessPolicy {
	var returns awsopensearchserverless.CfnAccessPolicy
	_jsii_.Get(
		j,
		"dataAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorCollection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VectorCollection) StandbyReplicas() VectorCollectionStandbyReplicas {
	var returns VectorCollectionStandbyReplicas
	_jsii_.Get(
		j,
		"standbyReplicas",
		&returns,
	)
	return returns
}


// Experimental.
func NewVectorCollection(scope constructs.Construct, id *string, props *VectorCollectionProps) VectorCollection {
	_init_.Initialize()

	if err := validateNewVectorCollectionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_VectorCollection{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewVectorCollection_Override(v VectorCollection, scope constructs.Construct, id *string, props *VectorCollectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollection",
		[]interface{}{scope, id, props},
		v,
	)
}

func (j *jsiiProxy_VectorCollection)SetAossPolicy(val awsiam.ManagedPolicy) {
	if err := j.validateSetAossPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aossPolicy",
		val,
	)
}

func (j *jsiiProxy_VectorCollection)SetCollectionArn(val *string) {
	if err := j.validateSetCollectionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionArn",
		val,
	)
}

func (j *jsiiProxy_VectorCollection)SetCollectionId(val *string) {
	if err := j.validateSetCollectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionId",
		val,
	)
}

func (j *jsiiProxy_VectorCollection)SetCollectionName(val *string) {
	if err := j.validateSetCollectionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionName",
		val,
	)
}

func (j *jsiiProxy_VectorCollection)SetDataAccessPolicy(val awsopensearchserverless.CfnAccessPolicy) {
	if err := j.validateSetDataAccessPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessPolicy",
		val,
	)
}

func (j *jsiiProxy_VectorCollection)SetStandbyReplicas(val VectorCollectionStandbyReplicas) {
	if err := j.validateSetStandbyReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"standbyReplicas",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func VectorCollection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVectorCollection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VectorCollection) GrantDataAccess(grantee awsiam.IRole) {
	if err := v.validateGrantDataAccessParameters(grantee); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"grantDataAccess",
		[]interface{}{grantee},
	)
}

func (v *jsiiProxy_VectorCollection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

