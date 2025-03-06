//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockDataAutomation) validateAddObservabilityToConstructParameters(props *BaseClassProps) error {
	return nil
}

func (b *jsiiProxy_BedrockDataAutomation) validateUpdateConstructUsageMetricCodeParameters(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) error {
	return nil
}

func (b *jsiiProxy_BedrockDataAutomation) validateUpdateEnvSuffixParameters(props *BaseClassProps) error {
	return nil
}

func validateBedrockDataAutomation_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockDataAutomation) validateSetEnablexrayParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_BedrockDataAutomation) validateSetFieldLogLevelParameters(val awsappsync.FieldLogLevel) error {
	return nil
}

func (j *jsiiProxy_BedrockDataAutomation) validateSetLambdaTracingParameters(val awslambda.Tracing) error {
	return nil
}

func (j *jsiiProxy_BedrockDataAutomation) validateSetRetentionParameters(val awslogs.RetentionDays) error {
	return nil
}

func (j *jsiiProxy_BedrockDataAutomation) validateSetStageParameters(val *string) error {
	return nil
}

func validateBedrockDataAutomation_SetUsageMetricMapParameters(val *map[string]*float64) error {
	return nil
}

func validateNewBedrockDataAutomationParameters(scope constructs.Construct, id *string, props *BedrockDataAutomationProps) error {
	return nil
}

