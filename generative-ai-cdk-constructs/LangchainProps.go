package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// LangchainProps.
// Experimental.
type LangchainProps struct {
	// The description the this Lambda Layer.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the layer.
	// Default: - A name will be generated.
	//
	// Experimental.
	LayerVersionName *string `field:"optional" json:"layerVersionName" yaml:"layerVersionName"`
	// The SPDX licence identifier or URL to the license file for this layer.
	// Default: - No license information will be recorded.
	//
	// Experimental.
	License *string `field:"optional" json:"license" yaml:"license"`
	// Whether to retain this version of the layer when a new version is added or when the stack is deleted.
	// Default: RemovalPolicy.DESTROY
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

