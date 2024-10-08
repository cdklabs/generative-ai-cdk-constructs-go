package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Deploys a Bedrock Knowledge Base and configures a backend by OpenSearch Serverless, Pinecone, Redis Enterprise Cloud or Amazon Aurora PostgreSQL.
// Experimental.
type KnowledgeBase interface {
	awscdk.Resource
	IKnowledgeBase
	// The description knowledge base.
	// Experimental.
	Description() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// A narrative instruction of the knowledge base.
	// Experimental.
	Instruction() *string
	// The ARN of the knowledge base.
	// Experimental.
	KnowledgeBaseArn() *string
	// The ID of the knowledge base.
	// Experimental.
	KnowledgeBaseId() *string
	// Instance of knowledge base.
	// Experimental.
	KnowledgeBaseInstance() awsbedrock.CfnKnowledgeBase
	// Specifies whether to use the knowledge base or not when sending an InvokeAgent request.
	// Experimental.
	KnowledgeBaseState() *string
	// The name of the knowledge base.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The role the Knowledge Base uses to access the vector store and data source.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The vector store for the knowledge base.
	// Experimental.
	VectorStore() interface{}
	// Add a Confluence data source to the knowledge base.
	// Experimental.
	AddConfluenceDataSource(props *ConfluenceDataSourceAssociationProps) ConfluenceDataSource
	// Add an S3 data source to the knowledge base.
	// Experimental.
	AddS3DataSource(props *S3DataSourceAssociationProps) S3DataSource
	// Add a Salesforce data source to the knowledge base.
	// Experimental.
	AddSalesforceDataSource(props *SalesforceDataSourceAssociationProps) SalesforceDataSource
	// Add a SharePoint data source to the knowledge base.
	// Experimental.
	AddSharePointDataSource(props *SharePointDataSourceAssociationProps) SharePointDataSource
	// Add a web crawler data source to the knowledge base.
	// Experimental.
	AddWebCrawlerDataSource(props *WebCrawlerDataSourceAssociationProps) WebCrawlerDataSource
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Associate knowledge base with an agent.
	// Experimental.
	AssociateToAgent(agent Agent)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KnowledgeBase
type jsiiProxy_KnowledgeBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IKnowledgeBase
}

func (j *jsiiProxy_KnowledgeBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) KnowledgeBaseInstance() awsbedrock.CfnKnowledgeBase {
	var returns awsbedrock.CfnKnowledgeBase
	_jsii_.Get(
		j,
		"knowledgeBaseInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) KnowledgeBaseState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KnowledgeBase) VectorStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorStore",
		&returns,
	)
	return returns
}


// Experimental.
func NewKnowledgeBase(scope constructs.Construct, id *string, props *KnowledgeBaseProps) KnowledgeBase {
	_init_.Initialize()

	if err := validateNewKnowledgeBaseParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KnowledgeBase{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKnowledgeBase_Override(k KnowledgeBase, scope constructs.Construct, id *string, props *KnowledgeBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		[]interface{}{scope, id, props},
		k,
	)
}

// Experimental.
func KnowledgeBase_FromKnowledgeBaseAttributes(scope constructs.Construct, id *string, attrs *KnowledgeBaseAttributes) IKnowledgeBase {
	_init_.Initialize()

	if err := validateKnowledgeBase_FromKnowledgeBaseAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IKnowledgeBase

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		"fromKnowledgeBaseAttributes",
		[]interface{}{scope, id, attrs},
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
func KnowledgeBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func KnowledgeBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func KnowledgeBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKnowledgeBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) AddConfluenceDataSource(props *ConfluenceDataSourceAssociationProps) ConfluenceDataSource {
	if err := k.validateAddConfluenceDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns ConfluenceDataSource

	_jsii_.Invoke(
		k,
		"addConfluenceDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) AddS3DataSource(props *S3DataSourceAssociationProps) S3DataSource {
	if err := k.validateAddS3DataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns S3DataSource

	_jsii_.Invoke(
		k,
		"addS3DataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) AddSalesforceDataSource(props *SalesforceDataSourceAssociationProps) SalesforceDataSource {
	if err := k.validateAddSalesforceDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns SalesforceDataSource

	_jsii_.Invoke(
		k,
		"addSalesforceDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) AddSharePointDataSource(props *SharePointDataSourceAssociationProps) SharePointDataSource {
	if err := k.validateAddSharePointDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns SharePointDataSource

	_jsii_.Invoke(
		k,
		"addSharePointDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) AddWebCrawlerDataSource(props *WebCrawlerDataSourceAssociationProps) WebCrawlerDataSource {
	if err := k.validateAddWebCrawlerDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns WebCrawlerDataSource

	_jsii_.Invoke(
		k,
		"addWebCrawlerDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := k.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (k *jsiiProxy_KnowledgeBase) AssociateToAgent(agent Agent) {
	if err := k.validateAssociateToAgentParameters(agent); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"associateToAgent",
		[]interface{}{agent},
	)
}

func (k *jsiiProxy_KnowledgeBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := k.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := k.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KnowledgeBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

