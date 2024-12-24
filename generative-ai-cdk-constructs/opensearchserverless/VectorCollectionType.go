package opensearchserverless


// The type of collection.
// Experimental.
type VectorCollectionType string

const (
	// Search – Full-text search that powers applications in your internal networks (content management systems, legal documents) and internet-facing applications, such as ecommerce website search and content search.
	// Experimental.
	VectorCollectionType_SEARCH VectorCollectionType = "SEARCH"
	// Time series – The log analytics segment that focuses on analyzing large volumes of semi-structured, machine-generated data in real-time for operational, security, user behavior, and business insights.
	// Experimental.
	VectorCollectionType_TIMESERIES VectorCollectionType = "TIMESERIES"
	// Vector search – Semantic search on vector embeddings that simplifies vector data management and powers machine learning (ML) augmented search experiences and generative AI applications, such as chatbots, personal assistants, and fraud detection.
	// Experimental.
	VectorCollectionType_VECTORSEARCH VectorCollectionType = "VECTORSEARCH"
)

