//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateAgentActionGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAgentActionGroupParameters(scope constructs.Construct, id *string, props *AgentActionGroupProps) error {
	return nil
}

