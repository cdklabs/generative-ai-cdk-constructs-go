package opensearchserverless


// TokenFilterType defines the available token filters for text analysis.
//
// Token filters process tokens after they have been created by the tokenizer.
// They can modify, add, or remove tokens based on specific rules.
// Experimental.
type TokenFilterType string

const (
	// Converts inflected Japanese words to their base form.
	// Experimental.
	TokenFilterType_KUROMOJI_BASEFORM TokenFilterType = "KUROMOJI_BASEFORM"
	// Tags words with their parts of speech in Japanese text analysis.
	// Experimental.
	TokenFilterType_KUROMOJI_PART_OF_SPEECH TokenFilterType = "KUROMOJI_PART_OF_SPEECH"
	// Reduces Japanese words to their stem form.
	// Experimental.
	TokenFilterType_KUROMOJI_STEMMER TokenFilterType = "KUROMOJI_STEMMER"
	// Normalizes CJK width differences by converting all characters to their fullwidth or halfwidth variants.
	// Experimental.
	TokenFilterType_CJK_WIDTH TokenFilterType = "CJK_WIDTH"
	// Removes Japanese stop words from text.
	// Experimental.
	TokenFilterType_JA_STOP TokenFilterType = "JA_STOP"
	// Converts all characters to lowercase.
	// Experimental.
	TokenFilterType_LOWERCASE TokenFilterType = "LOWERCASE"
	// Applies Unicode folding rules for better text matching.
	// Experimental.
	TokenFilterType_ICU_FOLDING TokenFilterType = "ICU_FOLDING"
	// Tags words with their parts of speech in Korean text analysis.
	// Experimental.
	TokenFilterType_NORI_PART_OF_SPEECH TokenFilterType = "NORI_PART_OF_SPEECH"
	// Converts Korean text to its reading form.
	// Experimental.
	TokenFilterType_NORI_READINGFORM TokenFilterType = "NORI_READINGFORM"
	// Normalizes Korean numbers to regular Arabic numbers.
	// Experimental.
	TokenFilterType_NORI_NUMBER TokenFilterType = "NORI_NUMBER"
)

