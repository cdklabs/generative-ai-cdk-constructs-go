package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Represents a Guardrail, either created with CDK or imported.
// Experimental.
type IGuardrail interface {
	awscdk.IResource
	// Grant the given identity permissions to apply the guardrail.
	// Experimental.
	GrantApply(grantee awsiam.IGrantable) awsiam.Grant
	// The ARN of the guardrail.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:guardrail/yympzo398ipq"@attributeundefined
	//
	// Experimental.
	GuardrailArn() *string
	// The ID of the guardrail.
	//
	// Example:
	//   "yympzo398ipq"@attributeundefined
	//
	// Experimental.
	GuardrailId() *string
	// The version of the guardrail.
	//
	// If no explicit version is created,
	// this will default to "DRAFT".
	// Experimental.
	GuardrailVersion() *string
	// Experimental.
	SetGuardrailVersion(g *string)
	// Optional KMS encryption key associated with this guardrail.
	// Experimental.
	KmsKey() awskms.IKey
}

// The jsii proxy for IGuardrail
type jsiiProxy_IGuardrail struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGuardrail) GrantApply(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantApplyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantApply",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IGuardrail) GuardrailArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail) GuardrailId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail) GuardrailVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guardrailVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGuardrail)SetGuardrailVersion(val *string) {
	if err := j.validateSetGuardrailVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guardrailVersion",
		val,
	)
}

func (j *jsiiProxy_IGuardrail) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

