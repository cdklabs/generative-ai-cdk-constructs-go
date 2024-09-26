package bedrock


// Defines the crawling filters for Salesforce data ingestion.
// Experimental.
type SalesforceCrawlingFilters struct {
	// The Salesforce object type to which this filter applies.
	// Experimental.
	ObjectType SalesforceObjectType `field:"required" json:"objectType" yaml:"objectType"`
	// Regular expression patterns to exclude specific content.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Regular expression patterns to include specific content.
	// Experimental.
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
}

