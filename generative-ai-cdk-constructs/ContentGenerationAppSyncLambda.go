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

// The ContentGenerationAppSyncLambda class.
// Experimental.
type ContentGenerationAppSyncLambda interface {
	BaseClass
	// Returns an instance of appsync.GraphqlApi created by the construct.
	// Experimental.
	CgLambdaFunction() awslambda.DockerImageFunction
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
	// Returns the instance of events.IEventBus used by the construct.
	// Experimental.
	GeneratedImageBus() awsevents.IEventBus
	// Returns an instance of appsync.GraphqlApi created by the construct.
	// Experimental.
	GraphqlApi() awsappsync.GraphqlApi
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
	// Returns an instance of s3.Bucket created by the construct. IMPORTANT: If existingGeneratedAssetsBucketObj was provided in Pattern Construct Props, this property will be undefined.
	// Experimental.
	S3GenerateAssetsBucket() awss3.Bucket
	// Returns an instance of s3.IBucket created by the construct.
	// Experimental.
	S3GenerateAssetsBucketInterface() awss3.IBucket
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

// The jsii proxy struct for ContentGenerationAppSyncLambda
type jsiiProxy_ContentGenerationAppSyncLambda struct {
	jsiiProxy_BaseClass
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) CgLambdaFunction() awslambda.DockerImageFunction {
	var returns awslambda.DockerImageFunction
	_jsii_.Get(
		j,
		"cgLambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) GeneratedImageBus() awsevents.IEventBus {
	var returns awsevents.IEventBus
	_jsii_.Get(
		j,
		"generatedImageBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) GraphqlApi() awsappsync.GraphqlApi {
	var returns awsappsync.GraphqlApi
	_jsii_.Get(
		j,
		"graphqlApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) S3GenerateAssetsBucket() awss3.Bucket {
	var returns awss3.Bucket
	_jsii_.Get(
		j,
		"s3GenerateAssetsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) S3GenerateAssetsBucketInterface() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"s3GenerateAssetsBucketInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) SecurityGroup() awsec2.ISecurityGroup {
	var returns awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}


// Constructs a new instance of the ContentGenerationAppSyncLambda class.
// Experimental.
func NewContentGenerationAppSyncLambda(scope constructs.Construct, id *string, props *ContentGenerationAppSyncLambdaProps) ContentGenerationAppSyncLambda {
	_init_.Initialize()

	if err := validateNewContentGenerationAppSyncLambdaParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContentGenerationAppSyncLambda{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the ContentGenerationAppSyncLambda class.
// Experimental.
func NewContentGenerationAppSyncLambda_Override(c ContentGenerationAppSyncLambda, scope constructs.Construct, id *string, props *ContentGenerationAppSyncLambdaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambda",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_ContentGenerationAppSyncLambda)SetStage(val *string) {
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
func ContentGenerationAppSyncLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContentGenerationAppSyncLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContentGenerationAppSyncLambda_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambda",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func ContentGenerationAppSyncLambda_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateContentGenerationAppSyncLambda_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambda",
		"usageMetricMap",
		val,
	)
}

func (c *jsiiProxy_ContentGenerationAppSyncLambda) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := c.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (c *jsiiProxy_ContentGenerationAppSyncLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContentGenerationAppSyncLambda) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := c.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (c *jsiiProxy_ContentGenerationAppSyncLambda) UpdateEnvSuffix(props *BaseClassProps) {
	if err := c.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

