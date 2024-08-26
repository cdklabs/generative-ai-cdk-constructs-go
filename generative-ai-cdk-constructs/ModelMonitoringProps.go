package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for the ModelMonitoringProps class.
// Experimental.
type ModelMonitoringProps struct {
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

