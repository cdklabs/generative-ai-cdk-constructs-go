package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/internal"
)

// Experimental.
type CustomSageMakerEndpoint interface {
	SageMakerEndpointBase
	awsiam.IGrantable
	// Experimental.
	CfnEndpoint() awssagemaker.CfnEndpoint
	// Experimental.
	CfnEndpointConfig() awssagemaker.CfnEndpointConfig
	// Experimental.
	CfnModel() awssagemaker.CfnModel
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
	// Experimental.
	EndpointArn() *string
	// Experimental.
	ErrorTopic() awssns.Topic
	// Default  log config for all constructs.
	// Experimental.
	FieldLogLevel() awsappsync.FieldLogLevel
	// Experimental.
	SetFieldLogLevel(val awsappsync.FieldLogLevel)
	// The principal to grant permissions to.
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Experimental.
	InstanceCount() *float64
	// Experimental.
	InstanceType() SageMakerInstanceType
	// enable disable lambda tracing.
	// Default: - Active.
	//
	// Experimental.
	LambdaTracing() awslambda.Tracing
	// Experimental.
	SetLambdaTracing(val awslambda.Tracing)
	// Experimental.
	ModelDataDownloadTimeoutInSeconds() *float64
	// Experimental.
	ModelDataUrl() *string
	// Experimental.
	ModelId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Default  log retention config for all constructs.
	// Experimental.
	Retention() awslogs.RetentionDays
	// Experimental.
	SetRetention(val awslogs.RetentionDays)
	// Experimental.
	Role() awsiam.Role
	// Experimental.
	ScalingPolicy() awsapplicationautoscaling.StepScalingPolicy
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
	// Experimental.
	SuccessTopic() awssns.Topic
	// Experimental.
	AddObservabilityToConstruct(props *BaseClassProps)
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Experimental.
	CreateSageMakerRole() awsiam.Role
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Experimental.
	UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction)
	// Experimental.
	UpdateEnvSuffix(props *BaseClassProps)
}

// The jsii proxy struct for CustomSageMakerEndpoint
type jsiiProxy_CustomSageMakerEndpoint struct {
	jsiiProxy_SageMakerEndpointBase
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_CustomSageMakerEndpoint) CfnEndpoint() awssagemaker.CfnEndpoint {
	var returns awssagemaker.CfnEndpoint
	_jsii_.Get(
		j,
		"cfnEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) CfnEndpointConfig() awssagemaker.CfnEndpointConfig {
	var returns awssagemaker.CfnEndpointConfig
	_jsii_.Get(
		j,
		"cfnEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) CfnModel() awssagemaker.CfnModel {
	var returns awssagemaker.CfnModel
	_jsii_.Get(
		j,
		"cfnModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) ErrorTopic() awssns.Topic {
	var returns awssns.Topic
	_jsii_.Get(
		j,
		"errorTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) InstanceType() SageMakerInstanceType {
	var returns SageMakerInstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) ModelDataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelDataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) Role() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) ScalingPolicy() awsapplicationautoscaling.StepScalingPolicy {
	var returns awsapplicationautoscaling.StepScalingPolicy
	_jsii_.Get(
		j,
		"scalingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomSageMakerEndpoint) SuccessTopic() awssns.Topic {
	var returns awssns.Topic
	_jsii_.Get(
		j,
		"successTopic",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomSageMakerEndpoint(scope constructs.Construct, id *string, props *CustomSageMakerEndpointProps) CustomSageMakerEndpoint {
	_init_.Initialize()

	if err := validateNewCustomSageMakerEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomSageMakerEndpoint{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomSageMakerEndpoint_Override(c CustomSageMakerEndpoint, scope constructs.Construct, id *string, props *CustomSageMakerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CustomSageMakerEndpoint)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_CustomSageMakerEndpoint)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_CustomSageMakerEndpoint)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_CustomSageMakerEndpoint)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_CustomSageMakerEndpoint)SetStage(val *string) {
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
func CustomSageMakerEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomSageMakerEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CustomSageMakerEndpoint_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpoint",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func CustomSageMakerEndpoint_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateCustomSageMakerEndpoint_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpoint",
		"usageMetricMap",
		val,
	)
}

func (c *jsiiProxy_CustomSageMakerEndpoint) AddObservabilityToConstruct(props *BaseClassProps) {
	if err := c.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (c *jsiiProxy_CustomSageMakerEndpoint) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := c.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CustomSageMakerEndpoint) CreateSageMakerRole() awsiam.Role {
	var returns awsiam.Role

	_jsii_.Invoke(
		c,
		"createSageMakerRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomSageMakerEndpoint) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomSageMakerEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomSageMakerEndpoint) UpdateConstructUsageMetricCode(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := c.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (c *jsiiProxy_CustomSageMakerEndpoint) UpdateEnvSuffix(props *BaseClassProps) {
	if err := c.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

