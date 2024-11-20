package bedrock


// Experimental.
type BedrockFoundationModelProps struct {
	// Bedrock Agents can use this model.
	// Default: - false.
	//
	// Experimental.
	SupportsAgents *bool `field:"optional" json:"supportsAgents" yaml:"supportsAgents"`
	// Can be used with a Cross-Region Inference Profile.
	// Default: - false.
	//
	// Experimental.
	SupportsCrossRegion *bool `field:"optional" json:"supportsCrossRegion" yaml:"supportsCrossRegion"`
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

