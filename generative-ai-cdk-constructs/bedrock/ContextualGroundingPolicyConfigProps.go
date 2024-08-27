package bedrock


// Experimental.
type ContextualGroundingPolicyConfigProps struct {
	// The filter details for the guardrails contextual grounding filter.
	//
	// GROUNDING: Validate if the model responses are grounded and factually correct based on the information provided in the reference source,
	// and block responses that are below the defined threshold of grounding.
	// RELEVANCE: Validate if the model responses are relevant to the user's query and block responses
	// that are below the defined threshold of relevance.
	// Experimental.
	FiltersConfigType ContextualGroundingFilterConfigType `field:"required" json:"filtersConfigType" yaml:"filtersConfigType"`
	// The threshold details for the guardrails contextual grounding filter.
	//
	// 0 blocks nothing, 0.99 blocks almost everything
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

