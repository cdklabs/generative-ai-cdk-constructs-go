package bedrock


// Properties for importing a knowledge base outside of this stack.
// Experimental.
type VectorKnowledgeBaseAttributes struct {
	// The Service Execution Role associated with the knowledge base.
	//
	// Example:
	//   "arn:aws:iam::123456789012:role/AmazonBedrockExecutionRoleForKnowledgeBaseawscdkbdgeBaseKB12345678"
	//
	// Experimental.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The ID of the knowledge base.
	//
	// Example:
	//   "KB12345678"
	//
	// Experimental.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The description of the knowledge base.
	// Default: - No description provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Instructions for agents based on the design and type of information of the Knowledge Base.
	//
	// This will impact how Agents interact with the Knowledge Base.
	// Default: - No description provided.
	//
	// Experimental.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// Specifies whether to use the knowledge base or not when sending an InvokeAgent request.
	// Default: - ENABLED.
	//
	// Experimental.
	KnowledgeBaseState *string `field:"optional" json:"knowledgeBaseState" yaml:"knowledgeBaseState"`
	// The vector store type for the knowledge base.
	// Experimental.
	VectorStoreType VectorStoreType `field:"required" json:"vectorStoreType" yaml:"vectorStoreType"`
}

