package bedrock


// Properties for configuring a Foundation Model parsing strategy.
// Experimental.
type FoundationModelParsingStategyProps struct {
	// The Foundation Model to use for parsing non-textual information.
	//
	// Currently supported models are Claude 3 Sonnet and Claude 3 Haiku.
	// Experimental.
	ParsingModel IInvokable `field:"required" json:"parsingModel" yaml:"parsingModel"`
	// Custom prompt to instruct the parser on how to interpret the document.
	// Default: - Uses the default instruction prompt as provided in the AWS Console.
	//
	// Experimental.
	ParsingPrompt *string `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

