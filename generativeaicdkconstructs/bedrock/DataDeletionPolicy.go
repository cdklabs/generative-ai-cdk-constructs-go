package bedrock


// Specifies the policy for handling data when a data source resource is deleted.
//
// This policy affects the vector embeddings created from the data source.
// Experimental.
type DataDeletionPolicy string

const (
	// Deletes all vector embeddings derived from the data source upon deletion of a data source resource.
	// Experimental.
	DataDeletionPolicy_DELETE DataDeletionPolicy = "DELETE"
	// Retains all vector embeddings derived from the data source even after deletion of a data source resource.
	// Experimental.
	DataDeletionPolicy_RETAIN DataDeletionPolicy = "RETAIN"
)

