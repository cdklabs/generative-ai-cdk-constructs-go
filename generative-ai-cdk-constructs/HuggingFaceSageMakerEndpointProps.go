package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
)

// Experimental.
type HuggingFaceSageMakerEndpointProps struct {
	// Experimental.
	Container ContainerImage `field:"required" json:"container" yaml:"container"`
	// Experimental.
	InstanceType SageMakerInstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Experimental.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Experimental.
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Experimental.
	Role awsiam.Role `field:"optional" json:"role" yaml:"role"`
	// Experimental.
	StartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"startupHealthCheckTimeoutInSeconds" yaml:"startupHealthCheckTimeoutInSeconds"`
	// Experimental.
	VpcConfig *awssagemaker.CfnModel_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

