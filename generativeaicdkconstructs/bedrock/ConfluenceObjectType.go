package bedrock


// Represents the different types of content objects in Confluence that can be crawled by the data source.
// Experimental.
type ConfluenceObjectType string

const (
	// Experimental.
	ConfluenceObjectType_SPACE ConfluenceObjectType = "SPACE"
	// Experimental.
	ConfluenceObjectType_PAGE ConfluenceObjectType = "PAGE"
	// Experimental.
	ConfluenceObjectType_BLOG ConfluenceObjectType = "BLOG"
	// Experimental.
	ConfluenceObjectType_COMMENT ConfluenceObjectType = "COMMENT"
	// Experimental.
	ConfluenceObjectType_ATTACHMENT ConfluenceObjectType = "ATTACHMENT"
)

