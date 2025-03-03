package bedrock


// Experimental.
type SessionSummaryMemoryProps struct {
	// Maximum number of recent session summaries to include (min 1).
	// Default: 20.
	//
	// Experimental.
	MaxRecentSessions *float64 `field:"optional" json:"maxRecentSessions" yaml:"maxRecentSessions"`
	// Duration in days for which session summaries are retained (1-365).
	// Default: 30.
	//
	// Experimental.
	MemoryDurationDays *float64 `field:"optional" json:"memoryDurationDays" yaml:"memoryDurationDays"`
}

