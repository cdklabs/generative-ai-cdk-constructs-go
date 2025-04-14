package mongodbatlas


// Interface for MongoDB Atlas vector store configuration.
// Experimental.
type MongoDBAtlasVectorStoreProps struct {
	// The name of the collection.
	// Experimental.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// The ARN of the secret containing MongoDB Atlas credentials.
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The endpoint URL for MongoDB Atlas.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The name of the endpoint service.
	// Experimental.
	EndpointServiceName *string `field:"required" json:"endpointServiceName" yaml:"endpointServiceName"`
	// The field mapping for MongoDB Atlas.
	// Experimental.
	FieldMapping *MongoDbAtlasFieldMapping `field:"required" json:"fieldMapping" yaml:"fieldMapping"`
	// The name of the vector index.
	// Experimental.
	VectorIndexName *string `field:"required" json:"vectorIndexName" yaml:"vectorIndexName"`
}

