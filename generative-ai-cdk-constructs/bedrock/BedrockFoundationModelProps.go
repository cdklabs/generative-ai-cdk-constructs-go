package bedrock


// Experimental.
type BedrockFoundationModelProps struct {
	// https://docs.aws.amazon.com/bedrock/latest/userguide/model-lifecycle.html A version is marked Legacy when there is a more recent version which provides superior performance. Amazon Bedrock sets an EOL date for Legacy versions.
	// Default: - false.
	//
	// Experimental.
	Legacy *bool `field:"optional" json:"legacy" yaml:"legacy"`
	// Currently, some of the offered models are optimized with prompts/parsers fine-tuned for integrating with the agents architecture.
	// Default: - false.
	//
	// Experimental.
	OptimizedForAgents *bool `field:"optional" json:"optimizedForAgents" yaml:"optimizedForAgents"`
	// Embeddings models have different supported vector types.
	// Experimental.
	SupportedVectorType *[]VectorType `field:"optional" json:"supportedVectorType" yaml:"supportedVectorType"`
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

