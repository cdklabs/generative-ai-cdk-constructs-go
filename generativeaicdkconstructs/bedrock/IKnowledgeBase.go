package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/bedrock/internal"
)

// Represents a Knowledge Base, either created with CDK or imported, of any type.
// Experimental.
type IKnowledgeBase interface {
	awscdk.IResource
	// Grant the given principal identity permissions to perform actions on this knowledge base.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to query the knowledge base.
	// Experimental.
	GrantQuery(grantee awsiam.IGrantable) awsiam.Grant
	// The description of the knowledge base.
	// Experimental.
	Description() *string
	// A narrative instruction of the knowledge base.
	//
	// A Bedrock Agent can use this instruction to determine if it should
	// query this Knowledge Base.
	// Experimental.
	Instruction() *string
	// The ARN of the knowledge base.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:knowledge-base/KB12345678"
	//
	// Experimental.
	KnowledgeBaseArn() *string
	// The ID of the knowledge base.
	//
	// Example:
	//   "KB12345678"
	//
	// Experimental.
	KnowledgeBaseId() *string
	// The role associated with the knowledge base.
	// Experimental.
	Role() awsiam.IRole
	// The type of knowledge base.
	// Experimental.
	Type() KnowledgeBaseType
}

// The jsii proxy for IKnowledgeBase
type jsiiProxy_IKnowledgeBase struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKnowledgeBase) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) GrantQuery(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantQueryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantQuery",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Type() KnowledgeBaseType {
	var returns KnowledgeBaseType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

