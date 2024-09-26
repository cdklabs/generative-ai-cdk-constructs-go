package bedrock


// Represents the authentication types available for connecting to a Salesforce data source.
// Experimental.
type SalesforceDataSourceAuthType string

const (
	// Your secret authentication credentials in AWS Secrets Manager should include: - `consumerKey` (app client ID) - `consumerSecret` (client secret) - `authenticationUrl`.
	// Experimental.
	SalesforceDataSourceAuthType_OAUTH2_CLIENT_CREDENTIALS SalesforceDataSourceAuthType = "OAUTH2_CLIENT_CREDENTIALS"
)

