package auroradsql


// Interface for multi-region cluster properties.
// Experimental.
type MultiRegionProperties struct {
	// The Region that serves as the witness Region for a multi-Region cluster.
	//
	// The witness Region helps maintain cluster consistency and quorum.
	// The witness Region receives data written to any Read-Write Region
	// but does not have an endpoint.
	// Experimental.
	WitnessRegion *string `field:"required" json:"witnessRegion" yaml:"witnessRegion"`
	// The set of peered clusters that form the multi-Region cluster configuration.
	//
	// Each peered cluster represents a database instance in a different Region.
	// Default: - No peered clusters (single region cluster).
	//
	// Experimental.
	Clusters *[]ICluster `field:"optional" json:"clusters" yaml:"clusters"`
}

