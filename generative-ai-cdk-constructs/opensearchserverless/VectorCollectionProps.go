package opensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for configuring the vector collection.
// Experimental.
type VectorCollectionProps struct {
	// The name of the collection.
	//
	// Must be between 3-32 characters long and contain only
	// lowercase letters, numbers, and hyphens.
	// Default: - A CDK generated name will be used.
	//
	// Experimental.
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// Type of vector collection.
	// Default: - VECTORSEARCH.
	//
	// Experimental.
	CollectionType VectorCollectionType `field:"optional" json:"collectionType" yaml:"collectionType"`
	// A user defined IAM policy that allows API access to the collection.
	// Experimental.
	CustomAossPolicy awsiam.ManagedPolicy `field:"optional" json:"customAossPolicy" yaml:"customAossPolicy"`
	// Description for the collection.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether to use standby replicas for the collection.
	// Default: VectorCollectionStandbyReplicas.ENABLED
	//
	// Experimental.
	StandbyReplicas VectorCollectionStandbyReplicas `field:"optional" json:"standbyReplicas" yaml:"standbyReplicas"`
	// A list of tags associated with the inference profile.
	// Experimental.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

