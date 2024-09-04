//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func validateAgentAlias_FromAliasArnParameters(scope constructs.Construct, id *string, aliasArn *string) error {
	return nil
}

func validateAgentAlias_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAgentAliasParameters(scope constructs.Construct, id *string, props *AgentAliasProps) error {
	return nil
}

