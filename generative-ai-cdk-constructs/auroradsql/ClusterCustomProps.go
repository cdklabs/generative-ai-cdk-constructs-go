package auroradsql

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Aurora DSQL cluster resource.
// Experimental.
type ClusterCustomProps struct {
	// KMS key to use for the cluster.
	// Default: - A new KMS key is created.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Defines the structure for multi-Region cluster configurations, containing the witness Region and peered cluster settings.
	//
	// If not provided, the cluster will be created in the same region as the stack (single region cluster).
	// Default: - No multi-Region cluster configurations.
	//
	// Experimental.
	MultiRegionProperties *MultiRegionProperties `field:"optional" json:"multiRegionProperties" yaml:"multiRegionProperties"`
	// The removal policy for the cluster.
	//
	// Only RemovalPolicy.DESTROY and RemovalPolicy.RETAIN are allowed.
	// Default: - RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Tags to apply to the cluster.
	// Default: - No tags.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

