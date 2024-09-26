package bedrock


// The filters (regular expression patterns) to include or exclude in the crawling in accordance with your scope.
// Experimental.
type CrawlingFilters struct {
	// Exclude paths.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Include patterns.
	// Experimental.
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
}

