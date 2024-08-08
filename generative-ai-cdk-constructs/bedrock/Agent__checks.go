//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_Agent) validateAddActionGroupParameters(actionGroup AgentActionGroup) error {
	if actionGroup == nil {
		return fmt.Errorf("parameter actionGroup is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Agent) validateAddActionGroupsParameters(actionGroups *[]AgentActionGroup) error {
	if actionGroups == nil {
		return fmt.Errorf("parameter actionGroups is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Agent) validateAddAliasParameters(props *AddAgentAliasProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_Agent) validateAddGuardrailParameters(guardrail Guardrail) error {
	if guardrail == nil {
		return fmt.Errorf("parameter guardrail is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Agent) validateAddKnowledgeBaseParameters(knowledgeBase KnowledgeBase) error {
	if knowledgeBase == nil {
		return fmt.Errorf("parameter knowledgeBase is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_Agent) validateAddKnowledgeBasesParameters(knowledgeBases *[]KnowledgeBase) error {
	if knowledgeBases == nil {
		return fmt.Errorf("parameter knowledgeBases is required, but nil was provided")
	}

	return nil
}

func validateAgent_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Agent) validateSetActionGroupsParameters(val *[]*awsbedrock.CfnAgent_AgentActionGroupProperty) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func (j *jsiiProxy_Agent) validateSetKnowledgeBasesParameters(val *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	for idx_97dfc6, v := range *val {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNewAgentParameters(scope constructs.Construct, id *string, props *AgentProps) error {
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

