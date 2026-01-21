package sagemakerdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs"
)

// The JumpStartSageMakerEndpoint class.
// Experimental.
type JumpStartSageMakerEndpoint interface {
	SageMakerEndpointBase
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
	// Default: true.
	//
	// Experimental.
	Enablexray() *bool
	// Experimental.
	SetEnablexray(val *bool)
	// Experimental.
	EndpointArn() *string
	// Default  log config for all constructs.
	// Experimental.
	FieldLogLevel() awsappsync.FieldLogLevel
	// Experimental.
	SetFieldLogLevel(val awsappsync.FieldLogLevel)
	// Experimental.
	GrantPrincipal() awsiam.IPrincipal
	// Experimental.
	InstanceCount() *float64
	// Experimental.
	InstanceType() SageMakerInstanceType
	// enable disable lambda tracing.
	// Default: Active.
	//
	// Experimental.
	LambdaTracing() awslambda.Tracing
	// Experimental.
	SetLambdaTracing(val awslambda.Tracing)
	// Experimental.
	Model() JumpStartModel
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
	// Value will be appended to resources name.
	// Default: _dev.
	//
	// Experimental.
	Stage() *string
	// Experimental.
	SetStage(val *string)
	// Experimental.
	AddObservabilityToConstruct(props *generativeaicdkconstructs.BaseClassProps)
	// Experimental.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Experimental.
	CreateSageMakerRole() awsiam.Role
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// update template description with construct usage metric and add AWS_SDK_UA_APP_ID to user agent on aws sdk.
	// Experimental.
	UpdateConstructUsageMetricCode(props *generativeaicdkconstructs.BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction)
	// Experimental.
	UpdateEnvSuffix(props *generativeaicdkconstructs.BaseClassProps)
}

// The jsii proxy struct for JumpStartSageMakerEndpoint
type jsiiProxy_JumpStartSageMakerEndpoint struct {
	jsiiProxy_SageMakerEndpointBase
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) CfnEndpoint() awssagemaker.CfnEndpoint {
	var returns awssagemaker.CfnEndpoint
	_jsii_.Get(
		j,
		"cfnEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) CfnEndpointConfig() awssagemaker.CfnEndpointConfig {
	var returns awssagemaker.CfnEndpointConfig
	_jsii_.Get(
		j,
		"cfnEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) CfnModel() awssagemaker.CfnModel {
	var returns awssagemaker.CfnModel
	_jsii_.Get(
		j,
		"cfnModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) InstanceType() SageMakerInstanceType {
	var returns SageMakerInstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) Model() JumpStartModel {
	var returns JumpStartModel
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) Role() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}


// Experimental.
func NewJumpStartSageMakerEndpoint(scope constructs.Construct, id *string, props *JumpStartSageMakerEndpointProps) JumpStartSageMakerEndpoint {
	_init_.Initialize()

	if err := validateNewJumpStartSageMakerEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_JumpStartSageMakerEndpoint{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewJumpStartSageMakerEndpoint_Override(j JumpStartSageMakerEndpoint, scope constructs.Construct, id *string, props *JumpStartSageMakerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpoint",
		[]interface{}{scope, id, props},
		j,
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint)SetStage(val *string) {
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
func JumpStartSageMakerEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJumpStartSageMakerEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func JumpStartSageMakerEndpoint_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpoint",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func JumpStartSageMakerEndpoint_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateJumpStartSageMakerEndpoint_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpoint",
		"usageMetricMap",
		val,
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) AddObservabilityToConstruct(props *generativeaicdkconstructs.BaseClassProps) {
	if err := j.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := j.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) CreateSageMakerRole() awsiam.Role {
	var returns awsiam.Role

	_jsii_.Invoke(
		j,
		"createSageMakerRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := j.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		j,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) UpdateConstructUsageMetricCode(props *generativeaicdkconstructs.BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := j.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (j *jsiiProxy_JumpStartSageMakerEndpoint) UpdateEnvSuffix(props *generativeaicdkconstructs.BaseClassProps) {
	if err := j.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

