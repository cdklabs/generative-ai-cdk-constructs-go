package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IInstanceAliase interface {
	// Experimental.
	Aliases() *map[string]*string
	// Experimental.
	SetAliases(a *map[string]*string)
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(r *string)
}

// The jsii proxy for IInstanceAliase
type jsiiProxy_IInstanceAliase struct {
	_ byte // padding
}

func (j *jsiiProxy_IInstanceAliase) Aliases() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceAliase)SetAliases(val *map[string]*string) {
	if err := j.validateSetAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aliases",
		val,
	)
}

func (j *jsiiProxy_IInstanceAliase) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceAliase)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

