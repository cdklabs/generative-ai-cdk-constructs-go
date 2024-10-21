package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// ****************************************************************************                     ATTRS FOR IMPORTED CONSTRUCT ***************************************************************************.
// Experimental.
type GuardrailAttributes struct {
	// The ARN of the guardrail.
	//
	// At least one of guardrailArn or guardrailId must be
	// defined in order to initialize a guardrail ref.
	// Experimental.
	GuardrailArn *string `field:"required" json:"guardrailArn" yaml:"guardrailArn"`
	// The version of the guardrail.
	// Default: "DRAFT".
	//
	// Experimental.
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
	// The KMS key of the guardrail if custom encryption is configured.
	// Default: undefined - Means data is encrypted by default with a AWS-managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

