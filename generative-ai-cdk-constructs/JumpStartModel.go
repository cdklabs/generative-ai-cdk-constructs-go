package generative-ai-cdk-constructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// Experimental.
type JumpStartModel interface {
	// Experimental.
	Bind() IJumpStartModelSpec
}

// The jsii proxy struct for JumpStartModel
type jsiiProxy_JumpStartModel struct {
	_ byte // padding
}

// Experimental.
func NewJumpStartModel(name *string) JumpStartModel {
	_init_.Initialize()

	if err := validateNewJumpStartModelParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_JumpStartModel{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		[]interface{}{name},
		&j,
	)

	return &j
}

// Experimental.
func NewJumpStartModel_Override(j JumpStartModel, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		[]interface{}{name},
		j,
	)
}

// Experimental.
func JumpStartModel_Of(name *string) JumpStartModel {
	_init_.Initialize()

	if err := validateJumpStartModel_OfParameters(name); err != nil {
		panic(err)
	}
	var returns JumpStartModel

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"of",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_BASE_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_BASE_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_V2_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_V2_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_V2_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_V2_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_LARGE_V3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_LARGE_V3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_MEDIUM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_MEDIUM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_MEDIUM_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_MEDIUM_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_MEDIUM_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_MEDIUM_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_MEDIUM_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_MEDIUM_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_SMALL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_SMALL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_SMALL_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_SMALL_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_SMALL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_SMALL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_SMALL_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_SMALL_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_TINY_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_TINY_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_TINY_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_TINY_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_TINY_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_TINY_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ASR_WHISPER_TINY_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ASR_WHISPER_TINY_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_MULTILINGUAL_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_CASED_WHOLE_WORD_MASKING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_MULTILINGUAL_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILBERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILROBERTA_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILROBERTA_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILROBERTA_BASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILROBERTA_BASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_DISTILROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_DISTILROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_BASE_OPENAI_DETECTOR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_LARGE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_LARGE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_EQA_ROBERTA_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_EQA_ROBERTA_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_FILLMASK_BERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_FILLMASK_BERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_FILLMASK_BERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_FILLMASK_BERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AHXT_LITELLAMA_460M_1T_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AHXT_LITELLAMA_460M_1T_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AHXT_LITELLAMA_460M_1T_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AHXT_LITELLAMA_460M_1T_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AI_FOREVER_MGPT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AI_FOREVER_MGPT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AI_FOREVER_MGPT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AI_FOREVER_MGPT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ALPINDALE_WIZARD_LM_2_8_22B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ALPINDALE_WIZARD_LM_2_8_22B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AMAZON_FALCONLITE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AMAZON_FALCONLITE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AMAZON_FALCONLITE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AMAZON_FALCONLITE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AMAZON_FALCONLITE2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AMAZON_FALCONLITE2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AMAZON_FALCONLITE2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AMAZON_FALCONLITE2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AMAZON_MISTRALLITE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AMAZON_MISTRALLITE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AMAZON_MISTRALLITE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AMAZON_MISTRALLITE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_AYA_101_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_AYA_101_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BAICHUAN2_7B_BASE_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BAICHUAN2_7B_BASE_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BAICHUAN2_7B_BASE_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BAICHUAN2_7B_BASE_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BERKELEY_NEST_STARLING_LM_7B_ALPHA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BERKELEY_NEST_STARLING_LM_7B_ALPHA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BERKELEY_NEST_STARLING_LM_7B_ALPHA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BERKELEY_NEST_STARLING_LM_7B_ALPHA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_BILINGUAL_RINNA_4B_INSTRUCTION_PPO_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CALM2_7B_CHAT_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_COGNITIVE_DOLPHIN_29_LLAMA3_8B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_COGNITIVE_DOLPHIN_29_LLAMA3_8B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_COHEREFORAI_C4AI_COMMAND_R_PLUS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_COHEREFORAI_C4AI_COMMAND_R_PLUS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_COHEREFORAI_C4AI_COMMAND_R_PLUS_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_COHEREFORAI_C4AI_COMMAND_R_PLUS_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CULTRIX_MISTRALTRIX_V1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CULTRIX_MISTRALTRIX_V1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_CULTRIX_MISTRALTRIX_V1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_CULTRIX_MISTRALTRIX_V1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DBRX_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DBRX_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DBRX_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DBRX_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DBRX_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DBRX_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DBRX_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DBRX_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DOLPHIN_2_2_1_MISTRAL_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DOLPHIN_2_2_1_MISTRAL_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DOLPHIN_2_2_1_MISTRAL_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DOLPHIN_2_2_1_MISTRAL_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DOLPHIN_2_5_MIXTRAL_8X7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DOLPHIN_2_5_MIXTRAL_8X7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DOLPHIN_2_5_MIXTRAL_8X7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DOLPHIN_2_5_MIXTRAL_8X7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DOLPHIN_2_7_MIXTRAL_8X7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DOLPHIN_2_7_MIXTRAL_8X7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_DOLPHIN_2_7_MIXTRAL_8X7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_DOLPHIN_2_7_MIXTRAL_8X7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_1_3B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_1_3B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_1_3B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_1_3B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_2_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_2_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_2_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_GPT_NEO_2_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_160M_DEDUPED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_160M_DEDUPED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_160M_DEDUPED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_160M_DEDUPED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_70M_DEDUPED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_70M_DEDUPED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_70M_DEDUPED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELEUTHERAI_PYTHIA_70M_DEDUPED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_13B_CHAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_13B_CHAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_13B_FAST_CHAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_13B_FAST_CHAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_CHAT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_CHAT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_CHAT_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_CHAT_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_FAST_CHAT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_FAST_CHAT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_FAST_CHAT_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ELYZA_JAPANESE_LLAMA_2_7B_FAST_CHAT_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_BF16_1_5_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_BF16_1_5_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_180B_CHAT_BF16_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_1_3_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_BF16_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_BF16_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_1_3_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_40B_INSTRUCT_BF16_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_2_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_BF16_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_BF16_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_2_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON_7B_INSTRUCT_BF16_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_FALCON2_11B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_FALCON2_11B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GARAGE_BAIND_PLATYPUS2_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GARAGE_BAIND_PLATYPUS2_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GARAGE_BAIND_PLATYPUS2_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GARAGE_BAIND_PLATYPUS2_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_2B_INSTRUCT_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GEMMA_7B_INSTRUCT_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GRADIENTAI_LLAMA_3_8B_INSTRUCT_262K_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GRADIENTAI_LLAMA_3_8B_INSTRUCT_262K_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_GRADIENTAI_LLAMA_3_8B_INSTRUCT_262K_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_GRADIENTAI_LLAMA_3_8B_INSTRUCT_262K_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_ALPHA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_ALPHA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_ALPHA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_ALPHA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_BETA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_BETA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_BETA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_MISTRAL_7B_SFT_BETA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_ALPHA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_ALPHA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_ALPHA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_ALPHA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_BETA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_BETA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_BETA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_STARCHAT_BETA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_ALPHA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_ALPHA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_ALPHA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_ALPHA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_BETA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_BETA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_BETA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_7B_BETA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_ORPO_141B_A35B_V01_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_HUGGINGFACEH4_ZEPHYR_ORPO_141B_A35B_V01_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_LLAMA_3_8B_INSTRUCT_GRADIENT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_LLAMA_3_8B_INSTRUCT_GRADIENT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_LLAMA_3_8B_INSTRUCT_GRADIENT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_LLAMA_3_8B_INSTRUCT_GRADIENT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_6_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_6_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_7_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_7_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_8_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_8_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_9_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_9_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_2_9_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_2_9_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_5_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_5_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_6_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_3_6_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_V3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_V3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_V3_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_V3_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_V3_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_INSTRUCT_V3_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_OPENORCA_GPTQ_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_OPENORCA_GPTQ_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_OPENORCA_GPTQ_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_OPENORCA_GPTQ_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_OPENORCA_GPTQ_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_OPENORCA_GPTQ_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_V3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_V3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRAL_7B_V3_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRAL_7B_V3_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MISTRALAI_MIXTRAL_8X22B_INSTRUCT_V0_1_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X22B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X22B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_10_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_10_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_6_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_6_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_7_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_7_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_8_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_8_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_9_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_9_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_1_9_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_1_9_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_10_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_10_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_6_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_6_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_7_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_7_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_8_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_8_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_9_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_9_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_9_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_1_9_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_MIXTRAL_8X7B_INSTRUCT_GPTQ_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NEXAAIDEV_OCTOPUS_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NEXAAIDEV_OCTOPUS_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NEXUSFLOW_STARLING_LM_7B_BETA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NEXUSFLOW_STARLING_LM_7B_BETA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_HERMES_2_PRO_LLAMA_3_8B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_HERMES_2_PRO_LLAMA_3_8B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_HERMES_2_PRO_LLAMA_3_8B_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_HERMES_2_PRO_LLAMA_3_8B_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_2_SOLAR_10_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_2_SOLAR_10_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_2_SOLAR_10_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_2_SOLAR_10_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA_2_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA_2_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA_2_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA_2_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA2_13B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA2_13B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA2_13B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_NOUS_HERMES_LLAMA2_13B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_YARN_MISTRAL_7B_128K_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_YARN_MISTRAL_7B_128K_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NOUSRESEARCH_YARN_MISTRAL_7B_128K_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NOUSRESEARCH_YARN_MISTRAL_7B_128K_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NVIDIA_LLAMA3_CHATQA_1_5_70B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NVIDIA_LLAMA3_CHATQA_1_5_70B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NVIDIA_LLAMA3_CHATQA_1_5_8B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NVIDIA_LLAMA3_CHATQA_1_5_8B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_NVIDIA_LLAMA3_CHATQA_1_5_8B_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_NVIDIA_LLAMA3_CHATQA_1_5_8B_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_OPENLM_RESEARCH_OPEN_LLAMA_7B_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_OPENLM_RESEARCH_OPEN_LLAMA_7B_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_OPENLM_RESEARCH_OPEN_LLAMA_7B_V2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_OPENLM_RESEARCH_OPEN_LLAMA_7B_V2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_PHI_2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_PHI_2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_PHI_2_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_PHI_2_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_PHI_3_MINI_128K_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_PHI_3_MINI_128K_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_PHI_3_MINI_128K_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_PHI_3_MINI_128K_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_PHI_3_MINI_4K_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_PHI_3_MINI_4K_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_PHI_3_MINI_4K_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_PHI_3_MINI_4K_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_QWEN2_0_5B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_QWEN2_0_5B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_QWEN2_0_5B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_QWEN2_0_5B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_QWEN2_1_5B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_QWEN2_1_5B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_QWEN2_1_5B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_QWEN2_1_5B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_QWEN2_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_QWEN2_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_QWEN2_7B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_QWEN2_7B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_RINNA_3_6B_INSTRUCTION_PPO_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SEALION_3B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SEALION_3B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SEALION_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SEALION_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SEALION_7B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SEALION_7B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SHENZHI_WANG_LLAMA3_8B_CHINESE_CHAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SHENZHI_WANG_LLAMA3_8B_CHINESE_CHAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SNOWFLAKE_ARCTIC_INSTRUCT_VLLM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SNOWFLAKE_ARCTIC_INSTRUCT_VLLM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SNOWFLAKE_ARCTIC_INSTRUCT_VLLM_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SNOWFLAKE_ARCTIC_INSTRUCT_VLLM_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_SNOWFLAKE_ARCTIC_INSTRUCT_VLLM_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_SNOWFLAKE_ARCTIC_INSTRUCT_VLLM_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_STARCODER_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_STARCODER_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_STARCODER_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_STARCODER_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_STARCODERBASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_STARCODERBASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_STARCODERBASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_STARCODERBASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TEKNIUM_OPENHERMES_2_MISTRAL_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TEKNIUM_OPENHERMES_2_MISTRAL_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TEKNIUM_OPENHERMES_2_MISTRAL_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TEKNIUM_OPENHERMES_2_MISTRAL_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_THEBLOKE_MISTRAL_7B_OPENORCA_AWQ_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_THEBLOKE_MISTRAL_7B_OPENORCA_AWQ_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_THEBLOKE_MISTRAL_7B_OPENORCA_AWQ_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_THEBLOKE_MISTRAL_7B_OPENORCA_AWQ_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_THEBLOKE_MISTRAL_7B_OPENORCA_AWQ_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_THEBLOKE_MISTRAL_7B_OPENORCA_AWQ_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TIIUAE_FALCON_RW_1B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TIIUAE_FALCON_RW_1B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TIIUAE_FALCON_RW_1B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TIIUAE_FALCON_RW_1B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TINYLLAMA_1_1B_INTERMEDIATE_STEP_1431K_3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TINYLLAMA_1_1B_INTERMEDIATE_STEP_1431K_3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TINYLLAMA_1_1B_INTERMEDIATE_STEP_1431K_3_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TINYLLAMA_1_1B_INTERMEDIATE_STEP_1431K_3_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V0_6_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V0_6_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V0_6_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V0_6_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V1_0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V1_0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V1_0_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_TINYLLAMA_TINYLLAMA_1_1B_CHAT_V1_0_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_WRITER_PALMYRA_SMALL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_WRITER_PALMYRA_SMALL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_WRITER_PALMYRA_SMALL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_WRITER_PALMYRA_SMALL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_34B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_34B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_34B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_34B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_34B_CHAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_34B_CHAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_34B_CHAT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_34B_CHAT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_6B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_6B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_6B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_6B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_6B_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_6B_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_6B_CHAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_6B_CHAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_6B_CHAT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_6B_CHAT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_6B_CHAT_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_6B_CHAT_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_9B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_9B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_9B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_9B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_9B_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_9B_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_9B_CHAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_9B_CHAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_9B_CHAT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_9B_CHAT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_YI_1_5_9B_CHAT_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_YI_1_5_9B_CHAT_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_LLM_ZEPHYR_7B_GEMMA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_LLM_ZEPHYR_7B_GEMMA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENG_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENG_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENGLISH_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENGLISH_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENGLISH_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENGLISH_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENGLISH_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_CASED_FINETUNED_CONLL03_ENGLISH_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENG_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENG_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENGLISH_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENGLISH_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENGLISH_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENGLISH_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENGLISH_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_NER_DISTILBERT_BASE_UNCASED_FINETUNED_CONLL03_ENGLISH_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_ALL_MINILM_L6_V2_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_V1_5_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_V1_5_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_V1_5_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_BASE_EN_V1_5_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_V1_5_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_V1_5_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_V1_5_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_EN_V1_5_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_ZH_V1_5_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_ZH_V1_5_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_ZH_V1_5_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_LARGE_ZH_V1_5_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_M3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_M3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_M3_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_M3_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_V1_5_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_V1_5_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_V1_5_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_BGE_SMALL_EN_V1_5_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_BASE_V2_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_LARGE_V2_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_E5_SMALL_V2_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_BASE_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_LARGE_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_GTE_SMALL_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_BASE_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SENTENCESIMILARITY_MULTILINGUAL_E5_LARGE_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_MULTILINGUAL_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_MULTILINGUAL_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILBERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_DISTILROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_DISTILROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_BASE_OPENAI_DETECTOR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_ROBERTA_LARGE_OPENAI_DETECTOR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_CLM_ENDE_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENDE_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_ENRO_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_TLM_XNLI15_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SPC_XLM_MLM_XNLI15_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BART_LARGE_CNN_SAMSUM_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_CNN_DAILYMAIL_SUMM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_CNN_DAILYMAIL_SUMM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BERT_SMALL2BERT_SMALL_FINETUNED_CNN_DAILY_MAIL_SUMMARIZATION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_ARXIV_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_BIGBIRD_PEGASUS_LARGE_PUBMED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_12_6_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_CNN_6_6_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_1_1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_SUMMARIZATION_DISTILBART_XSUM_12_3_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_MULTILINGUAL_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_CASED_WHOLE_WORD_MASKING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_BERT_LARGE_UNCASED_WHOLE_WORD_MASKING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_MULTILINGUAL_CASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILBERT_BASE_UNCASED_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILROBERTA_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILROBERTA_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILROBERTA_BASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILROBERTA_BASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_DISTILROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_DISTILROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_MODELS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_MODELS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_MODELS_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_MODELS_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_MODELS_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_MODELS_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_MODELS_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_MODELS_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_BASE_OPENAI_DETECTOR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_ROBERTA_LARGE_OPENAI_DETECTOR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_CLM_ENDE_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_CLM_ENDE_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_CLM_ENDE_1024_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_CLM_ENDE_1024_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_CLM_ENDE_1024_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_CLM_ENDE_1024_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_CLM_ENDE_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_CLM_ENDE_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENDE_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENDE_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENDE_1024_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENDE_1024_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENDE_1024_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENDE_1024_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENDE_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENDE_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENRO_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENRO_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENRO_1024_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENRO_1024_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENRO_1024_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENRO_1024_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_ENRO_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_ENRO_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TC_XLM_MLM_TLM_XNLI15_1024_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BART4CSC_BASE_CHINESE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_BNB_INT8_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_BNB_INT8_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_BNB_INT8_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_BNB_INT8_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_BNB_INT8_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_BNB_INT8_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_FP16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_FP16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_FP16_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_BIGSCIENCE_T0PP_FP16_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_2_5",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_BASE_SAMSUM_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_5",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_6() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_1_6",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_LARGE_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_2_5",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_3_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_1_3_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_SMALL_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_5",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_6() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_1_6",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XL_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_BNB_INT8_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_FP16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_FP16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_FP16_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_T5_XXL_FP16_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_FLAN_UL2_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_PEGASUS_PARAPHRASE_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_QCPG_SENTENCES_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXT2TEXT_T5_ONE_LINE_SUMMARY_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTEMBEDDING_ALL_MINILM_L6_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTEMBEDDING_ALL_MINILM_L6_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B1_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_1B7_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOM_560M_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOM_560M_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B1_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_1B7_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_BLOOMZ_560M_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_4_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_4_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_1_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DISTILGPT2_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DISTILGPT2_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_12B_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_3B_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_DOLLY_V2_7B_BF16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_FALCON_40B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_FALCON_40B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_FALCON_40B_INSTRUCT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_FALCON_40B_INSTRUCT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_FALCON_7B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_FALCON_7B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_FALCON_7B_INSTRUCT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_FALCON_7B_INSTRUCT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_4_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_4_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_1_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_1_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_GPT2_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_GPT2_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_MODELS_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_MODELS_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION_OPEN_LLAMA_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_176B_INT8_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_176B_INT8_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_176B_INT8_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_176B_INT8_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_176B_INT8_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_176B_INT8_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_3B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOM_7B1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_176B_FP16_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_3B_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_BLOOMZ_7B1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_2_XL_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_2_4",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_J_6B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_1_3B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_125M_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_GPT_NEO_2_7B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_LIGHTGPT_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_BF16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_INSTRUCT_BF16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_MPT_7B_STORYWRITER_BF16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_3B_V1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_BASE_7B_V1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_3B_V1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_CHAT_7B_V1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3B_V1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3BV1FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_3BV1FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B_V1_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B1FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION1_REDPAJAMA_INCITE_INSTRUCT_7B1FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOX_20B_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TEXTGENERATION2_GPT_NEOXT_CHAT_BASE_20B_FP16_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_EN_ES_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_EN_ES_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_EN_ES_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_EN_ES_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_EN_ES_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_EN_ES_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_EN_VI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_EN_VI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_EN_VI_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_EN_VI_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_EN_VI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_EN_VI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_OPUS_MT_MUL_EN_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_OPUS_MT_MUL_EN_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_BASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_BASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_LARGE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_LARGE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_LARGE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_LARGE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_LARGE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_LARGE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_SMALL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_SMALL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_SMALL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_SMALL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TRANSLATION_T5_SMALL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TRANSLATION_T5_SMALL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_22H_VINTEDOIS_DIFFUSION_V0_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_22H_VINTEDOIS_DIFFUSION_V0_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_22H_VINTEDOIS_DIFFUSION_V0_1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_22H_VINTEDOIS_DIFFUSION_V0_1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_22H_VINTEDOIS_DIFFUSION_V0_1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_22H_VINTEDOIS_DIFFUSION_V0_1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AKIKAGURA_MKGEN_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AKIKAGURA_MKGEN_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AKIKAGURA_MKGEN_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AKIKAGURA_MKGEN_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AKIKAGURA_MKGEN_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AKIKAGURA_MKGEN_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES_FASTDB_4800_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES_FASTDB_4800_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES_FASTDB_4800_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES_FASTDB_4800_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES_FASTDB_4800_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES_FASTDB_4800_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES9000_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES9000_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES9000_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES9000_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES9000_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ALXDFY_NOGGLES9000_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ANDITE_ANYTHING_V4_0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ANDITE_ANYTHING_V4_0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ANDITE_ANYTHING_V4_0_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ANDITE_ANYTHING_V4_0_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ANDITE_ANYTHING_V4_0_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ANDITE_ANYTHING_V4_0_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ASTRALITEHEART_PONY_DIFFUSION_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ASTRALITEHEART_PONY_DIFFUSION_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ASTRALITEHEART_PONY_DIFFUSION_V2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ASTRALITEHEART_PONY_DIFFUSION_V2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ASTRALITEHEART_PONY_DIFFUSION_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ASTRALITEHEART_PONY_DIFFUSION_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AVRIK_ABSTRACT_ANIM_SPRITESHEETS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AVRIK_ABSTRACT_ANIM_SPRITESHEETS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AVRIK_ABSTRACT_ANIM_SPRITESHEETS_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AVRIK_ABSTRACT_ANIM_SPRITESHEETS_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AVRIK_ABSTRACT_ANIM_SPRITESHEETS_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AVRIK_ABSTRACT_ANIM_SPRITESHEETS_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AYBEECEEDEE_KNOLLINGCASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AYBEECEEDEE_KNOLLINGCASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AYBEECEEDEE_KNOLLINGCASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AYBEECEEDEE_KNOLLINGCASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_AYBEECEEDEE_KNOLLINGCASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_AYBEECEEDEE_KNOLLINGCASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BINGSU_MY_K_ANYTHING_V3_0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BINGSU_MY_K_ANYTHING_V3_0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BINGSU_MY_K_ANYTHING_V3_0_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BINGSU_MY_K_ANYTHING_V3_0_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BINGSU_MY_K_ANYTHING_V3_0_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BINGSU_MY_K_ANYTHING_V3_0_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BINGSU_MY_KOREAN_STABLE_DIFFUSION_V1_5_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BINGSU_MY_KOREAN_STABLE_DIFFUSION_V1_5_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BINGSU_MY_KOREAN_STABLE_DIFFUSION_V1_5_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BINGSU_MY_KOREAN_STABLE_DIFFUSION_V1_5_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BINGSU_MY_KOREAN_STABLE_DIFFUSION_V1_5_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BINGSU_MY_KOREAN_STABLE_DIFFUSION_V1_5_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BUNTOPSIH_NOVGORANSTEFANOVSKI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BUNTOPSIH_NOVGORANSTEFANOVSKI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BUNTOPSIH_NOVGORANSTEFANOVSKI_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BUNTOPSIH_NOVGORANSTEFANOVSKI_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_BUNTOPSIH_NOVGORANSTEFANOVSKI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_BUNTOPSIH_NOVGORANSTEFANOVSKI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CLAUDFUEN_PHOTOREALISTIC_FUEN_V1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CLAUDFUEN_PHOTOREALISTIC_FUEN_V1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CLAUDFUEN_PHOTOREALISTIC_FUEN_V1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CLAUDFUEN_PHOTOREALISTIC_FUEN_V1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CLAUDFUEN_PHOTOREALISTIC_FUEN_V1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CLAUDFUEN_PHOTOREALISTIC_FUEN_V1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CODER119_VECTORARTZ_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CODER119_VECTORARTZ_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CODER119_VECTORARTZ_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CODER119_VECTORARTZ_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CODER119_VECTORARTZ_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CODER119_VECTORARTZ_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CONFLICTX_COMPLEX_LINEART_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CONFLICTX_COMPLEX_LINEART_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CONFLICTX_COMPLEX_LINEART_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CONFLICTX_COMPLEX_LINEART_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_CONFLICTX_COMPLEX_LINEART_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_CONFLICTX_COMPLEX_LINEART_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_CATS_MUSICAL_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_CATS_MUSICAL_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_CATS_MUSICAL_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_CATS_MUSICAL_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_CATS_MUSICAL_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_CATS_MUSICAL_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_JWST_DEEP_SPACE_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_JWST_DEEP_SPACE_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_JWST_DEEP_SPACE_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_JWST_DEEP_SPACE_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_JWST_DEEP_SPACE_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_JWST_DEEP_SPACE_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_TRON_LEGACY_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_TRON_LEGACY_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_TRON_LEGACY_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_TRON_LEGACY_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_TRON_LEGACY_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_TRON_LEGACY_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_VAN_GOGH_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_VAN_GOGH_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_VAN_GOGH_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_VAN_GOGH_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DALLINMACKAY_VAN_GOGH_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DALLINMACKAY_VAN_GOGH_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DGSPITZER_CYBERPUNK_ANIME_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DGSPITZER_CYBERPUNK_ANIME_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DGSPITZER_CYBERPUNK_ANIME_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DGSPITZER_CYBERPUNK_ANIME_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DGSPITZER_CYBERPUNK_ANIME_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DGSPITZER_CYBERPUNK_ANIME_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DREAMLIKE_ART_DREAMLIKE_DIFFUSION_1_0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DREAMLIKE_ART_DREAMLIKE_DIFFUSION_1_0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DREAMLIKE_ART_DREAMLIKE_DIFFUSION_1_0_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DREAMLIKE_ART_DREAMLIKE_DIFFUSION_1_0_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_DREAMLIKE_ART_DREAMLIKE_DIFFUSION_1_0_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_DREAMLIKE_ART_DREAMLIKE_DIFFUSION_1_0_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EIMISS_EIMISANIMEDIFFUSION_1_0V_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EIMISS_EIMISANIMEDIFFUSION_1_0V_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EIMISS_EIMISANIMEDIFFUSION_1_0V_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EIMISS_EIMISANIMEDIFFUSION_1_0V_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EIMISS_EIMISANIMEDIFFUSION_1_0V_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EIMISS_EIMISANIMEDIFFUSION_1_0V_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ENVVI_INKPUNK_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ENVVI_INKPUNK_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ENVVI_INKPUNK_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ENVVI_INKPUNK_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_ENVVI_INKPUNK_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_ENVVI_INKPUNK_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EVEL_YOSHIN_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EVEL_YOSHIN_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EVEL_YOSHIN_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EVEL_YOSHIN_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EVEL_YOSHIN_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EVEL_YOSHIN_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EXTRAPHY_MUSTAFA_KEMAL_ATATURK_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EXTRAPHY_MUSTAFA_KEMAL_ATATURK_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EXTRAPHY_MUSTAFA_KEMAL_ATATURK_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EXTRAPHY_MUSTAFA_KEMAL_ATATURK_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_EXTRAPHY_MUSTAFA_KEMAL_ATATURK_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_EXTRAPHY_MUSTAFA_KEMAL_ATATURK_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FFFILONI_MR_MEN_AND_LITTLE_MISSES_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FFFILONI_MR_MEN_AND_LITTLE_MISSES_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FFFILONI_MR_MEN_AND_LITTLE_MISSES_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FFFILONI_MR_MEN_AND_LITTLE_MISSES_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FFFILONI_MR_MEN_AND_LITTLE_MISSES_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FFFILONI_MR_MEN_AND_LITTLE_MISSES_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_ELRISITAS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_ELRISITAS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_ELRISITAS_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_ELRISITAS_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_ELRISITAS_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_ELRISITAS_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_MODEL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_MODEL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_MODEL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_MODEL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_MODEL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_BALLOONART_MODEL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICRO_MODEL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICRO_MODEL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICROSCOPIC_MODEL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICROSCOPIC_MODEL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICROSCOPIC_MODEL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICROSCOPIC_MODEL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICROSCOPIC_MODEL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_MICROSCOPIC_MODEL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_PAPERCUT_MODEL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_PAPERCUT_MODEL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_PAPERCUT_MODEL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_PAPERCUT_MODEL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_PAPERCUT_MODEL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_PAPERCUT_MODEL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_VOXELART_MODEL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_VOXELART_MODEL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_VOXELART_MODEL_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_VOXELART_MODEL_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_VOXELART_MODEL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_FICTIVERSE_STABLE_DIFFUSION_VOXELART_MODEL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_HAOR_EVT_V3_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_HAOR_EVT_V3_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_HAOR_EVT_V3_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_HAOR_EVT_V3_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_HAOR_EVT_V3_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_HAOR_EVT_V3_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_HASSANBLEND_HASSANBLEND1_4_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_HASSANBLEND_HASSANBLEND1_4_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_HASSANBLEND_HASSANBLEND1_4_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_HASSANBLEND_HASSANBLEND1_4_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_HASSANBLEND_HASSANBLEND1_4_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_HASSANBLEND_HASSANBLEND1_4_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_1B_CHINESE_EN_V01_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_1B_CHINESE_EN_V01_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_1B_CHINESE_V0_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_1B_CHINESE_V0_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_EN_V0_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_EN_V0_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_EN_V0_1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_EN_V0_1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_EN_V0_1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_EN_V0_1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_V0_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_V0_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_V0_1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_V0_1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_V0_1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IDEA_CCNL_TAIYI_STABLE_DIFFUSION_1B_CHINESE_V0_1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IFANSNEK_JOHNDIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IFANSNEK_JOHNDIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IFANSNEK_JOHNDIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IFANSNEK_JOHNDIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_IFANSNEK_JOHNDIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_IFANSNEK_JOHNDIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_JERSONM89_AVATAR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_JERSONM89_AVATAR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_JERSONM89_AVATAR_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_JERSONM89_AVATAR_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_JERSONM89_AVATAR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_JERSONM89_AVATAR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_JVKAPE_ICONSMI_APPICONSMODELFORSD_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_JVKAPE_ICONSMI_APPICONSMODELFORSD_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_JVKAPE_ICONSMI_APPICONSMODELFORSD_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_JVKAPE_ICONSMI_APPICONSMODELFORSD_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_JVKAPE_ICONSMI_APPICONSMODELFORSD_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_JVKAPE_ICONSMI_APPICONSMODELFORSD_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_KATAKANA_2D_MIX_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_KATAKANA_2D_MIX_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_KATAKANA_2D_MIX_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_KATAKANA_2D_MIX_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_KATAKANA_2D_MIX_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_KATAKANA_2D_MIX_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LACAMBRE_VULVINE_LOOK_V02_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LACAMBRE_VULVINE_LOOK_V02_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LACAMBRE_VULVINE_LOOK_V02_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LACAMBRE_VULVINE_LOOK_V02_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LACAMBRE_VULVINE_LOOK_V02_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LACAMBRE_VULVINE_LOOK_V02_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LANGBOAT_GUOHUA_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LANGBOAT_GUOHUA_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LANGBOAT_GUOHUA_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LANGBOAT_GUOHUA_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LANGBOAT_GUOHUA_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LANGBOAT_GUOHUA_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LINAQRUF_ANYTHING_V3_0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LINAQRUF_ANYTHING_V3_0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LINAQRUF_ANYTHING_V3_0_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LINAQRUF_ANYTHING_V3_0_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_LINAQRUF_ANYTHING_V3_0_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_LINAQRUF_ANYTHING_V3_0_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MIKESMODELS_WALTZ_WITH_BASHIR_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MIKESMODELS_WALTZ_WITH_BASHIR_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MIKESMODELS_WALTZ_WITH_BASHIR_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MIKESMODELS_WALTZ_WITH_BASHIR_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MIKESMODELS_WALTZ_WITH_BASHIR_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MIKESMODELS_WALTZ_WITH_BASHIR_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITCHTECH_KLINGON_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITCHTECH_KLINGON_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITCHTECH_KLINGON_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITCHTECH_KLINGON_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITCHTECH_KLINGON_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITCHTECH_KLINGON_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITCHTECH_VULCAN_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITCHTECH_VULCAN_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITCHTECH_VULCAN_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITCHTECH_VULCAN_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITCHTECH_VULCAN_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITCHTECH_VULCAN_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITSUA_MITSUA_DIFFUSION_CC0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITSUA_MITSUA_DIFFUSION_CC0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITSUA_MITSUA_DIFFUSION_CC0_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITSUA_MITSUA_DIFFUSION_CC0_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_MITSUA_MITSUA_DIFFUSION_CC0_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_MITSUA_MITSUA_DIFFUSION_CC0_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NACLBIT_TRINART_STABLE_DIFFUSION_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NACLBIT_TRINART_STABLE_DIFFUSION_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NACLBIT_TRINART_STABLE_DIFFUSION_V2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NACLBIT_TRINART_STABLE_DIFFUSION_V2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NACLBIT_TRINART_STABLE_DIFFUSION_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NACLBIT_TRINART_STABLE_DIFFUSION_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCANE_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCANE_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCANE_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCANE_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCANE_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCANE_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCHER_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCHER_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCHER_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCHER_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCHER_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ARCHER_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_CLASSIC_ANIM_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_CLASSIC_ANIM_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_CLASSIC_ANIM_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_CLASSIC_ANIM_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_CLASSIC_ANIM_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_CLASSIC_ANIM_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ELDEN_RING_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ELDEN_RING_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ELDEN_RING_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ELDEN_RING_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_ELDEN_RING_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_ELDEN_RING_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_FUTURE_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_FUTURE_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_FUTURE_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_FUTURE_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_FUTURE_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_FUTURE_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_GHIBLI_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_GHIBLI_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_GHIBLI_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_GHIBLI_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_GHIBLI_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_GHIBLI_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_MO_DI_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_MO_DI_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_MO_DI_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_MO_DI_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_MO_DI_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_MO_DI_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_NITRO_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_NITRO_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_NITRO_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_NITRO_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_NITRO_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_NITRO_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_REDSHIFT_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_REDSHIFT_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_REDSHIFT_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_REDSHIFT_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_REDSHIFT_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_REDSHIFT_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_SPIDER_VERSE_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_SPIDER_VERSE_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_SPIDER_VERSE_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_SPIDER_VERSE_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NITROSOCKE_SPIDER_VERSE_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NITROSOCKE_SPIDER_VERSE_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NOUSR_ROBO_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NOUSR_ROBO_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NOUSR_ROBO_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NOUSR_ROBO_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_NOUSR_ROBO_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_NOUSR_ROBO_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_OGKALU_COMIC_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_OGKALU_COMIC_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_OGKALU_COMIC_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_OGKALU_COMIC_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_OGKALU_COMIC_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_OGKALU_COMIC_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_OPENJOURNEY_OPENJOURNEY_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_OPENJOURNEY_OPENJOURNEY_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_OPENJOURNEY_OPENJOURNEY_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_OPENJOURNEY_OPENJOURNEY_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_OPENJOURNEY_OPENJOURNEY_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_OPENJOURNEY_OPENJOURNEY_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PIESPOSITO_OPENPOTIONBOTTLE_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PIESPOSITO_OPENPOTIONBOTTLE_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PIESPOSITO_OPENPOTIONBOTTLE_V2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PIESPOSITO_OPENPOTIONBOTTLE_V2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PIESPOSITO_OPENPOTIONBOTTLE_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PIESPOSITO_OPENPOTIONBOTTLE_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PLASMO_VOXEL_ISH_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PLASMO_VOXEL_ISH_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PLASMO_VOXEL_ISH_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PLASMO_VOXEL_ISH_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PLASMO_VOXEL_ISH_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PLASMO_VOXEL_ISH_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PLASMO_WOOLITIZE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PLASMO_WOOLITIZE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PLASMO_WOOLITIZE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PLASMO_WOOLITIZE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PLASMO_WOOLITIZE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PLASMO_WOOLITIZE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUND_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUND_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUND_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUND_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUND_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUND_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUNDDIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROGAMERGOV_MIN_ILLUST_BACKGROUNDDIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROMPTHERO_LINKEDIN_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROMPTHERO_LINKEDIN_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROMPTHERO_LINKEDIN_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROMPTHERO_LINKEDIN_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROMPTHERO_LINKEDIN_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROMPTHERO_LINKEDIN_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROMPTHERO_OPENJOURNEY_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROMPTHERO_OPENJOURNEY_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROMPTHERO_OPENJOURNEY_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROMPTHERO_OPENJOURNEY_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_PROMPTHERO_OPENJOURNEY_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_PROMPTHERO_OPENJOURNEY_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_QILEX_MAGIC_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_QILEX_MAGIC_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_QILEX_MAGIC_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_QILEX_MAGIC_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_QILEX_MAGIC_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_QILEX_MAGIC_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RABIDGREMLIN_SD_DB_EPIC_SPACE_MACHINE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RABIDGREMLIN_SD_DB_EPIC_SPACE_MACHINE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RABIDGREMLIN_SD_DB_EPIC_SPACE_MACHINE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RABIDGREMLIN_SD_DB_EPIC_SPACE_MACHINE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RABIDGREMLIN_SD_DB_EPIC_SPACE_MACHINE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RABIDGREMLIN_SD_DB_EPIC_SPACE_MACHINE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RAYHELL_POPUPBOOK_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RAYHELL_POPUPBOOK_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RAYHELL_POPUPBOOK_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RAYHELL_POPUPBOOK_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RAYHELL_POPUPBOOK_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RAYHELL_POPUPBOOK_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RUNWAYML_STABLE_DIFFUSION_V1_5_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RUNWAYML_STABLE_DIFFUSION_V1_5_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RUNWAYML_STABLE_DIFFUSION_V1_5_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RUNWAYML_STABLE_DIFFUSION_V1_5_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_RUNWAYML_STABLE_DIFFUSION_V1_5_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_RUNWAYML_STABLE_DIFFUSION_V1_5_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_S3NH_BEKSINSKI_STYLE_STABLE_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_S3NH_BEKSINSKI_STYLE_STABLE_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_S3NH_BEKSINSKI_STYLE_STABLE_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_S3NH_BEKSINSKI_STYLE_STABLE_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_S3NH_BEKSINSKI_STYLE_STABLE_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_S3NH_BEKSINSKI_STYLE_STABLE_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHAR_CYCLPS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHAR_CYCLPS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHARACTER_CYCLPS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHARACTER_CYCLPS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHARACTER_CYCLPS_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHARACTER_CYCLPS_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHARACTER_CYCLPS_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_ORIGINAL_CHARACTER_CYCLPS_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_STYLE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_STYLE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_STYLE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_STYLE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_STYLE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_PERSONA_5_SHIGENORI_STYLE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_SERAPHM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_SERAPHM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_SERAPHM_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_SERAPHM_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_SERAPHM_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SD_DREAMBOOTH_LIBRARY_SERAPHM_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SHIRAYU_SD_TOHOKU_V1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SHIRAYU_SD_TOHOKU_V1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SHIRAYU_SD_TOHOKU_V1_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SHIRAYU_SD_TOHOKU_V1_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_SHIRAYU_SD_TOHOKU_V1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_SHIRAYU_SD_TOHOKU_V1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_THELASTBEN_HRRZG_STYLE_768PX_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_THELASTBEN_HRRZG_STYLE_768PX_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_THELASTBEN_HRRZG_STYLE_768PX_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_THELASTBEN_HRRZG_STYLE_768PX_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_THELASTBEN_HRRZG_STYLE_768PX_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_THELASTBEN_HRRZG_STYLE_768PX_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TIMOTHEPEARCE_GINA_THE_CAT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TIMOTHEPEARCE_GINA_THE_CAT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TIMOTHEPEARCE_GINA_THE_CAT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TIMOTHEPEARCE_GINA_THE_CAT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TIMOTHEPEARCE_GINA_THE_CAT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TIMOTHEPEARCE_GINA_THE_CAT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TRYSTAR_CLONEDIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TRYSTAR_CLONEDIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TRYSTAR_CLONEDIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TRYSTAR_CLONEDIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TRYSTAR_CLONEDIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TRYSTAR_CLONEDIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TUWONGA_DBLUTH_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TUWONGA_DBLUTH_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TUWONGA_DBLUTH_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TUWONGA_DBLUTH_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TUWONGA_DBLUTH_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TUWONGA_DBLUTH_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TUWONGA_ROTOSCOPEE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TUWONGA_ROTOSCOPEE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TUWONGA_ROTOSCOPEE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TUWONGA_ROTOSCOPEE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_TUWONGA_ROTOSCOPEE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_TUWONGA_ROTOSCOPEE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_VOLRATH50_FANTASY_CARD_DIFFUSION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_VOLRATH50_FANTASY_CARD_DIFFUSION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_VOLRATH50_FANTASY_CARD_DIFFUSION_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_VOLRATH50_FANTASY_CARD_DIFFUSION_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_VOLRATH50_FANTASY_CARD_DIFFUSION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_VOLRATH50_FANTASY_CARD_DIFFUSION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_YAYAB_SD_ONEPIECE_DIFFUSERS4_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_YAYAB_SD_ONEPIECE_DIFFUSERS4_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_YAYAB_SD_ONEPIECE_DIFFUSERS4_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_YAYAB_SD_ONEPIECE_DIFFUSERS4_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_TXT2IMG_YAYAB_SD_ONEPIECE_DIFFUSERS4_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_TXT2IMG_YAYAB_SD_ONEPIECE_DIFFUSERS4_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DEBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DEBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DEBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DEBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DISTILROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DISTILROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DISTILROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_DISTILROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_MINILM2_L6_H768_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_MINILM2_L6_H768_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_MINILM2_L6_H768_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_MINILM2_L6_H768_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_ROBERTA_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_ROBERTA_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_ROBERTA_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_CROSS_ENCODER_NLI_ROBERTA_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_DIGITALEPIDEMIOLOGYLAB_COVID_TWIT_BERT2_MNLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_DIGITALEPIDEMIOLOGYLAB_COVID_TWIT_BERT2_MNLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_DIGITALEPIDEMIOLOGYLAB_COVID_TWITTER_BERT_V2_MNLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_DIGITALEPIDEMIOLOGYLAB_COVID_TWITTER_BERT_V2_MNLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_DIGITALEPIDEMIOLOGYLAB_COVID_TWITTER_BERT_V2_MNLI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_DIGITALEPIDEMIOLOGYLAB_COVID_TWITTER_BERT_V2_MNLI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_ELELDAR_THEME_CLASSIFICATION_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_ELELDAR_THEME_CLASSIFICATION_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_ELELDAR_THEME_CLASSIFICATION_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_ELELDAR_THEME_CLASSIFICATION_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_ALLNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_ALLNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_ALLNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_ALLNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_MULTINLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_MULTINLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_MULTINLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_MULTINLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_SNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_SNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_SNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_MULTILINGUAL_CASED_SNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_ALLNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_ALLNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_ALLNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_ALLNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_MULTINLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_MULTINLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_MULTINLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_MULTINLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_SNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_SNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_SNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERT_BASE_TURKISH_CASED_SNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERTBASE_MLING_CASED_ALLNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERTBASE_MLING_CASED_ALLNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_BERTBASE_MLING_CASED_MULTINLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_BERTBASE_MLING_CASED_MULTINLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CBERTBASE_TURKISH_MC4CASED_MULTINLITR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CBERTBASE_TURKISH_MC4CASED_MULTINLITR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CBERTBASE_TURKISH_MC4CASED_SNLITR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CBERTBASE_TURKISH_MC4CASED_SNLITR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CBERTBASE_TURKISHMC4_CASED_ALLNLITR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CBERTBASE_TURKISHMC4_CASED_ALLNLITR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_ALLNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_ALLNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_ALLNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_ALLNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_MULTINLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_MULTINLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_MULTINLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_MULTINLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_SNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_SNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_SNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_CONVBERT_BASE_TURKISH_MC4_CASED_SNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DBASE_TURKISH_CASED_ALLNLITR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DBASE_TURKISH_CASED_ALLNLITR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DBERTBASE_TURKISH_CASED_MULTINLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DBERTBASE_TURKISH_CASED_MULTINLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_ALLNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_ALLNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_ALLNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_ALLNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_MULTINLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_MULTINLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_MULTINLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_MULTINLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_SNLI_TR_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_SNLI_TR_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_SNLI_TR_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_EMRECAN_DISTILBERT_BASE_TURKISH_CASED_SNLI_TR_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_FACEBOOK_BART_LARGE_MNLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_FACEBOOK_BART_LARGE_MNLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_FACEBOOK_BART_LARGE_MNLI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_FACEBOOK_BART_LARGE_MNLI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_JIVA_XLM_ROBERTA_LARGE_IT_MNLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_JIVA_XLM_ROBERTA_LARGE_IT_MNLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_JIVA_XLM_ROBERTA_LARGE_IT_MNLI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_JIVA_XLM_ROBERTA_LARGE_IT_MNLI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_LIGHTETERNAL_NLI_XLM_R_GREEK_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_LIGHTETERNAL_NLI_XLM_R_GREEK_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_LIGHTETERNAL_NLI_XLM_R_GREEK_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_LIGHTETERNAL_NLI_XLM_R_GREEK_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_MORITZLAURER_DEBERTA_V3_LARGE_MNLI_FEVER_ANLI_LING_WANLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_MORITZLAURER_DEBERTA_V3_LARGE_MNLI_FEVER_ANLI_LING_WANLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_MORITZLAURER_DEBERTA_V3_LARGE_MNLI_FEVER_ANLI_LING_WANLI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_MORITZLAURER_DEBERTA_V3_LARGE_MNLI_FEVER_ANLI_LING_WANLI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_MORITZLAURER_DEBERTA3LARGE_MNLI_FEVER_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_MORITZLAURER_DEBERTA3LARGE_MNLI_FEVER_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_MORITZLAURER_MDEBERTA_V3_BASE_XNLI_MULTILINGUAL_NLI_2MIL7_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_MORITZLAURER_MDEBERTA_V3_BASE_XNLI_MULTILINGUAL_NLI_2MIL7_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_MORITZLAURER_MDEBERTA_V3_BASE_XNLI_MULTILINGUAL_NLI_2MIL7_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_MORITZLAURER_MDEBERTA_V3_BASE_XNLI_MULTILINGUAL_NLI_2MIL7_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_MORITZLAURER_MDEBERTA3BASE_XNLI_MLING_NLI_2M7_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_MORITZLAURER_MDEBERTA3BASE_XNLI_MLING_NLI_2M7_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_NARSIL_BART_LARGE_MNLI_OPTI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_NARSIL_BART_LARGE_MNLI_OPTI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_NARSIL_BART_LARGE_MNLI_OPTI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_NARSIL_BART_LARGE_MNLI_OPTI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_NARSIL_DEBERTA_LARGE_MNLI_ZERO_CLS_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_NARSIL_DEBERTA_LARGE_MNLI_ZERO_CLS_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_NARSIL_DEBERTA_LARGE_MNLI_ZERO_CLS_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_NARSIL_DEBERTA_LARGE_MNLI_ZERO_CLS_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_NAVTECA_BART_LARGE_MNLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_NAVTECA_BART_LARGE_MNLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_NAVTECA_BART_LARGE_MNLI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_NAVTECA_BART_LARGE_MNLI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_RECOGNAI_BERT_BASE_SPANISH_WWM_CASED_XNLI_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_RECOGNAI_BERT_BASE_SPANISH_WWM_CASED_XNLI_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_RECOGNAI_BERT_BASE_SPANISH_WWM_CASED_XNLI_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_RECOGNAI_BERT_BASE_SPANISH_WWM_CASED_XNLI_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_MEDIUM_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_MEDIUM_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_MEDIUM_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_MEDIUM_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_SMALL_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_SMALL_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_SMALL_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"HUGGINGFACE_ZSTC_RECOGNAI_ZEROSHOT_SELECTRA_SMALL_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TC_LLAMA_PROMPT_GUARD_86M_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TC_LLAMA_PROMPT_GUARD_86M_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_5",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_6() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_6",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_7() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_7",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_2_1_8() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_2_1_8",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_4_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_4_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_2_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_2_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_2_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_2_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_13B_F_4_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_13B_F_4_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_5",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_6() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_6",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_2_0_7() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_2_0_7",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_4_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_4_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_2_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_2_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_2_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_2_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_70B_F_4_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_70B_F_4_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_5() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_5",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_6() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_6",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_7() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_7",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_2_1_8() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_2_1_8",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_4_5_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_4_5_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_2_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_2_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_2_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_2_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_2_7B_F_4_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_2_7B_F_4_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_405B_FP8_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_405B_FP8_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_405B_FP8_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_405B_FP8_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_405B_FP8_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_405B_FP8_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_405B_INSTRUCT_FP8_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_405B_INSTRUCT_FP8_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_405B_INSTRUCT_FP8_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_405B_INSTRUCT_FP8_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_405B_INSTRUCT_FP8_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_405B_INSTRUCT_FP8_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_70B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_70B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_70B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_70B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_70B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_70B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_70B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_70B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_70B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_70B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_70B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_70B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_8B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_8B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_8B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_8B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_8B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_8B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_8B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_8B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_8B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_8B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_1_8B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_1_8B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_70B_INSTRUCT_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_3_8B_INSTRUCT_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_4_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_INSTRUCT_2_4_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_13B_PYTHON_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_4_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_INSTRUCT_2_4_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_34B_PYTHON_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_4_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_INSTRUCT_1_4_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_70B_PYTHON_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_4_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_INSTRUCT_2_4_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_2_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_3_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_3_1",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_4_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_CODELLAMA_7B_PYTHON_3_4_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_GUARD_3_8B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_GUARD_3_8B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_GUARD_7B_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_GUARD_7B_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_GUARD_7B_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_GUARD_7B_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_GUARD_7B_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_GUARD_7B_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_META_TEXTGENERATION_LLAMA_GUARD_7B_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"META_TEXTGENERATION_LLAMA_GUARD_7B_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_2_DEPTH_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_2_DEPTH_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_2_DEPTH_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_2_DEPTH_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V1_5_CONTROLNET_V1_1_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_DEPTH2IMG_STABLE_DIFFUSION_V2_1_CONTROLNET_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_IMAGEGENERATION_STABILITYAI_STABLE_DIFFUSION_V2_1_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_IMAGEGENERATION_STABILITYAI_STABLE_DIFFUSION_V2_1_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_IMAGEGENERATION_STABILITYAI_STABLE_DIFFUSION_XL_BASE_1_0_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_IMAGEGENERATION_STABILITYAI_STABLE_DIFFUSION_XL_BASE_1_0_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_IMAGEGENERATION_STABILITYAI_STABLE_DIFFUSION_XL_BASE_1_0_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_IMAGEGENERATION_STABILITYAI_STABLE_DIFFUSION_XL_BASE_1_0_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_RUNWAYML_STABLE_DIFFUSION_INPAINTING_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION_2_INPAINTING_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION2_INPAINTING_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_INPAINTING_STABILITYAI_STABLE_DIFFUSION2_INPAINTING_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TEXTGENERATIONJP_JAPANESE_STABLELM_INSTRUCT_ALPHA_7B_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TEXTGENERATIONJP_JAPANESE_STABLELM_INSTRUCT_ALPHA_7B_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_2",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_2_3",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_3_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_1_3_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V1_4_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_2_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_2_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_4() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_0_4",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_2",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_1_1_3",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_1_BASE_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_1() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_1",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_2() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_2",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_3() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_0_3",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_TXT2IMG_STABILITYAI_STABLE_DIFFUSION_V2_FP16_2_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_UPSCALING_STABILITYAI_STABLE_DIFFUSION_X4_UPSCALER_FP16_1_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_UPSCALING_STABILITYAI_STABLE_DIFFUSION_X4_UPSCALER_FP16_1_0_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_UPSCALING_STABILITYAI_STABLE_DIFFUSION_X4_UPSCALER_FP16_1_1_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_UPSCALING_STABILITYAI_STABLE_DIFFUSION_X4_UPSCALER_FP16_1_1_0",
		&returns,
	)
	return returns
}

func JumpStartModel_MODEL_UPSCALING_STABILITYAI_STABLE_DIFFUSION_X4_UPSCALER_FP16_2_0_0() JumpStartModel {
	_init_.Initialize()
	var returns JumpStartModel
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		"MODEL_UPSCALING_STABILITYAI_STABLE_DIFFUSION_X4_UPSCALER_FP16_2_0_0",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JumpStartModel) Bind() IJumpStartModelSpec {
	var returns IJumpStartModelSpec

	_jsii_.Invoke(
		j,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

