package bedrock


// Defines the crawling filters for SharePoint data ingestion.
//
// These filters allow
// you to specify which content should be included or excluded during the crawling process.
// If you specify an inclusion and exclusion filter and both match a document,
// the exclusion filter takes precedence and the document isn’t crawled.
// Experimental.
type SharePointCrawlingFilters struct {
	// The SharePoint object type this filter applies to.
	// Experimental.
	ObjectType SharePointObjectType `field:"required" json:"objectType" yaml:"objectType"`
	// Optional array of regular expression patterns to exclude specific content.
	//
	// Content matching these patterns will be skipped during crawling.
	//
	// Example:
	//   []*string{
	//   	".*private.*",
	//   	".*confidential.*",
	//   }
	//
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Optional array of regular expression patterns to include specific content.
	//
	// Only content matching these patterns will be crawled.
	//
	// Example:
	//   []*string{
	//   	".*public.*",
	//   	".*shared.*",
	//   }
	//
	// Experimental.
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
}

