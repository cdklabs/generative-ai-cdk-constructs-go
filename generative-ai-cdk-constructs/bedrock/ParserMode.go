package bedrock


// Specifies whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the promptType.
//
// If you set the field as OVERRIDEN, the
// overrideLambda field in the PromptOverrideConfiguration must be specified
// with the ARN of a Lambda function.
// Experimental.
type ParserMode string

const (
	// Experimental.
	ParserMode_DEFAULT ParserMode = "DEFAULT"
	// Experimental.
	ParserMode_OVERRIDDEN ParserMode = "OVERRIDDEN"
)

