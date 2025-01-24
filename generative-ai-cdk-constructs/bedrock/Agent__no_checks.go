//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Agent) validateAddActionGroupParameters(actionGroup AgentActionGroup) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddGuardrailParameters(guardrail IGuardrail) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddKnowledgeBaseParameters(knowledgeBase IKnowledgeBase) error {
	return nil
}

func (a *jsiiProxy_Agent) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Agent) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Agent) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAgent_FromAgentAttrsParameters(scope constructs.Construct, id *string, attrs *AgentAttributes) error {
	return nil
}

func validateAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAgent_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAgent_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_Agent) validateSetActionGroupsParameters(val *[]AgentActionGroup) error {
	return nil
}

func (j *jsiiProxy_Agent) validateSetAgentVersionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Agent) validateSetKnowledgeBasesParameters(val *[]IKnowledgeBase) error {
	return nil
}

func validateNewAgentParameters(scope constructs.Construct, id *string, props *AgentProps) error {
	return nil
}

