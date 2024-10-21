package bedrock


// Guardrail action when a sensitive entity is detected.
// Experimental.
type GuardrailAction string

const (
	// If sensitive information is detected in the prompt or response, the guardrail blocks all the content and returns a message that you configure.
	// Experimental.
	GuardrailAction_BLOCK GuardrailAction = "BLOCK"
	// If sensitive information is detected in the model response, the guardrail masks it with an identifier, the sensitive information is masked and replaced with identifier tags (for example: [NAME-1], [NAME-2], [EMAIL-1], etc.).
	// Experimental.
	GuardrailAction_ANONYMIZE GuardrailAction = "ANONYMIZE"
)

