package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type LlamaIndexDataLoaderProps struct {
	// Default: 'WARNING'.
	//
	// Experimental.
	ContainerLoggingLevel *string `field:"optional" json:"containerLoggingLevel" yaml:"containerLoggingLevel"`
	// The directory to build the Docker image.
	// Default: __dirname + '/docker'.
	//
	// Experimental.
	DockerImageAssetDirectory *string `field:"optional" json:"dockerImageAssetDirectory" yaml:"dockerImageAssetDirectory"`
	// The default memory.
	// Default: 2048.
	//
	// Experimental.
	MemoryLimitMiB *float64 `field:"optional" json:"memoryLimitMiB" yaml:"memoryLimitMiB"`
	// Enable observability.
	//
	// Warning: associated cost with the services
	// used. Best practive to enable by default.
	// Default: - true.
	//
	// Experimental.
	Observability *bool `field:"optional" json:"observability" yaml:"observability"`
	// Default: undefined.
	//
	// Experimental.
	OutputBucket awss3.Bucket `field:"optional" json:"outputBucket" yaml:"outputBucket"`
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Default: undefined.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

