package opensearchserverless


// Experimental.
type TokenizerType string

const (
	// Kuromoji tokenizer is used for Japanese text analysis and segmentation.
	// Experimental.
	TokenizerType_KUROMOJI_TOKENIZER TokenizerType = "KUROMOJI_TOKENIZER"
	// ICU tokenizer is used for Unicode text segmentation based on UAX #29 rules.
	// Experimental.
	TokenizerType_ICU_TOKENIZER TokenizerType = "ICU_TOKENIZER"
	// Nori tokenizer is used for Korean text analysis and segmentation.
	// Experimental.
	TokenizerType_NORI_TOKENIZER TokenizerType = "NORI_TOKENIZER"
)

