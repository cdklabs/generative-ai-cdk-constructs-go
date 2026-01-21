package sagemakerdeployment


// Experimental.
type AsyncInferenceConfig struct {
	// Experimental.
	FailurePath *string `field:"required" json:"failurePath" yaml:"failurePath"`
	// Experimental.
	OutputPath *string `field:"required" json:"outputPath" yaml:"outputPath"`
	// Experimental.
	MaxConcurrentInvocationsPerInstance *float64 `field:"optional" json:"maxConcurrentInvocationsPerInstance" yaml:"maxConcurrentInvocationsPerInstance"`
}

