package bedrock


// ****************************************************************************                       PROPS FOR NEW CONSTRUCT ***************************************************************************.
// Experimental.
type CrossRegionInferenceProfileProps struct {
	// The geographic region where the traffic is going to be distributed.
	//
	// Routing
	// factors in user traffic, demand and utilization of resources.
	// Experimental.
	GeoRegion CrossRegionInferenceProfileRegion `field:"required" json:"geoRegion" yaml:"geoRegion"`
	// A model supporting cross-region inference.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference-support.html
	//
	// Experimental.
	Model BedrockFoundationModel `field:"required" json:"model" yaml:"model"`
}

