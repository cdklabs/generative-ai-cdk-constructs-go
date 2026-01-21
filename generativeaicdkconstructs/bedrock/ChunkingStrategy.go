package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// Experimental.
type ChunkingStrategy interface {
	// The CloudFormation property representation of this configuration.
	// Experimental.
	Configuration() *awsbedrock.CfnDataSource_ChunkingConfigurationProperty
	// Experimental.
	SetConfiguration(val *awsbedrock.CfnDataSource_ChunkingConfigurationProperty)
}

// The jsii proxy struct for ChunkingStrategy
type jsiiProxy_ChunkingStrategy struct {
	_ byte // padding
}

func (j *jsiiProxy_ChunkingStrategy) Configuration() *awsbedrock.CfnDataSource_ChunkingConfigurationProperty {
	var returns *awsbedrock.CfnDataSource_ChunkingConfigurationProperty
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_ChunkingStrategy)SetConfiguration(val *awsbedrock.CfnDataSource_ChunkingConfigurationProperty) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

// Method for customizing a fixed sized chunking strategy.
// Experimental.
func ChunkingStrategy_FixedSize(props *awsbedrock.CfnDataSource_FixedSizeChunkingConfigurationProperty) ChunkingStrategy {
	_init_.Initialize()

	if err := validateChunkingStrategy_FixedSizeParameters(props); err != nil {
		panic(err)
	}
	var returns ChunkingStrategy

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"fixedSize",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Method for customizing a hierarchical chunking strategy.
//
// For custom chunking, the maximum token chunk size depends on the model.
// - Amazon Titan Text Embeddings: 8192
// - Cohere Embed models: 512.
// Experimental.
func ChunkingStrategy_Hierarchical(props *HierarchicalChunkingProps) ChunkingStrategy {
	_init_.Initialize()

	if err := validateChunkingStrategy_HierarchicalParameters(props); err != nil {
		panic(err)
	}
	var returns ChunkingStrategy

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"hierarchical",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Method for customizing a semantic chunking strategy.
//
// For custom chunking, the maximum token chunk size depends on the model.
// - Amazon Titan Text Embeddings: 8192
// - Cohere Embed models: 512.
// Experimental.
func ChunkingStrategy_Semantic(props *awsbedrock.CfnDataSource_SemanticChunkingConfigurationProperty) ChunkingStrategy {
	_init_.Initialize()

	if err := validateChunkingStrategy_SemanticParameters(props); err != nil {
		panic(err)
	}
	var returns ChunkingStrategy

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"semantic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func ChunkingStrategy_DEFAULT() ChunkingStrategy {
	_init_.Initialize()
	var returns ChunkingStrategy
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"DEFAULT",
		&returns,
	)
	return returns
}

func ChunkingStrategy_FIXED_SIZE() ChunkingStrategy {
	_init_.Initialize()
	var returns ChunkingStrategy
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"FIXED_SIZE",
		&returns,
	)
	return returns
}

func ChunkingStrategy_HIERARCHICAL_COHERE() ChunkingStrategy {
	_init_.Initialize()
	var returns ChunkingStrategy
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"HIERARCHICAL_COHERE",
		&returns,
	)
	return returns
}

func ChunkingStrategy_HIERARCHICAL_TITAN() ChunkingStrategy {
	_init_.Initialize()
	var returns ChunkingStrategy
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"HIERARCHICAL_TITAN",
		&returns,
	)
	return returns
}

func ChunkingStrategy_NONE() ChunkingStrategy {
	_init_.Initialize()
	var returns ChunkingStrategy
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"NONE",
		&returns,
	)
	return returns
}

func ChunkingStrategy_SEMANTIC() ChunkingStrategy {
	_init_.Initialize()
	var returns ChunkingStrategy
	_jsii_.StaticGet(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		"SEMANTIC",
		&returns,
	)
	return returns
}

