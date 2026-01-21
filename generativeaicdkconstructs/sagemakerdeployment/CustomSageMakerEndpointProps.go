package sagemakerdeployment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker"
)

// Experimental.
type CustomSageMakerEndpointProps struct {
	// Experimental.
	Container ContainerImage `field:"required" json:"container" yaml:"container"`
	// Experimental.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Experimental.
	InstanceType SageMakerInstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Experimental.
	ModelDataUrl *string `field:"required" json:"modelDataUrl" yaml:"modelDataUrl"`
	// Experimental.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Experimental.
	AsyncInference *AsyncInferenceConfig `field:"optional" json:"asyncInference" yaml:"asyncInference"`
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Experimental.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Experimental.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Experimental.
	ModelDataDownloadTimeoutInSeconds *float64 `field:"optional" json:"modelDataDownloadTimeoutInSeconds" yaml:"modelDataDownloadTimeoutInSeconds"`
	// Experimental.
	Role awsiam.Role `field:"optional" json:"role" yaml:"role"`
	// Experimental.
	StartupHealthCheckTimeoutInSeconds *float64 `field:"optional" json:"startupHealthCheckTimeoutInSeconds" yaml:"startupHealthCheckTimeoutInSeconds"`
	// Experimental.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
	// Experimental.
	VpcConfig *awssagemaker.CfnModel_VpcConfigProperty `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

