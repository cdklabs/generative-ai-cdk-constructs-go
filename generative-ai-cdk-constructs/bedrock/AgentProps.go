package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a Bedrock Agent.
// Experimental.
type AgentProps struct {
	// The Bedrock text foundation model for the agent to use.
	// Experimental.
	FoundationModel IInvokable `field:"required" json:"foundationModel" yaml:"foundationModel"`
	// A narrative instruction to provide the agent as context.
	// Experimental.
	Instruction *string `field:"required" json:"instruction" yaml:"instruction"`
	// AgentActionGroup to make available to the agent.
	// Default: - No AgentActionGroup  is used.
	//
	// Experimental.
	ActionGroups *[]AgentActionGroup `field:"optional" json:"actionGroups" yaml:"actionGroups"`
	// Name of the alias for the agent.
	// Default: - No alias is created.
	//
	// Experimental.
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// A description of the agent.
	// Default: - No description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Select whether the agent can prompt additional information from the user when it does not have enough information to respond to an utterance.
	// Default: - False.
	//
	// Experimental.
	EnableUserInput *bool `field:"optional" json:"enableUserInput" yaml:"enableUserInput"`
	// KMS encryption key to use for the agent.
	// Default: - An AWS managed key is used.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The existing IAM Role for the agent with a trust policy that allows the Bedrock service to assume the role.
	// Experimental.
	ExistingRole awsiam.Role `field:"optional" json:"existingRole" yaml:"existingRole"`
	// Guardrail configuration.
	//
	// Warning: If you provide a guardrail configuration through the constructor,
	// you will need to provide the correct permissions for your agent to access
	// the guardrails. If you want the permissions to be configured on your behalf,
	// use the addGuardrail method.
	// Default: - No guardrails associated to the agent.
	//
	// Experimental.
	GuardrailConfiguration *GuardrailConfiguration `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// How long sessions should be kept open for the agent.
	// Default: - 1 hour.
	//
	// Experimental.
	IdleSessionTTL awscdk.Duration `field:"optional" json:"idleSessionTTL" yaml:"idleSessionTTL"`
	// Knowledge Bases to make available to the agent.
	// Default: - No knowledge base is used.
	//
	// Experimental.
	KnowledgeBases *[]KnowledgeBase `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// The name of the agent.
	// Default: - A name is automatically generated.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Overrides for the agent.
	// Default: - No overrides are provided.
	//
	// Experimental.
	PromptOverrideConfiguration *PromptOverrideConfiguration `field:"optional" json:"promptOverrideConfiguration" yaml:"promptOverrideConfiguration"`
	// Whether to prepare the agent for use.
	// Default: - false.
	//
	// Experimental.
	ShouldPrepareAgent *bool `field:"optional" json:"shouldPrepareAgent" yaml:"shouldPrepareAgent"`
	// OPTIONAL: Tag (KEY-VALUE) bedrock agent resource.
	// Default: - false.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

