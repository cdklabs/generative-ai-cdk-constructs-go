package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// ****************************************************************************                     ATTRS FOR IMPORTED CONSTRUCT ***************************************************************************.
// Experimental.
type PromptAttributes struct {
	// The ARN of the prompt.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:prompt/PROMPT12345"
	//
	// Experimental.
	PromptArn *string `field:"required" json:"promptArn" yaml:"promptArn"`
	// Optional KMS encryption key associated with this prompt.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The version of the prompt.
	// Default: - "DRAFT".
	//
	// Experimental.
	PromptVersion *string `field:"optional" json:"promptVersion" yaml:"promptVersion"`
}

