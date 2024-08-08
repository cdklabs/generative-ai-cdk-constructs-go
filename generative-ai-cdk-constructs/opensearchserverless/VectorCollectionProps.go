package opensearchserverless


// Experimental.
type VectorCollectionProps struct {
	// The name of the collection.
	// Experimental.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// Indicates whether to use standby replicas for the collection.
	// Default: ENABLED.
	//
	// Experimental.
	StandbyReplicas VectorCollectionStandbyReplicas `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
}

