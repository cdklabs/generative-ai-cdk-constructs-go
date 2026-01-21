package bedrock


// Enum for collaborator's relay conversation history types.
// Experimental.
type RelayConversationHistoryType string

const (
	// Sending to the collaborator.
	// Experimental.
	RelayConversationHistoryType_TO_COLLABORATOR RelayConversationHistoryType = "TO_COLLABORATOR"
	// Disabling relay of conversation history to the collaborator.
	// Experimental.
	RelayConversationHistoryType_DISABLED RelayConversationHistoryType = "DISABLED"
)

