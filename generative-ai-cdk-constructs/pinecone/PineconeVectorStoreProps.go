package pinecone


// Properties for a PineconeVectorStore.
// Experimental.
type PineconeVectorStoreProps struct {
	// Connection string for your Pinecone index management page.
	// Experimental.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// ARN of the secret containing the API Key to use when connecting to the Pinecone database.
	//
	// Learn more in the link below.
	// See: https://www.pinecone.io/blog/amazon-bedrock-integration/
	//
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The name of the field in which Amazon Bedrock stores metadata about the vector store.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// The name of the field in which Amazon Bedrock stores the raw text from your data.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// If you encrypted your secret, provide the KMS key here so that Bedrock can decrypt it.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Name space that will be used for writing new data to your Pinecone database.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

