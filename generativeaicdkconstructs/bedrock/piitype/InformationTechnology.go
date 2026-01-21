package piitype


// Types of PII in the domain of IT (Information Technology).
// Experimental.
type InformationTechnology string

const (
	// A web address, such as www.example.com.
	// Experimental.
	InformationTechnology_URL InformationTechnology = "URL"
	// An IPv4 address, such as 198.51.100.0.
	// Experimental.
	InformationTechnology_IP_ADDRESS InformationTechnology = "IP_ADDRESS"
	// A media access control (MAC) address assigned to a network interface.
	// Experimental.
	InformationTechnology_MAC_ADDRESS InformationTechnology = "MAC_ADDRESS"
	// A unique identifier that's associated with a secret access key.
	//
	// You use
	// the access key ID and secret access key to sign programmatic AWS requests
	// cryptographically.
	// Experimental.
	InformationTechnology_AWS_ACCESS_KEY InformationTechnology = "AWS_ACCESS_KEY"
	// A unique identifier that's associated with a secret access key.
	//
	// You use
	// the access key ID and secret access key to sign programmatic AWS requests
	// cryptographically.
	// Experimental.
	InformationTechnology_AWS_SECRET_KEY InformationTechnology = "AWS_SECRET_KEY"
)

