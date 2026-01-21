package bedrock


// Interface to define a Word Filter.
// Experimental.
type WordFilter struct {
	// The text to filter.
	// Experimental.
	Text *string `field:"required" json:"text" yaml:"text"`
	// The action to take when a word is detected in the input.
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the word filter is enabled for input.
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when a word is detected in the output.
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the word filter is enabled for output.
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}

