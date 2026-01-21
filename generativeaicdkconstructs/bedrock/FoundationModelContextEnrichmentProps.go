package bedrock


// Properties for configuring a Foundation Model enrichment strategy.
// Experimental.
type FoundationModelContextEnrichmentProps struct {
	// The Bedrock Foundation Model configuration for context enrichment.
	// Experimental.
	EnrichmentModel IInvokable `field:"required" json:"enrichmentModel" yaml:"enrichmentModel"`
}

