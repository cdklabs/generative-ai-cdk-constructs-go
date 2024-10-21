package bedrock


// Interface to define a Contextual Grounding Filter.
// Experimental.
type ContextualGroundingFilter struct {
	// The threshold for the contextual grounding filter.
	//
	// - `0` (blocks nothing)
	// - `0.99` (blocks almost everything)
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The type of contextual grounding filter.
	// Experimental.
	Type ContextualGroundingFilterType `field:"required" json:"type" yaml:"type"`
}

