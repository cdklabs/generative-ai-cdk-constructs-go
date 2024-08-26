//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockCwDashboard) validateAddAllModelsMonitoringParameters(props *ModelMonitoringProps) error {
	return nil
}

func (b *jsiiProxy_BedrockCwDashboard) validateAddModelMonitoringParameters(modelName *string, modelId *string, props *ModelMonitoringProps) error {
	return nil
}

func validateBedrockCwDashboard_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBedrockCwDashboardParameters(scope constructs.Construct, id *string, props *BedrockCwDashboardProps) error {
	return nil
}

