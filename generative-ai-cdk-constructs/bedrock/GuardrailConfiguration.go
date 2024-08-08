package bedrock


// Details about the guardrail associated with the agent.
// Experimental.
type GuardrailConfiguration struct {
	// Experimental.
	GuardrailId *string `field:"optional" json:"guardrailId" yaml:"guardrailId"`
	// Experimental.
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
}

