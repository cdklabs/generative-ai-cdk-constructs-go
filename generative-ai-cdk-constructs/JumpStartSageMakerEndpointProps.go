package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
)

// Experimental.
type JumpStartSageMakerEndpointProps struct {
	// Experimental.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Experimental.
	Model JumpStartModel `field:"required" json:"model" yaml:"model"`
	// Experimental.
	AcceptEula *bool `field:"optional" json:"acceptEula" yaml:"acceptEula"`
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Experimental.
	InstanceType SageMakerInstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Experimental.
	Role awsiam.Role `field:"optional" json:"role" yaml:"role"`
	// Experimental.
	StartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"startupHealthCheckTimeoutInSeconds" yaml:"startupHealthCheckTimeoutInSeconds"`
	// Experimental.
	VpcConfig *awssagemaker.CfnModel_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

