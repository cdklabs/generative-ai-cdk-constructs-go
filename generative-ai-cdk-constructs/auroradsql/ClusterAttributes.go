package auroradsql

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes for specifying an imported Aurora DSQL cluster.
// Experimental.
type ClusterAttributes struct {
	// The ARN of the cluster.
	// Experimental.
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// VpcEndpointServiceName of the cluster.
	// Experimental.
	VpcEndpointServiceName *string `field:"required" json:"vpcEndpointServiceName" yaml:"vpcEndpointServiceName"`
	// The timestamp when the cluster was created, in ISO 8601 format.
	// Default: undefined - No creation time is provided.
	//
	// Experimental.
	CreationTime *string `field:"optional" json:"creationTime" yaml:"creationTime"`
	// KMS encryption key associated with this cluster.
	// Default: - no encryption key.
	//
	// Experimental.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The status of the cluster.
	// Default: undefined - No status is provided.
	//
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

