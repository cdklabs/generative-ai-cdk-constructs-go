package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/internal"
)

// A state machine fragment that creates a Bedrock Batch Inference Job and waits for its completion.
//
// Input schema:
// ```
// {
//   "job_name": string,        // Required. Name of the batch inference job
//   "manifest_keys": string[],    // Required. List of S3 keys of the input manifest files
//   "model_id": string        // Required. Model ID to use.
// }
// ```
//
// Output schema:
// ```
// {
//   "status": string,        // Required. Status of the batch job. One of: "Completed" or "PartiallyCompleted"
//   "bucket": string,        // Required. S3 bucket where the output is stored
//   "keys": string[]         // Required. Array of S3 keys for the output files
// }
// ```
//
// Error schema:
// ```
// {
//   "status": string,        // Required. Status will be one of: "Failed", "Stopped", or "Expired"
//   "error": string,         // Required. Error code
//   "cause": string          // Required. Error message
// }
// ```.
// Experimental.
type BedrockBatchSfn interface {
	awsstepfunctions.StateMachineFragment
	// The states to chain onto if this fragment is used.
	// Experimental.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	// Experimental.
	Id() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The start state of this state machine fragment.
	// Experimental.
	StartState() awsstepfunctions.State
	// Continue normal execution with the given state.
	// Experimental.
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Prefix the IDs of all states in this state machine fragment.
	//
	// Use this to avoid multiple copies of the state machine all having the
	// same state IDs.
	// Experimental.
	PrefixStates(prefix *string) awsstepfunctions.StateMachineFragment
	// Wrap all states in this state machine fragment up into a single state.
	//
	// This can be used to add retry or error handling onto this state
	// machine fragment.
	//
	// Be aware that this changes the result of the inner state machine
	// to be an array with the result of the state machine in it. Adjust
	// your paths accordingly. For example, change 'outputPath' to
	// '$[0]'.
	// Experimental.
	ToSingleState(options *awsstepfunctions.SingleStateOptions) awsstepfunctions.Parallel
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockBatchSfn
type jsiiProxy_BedrockBatchSfn struct {
	internal.Type__awsstepfunctionsStateMachineFragment
}

func (j *jsiiProxy_BedrockBatchSfn) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockBatchSfn) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockBatchSfn) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockBatchSfn) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}


// Experimental.
func NewBedrockBatchSfn(parent constructs.Construct, id *string, props *BedrockBatchSfnProps) BedrockBatchSfn {
	_init_.Initialize()

	if err := validateNewBedrockBatchSfnParameters(parent, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockBatchSfn{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.BedrockBatchSfn",
		[]interface{}{parent, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBedrockBatchSfn_Override(b BedrockBatchSfn, parent constructs.Construct, id *string, props *BedrockBatchSfnProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.BedrockBatchSfn",
		[]interface{}{parent, id, props},
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
func BedrockBatchSfn_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockBatchSfn_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.BedrockBatchSfn",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockBatchSfn) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := b.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		b,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockBatchSfn) PrefixStates(prefix *string) awsstepfunctions.StateMachineFragment {
	var returns awsstepfunctions.StateMachineFragment

	_jsii_.Invoke(
		b,
		"prefixStates",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockBatchSfn) ToSingleState(options *awsstepfunctions.SingleStateOptions) awsstepfunctions.Parallel {
	if err := b.validateToSingleStateParameters(options); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Parallel

	_jsii_.Invoke(
		b,
		"toSingleState",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockBatchSfn) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

