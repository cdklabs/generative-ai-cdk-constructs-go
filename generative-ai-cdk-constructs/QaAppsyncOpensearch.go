package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// The QaAppsyncOpensearch class.
// Experimental.
type QaAppsyncOpensearch interface {
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
	// Returns an instance of appsync.IGraphqlApi created by the construct.
	// Experimental.
	GraphqlApi() awsappsync.IGraphqlApi
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
	// Returns the instance of events.IEventBus used by the construct.
	// Experimental.
	QaBus() awsevents.IEventBus
	// Returns an instance of appsync.IGraphqlApi created by the construct.
	// Experimental.
	QaLambdaFunction() awslambda.DockerImageFunction
	// Default  log retention config for all constructs.
	// Experimental.
	Retention() awslogs.RetentionDays
	// Experimental.
	SetRetention(val awslogs.RetentionDays)
	// Returns an instance of s3.Bucket created by the construct. IMPORTANT: If existingInputAssetsBucketObj was provided in Pattern Construct Props, this property will be undefined.
	// Experimental.
	S3InputAssetsBucket() awss3.Bucket
	// Returns an instance of s3.IBucket created by the construct.
	// Experimental.
	S3InputAssetsBucketInterface() awss3.IBucket
	// Returns the instance of ec2.ISecurityGroup used by the construct.
	// Experimental.
	SecurityGroup() awsec2.ISecurityGroup
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
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

// The jsii proxy struct for QaAppsyncOpensearch
type jsiiProxy_QaAppsyncOpensearch struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_QaAppsyncOpensearch) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) GraphqlApi() awsappsync.IGraphqlApi {
	var returns awsappsync.IGraphqlApi
	_jsii_.Get(
		j,
		"graphqlApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) QaBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"qaBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) QaLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"qaLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) S3InputAssetsBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3InputAssetsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) S3InputAssetsBucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3InputAssetsBucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QaAppsyncOpensearch) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Constructs a new instance of the RagAppsyncStepfnOpensearch class.
// Experimental.
func NewQaAppsyncOpensearch(scope constructs.Construct, id *string, props *QaAppsyncOpensearchProps) QaAppsyncOpensearch {
	_init_.Initialize()

	if err := validateNewQaAppsyncOpensearchParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_QaAppsyncOpensearch{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the RagAppsyncStepfnOpensearch class.
// Experimental.
func NewQaAppsyncOpensearch_Override(q QaAppsyncOpensearch, scope constructs.Construct, id *string, props *QaAppsyncOpensearchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearch",
		[]interface{}{scope, id, props},
		q,
	)
}

func (j *jsiiProxy_QaAppsyncOpensearch)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_QaAppsyncOpensearch)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_QaAppsyncOpensearch)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_QaAppsyncOpensearch)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_QaAppsyncOpensearch)SetStage(val *string) {
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
func QaAppsyncOpensearch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQaAppsyncOpensearch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QaAppsyncOpensearch_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearch",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func QaAppsyncOpensearch_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateQaAppsyncOpensearch_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearch",
		"usageMetricMap",
		val,
	)
}

func (q *jsiiProxy_QaAppsyncOpensearch) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := q.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (q *jsiiProxy_QaAppsyncOpensearch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QaAppsyncOpensearch) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := q.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (q *jsiiProxy_QaAppsyncOpensearch) UpdateEnvSuffix(props *BaseClassProps) {
	if err := q.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

