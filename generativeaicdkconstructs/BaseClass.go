package generativeaicdkconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/internal"
)

// Experimental.
type BaseClass interface {
	constructs.Construct
	// construct usage metric , added in template description.
	// Experimental.
	ConstructUsageMetric() *string
	// enable disable xray tracing.
	// Default: true.
	//
	// Experimental.
	Enablexray() *bool
	// Experimental.
	SetEnablexray(val *bool)
	// Default  log config for all constructs.
	// Experimental.
	FieldLogLevel() awsappsync.FieldLogLevel
	// Experimental.
	SetFieldLogLevel(val awsappsync.FieldLogLevel)
	// enable disable lambda tracing.
	// Default: Active.
	//
	// Experimental.
	LambdaTracing() awslambda.Tracing
	// Experimental.
	SetLambdaTracing(val awslambda.Tracing)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Default  log retention config for all constructs.
	// Experimental.
	Retention() awslogs.RetentionDays
	// Experimental.
	SetRetention(val awslogs.RetentionDays)
	// Value will be appended to resources name.
	// Default: _dev.
	//
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
	// Experimental.
	AddObservabilityToConstruct(props *BaseClassProps)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// update template description with construct usage metric and add AWS_SDK_UA_APP_ID to user agent on aws sdk.
	// Experimental.
	UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction)
	// Experimental.
	UpdateEnvSuffix(props *BaseClassProps)
}

// The jsii proxy struct for BaseClass
type jsiiProxy_BaseClass struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BaseClass) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseClass) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseClass) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseClass) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseClass) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseClass) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseClass) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseClass(scope constructs.Construct, id *string) BaseClass {
	_init_.Initialize()

	if err := validateNewBaseClassParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_BaseClass{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.BaseClass",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewBaseClass_Override(b BaseClass, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.BaseClass",
		[]interface{}{scope, id},
		b,
	)
}

func (j *jsiiProxy_BaseClass)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_BaseClass)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_BaseClass)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_BaseClass)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_BaseClass)SetStage(val *string) {
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
func BaseClass_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBaseClass_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.BaseClass",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BaseClass_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.BaseClass",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func BaseClass_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateBaseClass_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.BaseClass",
		"usageMetricMap",
		val,
	)
}

func (b *jsiiProxy_BaseClass) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := b.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (b *jsiiProxy_BaseClass) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseClass) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := b.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (b *jsiiProxy_BaseClass) UpdateEnvSuffix(props *BaseClassProps) {
	if err := b.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

