package amazonaurora

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
)

// Properties for configuring an Amazon Aurora Vector Store.
// Experimental.
type AmazonAuroraVectorStoreProps struct {
	// The embeddings model dimension used for the Aurora Vector Store.
	//
	// The vector dimensions of the model must match the dimensions
	// used in the KnowledgeBase construct.
	// Experimental.
	EmbeddingsModelVectorDimension *float64 `field:"required" json:"embeddingsModelVectorDimension" yaml:"embeddingsModelVectorDimension"`
	// The field name for the metadata column in the Aurora Vector Store.
	// Experimental.
	MetadataField *string `field:"optional" json:"metadataField" yaml:"metadataField"`
	// The primary key field for the Aurora Vector Store table.
	// Experimental.
	PrimaryKeyField *string `field:"optional" json:"primaryKeyField" yaml:"primaryKeyField"`
	// The schema name for the Aurora Vector Store.
	// Experimental.
	SchemaName *string `field:"optional" json:"schemaName" yaml:"schemaName"`
	// The name of the table for the Aurora Vector Store.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// The field name for the text column in the Aurora Vector Store.
	// Experimental.
	TextField *string `field:"optional" json:"textField" yaml:"textField"`
	// The field name for the vector column in the Aurora Vector Store.
	// Experimental.
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
	// Cluster identifier.
	// Experimental.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The name of the database for the Aurora Vector Store.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The version of PostgreSQL to use for the Aurora Vector Store.
	//
	// By default, the latest supported version will be used.
	// Experimental.
	PostgreSQLVersion awsrds.AuroraPostgresEngineVersion `field:"optional" json:"postgreSQLVersion" yaml:"postgreSQLVersion"`
	// User's VPC in which they want to deploy Aurora Database.
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

