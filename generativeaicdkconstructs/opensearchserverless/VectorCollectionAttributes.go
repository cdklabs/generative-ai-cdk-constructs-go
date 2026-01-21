package opensearchserverless


// Attributes for importing an existing vector collection.
// Experimental.
type VectorCollectionAttributes struct {
	// The ARN of the collection.
	// Experimental.
	CollectionArn *string `field:"required" json:"collectionArn" yaml:"collectionArn"`
	// The ID of the collection.
	// Experimental.
	CollectionId *string `field:"required" json:"collectionId" yaml:"collectionId"`
	// The name of the collection.
	// Experimental.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// The type of collection.
	// Experimental.
	CollectionType VectorCollectionType `field:"required" json:"collectionType" yaml:"collectionType"`
	// The standby replicas configuration.
	// Experimental.
	StandbyReplicas VectorCollectionStandbyReplicas `field:"required" json:"standbyReplicas" yaml:"standbyReplicas"`
}

