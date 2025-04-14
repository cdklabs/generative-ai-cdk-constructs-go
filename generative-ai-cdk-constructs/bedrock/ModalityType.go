package bedrock


// The type of modality that can be used in content filters.
// Experimental.
type ModalityType string

const (
	// Text modality for content filters.
	// Experimental.
	ModalityType_TEXT ModalityType = "TEXT"
	// Image modality for content filters.
	// Experimental.
	ModalityType_IMAGE ModalityType = "IMAGE"
)

