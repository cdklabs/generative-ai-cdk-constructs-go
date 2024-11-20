package opensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Experimental.
type VectorCollectionProps struct {
	// The name of the collection.
	// Experimental.
	CollectionName *string `field:"required" json:"collectionName" yaml:"collectionName"`
	// A user defined IAM policy that allows API access to the collection.
	// Experimental.
	CustomAossPolicy awsiam.ManagedPolicy `field:"optional" json:"customAossPolicy" yaml:"customAossPolicy"`
	// Indicates whether to use standby replicas for the collection.
	// Default: ENABLED.
	//
	// Experimental.
	StandbyReplicas VectorCollectionStandbyReplicas `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
}

