package amazonaurora


// Experimental.
type AmazonAuroraDefaultVectorStoreProps struct {
	// The embeddings model vector dimension for the knowledge base.
	//
	// Must be identical as in the KnowledgeBase construct.
	// This is due to the factor that the embeddings models
	// have different vector dimensions and this construct
	// needs to know the vector dimensions to create the vector
	// index of appropriate dimensions in the Aurora database.
	// Experimental.
	EmbeddingsModelVectorDimension *float64 `field:"required" json:"embeddingsModelVectorDimension" yaml:"embeddingsModelVectorDimension"`
}

