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
	// The action to take when PII is detected in the input.
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the PII filter is enabled for input.
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when PII is detected in the output.
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the PII filter is enabled for output.
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}

