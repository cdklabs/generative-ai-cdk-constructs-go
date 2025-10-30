package sagemakerdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IJumpStartModelSpec interface {
	// Experimental.
	ArtifactKey() *string
	// Experimental.
	SetArtifactKey(a *string)
	// Experimental.
	DefaultInstanceType() *string
	// Experimental.
	SetDefaultInstanceType(d *string)
	// Experimental.
	Environment() *map[string]interface{}
	// Experimental.
	SetEnvironment(e *map[string]interface{})
	// Experimental.
	GatedBucket() *bool
	// Experimental.
	SetGatedBucket(g *bool)
	// Experimental.
	InstanceAliases() *[]IInstanceAliase
	// Experimental.
	SetInstanceAliases(i *[]IInstanceAliase)
	// Experimental.
	InstanceTypes() *[]*string
	// Experimental.
	SetInstanceTypes(i *[]*string)
	// Experimental.
	InstanceVariants() *[]IInstanceValiant
	// Experimental.
	SetInstanceVariants(i *[]IInstanceValiant)
	// Experimental.
	ModelId() *string
	// Experimental.
	SetModelId(m *string)
	// Experimental.
	ModelPackageArns() *map[string]*string
	// Experimental.
	SetModelPackageArns(m *map[string]*string)
	// Experimental.
	PrepackedArtifactKey() *string
	// Experimental.
	SetPrepackedArtifactKey(p *string)
	// Experimental.
	RequiresEula() *bool
	// Experimental.
	SetRequiresEula(r *bool)
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(v *string)
}

// The jsii proxy for IJumpStartModelSpec
type jsiiProxy_IJumpStartModelSpec struct {
	_ byte // padding
}

func (j *jsiiProxy_IJumpStartModelSpec) ArtifactKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetArtifactKey(val *string) {
	_jsii_.Set(
		j,
		"artifactKey",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) DefaultInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetDefaultInstanceType(val *string) {
	if err := j.validateSetDefaultInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceType",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) Environment() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetEnvironment(val *map[string]interface{}) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) GatedBucket() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"gatedBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetGatedBucket(val *bool) {
	if err := j.validateSetGatedBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatedBucket",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) InstanceAliases() *[]IInstanceAliase {
	var returns *[]IInstanceAliase
	_jsii_.Get(
		j,
		"instanceAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetInstanceAliases(val *[]IInstanceAliase) {
	_jsii_.Set(
		j,
		"instanceAliases",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) InstanceVariants() *[]IInstanceValiant {
	var returns *[]IInstanceValiant
	_jsii_.Get(
		j,
		"instanceVariants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetInstanceVariants(val *[]IInstanceValiant) {
	_jsii_.Set(
		j,
		"instanceVariants",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetModelId(val *string) {
	if err := j.validateSetModelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelId",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) ModelPackageArns() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"modelPackageArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetModelPackageArns(val *map[string]*string) {
	_jsii_.Set(
		j,
		"modelPackageArns",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) PrepackedArtifactKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prepackedArtifactKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetPrepackedArtifactKey(val *string) {
	_jsii_.Set(
		j,
		"prepackedArtifactKey",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) RequiresEula() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"requiresEula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetRequiresEula(val *bool) {
	if err := j.validateSetRequiresEulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresEula",
		val,
	)
}

func (j *jsiiProxy_IJumpStartModelSpec) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJumpStartModelSpec)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

