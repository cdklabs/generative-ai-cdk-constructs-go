package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Deploys a Bedrock Knowledge Base and configures a backend by OpenSearch Serverless, Pinecone, Redis Enterprise Cloud or Amazon Aurora PostgreSQL.
// Experimental.
type KnowledgeBase interface {
	constructs.Construct
	// The description knowledge base.
	// Experimental.
	Description() *string
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
	// The role the Knowledge Base uses to access the vector store and data source.
	// Experimental.
	Role() awsiam.Role
	// The vector store for the knowledge base.
	// Experimental.
	VectorStore() interface{}
	// Associate knowledge base with an agent.
	// Experimental.
	AssociateToAgent(agent Agent)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KnowledgeBase
type jsiiProxy_KnowledgeBase struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_KnowledgeBase) Role() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"role",
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

