package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Bedrock guardrail props.
// Experimental.
type GuardrailProps struct {
	// The message to return when the guardrail blocks a prompt.
	// Experimental.
	BlockedInputMessaging *string `field:"optional" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// The message to return when the guardrail blocks a model response.
	// Experimental.
	BlockedOutputsMessaging *string `field:"optional" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// Contextual grounding policy config for a guardrail.
	// Experimental.
	ContextualGroundingfiltersConfig *[]*ContextualGroundingPolicyConfigProps `field:"optional" json:"contextualGroundingfiltersConfig" yaml:"contextualGroundingfiltersConfig"`
	// The description of the guardrail.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of content filter configs in content policy.
	// Experimental.
	FiltersConfig *[]*ContentPolicyConfigProps `field:"optional" json:"filtersConfig" yaml:"filtersConfig"`
	// The ARN of the AWS KMS key used to encrypt the guardrail.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the guardrail.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// PII fields which needs to be masked.
	// Experimental.
	PiiConfig *[]*SensitiveInformationPolicyConfigProps `field:"optional" json:"piiConfig" yaml:"piiConfig"`
	// Metadata that you can assign to a guardrail as key-value pairs.
	// Experimental.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

