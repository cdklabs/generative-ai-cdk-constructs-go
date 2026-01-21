package bedrock


// Enum representing the type of context enrichment.
// Experimental.
type ContextEnrichmentType string

const (
	// Uses a Bedrock Foundation Model for context enrichment.
	// Experimental.
	ContextEnrichmentType_BEDROCK_FOUNDATION_MODEL ContextEnrichmentType = "BEDROCK_FOUNDATION_MODEL"
)

