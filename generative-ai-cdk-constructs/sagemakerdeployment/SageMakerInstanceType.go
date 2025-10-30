package sagemakerdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// Supported instance types for SageMaker instance-based production variants.
// Experimental.
type SageMakerInstanceType interface {
	// Return the instance type as a string.
	//
	// Returns: The instance type as a string.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SageMakerInstanceType
type jsiiProxy_SageMakerInstanceType struct {
	_ byte // padding
}

// Experimental.
func NewSageMakerInstanceType(instanceType *string) SageMakerInstanceType {
	_init_.Initialize()

	if err := validateNewSageMakerInstanceTypeParameters(instanceType); err != nil {
		panic(err)
	}
	j := jsiiProxy_SageMakerInstanceType{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		[]interface{}{instanceType},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerInstanceType_Override(s SageMakerInstanceType, instanceType *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		[]interface{}{instanceType},
		s,
	)
}

// Builds an InstanceType from a given string or token (such as a CfnParameter).
//
// Returns: A strongly typed InstanceType.
// Experimental.
func SageMakerInstanceType_Of(instanceType *string) SageMakerInstanceType {
	_init_.Initialize()

	if err := validateSageMakerInstanceType_OfParameters(instanceType); err != nil {
		panic(err)
	}
	var returns SageMakerInstanceType

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"of",
		[]interface{}{instanceType},
		&returns,
	)

	return returns
}

func SageMakerInstanceType_ML_C4_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C4_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C4_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C4_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C4_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C4_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C4_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C4_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C4_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C4_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5_18XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5_18XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5_9XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5_9XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5D_18XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5D_18XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5D_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5D_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5D_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5D_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5D_9XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5D_9XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5D_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5D_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C5D_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C5D_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_16XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_16XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_32XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_32XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_C6I_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_C6I_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G4DN_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G4DN_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G4DN_16XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G4DN_16XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G4DN_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G4DN_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G4DN_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G4DN_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G4DN_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G4DN_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G4DN_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G4DN_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_16XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_16XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_48XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_48XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_G5_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_G5_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF1_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF1_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF1_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF1_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF1_6XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF1_6XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF1_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF1_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF2_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF2_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF2_48XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF2_48XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF2_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF2_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_INF2_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_INF2_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M4_10XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M4_10XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M4_16XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M4_16XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M4_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M4_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M4_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M4_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M4_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M4_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5D_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5D_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5D_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5D_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5D_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5D_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5D_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5D_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5D_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5D_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_M5D_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_M5D_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P2_16XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P2_16XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P2_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P2_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P2_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P2_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P3_16XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P3_16XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P3_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P3_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P3_8XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P3_8XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_P4D_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_P4D_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5D_12XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5D_12XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5D_24XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5D_24XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5D_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5D_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5D_4XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5D_4XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5D_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5D_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_R5D_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_R5D_XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_T2_2XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_T2_2XLARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_T2_LARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_T2_LARGE",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_T2_MEDIUM() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_T2_MEDIUM",
		&returns,
	)
	return returns
}

func SageMakerInstanceType_ML_T2_XLARGE() SageMakerInstanceType {
	_init_.Initialize()
	var returns SageMakerInstanceType
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		"ML_T2_XLARGE",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SageMakerInstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

