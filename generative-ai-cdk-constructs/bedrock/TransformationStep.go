package bedrock


// Defines the step in the ingestion process where the custom transformation is applied.
// Experimental.
type TransformationStep string

const (
	// Processes documents after they have been converted into chunks.
	//
	// This allows for custom chunk-level metadata addition or custom post-chunking logic.
	// Experimental.
	TransformationStep_POST_CHUNKING TransformationStep = "POST_CHUNKING"
)

