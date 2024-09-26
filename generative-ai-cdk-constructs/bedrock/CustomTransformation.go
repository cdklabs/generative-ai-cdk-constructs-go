package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a custom transformation configuration for a data source ingestion.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-chunking-parsing.html#kb-custom-transformation
//
// Experimental.
type CustomTransformation interface {
	// The CloudFormation property representation of this custom transformation configuration.
	// Experimental.
	Configuration() *awsbedrock.CfnDataSource_CustomTransformationConfigurationProperty
	// Experimental.
	SetConfiguration(val *awsbedrock.CfnDataSource_CustomTransformationConfigurationProperty)
	// Experimental.
	GeneratePolicyStatements(scope constructs.Construct) *[]awsiam.PolicyStatement
}

// The jsii proxy struct for CustomTransformation
type jsiiProxy_CustomTransformation struct {
	_ byte // padding
}

func (j *jsiiProxy_CustomTransformation) Configuration() *awsbedrock.CfnDataSource_CustomTransformationConfigurationProperty {
	var returns *awsbedrock.CfnDataSource_CustomTransformationConfigurationProperty
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomTransformation_Override(c CustomTransformation) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomTransformation",
		nil, // no parameters
		c,
	)
}

func (j *jsiiProxy_CustomTransformation)SetConfiguration(val *awsbedrock.CfnDataSource_CustomTransformationConfigurationProperty) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

// This feature allows you to use a Lambda function to inject your own logic into the knowledge base ingestion process.
// See: https://github.com/aws-samples/amazon-bedrock-samples/blob/main/knowledge-bases/features-examples/02-optimizing-accuracy-retrieved-results/advanced_chunking_options.ipynb
//
// Experimental.
func CustomTransformation_Lambda(props *LambdaCustomTransformationProps) CustomTransformation {
	_init_.Initialize()

	if err := validateCustomTransformation_LambdaParameters(props); err != nil {
		panic(err)
	}
	var returns CustomTransformation

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomTransformation",
		"lambda",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomTransformation) GeneratePolicyStatements(scope constructs.Construct) *[]awsiam.PolicyStatement {
	if err := c.validateGeneratePolicyStatementsParameters(scope); err != nil {
		panic(err)
	}
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		c,
		"generatePolicyStatements",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

