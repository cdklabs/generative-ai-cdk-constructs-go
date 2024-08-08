//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_Guardrail) validateAddSensitiveInformationPolicyConfigParameters(props *[]*SensitiveInformationPolicyConfigProps, guardrailRegexesConfig *awsbedrock.CfnGuardrail_RegexConfigProperty) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	for idx_b26480, v := range *props {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter props[%#v]", idx_b26480) }); err != nil {
			return err
		}
	}

	if guardrailRegexesConfig == nil {
		return fmt.Errorf("parameter guardrailRegexesConfig is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(guardrailRegexesConfig, func() string { return "parameter guardrailRegexesConfig" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddTagsParameters(props *GuardrailProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddTopicPolicyConfigParameters(topic Topic) error {
	if topic == nil {
		return fmt.Errorf("parameter topic is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddVersionParameters(id *string) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateAddWordPolicyConfigParameters(wordsFilter *[]*awsbedrock.CfnGuardrail_WordConfigProperty) error {
	for idx_c24105, v := range *wordsFilter {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter wordsFilter[%#v]", idx_c24105) }); err != nil {
			return err
		}
	}

	return nil
}

func (g *jsiiProxy_Guardrail) validateUploadWordPolicyFromFileParameters(filePath *string) error {
	if filePath == nil {
		return fmt.Errorf("parameter filePath is required, but nil was provided")
	}

	return nil
}

func validateGuardrail_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewGuardrailParameters(scope constructs.Construct, id *string, props *GuardrailProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

