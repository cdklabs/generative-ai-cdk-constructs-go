package opensearchmanagedcluster


// Properties for an OpenSearchManagedClusterVectorStore.
// Experimental.
type OpenSearchManagedClusterVectorStoreProps struct {
	// The ARN of your OpenSearch Customer Managed Domain.
	// Experimental.
	DomainArn *string `field:"required" json:"domainArn" yaml:"domainArn"`
	// The url of your OpenSearch Managed cluster domain.
	// Experimental.
	DomainEndpoint *string `field:"required" json:"domainEndpoint" yaml:"domainEndpoint"`
	// Configuration for field mappings in the vector store.
	//
	// Bedrock uses these fields to store your data.
	// If you haven't configured these fields in your vector database, your Knowledge Base
	// will fail to be created.
	// Experimental.
	FieldMapping *OpenSearchFieldMapping `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// The vector index name of your OpenSearch Customer Managed Domain.
	// Experimental.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
}

