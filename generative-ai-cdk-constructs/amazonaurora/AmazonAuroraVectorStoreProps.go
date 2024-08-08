package amazonaurora


// Properties for a AmazonAuroraVectorStore.
// Experimental.
type AmazonAuroraVectorStoreProps struct {
	// The Secret ARN of your Amazon Aurora DB cluster.
	// Experimental.
	CredentialsSecretArn *string `field:"required" json:"credentialsSecretArn" yaml:"credentialsSecretArn"`
	// The name of your Database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Provide the metadata field that you configured in Amazon Aurora.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// Provide the primary key that you configured in Amazon Aurora.
	// Experimental.
	PrimaryKeyField *string `field:"required" json:"primaryKeyField" yaml:"primaryKeyField"`
	// The ARN of your Amazon Aurora DB cluster.
	// Experimental.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The Table Name of your Amazon Aurora DB cluster.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Provide the text field that you configured in Amazon Aurora.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// Provide the vector field that you configured in Amazon Aurora.
	// Experimental.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
}

