package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Interface to create a new S3 Data Source object.
// Experimental.
type S3DataSourceProps struct {
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
	ParsingStrategy ParsingStrategy `field:"optional" json:"parsingStrategy" yaml:"parsingStrategy"`
	// The bucket that contains the data source.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The prefixes of the objects in the bucket that should be included in the data source.
	// Default: - All objects in the bucket.
	//
	// Experimental.
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
	// The knowledge base to associate with the data source.
	// Experimental.
	KnowledgeBase IKnowledgeBase `field:"required" json:"knowledgeBase" yaml:"knowledgeBase"`
}

