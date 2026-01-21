package bedrock


// A Regular expression (regex) filter for sensitive information.
// Experimental.
type RegexFilter struct {
	// The action to take when a regex match is detected.
	// Experimental.
	Action GuardrailAction `field:"required" json:"action" yaml:"action"`
	// The name of the regex filter.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The regular expression pattern to match.
	// Experimental.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The description of the regex filter.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The action to take when a regex match is detected in the input.
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the regex filter is enabled for input.
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when a regex match is detected in the output.
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the regex filter is enabled for output.
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}

