package bedrock


// Experimental.
type BedrockFoundationModelProps struct {
	// Bedrock Agents can use this model.
	// Default: - false.
	//
	// Experimental.
	SupportsAgents *bool `field:"optional" json:"supportsAgents" yaml:"supportsAgents"`
	// Bedrock Knowledge Base can use this model.
	// Default: - false.
	//
	// Experimental.
	SupportsKnowledgeBase *bool `field:"optional" json:"supportsKnowledgeBase" yaml:"supportsKnowledgeBase"`
	// Embedding models have different vector dimensions.
	//
	// Only applicable for embedding models.
	// Experimental.
	VectorDimensions *float64 `field:"optional" json:"vectorDimensions" yaml:"vectorDimensions"`
}

