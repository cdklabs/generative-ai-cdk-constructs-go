package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// The properties for the LangchainLayerProps class.
// Experimental.
type LangchainLayerProps struct {
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
	// Required.
	//
	// Lambda function architecture compatible with this Layer.
	// Experimental.
	Architecture awslambda.Architecture `field:"required" json:"architecture" yaml:"architecture"`
	// Required.
	//
	// Lambda function runtime compatible with this Layer.
	// Experimental.
	Runtime awslambda.Runtime `field:"required" json:"runtime" yaml:"runtime"`
	// A prop allowing additional python pip libraries to be installed with this langchain layer.
	// Default: - none.
	//
	// Experimental.
	AdditionalPackages *[]*string `field:"optional" json:"additionalPackages" yaml:"additionalPackages"`
	// Optional: Add '--upgrade' to pip install requirements.txt In case of a LangchainCommonLayer, this parameter is not used.
	// Default: - false.
	//
	// Experimental.
	AutoUpgrade *bool `field:"optional" json:"autoUpgrade" yaml:"autoUpgrade"`
	// Optional: Local compute will be used when installing requirements.txt. By default, a docker container will be spun up to install requirements. To override this behavior, use the python alias string of `python` or `python3` The string value will be the python alias used to install requirements.
	// Default: - none.
	//
	// Experimental.
	Local *string `field:"optional" json:"local" yaml:"local"`
}

