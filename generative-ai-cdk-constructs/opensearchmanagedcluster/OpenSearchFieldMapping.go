package opensearchmanagedcluster


// Field mapping configuration for OpenSearch vector store.
// Experimental.
type OpenSearchFieldMapping struct {
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text in chunks from your data.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// The name of the field in which Amazon Bedrock stores the vector embeddings.
	// Experimental.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
}

