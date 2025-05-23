package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Interface to add a new data source to an existing KB.
// Experimental.
type WebCrawlerDataSourceAssociationProps struct {
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
	// The source urls in the format `https://www.sitename.com`. Maximum of 100 URLs.
	// Experimental.
	SourceUrls *[]*string `field:"required" json:"sourceUrls" yaml:"sourceUrls"`
	// The max rate at which pages are crawled, up to 300 per minute per host.
	//
	// Higher values will decrease sync time but increase the load on the host.
	// Default: 300.
	//
	// Experimental.
	CrawlingRate *float64 `field:"optional" json:"crawlingRate" yaml:"crawlingRate"`
	// The scope of the crawling.
	// Default: - CrawlingScope.DEFAULT
	//
	// Experimental.
	CrawlingScope CrawlingScope `field:"optional" json:"crawlingScope" yaml:"crawlingScope"`
	// The filters (regular expression patterns) for the crawling.
	//
	// If there's a conflict, the exclude pattern takes precedence.
	// Default: None.
	//
	// Experimental.
	Filters *CrawlingFilters `field:"optional" json:"filters" yaml:"filters"`
	// The maximum number of pages to crawl.
	//
	// The max number of web pages crawled from your source URLs,
	// up to 25,000 pages. If the web pages exceed this limit, the data source sync will fail and
	// no web pages will be ingested.
	// Default: - No limit.
	//
	// Experimental.
	MaxPages *float64 `field:"optional" json:"maxPages" yaml:"maxPages"`
	// The user agent string to use when crawling.
	// Default: - Default user agent string.
	//
	// Experimental.
	UserAgent *string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// The user agent header to use when crawling.
	//
	// A string used for identifying
	// the crawler or bot when it accesses a web server. The user agent header value
	// consists of the bedrockbot, UUID, and a user agent suffix for your crawler (if one is provided).
	// By default, it is set to bedrockbot_UUID. You can optionally append a custom suffix to bedrockbot_UUID
	// to allowlist a specific user agent permitted to access your source URLs.
	// Default: - Default user agent header (bedrockbot_UUID).
	//
	// Experimental.
	UserAgentHeader *string `field:"optional" json:"userAgentHeader" yaml:"userAgentHeader"`
}

