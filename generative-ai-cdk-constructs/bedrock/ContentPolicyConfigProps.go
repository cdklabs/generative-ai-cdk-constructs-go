package bedrock


// Experimental.
type ContentPolicyConfigProps struct {
	// Experimental.
	FiltersConfigType FiltersConfigType `field:"required" json:"filtersConfigType" yaml:"filtersConfigType"`
	// Experimental.
	InputStrength FiltersConfigStrength `field:"optional" json:"inputStrength" yaml:"inputStrength"`
	// Experimental.
	OutputStrength FiltersConfigStrength `field:"optional" json:"outputStrength" yaml:"outputStrength"`
}

