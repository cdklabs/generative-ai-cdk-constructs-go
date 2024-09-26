package bedrock


// Properties for importing a knowledge base outside of this stack.
// Experimental.
type KnowledgeBaseAttributes struct {
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
}

