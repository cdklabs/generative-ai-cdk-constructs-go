package bedrock


// Experimental.
type CrossRegionInferenceProfileRegion string

const (
	// Cross-region Inference Identifier for the European area.
	//
	// According to the model chosen, this might include:
	// - Frankfurt (`eu-central-1`)
	// - Ireland (`eu-west-1`)
	// - Paris (`eu-west-3`).
	// Experimental.
	CrossRegionInferenceProfileRegion_EU CrossRegionInferenceProfileRegion = "EU"
	// Cross-region Inference Identifier for the United States area.
	//
	// According to the model chosen, this might include:
	// - N. Virginia (`us-east-1`)
	// - Oregon (`us-west-2`)
	// - Ohio (`us-east-2`).
	// Experimental.
	CrossRegionInferenceProfileRegion_US CrossRegionInferenceProfileRegion = "US"
	// Cross-region Inference Identifier for the Asia-Pacific area.
	//
	// According to the model chosen, this might include:
	// - Tokyo (`ap-northeast-1`)
	// - Seoul (`ap-northeast-2`)
	// - Mumbai (`ap-south-1`)
	// - Singapore (`ap-southeast-1`)
	// - Sydney (`ap-southeast-2`).
	// Experimental.
	CrossRegionInferenceProfileRegion_APAC CrossRegionInferenceProfileRegion = "APAC"
)

