package bedrock


// Experimental.
type VectorFieldMapping struct {
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// Default: "AMAZON_BEDROCK_METADATA".
	//
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	//
	// The text is split according to the chunking strategy you choose.
	// Default: "AMAZON_BEDROCK_TEXT".
	//
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
}

