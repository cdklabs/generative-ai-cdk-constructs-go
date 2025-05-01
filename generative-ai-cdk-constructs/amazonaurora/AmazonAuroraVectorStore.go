package amazonaurora

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/amazonaurora/internal"
)

// Experimental.
type AmazonAuroraVectorStore interface {
	constructs.Construct
	// The Secret ARN of your Amazon Aurora DB cluster.
	// Experimental.
	CredentialsSecretArn() *string
	// The name of the database for the Aurora Vector Store.
	// Experimental.
	DatabaseName() *string
	// The embeddings model dimension used for the Aurora Vector Store.
	//
	// The vector dimensions of the model must match the dimensions
	// used in the KnowledgeBase construct.
	// Experimental.
	EmbeddingsModelVectorDimension() *float64
	// The field name for the metadata column in the Aurora Vector Store.
	// Experimental.
	MetadataField() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The primary key field for the Aurora Vector Store table.
	// Experimental.
	PrimaryKeyField() *string
	// The ARN of your Amazon Aurora DB cluster.
	// Experimental.
	ResourceArn() *string
	// The schema name for the Aurora Vector Store.
	// Experimental.
	SchemaName() *string
	// The name of the table for the Aurora Vector Store.
	// Experimental.
	TableName() *string
	// The field name for the text column in the Aurora Vector Store.
	// Experimental.
	TextField() *string
	// The field name for the vector column in the Aurora Vector Store.
	// Experimental.
	VectorField() *string
	// The VPC of your Amazon Aurora DB cluster.
	// Experimental.
	Vpc() awsec2.IVpc
	// Experimental.
	AddIngressRuleToAuroraSecurityGroup(lambdaSecurityGroup awsec2.ISecurityGroup, auroraSecurityGroup awsec2.ISecurityGroup)
	// Experimental.
	CreateAuroraPgCRPolicy(clusterIdentifier *string) awsiam.ManagedPolicy
	// Experimental.
	CreateLambdaSecurityGroup(vpc awsec2.IVpc) awsec2.SecurityGroup
	// Experimental.
	GenerateResourceArn(clusterIdentifier *string) *string
	// Experimental.
	SetupCustomResource(databaseClusterResources *DatabaseClusterResources, lambdaSecurityGroup awsec2.SecurityGroup, auroraPgCRPolicy awsiam.ManagedPolicy) awscdk.CustomResource
	// Experimental.
	SetupDatabaseClusterResources(vpc awsec2.IVpc, secret awssecretsmanager.ISecret, clusterIdentifier *string, auroraSecurityGroup awsec2.ISecurityGroup) *DatabaseClusterResources
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AmazonAuroraVectorStore
type jsiiProxy_AmazonAuroraVectorStore struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AmazonAuroraVectorStore) CredentialsSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) EmbeddingsModelVectorDimension() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"embeddingsModelVectorDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) MetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) PrimaryKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) SchemaName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) TextField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) VectorField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Experimental.
func NewAmazonAuroraVectorStore(scope constructs.Construct, id *string, props *AmazonAuroraVectorStoreProps) AmazonAuroraVectorStore {
	_init_.Initialize()

	if err := validateNewAmazonAuroraVectorStoreParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonAuroraVectorStore{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAmazonAuroraVectorStore_Override(a AmazonAuroraVectorStore, scope constructs.Construct, id *string, props *AmazonAuroraVectorStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		[]interface{}{scope, id, props},
		a,
	)
}

// Creates an instance of AmazonAuroraVectorStore using existing Aurora Vector Store properties.
//
// You need to provide your existing Aurora Vector Store properties
// such as `databaseName`, `clusterIdentifier`, `vpc` where database is deployed,
// `secret` containing username and password for authentication to database,
// and `auroraSecurityGroup` with the ecurity group that was used for the database.
//
// Returns: An instance of AmazonAuroraVectorStore.
// Experimental.
func AmazonAuroraVectorStore_FromExistingAuroraVectorStore(scope constructs.Construct, id *string, props *ExistingAmazonAuroraVectorStoreProps) ExistingAmazonAuroraVectorStore {
	_init_.Initialize()

	if err := validateAmazonAuroraVectorStore_FromExistingAuroraVectorStoreParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns ExistingAmazonAuroraVectorStore

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		"fromExistingAuroraVectorStore",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
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
func AmazonAuroraVectorStore_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAmazonAuroraVectorStore_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonAuroraVectorStore) AddIngressRuleToAuroraSecurityGroup(lambdaSecurityGroup awsec2.ISecurityGroup, auroraSecurityGroup awsec2.ISecurityGroup) {
	if err := a.validateAddIngressRuleToAuroraSecurityGroupParameters(lambdaSecurityGroup, auroraSecurityGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addIngressRuleToAuroraSecurityGroup",
		[]interface{}{lambdaSecurityGroup, auroraSecurityGroup},
	)
}

func (a *jsiiProxy_AmazonAuroraVectorStore) CreateAuroraPgCRPolicy(clusterIdentifier *string) awsiam.ManagedPolicy {
	if err := a.validateCreateAuroraPgCRPolicyParameters(clusterIdentifier); err != nil {
		panic(err)
	}
	var returns awsiam.ManagedPolicy

	_jsii_.Invoke(
		a,
		"createAuroraPgCRPolicy",
		[]interface{}{clusterIdentifier},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonAuroraVectorStore) CreateLambdaSecurityGroup(vpc awsec2.IVpc) awsec2.SecurityGroup {
	if err := a.validateCreateLambdaSecurityGroupParameters(vpc); err != nil {
		panic(err)
	}
	var returns awsec2.SecurityGroup

	_jsii_.Invoke(
		a,
		"createLambdaSecurityGroup",
		[]interface{}{vpc},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonAuroraVectorStore) GenerateResourceArn(clusterIdentifier *string) *string {
	if err := a.validateGenerateResourceArnParameters(clusterIdentifier); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"generateResourceArn",
		[]interface{}{clusterIdentifier},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonAuroraVectorStore) SetupCustomResource(databaseClusterResources *DatabaseClusterResources, lambdaSecurityGroup awsec2.SecurityGroup, auroraPgCRPolicy awsiam.ManagedPolicy) awscdk.CustomResource {
	if err := a.validateSetupCustomResourceParameters(databaseClusterResources, lambdaSecurityGroup, auroraPgCRPolicy); err != nil {
		panic(err)
	}
	var returns awscdk.CustomResource

	_jsii_.Invoke(
		a,
		"setupCustomResource",
		[]interface{}{databaseClusterResources, lambdaSecurityGroup, auroraPgCRPolicy},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonAuroraVectorStore) SetupDatabaseClusterResources(vpc awsec2.IVpc, secret awssecretsmanager.ISecret, clusterIdentifier *string, auroraSecurityGroup awsec2.ISecurityGroup) *DatabaseClusterResources {
	if err := a.validateSetupDatabaseClusterResourcesParameters(vpc, secret, clusterIdentifier, auroraSecurityGroup); err != nil {
		panic(err)
	}
	var returns *DatabaseClusterResources

	_jsii_.Invoke(
		a,
		"setupDatabaseClusterResources",
		[]interface{}{vpc, secret, clusterIdentifier, auroraSecurityGroup},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AmazonAuroraVectorStore) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

