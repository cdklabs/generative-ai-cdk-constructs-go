//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Agent) validateAddActionGroupParameters(actionGroup AgentActionGroup) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddActionGroupsParameters(actionGroups *[]AgentActionGroup) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddAliasParameters(props *AddAgentAliasProps) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddGuardrailParameters(guardrail Guardrail) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddKnowledgeBaseParameters(knowledgeBase KnowledgeBase) error {
	return nil
}

func (a *jsiiProxy_Agent) validateAddKnowledgeBasesParameters(knowledgeBases *[]KnowledgeBase) error {
	return nil
}

func validateAgent_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Agent) validateSetActionGroupsParameters(val *[]*awsbedrock.CfnAgent_AgentActionGroupProperty) error {
	return nil
}

func (j *jsiiProxy_Agent) validateSetKnowledgeBasesParameters(val *[]*awsbedrock.CfnAgent_AgentKnowledgeBaseProperty) error {
	return nil
}

func validateNewAgentParameters(scope constructs.Construct, id *string, props *AgentProps) error {
	return nil
}

