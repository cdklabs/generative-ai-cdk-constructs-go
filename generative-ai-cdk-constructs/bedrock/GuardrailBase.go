package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Abstract base class for a Guardrail.
//
// Contains methods and attributes valid for Guardrails either created with CDK or imported.
// Experimental.
type GuardrailBase interface {
	awscdk.Resource
	IGuardrail
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	// Experimental.
	Env() *awscdk.ResourceEnvironment
	// The ARN of the guardrail.
	// Experimental.
	GuardrailArn() *string
	// The ID of the guardrail.
	// Experimental.
	GuardrailId() *string
	// The ID of the guardrail.
	// Experimental.
	GuardrailVersion() *string
	// Experimental.
	SetGuardrailVersion(val *string)
	// The KMS key of the guardrail if custom encryption is configured.
	// Experimental.
	KmsKey() awskms.IKey
	// When this guardrail was last updated.
	// Experimental.
	LastUpdated() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	// Experimental.
	PhysicalName() *string
	// The stack in which this resource is defined.
	// Experimental.
	Stack() awscdk.Stack
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Experimental.
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	// Experimental.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	// Experimental.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given principal identity permissions to perform actions on this agent alias.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to apply the guardrail.
	// Experimental.
	GrantApply(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this guardrail.
	//
	// By default, the metric will be calculated as a sum over a period of 5 minutes.
	// You can customize this by using the `statistic` and `period` properties.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation client errors metric for this guardrail.
	// Experimental.
	MetricInvocationClientErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation latency metric for this guardrail.
	// Experimental.
	MetricInvocationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocations metric for this guardrail.
	// Experimental.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation server errors metric for this guardrail.
	// Experimental.
	MetricInvocationServerErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocations intervened metric for this guardrail.
	// Experimental.
	MetricInvocationsIntervened(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the invocation throttles metric for this guardrail.
	// Experimental.
	MetricInvocationThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the text unit count metric for this guardrail.
	// Experimental.
	MetricTextUnitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GuardrailBase
type jsiiProxy_GuardrailBase struct {
	internal.Type__awscdkResource
	jsiiProxy_IGuardrail
}

func (j *jsiiProxy_GuardrailBase) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) GuardrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) GuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) GuardrailVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) LastUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GuardrailBase) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Experimental.
func NewGuardrailBase_Override(g GuardrailBase, scope constructs.Construct, id *string, props *awscdk.ResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GuardrailBase)SetGuardrailVersion(val *string) {
	if err := j.validateSetGuardrailVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrailVersion",
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
func GuardrailBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGuardrailBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
// Experimental.
func GuardrailBase_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGuardrailBase_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
// Experimental.
func GuardrailBase_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateGuardrailBase_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for all guardrails.
//
// By default, the metric will be calculated as a sum over a period of 5 minutes.
// You can customize this by using the `statistic` and `period` properties.
// Experimental.
func GuardrailBase_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGuardrailBase_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Return the invocation latency metric for all guardrails.
// Experimental.
func GuardrailBase_MetricAllInvocationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGuardrailBase_MetricAllInvocationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"metricAllInvocationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Return the invocations metric for all guardrails.
// Experimental.
func GuardrailBase_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGuardrailBase_MetricAllInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Return the invocations intervened metric for all guardrails.
// Experimental.
func GuardrailBase_MetricAllInvocationsIntervened(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGuardrailBase_MetricAllInvocationsIntervenedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"metricAllInvocationsIntervened",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Return the text unit count metric for all guardrails.
// Experimental.
func GuardrailBase_MetricAllTextUnitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateGuardrailBase_MetricAllTextUnitCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		"metricAllTextUnitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := g.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (g *jsiiProxy_GuardrailBase) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := g.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) GetResourceNameAttribute(nameAttr *string) *string {
	if err := g.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := g.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) GrantApply(grantee awsiam.IGrantable) awsiam.Grant {
	if err := g.validateGrantApplyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		g,
		"grantApply",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricInvocationClientErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationClientErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocationClientErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricInvocationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricInvocationServerErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationServerErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocationServerErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricInvocationsIntervened(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationsIntervenedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocationsIntervened",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricInvocationThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricInvocationThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricInvocationThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) MetricTextUnitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricTextUnitCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTextUnitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GuardrailBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

