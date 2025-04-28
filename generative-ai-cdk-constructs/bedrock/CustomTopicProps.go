package bedrock


// Interface for creating a custom Topic.
// Experimental.
type CustomTopicProps struct {
	// Provide a clear definition to detect and block user inputs and FM responses that fall into this topic.
	//
	// Avoid starting with "don't".
	//
	// Example:
	//   `Investment advice refers to inquiries, guidance, or recommendations
	//   regarding the management or allocation of funds or assets with the goal of
	//   generating returns or achieving specific financial objectives.`
	//
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Representative phrases that refer to the topic.
	//
	// These phrases can represent
	// a user input or a model response. Add up to 5 phrases, up to 100 characters
	// each.
	//
	// Example:
	//   "Where should I invest my money?"
	//
	// Experimental.
	Examples *[]*string `field:"required" json:"examples" yaml:"examples"`
	// The name of the topic to deny.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to take when a topic is detected in the input.
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the topic filter is enabled for input.
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when a topic is detected in the output.
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the topic filter is enabled for output.
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}

