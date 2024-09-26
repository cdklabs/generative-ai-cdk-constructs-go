package bedrock


// The different authentication types available to connect to your Confluence instance.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/confluence-data-source-connector.html#configuration-confluence-connector
//
// Experimental.
type ConfluenceDataSourceAuthType string

const (
	// Your secret authentication credentials in AWS Secrets Manager should include: - `confluenceAppKey` - `confluenceAppSecret` - `confluenceAccessToken` - `confluenceRefreshToken`.
	// Experimental.
	ConfluenceDataSourceAuthType_OAUTH2_CLIENT_CREDENTIALS ConfluenceDataSourceAuthType = "OAUTH2_CLIENT_CREDENTIALS"
	// Your secret authentication credentials in AWS Secrets Manager should include:  - `username` (email of admin account)  - `password` (API token).
	// Experimental.
	ConfluenceDataSourceAuthType_BASIC ConfluenceDataSourceAuthType = "BASIC"
)

