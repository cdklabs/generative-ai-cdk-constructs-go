package opensearchserverless


// Configuration for standby replicas in a vector collection.
// Experimental.
type VectorCollectionStandbyReplicas string

const (
	// Enable standby replicas for high availability.
	// Experimental.
	VectorCollectionStandbyReplicas_ENABLED VectorCollectionStandbyReplicas = "ENABLED"
	// Disable standby replicas to reduce costs.
	// Experimental.
	VectorCollectionStandbyReplicas_DISABLED VectorCollectionStandbyReplicas = "DISABLED"
)

