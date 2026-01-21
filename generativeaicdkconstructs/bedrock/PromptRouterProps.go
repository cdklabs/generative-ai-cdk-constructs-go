package bedrock


// Experimental.
type PromptRouterProps struct {
	// Prompt Router Id.
	// Experimental.
	PromptRouterId *string `field:"required" json:"promptRouterId" yaml:"promptRouterId"`
	// The foundation models this router will route to.
	// Experimental.
	RoutingModels *[]BedrockFoundationModel `field:"required" json:"routingModels" yaml:"routingModels"`
}

