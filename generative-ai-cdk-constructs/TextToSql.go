package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type TextToSql interface {
	BaseClass
	// Returns the instance of s3.IBucket used by the construct.
	// Experimental.
	ConfigAssetBucket() awss3.IBucket
	// construct usage metric , added in template description.
	// Experimental.
	ConstructUsageMetric() *string
	// Returns the instance of ec2.ISecurityGroup used by the construct.
	// Experimental.
	DbSecurityGroup() awsec2.SecurityGroup
	// enable disable xray tracing.
	// Default: - True.
	//
	// Experimental.
	Enablexray() *bool
	// Experimental.
	SetEnablexray(val *bool)
	// Returns the instance of EventBus used by the construct.
	// Experimental.
	EventBus() awsevents.IEventBus
	// Returns the instance of EventBus used by the construct.
	// Experimental.
	EventsRule() awsevents.Rule
	// Returns the instance of feedback Queue  used by the construct.
	// Experimental.
	FeedbackQueue() awssqs.Queue
	// Default  log config for all constructs.
	// Experimental.
	FieldLogLevel() awsappsync.FieldLogLevel
	// Experimental.
	SetFieldLogLevel(val awsappsync.FieldLogLevel)
	// Returns the instance of ec2.ISecurityGroup used by the construct.
	// Experimental.
	LambdaSecurityGroup() awsec2.SecurityGroup
	// enable disable lambda tracing.
	// Default: - Active.
	//
	// Experimental.
	LambdaTracing() awslambda.Tracing
	// Experimental.
	SetLambdaTracing(val awslambda.Tracing)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns the instance of output Queue used by the construct.
	// Experimental.
	OutputQueue() awssqs.Queue
	// Default  log retention config for all constructs.
	// Experimental.
	Retention() awslogs.RetentionDays
	// Experimental.
	SetRetention(val awslogs.RetentionDays)
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
	// Returns the Instance of stepfunction created by the construct.
	// Experimental.
	StepFunction() awsstepfunctions.StateMachine
	// Returns the instance of subnet group used by the construct.
	// Experimental.
	SubnetGroup() awsrds.SubnetGroup
	// Returns the instance of ec2.IVpc used by the construct.
	// Experimental.
	Vpc() awsec2.IVpc
	// Experimental.
	AddObservabilityToConstruct(props *BaseClassProps)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction)
	// Experimental.
	UpdateEnvSuffix(props *BaseClassProps)
}

// The jsii proxy struct for TextToSql
type jsiiProxy_TextToSql struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_TextToSql) ConfigAssetBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"configAssetBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) DbSecurityGroup() awsec2.SecurityGroup {
	var returns awsec2.SecurityGroup
	_jsii_.Get(
		j,
		"dbSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) EventBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) EventsRule() awsevents.Rule {
	var returns awsevents.Rule
	_jsii_.Get(
		j,
		"eventsRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) FeedbackQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"feedbackQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) LambdaSecurityGroup() awsec2.SecurityGroup {
	var returns awsec2.SecurityGroup
	_jsii_.Get(
		j,
		"lambdaSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) OutputQueue() awssqs.Queue {
	var returns awssqs.Queue
	_jsii_.Get(
		j,
		"outputQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) StepFunction() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stepFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) SubnetGroup() awsrds.SubnetGroup {
	var returns awsrds.SubnetGroup
	_jsii_.Get(
		j,
		"subnetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TextToSql) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Constructs a new instance of the TextToSql class.
// Experimental.
func NewTextToSql(scope constructs.Construct, id *string, props *TextToSqlProps) TextToSql {
	_init_.Initialize()

	if err := validateNewTextToSqlParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_TextToSql{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.TextToSql",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the TextToSql class.
// Experimental.
func NewTextToSql_Override(t TextToSql, scope constructs.Construct, id *string, props *TextToSqlProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.TextToSql",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TextToSql)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_TextToSql)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_TextToSql)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_TextToSql)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_TextToSql)SetStage(val *string) {
	if err := j.validateSetStageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stage",
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
func TextToSql_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTextToSql_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.TextToSql",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TextToSql_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.TextToSql",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func TextToSql_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateTextToSql_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.TextToSql",
		"usageMetricMap",
		val,
	)
}

func (t *jsiiProxy_TextToSql) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := t.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (t *jsiiProxy_TextToSql) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TextToSql) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := t.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (t *jsiiProxy_TextToSql) UpdateEnvSuffix(props *BaseClassProps) {
	if err := t.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

