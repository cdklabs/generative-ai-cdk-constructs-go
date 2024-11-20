package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Cross-region inference enables you to seamlessly manage unplanned traffic bursts by utilizing compute across different AWS Regions.
//
// With cross-region
// inference, you can distribute traffic across multiple AWS Regions, enabling
// higher throughput and enhanced resilience during periods of peak demands.
// See: https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html
//
// Experimental.
type CrossRegionInferenceProfile interface {
	IInferenceProfile
	IInvokable
	// The ARN of the application inference profile.
	//
	// Example:
	//   'arn:aws:bedrock:us-east-1:123456789012:inference-profile/us.anthropic.claude-3-5-sonnet-20240620-v1:0'
	//
	// Experimental.
	InferenceProfileArn() *string
	// The unique identifier of the inference profile.
	//
	// Example:
	//   'us.anthropic.claude-3-5-sonnet-20240620-v1:0'
	//
	// Experimental.
	InferenceProfileId() *string
	// The underlying model supporting cross-region inference.
	// Experimental.
	InferenceProfileModel() BedrockFoundationModel
	// This equals to the inferenceProfileArn property, useful just to implement IInvokable interface.
	// Experimental.
	InvokableArn() *string
	// The type of inference profile.
	//
	// Example:
	//   InferenceProfileType.SYSTEM_DEFINED
	//
	// Experimental.
	Type() InferenceProfileType
	// Gives the appropriate policies to invoke and use the Foundation Model.
	// Experimental.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grants appropriate permissions to use the cross-region inference profile.
	//
	// Does not grant permissions to use the model in the profile.
	// Experimental.
	GrantProfileUsage(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy struct for CrossRegionInferenceProfile
type jsiiProxy_CrossRegionInferenceProfile struct {
	jsiiProxy_IInferenceProfile
	jsiiProxy_IInvokable
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InferenceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InferenceProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InferenceProfileModel() BedrockFoundationModel {
	var returns BedrockFoundationModel
	_jsii_.Get(
		j,
		"inferenceProfileModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) InvokableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CrossRegionInferenceProfile) Type() InferenceProfileType {
	var returns InferenceProfileType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func CrossRegionInferenceProfile_FromConfig(config *CrossRegionInferenceProfileProps) CrossRegionInferenceProfile {
	_init_.Initialize()

	if err := validateCrossRegionInferenceProfile_FromConfigParameters(config); err != nil {
		panic(err)
	}
	var returns CrossRegionInferenceProfile

	_jsii_.StaticInvoke(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CrossRegionInferenceProfile",
		"fromConfig",
		[]interface{}{config},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossRegionInferenceProfile) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CrossRegionInferenceProfile) GrantProfileUsage(grantee awsiam.IGrantable) awsiam.Grant {
	if err := c.validateGrantProfileUsageParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		c,
		"grantProfileUsage",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

