package amazonaurora

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for an existing Aurora Vector Store.
//
// You database must have TCP/IP port that the
// database will use for application connections
// set up for `5432`.
// Experimental.
type ExistingAmazonAuroraVectorStoreProps struct {
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
	// The id of the security group associated with the RDS Aurora instance.
	//
	// This security group allows access to the Aurora Vector Store from Lambda's
	// custom resource running pgVector SQL commands.
	// Experimental.
	AuroraSecurityGroupId *string `field:"required" json:"auroraSecurityGroupId" yaml:"auroraSecurityGroupId"`
	// The unique cluster identifier of your Aurora RDS cluster.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the database for the Aurora Vector Store.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The secret containing the database credentials.
	//
	// The secret must contain `host`, `port`, `username`,
	// `password` and `dbname` values.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// The VPC in which the existing Aurora Vector Store is located.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

