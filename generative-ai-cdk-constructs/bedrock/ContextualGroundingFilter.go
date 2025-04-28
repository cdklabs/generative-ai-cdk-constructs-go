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
	// The action to take when contextual grounding is detected.
	// Experimental.
	Action GuardrailAction `field:"optional" json:"action" yaml:"action"`
	// Whether the contextual grounding filter is enabled.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}

