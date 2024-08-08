package generative-ai-cdk-constructs


// Experimental.
type CrawlerTarget struct {
	// Type of URL to be crawled.
	// Experimental.
	TargetType CrawlerTargetType `field:"required" json:"targetType" yaml:"targetType"`
	// Target URL to be crawled.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Schedule the crawler to run every N hours following the completion of the previous job.
	// Default: - not scheduled.
	//
	// Experimental.
	CrawlIntervalHours *float64 `field:"optional" json:"crawlIntervalHours" yaml:"crawlIntervalHours"`
	// Download files from the web site.
	// Default: - true.
	//
	// Experimental.
	DownloadFiles *bool `field:"optional" json:"downloadFiles" yaml:"downloadFiles"`
	// File types (extensions) to be downloaded.
	// Default: - all file types.
	//
	// Experimental.
	FileTypes *[]*string `field:"optional" json:"fileTypes" yaml:"fileTypes"`
	// Maximum number of files to be downloaded.
	// Default: - crawler limit.
	//
	// Experimental.
	MaxFiles *float64 `field:"optional" json:"maxFiles" yaml:"maxFiles"`
	// Maximum number of requests to be made by crawler.
	// Default: - crawler limit.
	//
	// Experimental.
	MaxRequests *float64 `field:"optional" json:"maxRequests" yaml:"maxRequests"`
}

