package bedrock


// Interface to define a PII Filter.
// Experimental.
type PIIFilter struct {
	// The action to take when PII is detected.
	// Experimental.
	Action GuardrailAction `field:"required" json:"action" yaml:"action"`
	// The type of PII to filter.
	// Experimental.
	Type interface{} `field:"required" json:"type" yaml:"type"`
}

