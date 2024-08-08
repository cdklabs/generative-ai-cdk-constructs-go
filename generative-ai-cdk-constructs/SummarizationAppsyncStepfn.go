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
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type SummarizationAppsyncStepfn interface {
	BaseClass
	// construct usage metric , added in template description.
	// Experimental.
	ConstructUsageMetric() *string
	// Returns an instance of lambda.DockerImageFunction used for the document reading job created by the construct.
	// Experimental.
	DocumentReaderLambdaFunction() awslambda.DockerImageFunction
	// enable disable xray tracing.
	// Default: - True.
	//
	// Experimental.
	Enablexray() *bool
	// Experimental.
	SetEnablexray(val *bool)
	// Returns an instance of events.IEventBus created by the construct.
	// Experimental.
	EventBridgeBus() awsevents.IEventBus
	// Default  log config for all constructs.
	// Experimental.
	FieldLogLevel() awsappsync.FieldLogLevel
	// Experimental.
	SetFieldLogLevel(val awsappsync.FieldLogLevel)
	// Returns an instance of appsync.CfnGraphQLApi for summary created by the construct.
	// Experimental.
	GraphqlApi() awsappsync.IGraphqlApi
	// Graphql Api Id value.
	// Experimental.
	GraphqlApiId() *string
	// Graphql Url value.
	// Experimental.
	GraphqlUrl() *string
	// Returns the instance of s3.IBucket used by the construct.
	// Experimental.
	InputAssetBucket() awss3.IBucket
	// Returns an instance of lambda.DockerImageFunction used for the input validation job created by the construct.
	// Experimental.
	InputValidationLambdaFunction() awslambda.DockerImageFunction
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
	// Returns the instance of s3.IBucket used by the construct.
	// Experimental.
	ProcessedAssetBucket() awss3.IBucket
	// Default  log retention config for all constructs.
	// Experimental.
	Retention() awslogs.RetentionDays
	// Experimental.
	SetRetention(val awslogs.RetentionDays)
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
	// Step function.
	// Default: - fieldLogLevel - None.
	//
	// Experimental.
	StateMachine() awsstepfunctions.StateMachine
	// Returns an instance of lambda.DockerImageFunction used for the summary generation job created by the construct.
	// Experimental.
	SummaryGeneratorLambdaFunction() awslambda.DockerImageFunction
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

// The jsii proxy struct for SummarizationAppsyncStepfn
type jsiiProxy_SummarizationAppsyncStepfn struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) DocumentReaderLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"documentReaderLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) EventBridgeBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"eventBridgeBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) GraphqlApi() awsappsync.IGraphqlApi {
	var returns awsappsync.IGraphqlApi
	_jsii_.Get(
		j,
		"graphqlApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) GraphqlApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphqlApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) GraphqlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphqlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) InputAssetBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"inputAssetBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) InputValidationLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"inputValidationLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) ProcessedAssetBucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"processedAssetBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) SummaryGeneratorLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"summaryGeneratorLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SummarizationAppsyncStepfn) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Constructs a new instance of the SummarizationAppsyncStepfn class.
// Experimental.
func NewSummarizationAppsyncStepfn(scope constructs.Construct, id *string, props *SummarizationAppsyncStepfnProps) SummarizationAppsyncStepfn {
	_init_.Initialize()

	if err := validateNewSummarizationAppsyncStepfnParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SummarizationAppsyncStepfn{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfn",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the SummarizationAppsyncStepfn class.
// Experimental.
func NewSummarizationAppsyncStepfn_Override(s SummarizationAppsyncStepfn, scope constructs.Construct, id *string, props *SummarizationAppsyncStepfnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfn",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SummarizationAppsyncStepfn)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_SummarizationAppsyncStepfn)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_SummarizationAppsyncStepfn)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_SummarizationAppsyncStepfn)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_SummarizationAppsyncStepfn)SetStage(val *string) {
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
func SummarizationAppsyncStepfn_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSummarizationAppsyncStepfn_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfn",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SummarizationAppsyncStepfn_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfn",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func SummarizationAppsyncStepfn_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateSummarizationAppsyncStepfn_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfn",
		"usageMetricMap",
		val,
	)
}

func (s *jsiiProxy_SummarizationAppsyncStepfn) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := s.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (s *jsiiProxy_SummarizationAppsyncStepfn) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SummarizationAppsyncStepfn) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := s.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (s *jsiiProxy_SummarizationAppsyncStepfn) UpdateEnvSuffix(props *BaseClassProps) {
	if err := s.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

