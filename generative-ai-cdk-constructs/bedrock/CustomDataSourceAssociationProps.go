package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Interface to add a new CustomDataSource to an existing KB.
// Experimental.
type CustomDataSourceAssociationProps struct {
	// The chunking stategy to use for splitting your documents or content.
	//
	// The chunks are then converted to embeddings and written to the vector
	// index allowing for similarity search and retrieval of the content.
	// Default: ChunkingStrategy.DEFAULT
	//
	// Experimental.
	ChunkingStrategy ChunkingStrategy `field:"optional" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// The context enrichment configuration to use.
	// Default: - No context enrichment is used.
	//
	// Experimental.
	ContextEnrichment ContextEnrichment `field:"optional" json:"contextEnrichment" yaml:"contextEnrichment"`
	// The custom transformation strategy to use.
	// Default: - No custom transformation is used.
	//
	// Experimental.
	CustomTransformation CustomTransformation `field:"optional" json:"customTransformation" yaml:"customTransformation"`
	// The data deletion policy to apply to the data source.
	// Default: - Sets the data deletion policy to the default of the data source type.
	//
	// Experimental.
	DataDeletionPolicy DataDeletionPolicy `field:"optional" json:"dataDeletionPolicy" yaml:"dataDeletionPolicy"`
	// The name of the data source.
	// Default: - A new name will be generated.
	//
	// Experimental.
	DataSourceName *string `field:"optional" json:"dataSourceName" yaml:"dataSourceName"`
	// A description of the data source.
	// Default: - No description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key to use to encrypt the data source.
	// Default: - Service owned and managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The parsing strategy to use.
	// Default: - No Parsing Stategy is used.
	//
	// Experimental.
	ParsingStrategy ParsingStategy `field:"optional" json:"parsingStrategy" yaml:"parsingStrategy"`
}

