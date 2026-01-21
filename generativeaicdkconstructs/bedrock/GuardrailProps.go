package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Guardrail.
// Experimental.
type GuardrailProps struct {
	// The name of the guardrail.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The message to return when the guardrail blocks a prompt.
	// Default: "Sorry, your query violates our usage policy."
	//
	// Experimental.
	BlockedInputMessaging *string `field:"optional" json:"blockedInputMessaging" yaml:"blockedInputMessaging"`
	// The message to return when the guardrail blocks a model response.
	// Default: "Sorry, I am unable to answer your question because of our usage policy."
	//
	// Experimental.
	BlockedOutputsMessaging *string `field:"optional" json:"blockedOutputsMessaging" yaml:"blockedOutputsMessaging"`
	// The content filters to apply to the guardrail.
	//
	// Note, if one of.
	// Experimental.
	ContentFilters *[]*ContentFilter `field:"optional" json:"contentFilters" yaml:"contentFilters"`
	// The contextual grounding filters to apply to the guardrail.
	// Experimental.
	ContextualGroundingFilters *[]*ContextualGroundingFilter `field:"optional" json:"contextualGroundingFilters" yaml:"contextualGroundingFilters"`
	// Up to 30 denied topics to block user inputs or model responses associated with the topic.
	// Experimental.
	DeniedTopics *[]Topic `field:"optional" json:"deniedTopics" yaml:"deniedTopics"`
	// The description of the guardrail.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A custom KMS key to use for encrypting data.
	// Default: "Your data is encrypted by default with a key that AWS owns and manages for you."
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The managed word filters to apply to the guardrail.
	// Experimental.
	ManagedWordListFilters *[]*ManagedWordFilter `field:"optional" json:"managedWordListFilters" yaml:"managedWordListFilters"`
	// The PII filters to apply to the guardrail.
	// Experimental.
	PiiFilters *[]*PIIFilter `field:"optional" json:"piiFilters" yaml:"piiFilters"`
	// The regular expression (regex) filters to apply to the guardrail.
	// Experimental.
	RegexFilters *[]*RegexFilter `field:"optional" json:"regexFilters" yaml:"regexFilters"`
	// The word filters to apply to the guardrail.
	// Experimental.
	WordFilters *[]*WordFilter `field:"optional" json:"wordFilters" yaml:"wordFilters"`
}

