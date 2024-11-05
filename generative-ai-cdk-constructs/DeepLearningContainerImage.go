package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type DeepLearningContainerImage interface {
	ContainerImage
	// Experimental.
	Bind(scope constructs.Construct, grantable awsiam.IGrantable) *ContainerImageConfig
}

// The jsii proxy struct for DeepLearningContainerImage
type jsiiProxy_DeepLearningContainerImage struct {
	jsiiProxy_ContainerImage
}

// Experimental.
func NewDeepLearningContainerImage(repositoryName *string, tag *string, accountId *string) DeepLearningContainerImage {
	_init_.Initialize()

	if err := validateNewDeepLearningContainerImageParameters(repositoryName, tag); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeepLearningContainerImage{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		[]interface{}{repositoryName, tag, accountId},
		&j,
	)

	return &j
}

// Experimental.
func NewDeepLearningContainerImage_Override(d DeepLearningContainerImage, repositoryName *string, tag *string, accountId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		[]interface{}{repositoryName, tag, accountId},
		d,
	)
}

// Experimental.
func DeepLearningContainerImage_FromAsset(directory *string, options *awsecrassets.DockerImageAssetOptions) ContainerImage {
	_init_.Initialize()

	if err := validateDeepLearningContainerImage_FromAssetParameters(directory, options); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"fromAsset",
		[]interface{}{directory, options},
		&returns,
	)

	return returns
}

// Experimental.
func DeepLearningContainerImage_FromDeepLearningContainerImage(repositoryName *string, tag *string, accountId *string) ContainerImage {
	_init_.Initialize()

	if err := validateDeepLearningContainerImage_FromDeepLearningContainerImageParameters(repositoryName, tag); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"fromDeepLearningContainerImage",
		[]interface{}{repositoryName, tag, accountId},
		&returns,
	)

	return returns
}

// Experimental.
func DeepLearningContainerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) ContainerImage {
	_init_.Initialize()

	if err := validateDeepLearningContainerImage_FromEcrRepositoryParameters(repository); err != nil {
		panic(err)
	}
	var returns ContainerImage

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_19_0_DEEPSPEED0_7_3_CU113() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_19_0_DEEPSPEED0_7_3_CU113",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_20_0_DEEPSPEED0_7_5_CU116() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_20_0_DEEPSPEED0_7_5_CU116",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_21_0_DEEPSPEED0_8_0_CU117() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_21_0_DEEPSPEED0_8_0_CU117",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_21_0_DEEPSPEED0_8_3_CU117() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_21_0_DEEPSPEED0_8_3_CU117",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_21_0_FASTERTRANSFORMER5_3_0_CU117() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_21_0_FASTERTRANSFORMER5_3_0_CU117",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_22_1_DEEPSPEED0_8_3_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_22_1_DEEPSPEED0_8_3_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_22_1_DEEPSPEED0_9_2_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_22_1_DEEPSPEED0_9_2_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_22_1_FASTERTRANSFORMER5_3_0_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_22_1_FASTERTRANSFORMER5_3_0_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_22_1_NEURONX_SDK2_10_0() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_22_1_NEURONX_SDK2_10_0",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_22_1_NEURONX_SDK2_9_0() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_22_1_NEURONX_SDK2_9_0",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_23_0_DEEPSPEED0_9_5_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_23_0_DEEPSPEED0_9_5_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_23_0_FASTERTRANSFORMER5_3_0_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_23_0_FASTERTRANSFORMER5_3_0_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_23_0_NEURONX_SDK2_12_0() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_23_0_NEURONX_SDK2_12_0",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_24_0_DEEPSPEED0_10_0_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_24_0_DEEPSPEED0_10_0_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_24_0_FASTERTRANSFORMER5_3_0_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_24_0_FASTERTRANSFORMER5_3_0_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_24_0_NEURONX_SDK2_14_1() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_24_0_NEURONX_SDK2_14_1",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_25_0_DEEPSPEED0_11_0_CU118() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_25_0_DEEPSPEED0_11_0_CU118",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_25_0_NEURONX_SDK2_15_0() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_25_0_NEURONX_SDK2_15_0",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_26_0_DEEPSPEED0_12_6_CU121() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_26_0_DEEPSPEED0_12_6_CU121",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_26_0_NEURONX_SDK2_16_0() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_26_0_NEURONX_SDK2_16_0",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_27_0_DEEPSPEED0_12_6_CU121() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_27_0_DEEPSPEED0_12_6_CU121",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_27_0_NEURONX_SDK2_18_1() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_27_0_NEURONX_SDK2_18_1",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_28_0_NEURONX_SDK2_18_2() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_28_0_NEURONX_SDK2_18_2",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_DJL_INFERENCE_0_29_0_NEURONX_SDK2_19_1() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"DJL_INFERENCE_0_29_0_NEURONX_SDK2_19_1",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_10_2_TRANSFORMERS4_17_0_CPU_PY38_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_10_2_TRANSFORMERS4_17_0_CPU_PY38_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_10_2_TRANSFORMERS4_17_0_GPU_PY38_CU113_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_10_2_TRANSFORMERS4_17_0_GPU_PY38_CU113_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_13_1_TRANSFORMERS4_26_0_CPU_PY39_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_13_1_TRANSFORMERS4_26_0_CPU_PY39_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_13_1_TRANSFORMERS4_26_0_GPU_PY39_CU117_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_13_1_TRANSFORMERS4_26_0_GPU_PY39_CU117_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_7_1_TRANSFORMERS4_6_1_CPU_PY36_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_7_1_TRANSFORMERS4_6_1_CPU_PY36_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_7_1_TRANSFORMERS4_6_1_GPU_PY36_CU110_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_7_1_TRANSFORMERS4_6_1_GPU_PY36_CU110_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_10_2_CPU_PY36_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_10_2_CPU_PY36_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_10_2_GPU_PY36_CU111_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_10_2_GPU_PY36_CU111_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_6_1_CPU_PY36_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_6_1_CPU_PY36_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_6_1_GPU_PY36_CU111_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_8_1_TRANSFORMERS4_6_1_GPU_PY36_CU111_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_10_2_CPU_PY38_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_10_2_CPU_PY38_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_10_2_GPU_PY38_CU111_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_10_2_GPU_PY38_CU111_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_11_0_CPU_PY38_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_11_0_CPU_PY38_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_11_0_GPU_PY38_CU111_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_9_0_TRANSFORMERS4_11_0_GPU_PY38_CU111_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_9_1_TRANSFORMERS4_12_3_CPU_PY38_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_9_1_TRANSFORMERS4_12_3_CPU_PY38_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_1_9_1_TRANSFORMERS4_12_3_GPU_PY38_CU111_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_1_9_1_TRANSFORMERS4_12_3_GPU_PY38_CU111_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_2_0_0_TRANSFORMERS4_28_1_CPU_PY310_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_2_0_0_TRANSFORMERS4_28_1_CPU_PY310_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_2_0_0_TRANSFORMERS4_28_1_GPU_PY310_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_2_0_0_TRANSFORMERS4_28_1_GPU_PY310_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_2_1_0_TRANSFORMERS4_37_0_CPU_PY310_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_2_1_0_TRANSFORMERS4_37_0_CPU_PY310_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_2_1_0_TRANSFORMERS4_37_0_GPU_PY310_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_2_1_0_TRANSFORMERS4_37_0_GPU_PY310_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_1_13_0_TRANSFORMERS4_28_1_NEURONX_PY38_SDK2_9_1_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_1_13_0_TRANSFORMERS4_28_1_NEURONX_PY38_SDK2_9_1_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_1_13_1_TRANSFORMERS4_34_1_NEURONX_PY310_SDK2_15_0_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_1_13_1_TRANSFORMERS4_34_1_NEURONX_PY310_SDK2_15_0_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_1_13_1_TRANSFORMERS4_36_2_NEURONX_PY310_SDK2_16_1_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_1_13_1_TRANSFORMERS4_36_2_NEURONX_PY310_SDK2_16_1_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_2_1_2_TRANSFORMERS4_36_2_NEURONX_PY310_SDK2_18_0_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_2_1_2_TRANSFORMERS4_36_2_NEURONX_PY310_SDK2_18_0_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_2_1_2_TRANSFORMERS4_41_1_NEURONX_PY310_SDK2_19_1_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_2_1_2_TRANSFORMERS4_41_1_NEURONX_PY310_SDK2_19_1_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_2_1_2_TRANSFORMERS4_43_2_NEURONX_PY310_SDK2_20_0_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_INFERENCE_NEURONX_2_1_2_TRANSFORMERS4_43_2_NEURONX_PY310_SDK2_20_0_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_0_TGI0_6_0_GPU_PY39_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_0_TGI0_6_0_GPU_PY39_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_0_TGI0_8_2_GPU_PY39_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_0_TGI0_8_2_GPU_PY39_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_1_TGI0_9_3_GPU_PY39_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_1_TGI0_9_3_GPU_PY39_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_1_TGI1_0_3_GPU_PY39_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_1_TGI1_0_3_GPU_PY39_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_1_TGI1_1_0_GPU_PY39_CU118_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_0_1_TGI1_1_0_GPU_PY39_CU118_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_2_0_GPU_PY310_CU121_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_2_0_GPU_PY310_CU121_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_3_1_GPU_PY310_CU121_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_3_1_GPU_PY310_CU121_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_3_3_GPU_PY310_CU121_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_3_3_GPU_PY310_CU121_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_4_0_GPU_PY310_CU121_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_4_0_GPU_PY310_CU121_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_4_2_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_4_2_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_4_5_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI1_4_5_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI2_0_0_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI2_0_0_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI2_0_1_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_1_1_TGI2_0_1_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_3_0_TGI2_0_2_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_3_0_TGI2_0_2_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_3_0_TGI2_0_3_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_3_0_TGI2_0_3_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_3_0_TGI2_2_0_GPU_PY310_CU121_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_3_0_TGI2_2_0_GPU_PY310_CU121_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_4_0_TGI2_3_1_GPU_PY311_CU124_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_4_0_TGI2_3_1_GPU_PY311_CU124_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_4_0_TGI2_4_0_GPU_PY311_CU124_UBUNTU22_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_PYTORCH_TGI_INFERENCE_2_4_0_TGI2_4_0_GPU_PY311_CU124_UBUNTU22_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_0_TRANSFORMERS4_26_0_CPU_PY39_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_0_TRANSFORMERS4_26_0_CPU_PY39_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_0_TRANSFORMERS4_26_0_GPU_PY39_CU112_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_0_TRANSFORMERS4_26_0_GPU_PY39_CU112_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_1_TRANSFORMERS4_26_0_CPU_PY39_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_1_TRANSFORMERS4_26_0_CPU_PY39_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_1_TRANSFORMERS4_26_0_GPU_PY39_CU112_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_11_1_TRANSFORMERS4_26_0_GPU_PY39_CU112_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_1_TRANSFORMERS4_6_1_CPU_PY37_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_1_TRANSFORMERS4_6_1_CPU_PY37_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_1_TRANSFORMERS4_6_1_GPU_PY37_CU110_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_1_TRANSFORMERS4_6_1_GPU_PY37_CU110_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_3_TRANSFORMERS4_10_2_CPU_PY37_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_3_TRANSFORMERS4_10_2_CPU_PY37_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_3_TRANSFORMERS4_10_2_GPU_PY37_CU110_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_4_3_TRANSFORMERS4_10_2_GPU_PY37_CU110_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_10_2_CPU_PY37_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_10_2_CPU_PY37_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_10_2_GPU_PY37_CU112_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_10_2_GPU_PY37_CU112_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_11_0_CPU_PY37_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_11_0_CPU_PY37_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_11_0_GPU_PY37_CU112_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_11_0_GPU_PY37_CU112_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_12_3_CPU_PY37_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_12_3_CPU_PY37_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_12_3_GPU_PY37_CU112_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_1_TRANSFORMERS4_12_3_GPU_PY37_CU112_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_3_TRANSFORMERS4_12_3_CPU_PY37_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_3_TRANSFORMERS4_12_3_CPU_PY37_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_3_TRANSFORMERS4_12_3_GPU_PY37_CU112_UBUNTU18_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_5_3_TRANSFORMERS4_12_3_GPU_PY37_CU112_UBUNTU18_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_6_3_TRANSFORMERS4_17_0_CPU_PY38_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_6_3_TRANSFORMERS4_17_0_CPU_PY38_UBUNTU20_04",
		&returns,
	)
	return returns
}

func DeepLearningContainerImage_HUGGINGFACE_TENSORFLOW_INFERENCE_2_6_3_TRANSFORMERS4_17_0_GPU_PY38_CU112_UBUNTU20_04() ContainerImage {
	_init_.Initialize()
	var returns ContainerImage
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		"HUGGINGFACE_TENSORFLOW_INFERENCE_2_6_3_TRANSFORMERS4_17_0_GPU_PY38_CU112_UBUNTU20_04",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DeepLearningContainerImage) Bind(scope constructs.Construct, grantable awsiam.IGrantable) *ContainerImageConfig {
	if err := d.validateBindParameters(scope, grantable); err != nil {
		panic(err)
	}
	var returns *ContainerImageConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{scope, grantable},
		&returns,
	)

	return returns
}

