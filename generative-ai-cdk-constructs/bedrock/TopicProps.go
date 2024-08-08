package bedrock


// Experimental.
type TopicProps struct {
	// Experimental.
	Definition *string `field:"optional" json:"definition" yaml:"definition"`
	// Experimental.
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

