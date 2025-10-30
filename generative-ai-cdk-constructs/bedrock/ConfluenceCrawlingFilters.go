package bedrock


// Defines filters for crawling Confluence content.
//
// These filters allow you to include or exclude specific content based on object types and patterns.
//
// - For Spaces: Use the unique space key
// - For Pages: Use the main page title
// - For Blogs: Use the main blog title
// - For Comments: Use "Re: Page/Blog Title"
// - For Attachments: Use the filename with extension.
// Experimental.
type ConfluenceCrawlingFilters struct {
	// The type of Confluence object to apply the filters to.
	// Experimental.
	ObjectType ConfluenceObjectType `field:"required" json:"objectType" yaml:"objectType"`
	// Regular expression patterns to exclude content.
	//
	// Content matching these patterns will not be crawled, even if it matches an include pattern.
	// Experimental.
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Regular expression patterns to include content.
	//
	// If specified, only content matching these patterns will be crawled.
	// Experimental.
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
}

