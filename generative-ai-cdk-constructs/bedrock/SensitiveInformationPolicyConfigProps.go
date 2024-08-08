package bedrock


// Experimental.
type SensitiveInformationPolicyConfigProps struct {
	// Experimental.
	Action PiiEntitiesConfigAction `field:"required" json:"action" yaml:"action"`
	// Experimental.
	Type interface{} `field:"required" json:"type" yaml:"type"`
}

