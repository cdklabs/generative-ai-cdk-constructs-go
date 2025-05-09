package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Interface to create a new standalone data source object.
// Experimental.
type SharePointDataSourceProps struct {
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
	// The AWS Secrets Manager secret that stores your authentication credentials for your Sharepoint instance URL.
	//
	// Secret must start with "AmazonBedrock-".
	// Experimental.
	AuthSecret awssecretsmanager.ISecret `field:"required" json:"authSecret" yaml:"authSecret"`
	// The domain of your SharePoint instance or site URL/URLs.
	//
	// Example:
	//   "yourdomain"
	//
	// Experimental.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The SharePoint site URL/URLs.
	//
	// Must start with “https”. All URLs must start with same protocol.
	//
	// Example:
	//   ["https://yourdomain.sharepoint.com/sites/mysite"]
	//
	// Experimental.
	SiteUrls *[]*string `field:"required" json:"siteUrls" yaml:"siteUrls"`
	// The identifier of your Microsoft 365 tenant.
	//
	// Example:
	//   "d1c035a6-1dcf-457d-97e3"
	//
	// Experimental.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The filters (regular expression patterns) for the crawling.
	//
	// If there's a conflict, the exclude pattern takes precedence.
	// Default: None - all your content is crawled.
	//
	// Experimental.
	Filters *[]*SharePointCrawlingFilters `field:"optional" json:"filters" yaml:"filters"`
	// The knowledge base to associate with the data source.
	// Experimental.
	KnowledgeBase IKnowledgeBase `field:"required" json:"knowledgeBase" yaml:"knowledgeBase"`
}

