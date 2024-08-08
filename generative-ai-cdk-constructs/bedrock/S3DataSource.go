package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Sets up a data source to be added to a knowledge base.
// Experimental.
type S3DataSource interface {
	constructs.Construct
	// The Data Source cfn resource.
	// Experimental.
	DataSource() awsbedrock.CfnDataSource
	// The unique identifier of the data source.
	// Experimental.
	DataSourceId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3DataSource
type jsiiProxy_S3DataSource struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_S3DataSource) DataSource() awsbedrock.CfnDataSource {
	var returns awsbedrock.CfnDataSource
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DataSource) DataSourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3DataSource) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3DataSource(scope constructs.Construct, id *string, props *S3DataSourceProps) S3DataSource {
	_init_.Initialize()

	if err := validateNewS3DataSourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3DataSource{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3DataSource_Override(s S3DataSource, scope constructs.Construct, id *string, props *S3DataSourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSource",
		[]interface{}{scope, id, props},
		s,
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
func S3DataSource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3DataSource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3DataSource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

