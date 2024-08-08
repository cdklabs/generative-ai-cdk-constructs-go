package opensearchserverless


// Experimental.
type TokenFilterType string

const (
	// Experimental.
	TokenFilterType_KUROMOJI_BASEFORM TokenFilterType = "KUROMOJI_BASEFORM"
	// Experimental.
	TokenFilterType_KUROMOJI_PART_OF_SPEECH TokenFilterType = "KUROMOJI_PART_OF_SPEECH"
	// Experimental.
	TokenFilterType_KUROMOJI_STEMMER TokenFilterType = "KUROMOJI_STEMMER"
	// Experimental.
	TokenFilterType_CJK_WIDTH TokenFilterType = "CJK_WIDTH"
	// Experimental.
	TokenFilterType_JA_STOP TokenFilterType = "JA_STOP"
	// Experimental.
	TokenFilterType_LOWERCASE TokenFilterType = "LOWERCASE"
	// Experimental.
	TokenFilterType_ICU_FOLDING TokenFilterType = "ICU_FOLDING"
)

