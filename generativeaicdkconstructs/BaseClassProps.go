package generativeaicdkconstructs


// Experimental.
type BaseClassProps struct {
	// construct id.
	// Experimental.
	ConstructId *string `field:"required" json:"constructId" yaml:"constructId"`
	// name of the construct.
	// Experimental.
	ConstructName ConstructName `field:"required" json:"constructName" yaml:"constructName"`
	// Enable observability.
	//
	// Warning: associated cost with the services
	// used. Best practice to enable by default.
	// Default: true.
	//
	// Experimental.
	Observability *bool `field:"optional" json:"observability" yaml:"observability"`
	// Value will be appended to resources name.
	// Default: _dev.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
}

