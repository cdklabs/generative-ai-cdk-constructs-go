package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for an S3 Data Source.
// Experimental.
type S3DataSourceProps struct {
	// The bucket that contains the data source.
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the data source.
	// Experimental.
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// The knowledge base that this data source belongs to.
	// Experimental.
	KnowledgeBase KnowledgeBase `field:"required" json:"knowledgeBase" yaml:"knowledgeBase"`
	// The chunking strategy to use.
	// Default: ChunkingStrategy.DEFAULT
	//
	// Experimental.
	ChunkingStrategy ChunkingStrategy `field:"optional" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// The prefixes of the objects in the bucket that should be included in the data source.
	// Default: - All objects in the bucket.
	//
	// Experimental.
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
	// The KMS key to use to encrypt the data source.
	// Default: Amazon Bedrock encrypts your data with a key that AWS owns and manages.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The maximum number of tokens to use in a chunk.
	// Default: 300.
	//
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The percentage of overlap to use in a chunk.
	// Default: 20.
	//
	// Experimental.
	OverlapPercentage *float64 `field:"optional" json:"overlapPercentage" yaml:"overlapPercentage"`
}

