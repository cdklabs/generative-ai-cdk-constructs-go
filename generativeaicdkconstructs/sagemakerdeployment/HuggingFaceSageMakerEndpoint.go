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
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/sagemakerdeployment/internal"
)

// The HuggingFaceSageMakerEndpoint class.
// Experimental.
type HuggingFaceSageMakerEndpoint interface {
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
	// The principal to grant permissions to.
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

// The jsii proxy struct for HuggingFaceSageMakerEndpoint
type jsiiProxy_HuggingFaceSageMakerEndpoint struct {
	jsiiProxy_SageMakerEndpointBase
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) CfnEndpoint() awssagemaker.CfnEndpoint {
	var returns awssagemaker.CfnEndpoint
	_jsii_.Get(
		j,
		"cfnEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) CfnEndpointConfig() awssagemaker.CfnEndpointConfig {
	var returns awssagemaker.CfnEndpointConfig
	_jsii_.Get(
		j,
		"cfnEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) CfnModel() awssagemaker.CfnModel {
	var returns awssagemaker.CfnModel
	_jsii_.Get(
		j,
		"cfnModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) ConstructUsageMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"constructUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) Enablexray() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enablexray",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) FieldLogLevel() awsappsync.FieldLogLevel {
	var returns awsappsync.FieldLogLevel
	_jsii_.Get(
		j,
		"fieldLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) InstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) InstanceType() SageMakerInstanceType {
	var returns SageMakerInstanceType
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) LambdaTracing() awslambda.Tracing {
	var returns awslambda.Tracing
	_jsii_.Get(
		j,
		"lambdaTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) Retention() awslogs.RetentionDays {
	var returns awslogs.RetentionDays
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) Role() awsiam.Role {
	var returns awsiam.Role
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint) Stage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stage",
		&returns,
	)
	return returns
}


// Experimental.
func NewHuggingFaceSageMakerEndpoint(scope constructs.Construct, id *string, props *HuggingFaceSageMakerEndpointProps) HuggingFaceSageMakerEndpoint {
	_init_.Initialize()

	if err := validateNewHuggingFaceSageMakerEndpointParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HuggingFaceSageMakerEndpoint{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHuggingFaceSageMakerEndpoint_Override(h HuggingFaceSageMakerEndpoint, scope constructs.Construct, id *string, props *HuggingFaceSageMakerEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpoint",
		[]interface{}{scope, id, props},
		h,
	)
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint)SetEnablexray(val *bool) {
	if err := j.validateSetEnablexrayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablexray",
		val,
	)
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint)SetFieldLogLevel(val awsappsync.FieldLogLevel) {
	if err := j.validateSetFieldLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldLogLevel",
		val,
	)
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint)SetLambdaTracing(val awslambda.Tracing) {
	if err := j.validateSetLambdaTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaTracing",
		val,
	)
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint)SetRetention(val awslogs.RetentionDays) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_HuggingFaceSageMakerEndpoint)SetStage(val *string) {
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
func HuggingFaceSageMakerEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHuggingFaceSageMakerEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HuggingFaceSageMakerEndpoint_UsageMetricMap() *map[string]*float64 {
	_init_.Initialize()
	var returns *map[string]*float64
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpoint",
		"usageMetricMap",
		&returns,
	)
	return returns
}

func HuggingFaceSageMakerEndpoint_SetUsageMetricMap(val *map[string]*float64) {
	_init_.Initialize()
	if err := validateHuggingFaceSageMakerEndpoint_SetUsageMetricMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.StaticSet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpoint",
		"usageMetricMap",
		val,
	)
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) AddObservabilityToConstruct(props *generativeaicdkconstructs.BaseClassProps) {
	if err := h.validateAddObservabilityToConstructParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addObservabilityToConstruct",
		[]interface{}{props},
	)
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := h.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) CreateSageMakerRole() awsiam.Role {
	var returns awsiam.Role

	_jsii_.Invoke(
		h,
		"createSageMakerRole",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := h.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		h,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) UpdateConstructUsageMetricCode(props *generativeaicdkconstructs.BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) {
	if err := h.validateUpdateConstructUsageMetricCodeParameters(props, scope, lambdaFunctions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"updateConstructUsageMetricCode",
		[]interface{}{props, scope, lambdaFunctions},
	)
}

func (h *jsiiProxy_HuggingFaceSageMakerEndpoint) UpdateEnvSuffix(props *generativeaicdkconstructs.BaseClassProps) {
	if err := h.validateUpdateEnvSuffixParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"updateEnvSuffix",
		[]interface{}{props},
	)
}

