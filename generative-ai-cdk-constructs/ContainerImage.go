package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// https://github.com/aws/deep-learning-containers/blob/master/available_images.md.
// Experimental.
type ContainerImage interface {
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *ContainerImageConfig
}

// The jsii proxy struct for ContainerImage
type jsiiProxy_ContainerImage struct {
	_ byte // padding
}

// Experimental.
func NewContainerImage_Override(c ContainerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.ContainerImage",
		nil, // no parameters
		c,
	)
}

// Experimental.
func ContainerImage_FromAsset(directory *string, options *awsecrassets.DockerImageAssetOptions) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromAssetParameters(directory, options); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.ContainerImage",
		"fromAsset",
		[]interface{}{directory, options},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) ContainerImage {
	_init_.Initialize()

	if err := validateContainerImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.ContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerImage) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *ContainerImageConfig {
	if err := c.validateBindParameters(scope, grantable); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

