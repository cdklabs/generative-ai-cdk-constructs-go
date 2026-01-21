package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type CustomParserProps struct {
	// Experimental.
	Parser awslambda.IFunction `field:"optional" json:"parser" yaml:"parser"`
	// Experimental.
	Steps *[]*PromptStepConfigurationCustomParser `field:"optional" json:"steps" yaml:"steps"`
}

