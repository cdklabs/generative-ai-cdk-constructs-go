package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type WebCrawler interface {
	BaseClass
	// construct usage metric , added in template description.
	// Experimental.
	ConstructUsageMetric() *string
	// Returns the instance of S3 bucket used by the construct.
	// Experimental.
	DataBucket() awss3.IBucket
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
	// Returns the instance of JobQueue used by the construct.
	// Experimental.
	JobQueue() awsbatch.IJobQueue
	// Returns the instance of Jobs DynamoDB table.
	// Experimental.
	JobsTable() awsdynamodb.ITable
	// Lambda crawler.
	// Experimental.
	LambdaCrawler() awslambda.IFunction
	// Lambda crawler API schema path.
	// Experimental.
	LambdaCrawlerApiSchemaPath() *string
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
	// Default  log retention config for all constructs.
	// Experimental.
	Retention() awslogs.RetentionDays
	// Experimental.
	SetRetention(val awslogs.RetentionDays)
	// Returns the instance of SNS Topic used by the construct.
	// Experimental.
	SnsTopic() awssns.ITopic
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
	// Returns the instance of Targets DynamoDB table.
	// Experimental.
	TargetsTable() awsdynamodb.ITable
	// Returns the instance of ec2.IVpc used by the construct.
	// Experimental.
	Vpc() awsec2.IVpc
	// Returns the instance of JobDefinition used by the construct.
	// Experimental.
	WebCrawlerJobDefinition() awsbatch.IJobDefinition
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

// The jsii proxy struct for WebCrawler
type jsiiProxy_WebCrawler struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_WebCrawler) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) DataBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"dataBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) JobQueue() awsbatch.IJobQueue {
	var returns awsbatch.IJobQueue
	_jsii_.Get(
		j,
		"jobQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) JobsTable() awsdynamodb.ITable {
	var returns awsdynamodb.ITable
	_jsii_.Get(
		j,
		"jobsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) LambdaCrawler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaCrawler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) LambdaCrawlerApiSchemaPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaCrawlerApiSchemaPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) SnsTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"snsTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) TargetsTable() awsdynamodb.ITable {
	var returns awsdynamodb.ITable
	_jsii_.Get(
		j,
		"targetsTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebCrawler) WebCrawlerJobDefinition() awsbatch.IJobDefinition {
	var returns awsbatch.IJobDefinition
	_jsii_.Get(
		j,
		"webCrawlerJobDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the WebCrawler class.
// Experimental.
func NewWebCrawler(scope constructs.Construct, id *string, props *WebCrawlerProps) WebCrawler {
	_init_.Initialize()

	if err := validateNewWebCrawlerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebCrawler{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawler",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the WebCrawler class.
// Experimental.
func NewWebCrawler_Override(w WebCrawler, scope constructs.Construct, id *string, props *WebCrawlerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawler",
		[]interface{}{scope, id, props},
		w,
	)
}

func (j *jsiiProxy_WebCrawler)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_WebCrawler)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_WebCrawler)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_WebCrawler)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_WebCrawler)SetStage(val *string) {
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
func WebCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWebCrawler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WebCrawler_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawler",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func WebCrawler_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateWebCrawler_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawler",
		"usageMetricMap",
		val,
	)
}

func (w *jsiiProxy_WebCrawler) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := w.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (w *jsiiProxy_WebCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebCrawler) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := w.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (w *jsiiProxy_WebCrawler) UpdateEnvSuffix(props *BaseClassProps) {
	if err := w.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

