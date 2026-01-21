package opensearchvectorindex

import (
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/opensearchserverless"
)

// Properties for the Analyzer.
// Experimental.
type Analyzer struct {
	// The analyzers to use.
	// Experimental.
	CharacterFilters *[]opensearchserverless.CharacterFilterType `field:"required" json:"characterFilters" yaml:"characterFilters"`
	// The token filters to use.
	// Experimental.
	TokenFilters *[]opensearchserverless.TokenFilterType `field:"required" json:"tokenFilters" yaml:"tokenFilters"`
	// The tokenizer to use.
	// Experimental.
	Tokenizer opensearchserverless.TokenizerType `field:"required" json:"tokenizer" yaml:"tokenizer"`
}

