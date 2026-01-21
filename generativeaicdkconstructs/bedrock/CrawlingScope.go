package bedrock


// The scope of the crawling.
// Experimental.
type CrawlingScope string

const (
	// Crawls only web pages that belong to the same host or primary domain.
	// Experimental.
	CrawlingScope_HOST_ONLY CrawlingScope = "HOST_ONLY"
	// Includes subdomains in addition to the host or primary domain, i.e. web pages that contain "aws.amazon.com" can also include sub domain "docs.aws.amazon.com".
	// Experimental.
	CrawlingScope_SUBDOMAINS CrawlingScope = "SUBDOMAINS"
	// Limit crawling to web pages that belong to the same host and with the same initial URL path.
	// Experimental.
	CrawlingScope_DEFAULT CrawlingScope = "DEFAULT"
)

