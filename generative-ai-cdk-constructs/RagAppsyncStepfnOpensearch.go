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

// The RagAppsyncStepfnOpensearch class.
// Experimental.
type RagAppsyncStepfnOpensearch interface {
	BaseClass
	// construct usage metric , added in template description.
	// Experimental.
	ConstructUsageMetric() *string
	// Returns an instance of lambda.DockerImageFunction used for the embeddings job created by the construct.
	// Experimental.
	EmbeddingsLambdaFunction() awslambda.DockerImageFunction
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
	// Returns an instance of lambda.DockerImageFunction used for the file transformer job created by the construct.
	// Experimental.
	FileTransformerLambdaFunction() awslambda.DockerImageFunction
	// Returns an instance of appsync.IGraphqlApi created by the construct.
	// Experimental.
	GraphqlApi() awsappsync.IGraphqlApi
	// Returns the instance of events.IEventBus used by the construct.
	// Experimental.
	IngestionBus() awsevents.IEventBus
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
	// Returns an instance of s3.Bucket created by the construct. IMPORTANT: If existingProcessedAssetsBucketObj was provided in Pattern Construct Props, this property will be undefined.
	// Experimental.
	S3ProcessedAssetsBucket() awss3.Bucket
	// Returns an instance of s3.IBucket created by the construct.
	// Experimental.
	S3ProcessedAssetsBucketInterface() awss3.IBucket
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
	// Returns an instance of stepfn.StateMachine created by the construct.
	// Experimental.
	StateMachine() awsstepfunctions.StateMachine
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

// The jsii proxy struct for RagAppsyncStepfnOpensearch
type jsiiProxy_RagAppsyncStepfnOpensearch struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) EmbeddingsLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"embeddingsLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) FileTransformerLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"fileTransformerLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) GraphqlApi() awsappsync.IGraphqlApi {
	var returns awsappsync.IGraphqlApi
	_jsii_.Get(
		j,
		"graphqlApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) IngestionBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"ingestionBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) InputValidationLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"inputValidationLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) S3InputAssetsBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3InputAssetsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) S3InputAssetsBucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3InputAssetsBucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) S3ProcessedAssetsBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3ProcessedAssetsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) S3ProcessedAssetsBucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3ProcessedAssetsBucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) StateMachine() awsstepfunctions.StateMachine {
	var returns awsstepfunctions.StateMachine
	_jsii_.Get(
		j,
		"stateMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch) Vpc() awsec2.IVpc {
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
func NewRagAppsyncStepfnOpensearch(scope constructs.Construct, id *string, props *RagAppsyncStepfnOpensearchProps) RagAppsyncStepfnOpensearch {
	_init_.Initialize()

	if err := validateNewRagAppsyncStepfnOpensearchParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RagAppsyncStepfnOpensearch{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearch",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the RagAppsyncStepfnOpensearch class.
// Experimental.
func NewRagAppsyncStepfnOpensearch_Override(r RagAppsyncStepfnOpensearch, scope constructs.Construct, id *string, props *RagAppsyncStepfnOpensearchProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearch",
		[]interface{}{scope, id, props},
		r,
	)
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_RagAppsyncStepfnOpensearch)SetStage(val *string) {
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
func RagAppsyncStepfnOpensearch_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRagAppsyncStepfnOpensearch_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearch",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RagAppsyncStepfnOpensearch_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearch",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func RagAppsyncStepfnOpensearch_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateRagAppsyncStepfnOpensearch_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearch",
		"usageMetricMap",
		val,
	)
}

func (r *jsiiProxy_RagAppsyncStepfnOpensearch) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := r.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (r *jsiiProxy_RagAppsyncStepfnOpensearch) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RagAppsyncStepfnOpensearch) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := r.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (r *jsiiProxy_RagAppsyncStepfnOpensearch) UpdateEnvSuffix(props *BaseClassProps) {
	if err := r.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

