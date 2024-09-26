package bedrock


// Represents the SharePoint object types that can be accessed by the data source connector.
// Experimental.
type SharePointObjectType string

const (
	// Represents a SharePoint page, which typically contains web parts and content.
	// Experimental.
	SharePointObjectType_PAGE SharePointObjectType = "PAGE"
	// Represents a calendar event in SharePoint.
	// Experimental.
	SharePointObjectType_EVENT SharePointObjectType = "EVENT"
	// Represents a file stored in SharePoint document libraries.
	// Experimental.
	SharePointObjectType_FILE SharePointObjectType = "FILE"
)

