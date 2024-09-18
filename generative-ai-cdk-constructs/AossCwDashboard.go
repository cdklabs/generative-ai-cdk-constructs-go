package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/internal"
)

// The BedrockCwDashboard class.
// Experimental.
type AossCwDashboard interface {
	constructs.Construct
	// Returns the instance of CloudWatch dashboard used by the construct.
	// Experimental.
	Dashboard() awscloudwatch.Dashboard
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	AddCollectionMonitoringbyAttributes(collectionName *string, collectionId *string, props *CollectionMonitoringProps)
	// Experimental.
	AddCollectionMonitoringByCollection(collection awsopensearchserverless.CfnCollection, props *CollectionMonitoringProps)
	// Experimental.
	AddIndexMonitoringByAtributes(collectionName *string, collectionId *string, IndexName *string, IndexId *string, props *IndexMonitoringProps)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AossCwDashboard
type jsiiProxy_AossCwDashboard struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AossCwDashboard) Dashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AossCwDashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Constructs a new instance of the AossCwDashboard class.
// Experimental.
func NewAossCwDashboard(scope constructs.Construct, id *string, props *AossCwDashboardProps) AossCwDashboard {
	_init_.Initialize()

	if err := validateNewAossCwDashboardParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AossCwDashboard{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.AossCwDashboard",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the AossCwDashboard class.
// Experimental.
func NewAossCwDashboard_Override(a AossCwDashboard, scope constructs.Construct, id *string, props *AossCwDashboardProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.AossCwDashboard",
		[]interface{}{scope, id, props},
		a,
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
func AossCwDashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAossCwDashboard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.AossCwDashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AossCwDashboard) AddCollectionMonitoringbyAttributes(collectionName *string, collectionId *string, props *CollectionMonitoringProps) {
	if err := a.validateAddCollectionMonitoringbyAttributesParameters(collectionName, collectionId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addCollectionMonitoringbyAttributes",
		[]interface{}{collectionName, collectionId, props},
	)
}

func (a *jsiiProxy_AossCwDashboard) AddCollectionMonitoringByCollection(collection awsopensearchserverless.CfnCollection, props *CollectionMonitoringProps) {
	if err := a.validateAddCollectionMonitoringByCollectionParameters(collection, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addCollectionMonitoringByCollection",
		[]interface{}{collection, props},
	)
}

func (a *jsiiProxy_AossCwDashboard) AddIndexMonitoringByAtributes(collectionName *string, collectionId *string, IndexName *string, IndexId *string, props *IndexMonitoringProps) {
	if err := a.validateAddIndexMonitoringByAtributesParameters(collectionName, collectionId, IndexName, IndexId, props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addIndexMonitoringByAtributes",
		[]interface{}{collectionName, collectionId, IndexName, IndexId, props},
	)
}

func (a *jsiiProxy_AossCwDashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

