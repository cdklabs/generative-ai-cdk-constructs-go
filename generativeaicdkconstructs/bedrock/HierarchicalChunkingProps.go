package bedrock


// Experimental.
type HierarchicalChunkingProps struct {
	// Maximum number of tokens that a child chunk can contain.
	//
	// Keep in mind the maximum chunk size depends on the embedding model chosen.
	// Experimental.
	MaxChildTokenSize *float64 `field:"required" json:"maxChildTokenSize" yaml:"maxChildTokenSize"`
	// Maximum number of tokens that a parent chunk can contain.
	//
	// Keep in mind the maximum chunk size depends on the embedding model chosen.
	// Experimental.
	MaxParentTokenSize *float64 `field:"required" json:"maxParentTokenSize" yaml:"maxParentTokenSize"`
	// The overlap tokens between adjacent chunks.
	// Experimental.
	OverlapTokens *float64 `field:"required" json:"overlapTokens" yaml:"overlapTokens"`
}

