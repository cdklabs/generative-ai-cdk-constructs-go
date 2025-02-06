package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/kendra"
)

// ****************************************************************************                             ABSTRACT CLASS ***************************************************************************.
// Experimental.
type KendraKnowledgeBaseBase interface {
	KnowledgeBaseBase
	// The description of the knowledge base.
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
	//
	// A Bedrock Agent can use this instruction to determine if it should
	// query this Knowledge Base.
	// Experimental.
	Instruction() *string
	// Experimental.
	KendraIndex() kendra.IKendraGenAiIndex
	// The ARN of the knowledge base.
	// Experimental.
	KnowledgeBaseArn() *string
	// The ID of the knowledge base.
	// Experimental.
	KnowledgeBaseId() *string
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
	// The role associated with the knowledge base.
	// Experimental.
	Role() awsiam.IRole
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// The type of knowledge base.
	// Experimental.
	Type() KnowledgeBaseType
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
	// Grant the given principal identity permissions to perform actions on this knowledge base.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to query the knowledge base.
	//
	// This contains:
	// - Retrieve
	// - RetrieveAndGenerate.
	// Experimental.
	GrantQuery(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to retrieve content from the knowledge base.
	// Experimental.
	GrantRetrieve(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to retrieve content from the knowledge base.
	// Experimental.
	GrantRetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraKnowledgeBaseBase
type jsiiProxy_KendraKnowledgeBaseBase struct {
	jsiiProxy_KnowledgeBaseBase
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) KendraIndex() kendra.IKendraGenAiIndex {
	var returns kendra.IKendraGenAiIndex
	_jsii_.Get(
		j,
		"kendraIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraKnowledgeBaseBase) Type() KnowledgeBaseType {
	var returns KnowledgeBaseType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewKendraKnowledgeBaseBase_Override(k KendraKnowledgeBaseBase, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseBase",
		[]interface{}{scope, id},
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
func KendraKnowledgeBaseBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKendraKnowledgeBaseBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func KendraKnowledgeBaseBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKendraKnowledgeBaseBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func KendraKnowledgeBaseBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateKendraKnowledgeBaseBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := k.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (k *jsiiProxy_KendraKnowledgeBaseBase) GetResourceNameAttribute(nameAttr *string) *string {
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

func (k *jsiiProxy_KendraKnowledgeBaseBase) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := k.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) GrantQuery(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateGrantQueryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"grantQuery",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) GrantRetrieve(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateGrantRetrieveParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"grantRetrieve",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) GrantRetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant {
	if err := k.validateGrantRetrieveAndGenerateParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		k,
		"grantRetrieveAndGenerate",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraKnowledgeBaseBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

