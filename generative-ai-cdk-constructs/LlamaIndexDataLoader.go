package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type LlamaIndexDataLoader interface {
	BaseClass
	// construct usage metric , added in template description.
	// Experimental.
	ConstructUsageMetric() *string
	// enable disable xray tracing.
	// Default: - True.
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
	// Default: - Active.
	//
	// Experimental.
	LambdaTracing() awslambda.Tracing
	// Experimental.
	SetLambdaTracing(val awslambda.Tracing)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutputBucket() awss3.Bucket
	// Experimental.
	QueueProcessingFargateService() awsecspatterns.QueueProcessingFargateService
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

// The jsii proxy struct for LlamaIndexDataLoader
type jsiiProxy_LlamaIndexDataLoader struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_LlamaIndexDataLoader) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) OutputBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"outputBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) QueueProcessingFargateService() awsecspatterns.QueueProcessingFargateService {
	var returns awsecspatterns.QueueProcessingFargateService
	_jsii_.Get(
		j,
		"queueProcessingFargateService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LlamaIndexDataLoader) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}


// Experimental.
func NewLlamaIndexDataLoader(scope constructs.Construct, id *string, props *LlamaIndexDataLoaderProps) LlamaIndexDataLoader {
	_init_.Initialize()

	if err := validateNewLlamaIndexDataLoaderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LlamaIndexDataLoader{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.LlamaIndexDataLoader",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLlamaIndexDataLoader_Override(l LlamaIndexDataLoader, scope constructs.Construct, id *string, props *LlamaIndexDataLoaderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.LlamaIndexDataLoader",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LlamaIndexDataLoader)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_LlamaIndexDataLoader)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_LlamaIndexDataLoader)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_LlamaIndexDataLoader)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_LlamaIndexDataLoader)SetStage(val *string) {
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
func LlamaIndexDataLoader_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLlamaIndexDataLoader_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.LlamaIndexDataLoader",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LlamaIndexDataLoader_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.LlamaIndexDataLoader",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func LlamaIndexDataLoader_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateLlamaIndexDataLoader_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.LlamaIndexDataLoader",
		"usageMetricMap",
		val,
	)
}

func (l *jsiiProxy_LlamaIndexDataLoader) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := l.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (l *jsiiProxy_LlamaIndexDataLoader) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LlamaIndexDataLoader) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := l.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (l *jsiiProxy_LlamaIndexDataLoader) UpdateEnvSuffix(props *BaseClassProps) {
	if err := l.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

