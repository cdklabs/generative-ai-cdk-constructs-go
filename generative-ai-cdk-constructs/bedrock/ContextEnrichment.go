package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Abstract class representing a context enrichment strategy.
//
// The enrichment stategy used to provide additional context.
// For example, Neptune GraphRAG uses Amazon Bedrock foundation
// models to perform chunk entity extraction.
// Experimental.
type ContextEnrichment interface {
	// The CloudFormation property representation of this configuration.
	// Experimental.
	Configuration() *awsbedrock.CfnDataSource_ContextEnrichmentConfigurationProperty
	// Experimental.
	SetConfiguration(val *awsbedrock.CfnDataSource_ContextEnrichmentConfigurationProperty)
	// Experimental.
	GeneratePolicyStatements() *[]awsiam.PolicyStatement
}

// The jsii proxy struct for ContextEnrichment
type jsiiProxy_ContextEnrichment struct {
	_ byte // padding
}

func (j *jsiiProxy_ContextEnrichment) Configuration() *awsbedrock.CfnDataSource_ContextEnrichmentConfigurationProperty {
	var returns *awsbedrock.CfnDataSource_ContextEnrichmentConfigurationProperty
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}


// Experimental.
func NewContextEnrichment_Override(c ContextEnrichment) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextEnrichment",
		nil, // no parameters
		c,
	)
}

func (j *jsiiProxy_ContextEnrichment)SetConfiguration(val *awsbedrock.CfnDataSource_ContextEnrichmentConfigurationProperty) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

// Creates a Foundation Model-based enrichment strategy used to provide additional context to the RAG application.
// Experimental.
func ContextEnrichment_FoundationModel(props *FoundationModelContextEnrichmentProps) ContextEnrichment {
	_init_.Initialize()

	if err := validateContextEnrichment_FoundationModelParameters(props); err != nil {
		panic(err)
	}
	var returns ContextEnrichment

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextEnrichment",
		"foundationModel",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContextEnrichment) GeneratePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		c,
		"generatePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

