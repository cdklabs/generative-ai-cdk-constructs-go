package bedrock


// Experimental.
type ContextualGroundingFilterConfigType string

const (
	// Grounding score represents the confidence that the model response is factually correct and grounded in the source.
	//
	// If the model response has a lower score than the defined threshold, the response will be blocked and the configured
	// blocked message will be returned to the user. A higher threshold level blocks more responses.
	// Experimental.
	ContextualGroundingFilterConfigType_GROUNDING ContextualGroundingFilterConfigType = "GROUNDING"
	// Relevance score represents the confidence that the model response is relevant to the user's query.
	//
	// If the model response has a lower score than the defined threshold, the response will be blocked and
	// the configured blocked message will be returned to the user. A higher threshold level blocks more responses
	// Experimental.
	ContextualGroundingFilterConfigType_RELEVANCE ContextualGroundingFilterConfigType = "RELEVANCE"
)

