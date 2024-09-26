package bedrock


// Represents the authentication types available for connecting to a SharePoint data source.
// Experimental.
type SharePointDataSourceAuthType string

const (
	// OAuth 2.0 Client Credentials flow for authentication with SharePoint. Your secret authentication credentials in AWS Secrets Manager should include: - `username`: The admin username for SharePoint authentication - `password`: The admin password associated with the username - `clientId`: The client ID (also known as application ID) - `clientSecret`: The client secret.
	// Experimental.
	SharePointDataSourceAuthType_OAUTH2_CLIENT_CREDENTIALS SharePointDataSourceAuthType = "OAUTH2_CLIENT_CREDENTIALS"
)

