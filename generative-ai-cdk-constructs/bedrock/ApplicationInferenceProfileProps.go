package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for creating a ApplicationInferenceProfile.
// Experimental.
type ApplicationInferenceProfileProps struct {
	// The name of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-inferenceprofilename
	//
	// Experimental.
	InferenceProfileName *string `field:"required" json:"inferenceProfileName" yaml:"inferenceProfileName"`
	// To create an application inference profile for one Region, specify a foundation model.
	//
	// Usage and costs for requests made to that Region with that model will be tracked.
	//
	// To create an application inference profile for multiple Regions,
	// specify a cross region (system-defined) inference profile.
	// The inference profile will route requests to the Regions defined in
	// the cross region (system-defined) inference profile that you choose.
	// Usage and costs for requests made to the Regions in the inference profile will be tracked.
	// Experimental.
	ModelSource IInvokable `field:"required" json:"modelSource" yaml:"modelSource"`
	// Description of the inference profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-applicationinferenceprofile.html#cfn-bedrock-applicationinferenceprofile-description
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags associated with the inference profile.
	// Experimental.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

