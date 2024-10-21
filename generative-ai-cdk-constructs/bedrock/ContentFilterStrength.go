package bedrock


// The strength of the content filter.
//
// As you increase the filter strength,
// the likelihood of filtering harmful content increases and the probability
// of seeing harmful content in your application reduces.
// Experimental.
type ContentFilterStrength string

const (
	// Experimental.
	ContentFilterStrength_NONE ContentFilterStrength = "NONE"
	// Experimental.
	ContentFilterStrength_LOW ContentFilterStrength = "LOW"
	// Experimental.
	ContentFilterStrength_MEDIUM ContentFilterStrength = "MEDIUM"
	// Experimental.
	ContentFilterStrength_HIGH ContentFilterStrength = "HIGH"
)

