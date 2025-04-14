package bedrock


// Interface to declare a content filter.
// Experimental.
type ContentFilter struct {
	// The strength of the content filter to apply to prompts / user input.
	// Experimental.
	InputStrength ContentFilterStrength `field:"required" json:"inputStrength" yaml:"inputStrength"`
	// The strength of the content filter to apply to model responses.
	// Experimental.
	OutputStrength ContentFilterStrength `field:"required" json:"outputStrength" yaml:"outputStrength"`
	// The type of harmful category that the content filter is applied to.
	// Experimental.
	Type ContentFilterType `field:"required" json:"type" yaml:"type"`
	// The input modalities to apply the content filter to.
	// Default: undefined - Applies to text modality.
	//
	// Experimental.
	InputModalities *[]ModalityType `field:"optional" json:"inputModalities" yaml:"inputModalities"`
	// The output modalities to apply the content filter to.
	// Default: undefined - Applies to text modality.
	//
	// Experimental.
	OutputModalities *[]ModalityType `field:"optional" json:"outputModalities" yaml:"outputModalities"`
}

