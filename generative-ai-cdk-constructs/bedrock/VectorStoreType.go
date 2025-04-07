package bedrock


// Knowledge base can be backed by different vector databases. This enum represents the different vector databases that can be used.
//
// `OPENSEARCH_SERVERLESS` is the default vector database.
// `PINECONE` is the vector database for Pinecone.
// `AMAZON_AURORA` is the vector database for Amazon Aurora PostgreSQL.
// Experimental.
type VectorStoreType string

const (
	// `OPENSEARCH_SERVERLESS` is the vector store for OpenSearch Serverless.
	// Experimental.
	VectorStoreType_OPENSEARCH_SERVERLESS VectorStoreType = "OPENSEARCH_SERVERLESS"
	// `PINECONE` is the vector store for Pinecone.
	// Experimental.
	VectorStoreType_PINECONE VectorStoreType = "PINECONE"
	// `RDS` is the vector store for Amazon Aurora.
	// Experimental.
	VectorStoreType_AMAZON_AURORA VectorStoreType = "AMAZON_AURORA"
	// `MONGO_DB_ATLAS` is the vector store for MongoDB Atlas.
	// Experimental.
	VectorStoreType_MONGO_DB_ATLAS VectorStoreType = "MONGO_DB_ATLAS"
	// `NEPTUNE_ANALYTICS` is the vector store for Amazon Neptune Analytics.
	// Experimental.
	VectorStoreType_NEPTUNE_ANALYTICS VectorStoreType = "NEPTUNE_ANALYTICS"
)

