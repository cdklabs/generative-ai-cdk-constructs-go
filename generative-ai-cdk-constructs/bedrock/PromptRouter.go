package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Experimental.
type PromptRouter interface {
	IInvokable
	IPromptRouter
	// The ARN of the Bedrock invokable abstraction.
	// Experimental.
	InvokableArn() *string
	// The ARN of the prompt router.
	// Experimental.
	PromptRouterArn() *string
	// The Id of the prompt router.
	// Experimental.
	PromptRouterId() *string
	// The foundation models / profiles this router will route to.
	// Experimental.
	RoutingEndpoints() *[]IInvokable
	// Gives the appropriate policies to invoke and use the invokable abstraction.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for PromptRouter
type jsiiProxy_PromptRouter struct {
	jsiiProxy_IInvokable
	jsiiProxy_IPromptRouter
}

func (j *jsiiProxy_PromptRouter) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptRouter) PromptRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptRouter) PromptRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptRouter) RoutingEndpoints() *[]IInvokable {
	var returns *[]IInvokable
	_jsii_.Get(
		j,
		"routingEndpoints",
		&returns,
	)
	return returns
}


// Experimental.
func NewPromptRouter(props *PromptRouterProps, region *string) PromptRouter {
	_init_.Initialize()

	if err := validateNewPromptRouterParameters(props, region); err != nil {
		panic(err)
	}
	j := jsiiProxy_PromptRouter{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptRouter",
		[]interface{}{props, region},
		&j,
	)

	return &j
}

// Experimental.
func NewPromptRouter_Override(p PromptRouter, props *PromptRouterProps, region *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptRouter",
		[]interface{}{props, region},
		p,
	)
}

// Experimental.
func PromptRouter_FromDefaultId(defaultRouter DefaultPromptRouterIdentifier, region *string) PromptRouter {
	_init_.Initialize()

	if err := validatePromptRouter_FromDefaultIdParameters(defaultRouter, region); err != nil {
		panic(err)
	}
	var returns PromptRouter

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptRouter",
		"fromDefaultId",
		[]interface{}{defaultRouter, region},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PromptRouter) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := p.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		p,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

