package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a CDK managed Bedrock Agent.
// Experimental.
type AgentProps struct {
	// The foundation model used for orchestration by the agent.
	// Experimental.
	FoundationModel IInvokable `field:"required" json:"foundationModel" yaml:"foundationModel"`
	// The instruction used by the agent.
	//
	// This determines how the agent will perform his task.
	// This instruction must have a minimum of 40 characters.
	// Experimental.
	Instruction *string `field:"required" json:"instruction" yaml:"instruction"`
	// The Action Groups associated with the agent.
	// Experimental.
	ActionGroups *[]AgentActionGroup `field:"optional" json:"actionGroups" yaml:"actionGroups"`
	// The collaboration type for the agent.
	// Default: - No collaboration (AgentCollaboratorType.DISABLED).
	//
	// Experimental.
	AgentCollaboration AgentCollaboratorType `field:"optional" json:"agentCollaboration" yaml:"agentCollaboration"`
	// Collaborators that this agent will work with.
	// Default: - No collaborators.
	//
	// Experimental.
	AgentCollaborators *[]AgentCollaborator `field:"optional" json:"agentCollaborators" yaml:"agentCollaborators"`
	// Select whether the agent can generate, run, and troubleshoot code when trying to complete a task.
	// Default: - false.
	//
	// Experimental.
	CodeInterpreterEnabled *bool `field:"optional" json:"codeInterpreterEnabled" yaml:"codeInterpreterEnabled"`
	// Details of custom orchestration for the agent.
	// Default: - Standard orchestration.
	//
	// Experimental.
	CustomOrchestration *CustomOrchestration `field:"optional" json:"customOrchestration" yaml:"customOrchestration"`
	// A description of the agent.
	// Default: - No description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The existing IAM Role for the agent to use.
	//
	// Ensure the role has a trust policy that allows the Bedrock service to assume the role.
	// Default: - A new role is created for you.
	//
	// Experimental.
	ExistingRole awsiam.IRole `field:"optional" json:"existingRole" yaml:"existingRole"`
	// Whether to delete the resource even if it's in use.
	// Default: - true.
	//
	// Experimental.
	ForceDelete *bool `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// The guardrail that will be associated with the agent.
	// Experimental.
	Guardrail IGuardrail `field:"optional" json:"guardrail" yaml:"guardrail"`
	// How long sessions should be kept open for the agent.
	//
	// If no conversation occurs
	// during this time, the session expires and Amazon Bedrock deletes any data
	// provided before the timeout.
	// Default: - 1 hour.
	//
	// Experimental.
	IdleSessionTTL awscdk.Duration `field:"optional" json:"idleSessionTTL" yaml:"idleSessionTTL"`
	// The KMS key of the agent if custom encryption is configured.
	// Default: - An AWS managed key is used.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The KnowledgeBases associated with the agent.
	// Experimental.
	KnowledgeBases *[]IKnowledgeBase `field:"optional" json:"knowledgeBases" yaml:"knowledgeBases"`
	// The type and configuration of the memory to maintain context across multiple sessions and recall past interactions.
	//
	// This can be useful for maintaining continuity in multi-turn conversations and recalling user preferences
	// or past interactions.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/agents-memory.html
	//
	// Default: - No memory will be used. Agents will retain context from the current session only.
	//
	// Experimental.
	Memory Memory `field:"optional" json:"memory" yaml:"memory"`
	// The name of the agent.
	// Default: - A name is generated by CDK.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of orchestration to use for the agent.
	// Default: - STANDARD.
	//
	// Experimental.
	OrchestrationType OrchestrationType `field:"optional" json:"orchestrationType" yaml:"orchestrationType"`
	// Overrides some prompt templates in different parts of an agent sequence configuration.
	// Default: - No overrides are provided.
	//
	// Experimental.
	PromptOverrideConfiguration PromptOverrideConfiguration `field:"optional" json:"promptOverrideConfiguration" yaml:"promptOverrideConfiguration"`
	// Specifies whether to automatically update the `DRAFT` version of the agent after making changes to the agent.
	//
	// The `DRAFT` version can be continually iterated
	// upon during internal development.
	// Default: - false.
	//
	// Experimental.
	ShouldPrepareAgent *bool `field:"optional" json:"shouldPrepareAgent" yaml:"shouldPrepareAgent"`
	// Select whether the agent can prompt additional information from the user when it does not have enough information to respond to an utterance.
	// Default: - false.
	//
	// Experimental.
	UserInputEnabled *bool `field:"optional" json:"userInputEnabled" yaml:"userInputEnabled"`
}

