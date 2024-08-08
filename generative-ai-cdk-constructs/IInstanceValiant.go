package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IInstanceValiant interface {
	// Experimental.
	Environment() *map[string]*string
	// Experimental.
	SetEnvironment(e *map[string]*string)
	// Experimental.
	ImageUri() *string
	// Experimental.
	SetImageUri(i *string)
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(i *string)
}

// The jsii proxy for IInstanceValiant
type jsiiProxy_IInstanceValiant struct {
	_ byte // padding
}

func (j *jsiiProxy_IInstanceValiant) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceValiant)SetEnvironment(val *map[string]*string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_IInstanceValiant) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceValiant)SetImageUri(val *string) {
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_IInstanceValiant) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceValiant)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

