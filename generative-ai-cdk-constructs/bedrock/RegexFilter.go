package bedrock


// A Regular expression (regex) filter for sensitive information.
//
// Example:
//   const regexFilter: RegexFilter = {
//     name: "my-custom-filter",
//     action: SensitiveInfoGuardrailAction.BLOCK,
//     pattern: "\\b(?:\\d{1,3}\\.){3}\\d{1,3}\\b",
//   };
//
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
}

