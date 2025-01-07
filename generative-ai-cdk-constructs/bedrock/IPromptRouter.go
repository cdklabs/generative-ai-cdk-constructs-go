package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IPromptRouter interface {
	// The ARN of the prompt router.
	// Experimental.
	PromptRouterArn() *string
	// The Id of the prompt router.
	// Experimental.
	PromptRouterId() *string
	// The foundation models / profiles this router will route to.
	// Experimental.
	RoutingEndpoints() *[]IInvokable
}

// The jsii proxy for IPromptRouter
type jsiiProxy_IPromptRouter struct {
	_ byte // padding
}

func (j *jsiiProxy_IPromptRouter) PromptRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptRouter) PromptRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptRouter) RoutingEndpoints() *[]IInvokable {
	var returns *[]IInvokable
	_jsii_.Get(
		j,
		"routingEndpoints",
		&returns,
	)
	return returns
}

