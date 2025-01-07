package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Experimental.
type ToolConfiguration struct {
	// Experimental.
	ToolChoice ToolChoice `field:"required" json:"toolChoice" yaml:"toolChoice"`
	// Experimental.
	Tools *[]*awsbedrock.CfnPrompt_ToolProperty `field:"required" json:"tools" yaml:"tools"`
}

