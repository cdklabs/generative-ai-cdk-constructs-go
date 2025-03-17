package bedrock


// Configuration for custom orchestration of the agent.
// Experimental.
type CustomOrchestration struct {
	// The Lambda function to use for custom orchestration.
	// Experimental.
	Executor OrchestrationExecutor `field:"required" json:"executor" yaml:"executor"`
}

