package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/kendra"
)

// Properties for creating a Kendra Index Knowledge Base.
// Experimental.
type KendraKnowledgeBaseProps struct {
	// The description of the knowledge base.
	// Default: - No description provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Existing IAM role with policy statements granting appropriate permissions to invoke the specific embeddings models.
	//
	// Any entity (e.g., an AWS service or application) that assumes
	// this role will be able to invoke or use the
	// specified embeddings model within the Bedrock service.
	// Experimental.
	ExistingRole awsiam.IRole `field:"optional" json:"existingRole" yaml:"existingRole"`
	// A narrative description of the knowledge base.
	//
	// A Bedrock Agent can use this instruction to determine if it should
	// query this Knowledge Base.
	// Default: - No description provided.
	//
	// Experimental.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// The name of the knowledge base.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Kendra Index to use for the knowledge base.
	// Experimental.
	KendraIndex kendra.IKendraGenAiIndex `field:"required" json:"kendraIndex" yaml:"kendraIndex"`
}

