package bedrockcwdashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/bedrock"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/bedrockcwdashboard/internal"
)

// The BedrockCwDashboard class.
// Experimental.
type BedrockCwDashboard interface {
	constructs.Construct
	// Returns the instance of CloudWatch dashboard used by the construct.
	// Experimental.
	Dashboard() awscloudwatch.Dashboard
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Add guardrail monitoring to the dashboard.
	// Experimental.
	AddAllGuardrailsMonitoring()
	// Experimental.
	AddAllModelsMonitoring(props *ModelMonitoringProps)
	// Add guardrail monitoring to the dashboard.
	// Experimental.
	AddGuardrailMonitoring(guardrail bedrock.IGuardrail)
	// Experimental.
	AddModelMonitoring(modelName *string, modelId *string, props *ModelMonitoringProps)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockCwDashboard
type jsiiProxy_BedrockCwDashboard struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BedrockCwDashboard) Dashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCwDashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Constructs a new instance of the BedrockCwDashboard class.
// Experimental.
func NewBedrockCwDashboard(scope constructs.Construct, id *string, props *BedrockCwDashboardProps) BedrockCwDashboard {
	_init_.Initialize()

	if err := validateNewBedrockCwDashboardParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockCwDashboard{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrockcwdashboard.BedrockCwDashboard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the BedrockCwDashboard class.
// Experimental.
func NewBedrockCwDashboard_Override(b BedrockCwDashboard, scope constructs.Construct, id *string, props *BedrockCwDashboardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrockcwdashboard.BedrockCwDashboard",
		[]interface{}{scope, id, props},
		b,
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
func BedrockCwDashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockCwDashboard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrockcwdashboard.BedrockCwDashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCwDashboard) AddAllGuardrailsMonitoring() {
	_jsii_.InvokeVoid(
		b,
		"addAllGuardrailsMonitoring",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockCwDashboard) AddAllModelsMonitoring(props *ModelMonitoringProps) {
	if err := b.validateAddAllModelsMonitoringParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addAllModelsMonitoring",
		[]interface{}{props},
	)
}

func (b *jsiiProxy_BedrockCwDashboard) AddGuardrailMonitoring(guardrail bedrock.IGuardrail) {
	if err := b.validateAddGuardrailMonitoringParameters(guardrail); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addGuardrailMonitoring",
		[]interface{}{guardrail},
	)
}

func (b *jsiiProxy_BedrockCwDashboard) AddModelMonitoring(modelName *string, modelId *string, props *ModelMonitoringProps) {
	if err := b.validateAddModelMonitoringParameters(modelName, modelId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addModelMonitoring",
		[]interface{}{modelName, modelId, props},
	)
}

func (b *jsiiProxy_BedrockCwDashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

