package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Interface to add a new data source to an existing KB.
// Experimental.
type ConfluenceDataSourceAssociationProps struct {
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
	// The AWS Secrets Manager secret that stores your authentication credentials for your Confluence instance URL.
	//
	// Secret must start with "AmazonBedrock-".
	// Experimental.
	AuthSecret awssecretsmanager.ISecret `field:"required" json:"authSecret" yaml:"authSecret"`
	// The Confluence host URL or instance URL.
	//
	// Example:
	//   https://example.atlassian.net
	//
	// Experimental.
	ConfluenceUrl *string `field:"required" json:"confluenceUrl" yaml:"confluenceUrl"`
	// The supported authentication method to connect to the data source.
	// Default: ConfluenceDataSourceAuthType.OAUTH2_CLIENT_CREDENTIALS
	//
	// Experimental.
	AuthType ConfluenceDataSourceAuthType `field:"optional" json:"authType" yaml:"authType"`
	// The filters (regular expression patterns) for the crawling.
	//
	// If there's a conflict, the exclude pattern takes precedence.
	// Default: None - all your content is crawled.
	//
	// Experimental.
	Filters *[]*ConfluenceCrawlingFilters `field:"optional" json:"filters" yaml:"filters"`
}

