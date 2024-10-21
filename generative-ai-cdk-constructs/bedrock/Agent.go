package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Deploy a Bedrock Agent.
// Experimental.
type Agent interface {
	constructs.Construct
	// A list of action groups associated with the agent.
	// Experimental.
	ActionGroups() *[]*awsbedrock.CfnAgent_AgentActionGroupProperty
	// Experimental.
	SetActionGroups(val *[]*awsbedrock.CfnAgent_AgentActionGroupProperty)
	// The ARN of the agent.
	// Experimental.
	AgentArn() *string
	// The unique identifier of the agent.
	// Experimental.
	AgentId() *string
	// Instance of Agent.
	// Experimental.
	AgentInstance() awsbedrock.CfnAgent
	// The version for the agent.
	// Experimental.
	Agentversion() *string
	// The ARN of the agent alias.
	// Experimental.
	AliasArn() *string
	// The unique identifier of the agent alias.
	// Experimental.
	AliasId() *string
	// The name for the agent alias.
	// Experimental.
	AliasName() *string
	// A list of KnowledgeBases associated with the agent.
	// Default: - No knowledge base is used.
	//
	// Experimental.
	KnowledgeBases() *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty
	// Experimental.
	SetKnowledgeBases(val *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty)
	// The name of the agent.
	// Experimental.
	Name() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The IAM role for the agent.
	// Experimental.
	Role() awsiam.Role
	// Add action group to the agent.
	// Experimental.
	AddActionGroup(actionGroup AgentActionGroup)
	// Add action groups to the agent.
	// Experimental.
	AddActionGroups(actionGroups *[]AgentActionGroup)
	// Add an alias to the agent.
	// Experimental.
	AddAlias(props *AddAgentAliasProps) AgentAlias
	// Add guardrail to the agent.
	// Experimental.
	AddGuardrail(guardrail IGuardrail)
	// Add knowledge base to the agent.
	// Experimental.
	AddKnowledgeBase(knowledgeBase KnowledgeBase)
	// Add knowledge bases to the agent.
	// Experimental.
	AddKnowledgeBases(knowledgeBases *[]KnowledgeBase)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Agent
type jsiiProxy_Agent struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_Agent) ActionGroups() *[]*awsbedrock.CfnAgent_AgentActionGroupProperty {
	var returns *[]*awsbedrock.CfnAgent_AgentActionGroupProperty
	_jsii_.Get(
		j,
		"actionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AgentInstance() awsbedrock.CfnAgent {
	var returns awsbedrock.CfnAgent
	_jsii_.Get(
		j,
		"agentInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Agentversion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentversion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AliasArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AliasId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) AliasName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) KnowledgeBases() *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty {
	var returns *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty
	_jsii_.Get(
		j,
		"knowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Agent) Role() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgent(scope constructs.Construct, id *string, props *AgentProps) Agent {
	_init_.Initialize()

	if err := validateNewAgentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Agent{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Agent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAgent_Override(a Agent, scope constructs.Construct, id *string, props *AgentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Agent",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_Agent)SetActionGroups(val *[]*awsbedrock.CfnAgent_AgentActionGroupProperty) {
	if err := j.validateSetActionGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroups",
		val,
	)
}

func (j *jsiiProxy_Agent)SetKnowledgeBases(val *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty) {
	if err := j.validateSetKnowledgeBasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knowledgeBases",
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
func Agent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Agent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) AddActionGroup(actionGroup AgentActionGroup) {
	if err := a.validateAddActionGroupParameters(actionGroup); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addActionGroup",
		[]interface{}{actionGroup},
	)
}

func (a *jsiiProxy_Agent) AddActionGroups(actionGroups *[]AgentActionGroup) {
	if err := a.validateAddActionGroupsParameters(actionGroups); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addActionGroups",
		[]interface{}{actionGroups},
	)
}

func (a *jsiiProxy_Agent) AddAlias(props *AddAgentAliasProps) AgentAlias {
	if err := a.validateAddAliasParameters(props); err != nil {
		panic(err)
	}
	var returns AgentAlias

	_jsii_.Invoke(
		a,
		"addAlias",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Agent) AddGuardrail(guardrail IGuardrail) {
	if err := a.validateAddGuardrailParameters(guardrail); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addGuardrail",
		[]interface{}{guardrail},
	)
}

func (a *jsiiProxy_Agent) AddKnowledgeBase(knowledgeBase KnowledgeBase) {
	if err := a.validateAddKnowledgeBaseParameters(knowledgeBase); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addKnowledgeBase",
		[]interface{}{knowledgeBase},
	)
}

func (a *jsiiProxy_Agent) AddKnowledgeBases(knowledgeBases *[]KnowledgeBase) {
	if err := a.validateAddKnowledgeBasesParameters(knowledgeBases); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addKnowledgeBases",
		[]interface{}{knowledgeBases},
	)
}

func (a *jsiiProxy_Agent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

