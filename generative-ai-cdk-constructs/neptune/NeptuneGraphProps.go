package neptune


// Properties for creating a new Neptune Graph database.
// Experimental.
type NeptuneGraphProps struct {
	// Indicates whether the Graph should have deletion protection enabled.
	// Default: false.
	//
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The graph name.
	//
	// The name must contain from 1 to 63 letters, numbers, or hyphens, and its first character must be a letter.
	// It cannot end with a hyphen or contain two consecutive hyphens.
	// Default: - A unique graph name is generated for you using the prefix `graph-for-${StackName}-${UUID}`.
	//
	// Experimental.
	GraphName *string `field:"optional" json:"graphName" yaml:"graphName"`
	// The amount of memory (in Neptune Capacity Units m-NCUs) to use for the graph.
	// Default: 16.
	//
	// Experimental.
	ProvisionedMemoryNCUs *float64 `field:"optional" json:"provisionedMemoryNCUs" yaml:"provisionedMemoryNCUs"`
	// Specifies whether or not the graph can be reachable over the internet.
	//
	// All access to graphs is IAM authenticated.
	// When the graph is publicly available, its domain name system (DNS) endpoint resolves to the public IP address from the internet.
	// When the graph isn't publicly available, you need to create a PrivateGraphEndpoint in a given VPC to ensure the DNS name
	// resolves to a private IP address that is reachable from the VPC.
	// Default: true.
	//
	// Experimental.
	PublicConnectivity *bool `field:"optional" json:"publicConnectivity" yaml:"publicConnectivity"`
	// The number of replicas in other AZs.
	//
	// Replicas are typically only needed for production critical workloads with strict availability requirements.
	// **Note: Each replica incurs additional cost as it maintains a full copy of the graph data.**
	// Default: 0.
	//
	// Experimental.
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// The Dimension used to save vectors.
	// Experimental.
	VectorSearchDimension *float64 `field:"optional" json:"vectorSearchDimension" yaml:"vectorSearchDimension"`
}

