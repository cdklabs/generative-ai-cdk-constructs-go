package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents an advanced parsing strategy configuration for Knowledge Base ingestion.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-chunking-parsing.html#kb-advanced-parsing
//
// Experimental.
type ParsingStategy interface {
	// The CloudFormation property representation of this configuration.
	// Experimental.
	Configuration() *awsbedrock.CfnDataSource_ParsingConfigurationProperty
	// Experimental.
	SetConfiguration(val *awsbedrock.CfnDataSource_ParsingConfigurationProperty)
	// Experimental.
	GeneratePolicyStatements() *[]awsiam.PolicyStatement
}

// The jsii proxy struct for ParsingStategy
type jsiiProxy_ParsingStategy struct {
	_ byte // padding
}

func (j *jsiiProxy_ParsingStategy) Configuration() *awsbedrock.CfnDataSource_ParsingConfigurationProperty {
	var returns *awsbedrock.CfnDataSource_ParsingConfigurationProperty
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}


// Experimental.
func NewParsingStategy_Override(p ParsingStategy) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParsingStategy",
		nil, // no parameters
		p,
	)
}

func (j *jsiiProxy_ParsingStategy)SetConfiguration(val *awsbedrock.CfnDataSource_ParsingConfigurationProperty) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

// Creates a Foundation Model-based parsing strategy for extracting non-textual information from documents such as tables and charts.
//
// - Additional costs apply when using advanced parsing due to foundation model usage.
// - There are limits on file types (PDF) and total data that can be parsed using advanced parsing.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-ds.html#kb-ds-supported-doc-formats-limits
//
// Experimental.
func ParsingStategy_FoundationModel(props *FoundationModelParsingStategyProps) ParsingStategy {
	_init_.Initialize()

	if err := validateParsingStategy_FoundationModelParameters(props); err != nil {
		panic(err)
	}
	var returns ParsingStategy

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParsingStategy",
		"foundationModel",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ParsingStategy) GeneratePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		p,
		"generatePolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

