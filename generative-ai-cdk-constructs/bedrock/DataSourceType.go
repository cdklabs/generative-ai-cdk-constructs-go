package bedrock


// Represents the types of data sources that can be associated to an Knowledge Base.
// Experimental.
type DataSourceType string

const (
	// Amazon S3 Bucket data source.
	// Experimental.
	DataSourceType_S3 DataSourceType = "S3"
	// Confluence Cloud Instance data source.
	// Experimental.
	DataSourceType_CONFLUENCE DataSourceType = "CONFLUENCE"
	// Salesforce instance data source.
	// Experimental.
	DataSourceType_SALESFORCE DataSourceType = "SALESFORCE"
	// Microsoft SharePoint instance data source.
	// Experimental.
	DataSourceType_SHAREPOINT DataSourceType = "SHAREPOINT"
	// Web Crawler data source.
	//
	// Extracts content from authorized public web pages using a crawler.
	// Experimental.
	DataSourceType_WEB_CRAWLER DataSourceType = "WEB_CRAWLER"
)

