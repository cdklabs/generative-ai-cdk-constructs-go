package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for the ModelMonitoringProps class.
// Experimental.
type ModelMonitoringProps struct {
	// Experimental.
	BucketedStepSize *string `field:"optional" json:"bucketedStepSize" yaml:"bucketedStepSize"`
	// Experimental.
	ImageSize *string `field:"optional" json:"imageSize" yaml:"imageSize"`
	// Experimental.
	InputTokenPrice *float64 `field:"optional" json:"inputTokenPrice" yaml:"inputTokenPrice"`
	// Experimental.
	OutputTokenPrice *float64 `field:"optional" json:"outputTokenPrice" yaml:"outputTokenPrice"`
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

