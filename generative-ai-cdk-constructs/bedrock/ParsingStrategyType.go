package bedrock


// Enum representing the types of parsing strategies available for Amazon Bedrock Knowledge Bases.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-advanced-parsing.html
//
// Experimental.
type ParsingStrategyType string

const (
	// Uses a Bedrock Foundation Model for advanced parsing of non-textual information from documents.
	// Experimental.
	ParsingStrategyType_FOUNDATION_MODEL ParsingStrategyType = "FOUNDATION_MODEL"
	// Processes multimodal data using Bedrock Data Automation (BDA).
	//
	// It leverages
	// generative AI to automate the transformation of multi-modal data into structured formats.
	// If you choose a foundation model or Amazon Bedrock Data Automation for parsing and it fails
	// to parse a file, the Amazon Bedrock default parser is used instead.
	// Experimental.
	ParsingStrategyType_DATA_AUTOMATION ParsingStrategyType = "DATA_AUTOMATION"
)

