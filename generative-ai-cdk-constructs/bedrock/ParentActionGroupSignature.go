package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// AWS Defined signatures for enabling certain capabilities in your agent.
// Experimental.
type ParentActionGroupSignature interface {
	// Experimental.
	Value() *string
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ParentActionGroupSignature
type jsiiProxy_ParentActionGroupSignature struct {
	_ byte // padding
}

func (j *jsiiProxy_ParentActionGroupSignature) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Constructor should be used as a temporary solution when a new signature is supported but its implementation in CDK hasn't been added yet.
// Experimental.
func NewParentActionGroupSignature(value *string) ParentActionGroupSignature {
	_init_.Initialize()

	if err := validateNewParentActionGroupSignatureParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_ParentActionGroupSignature{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParentActionGroupSignature",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Constructor should be used as a temporary solution when a new signature is supported but its implementation in CDK hasn't been added yet.
// Experimental.
func NewParentActionGroupSignature_Override(p ParentActionGroupSignature, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParentActionGroupSignature",
		[]interface{}{value},
		p,
	)
}

func ParentActionGroupSignature_CODE_INTERPRETER() ParentActionGroupSignature {
	_init_.Initialize()
	var returns ParentActionGroupSignature
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParentActionGroupSignature",
		"CODE_INTERPRETER",
		&returns,
	)
	return returns
}

func ParentActionGroupSignature_USER_INPUT() ParentActionGroupSignature {
	_init_.Initialize()
	var returns ParentActionGroupSignature
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParentActionGroupSignature",
		"USER_INPUT",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_ParentActionGroupSignature) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

