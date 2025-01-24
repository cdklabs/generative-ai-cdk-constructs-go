package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to create an Agent Alias with CDK.
// Experimental.
type AgentAlias interface {
	AgentAliasBase
	// The underlying agent for this alias.
	// Experimental.
	Agent() IAgent
	// The ARN of the agent alias.
	// Experimental.
	AliasArn() *string
	// The unique identifier of the agent alias.
	// Experimental.
	AliasId() *string
	// Experimental.
	AliasName() *string
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
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
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
	// Grant the given principal identity permissions to perform actions on this agent alias.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to invoke the agent alias.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Define an EventBridge rule that triggers when something happens to this agent alias.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	// Experimental.
	OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentAlias
type jsiiProxy_AgentAlias struct {
	jsiiProxy_AgentAliasBase
}

func (j *jsiiProxy_AgentAlias) Agent() IAgent {
	var returns IAgent
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentAlias) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgentAlias(scope constructs.Construct, id *string, props *AgentAliasProps) AgentAlias {
	_init_.Initialize()

	if err := validateNewAgentAliasParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentAlias{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgentAlias_Override(a AgentAlias, scope constructs.Construct, id *string, props *AgentAliasProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		[]interface{}{scope, id, props},
		a,
	)
}

// Brings an Agent Alias from an existing one created outside of CDK.
// Experimental.
func AgentAlias_FromAttributes(scope constructs.Construct, id *string, attrs *AgentAliasAttributes) IAgentAlias {
	_init_.Initialize()

	if err := validateAgentAlias_FromAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IAgentAlias

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		"fromAttributes",
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
func AgentAlias_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgentAlias_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func AgentAlias_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAgentAlias_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func AgentAlias_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAgentAlias_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AgentAlias) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := a.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := a.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		a,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) OnCloudTrailEvent(id *string, options *awsevents.OnEventOptions) awsevents.Rule {
	if err := a.validateOnCloudTrailEventParameters(id, options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		a,
		"onCloudTrailEvent",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentAlias) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

