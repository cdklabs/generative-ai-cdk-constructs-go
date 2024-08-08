//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BaseClass) validateAddObservabilityToConstructParameters(props *BaseClassProps) error {
	return nil
}

func (b *jsiiProxy_BaseClass) validateUpdateConstructUsageMetricCodeParameters(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) error {
	return nil
}

func (b *jsiiProxy_BaseClass) validateUpdateEnvSuffixParameters(props *BaseClassProps) error {
	return nil
}

func validateBaseClass_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BaseClass) validateSetEnablexrayParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_BaseClass) validateSetFieldLogLevelParameters(val awsappsync.FieldLogLevel) error {
	return nil
}

func (j *jsiiProxy_BaseClass) validateSetLambdaTracingParameters(val awslambda.Tracing) error {
	return nil
}

func (j *jsiiProxy_BaseClass) validateSetRetentionParameters(val awslogs.RetentionDays) error {
	return nil
}

func (j *jsiiProxy_BaseClass) validateSetStageParameters(val *string) error {
	return nil
}

func validateBaseClass_SetUsageMetricMapParameters(val *map[string]*float64) error {
	return nil
}

func validateNewBaseClassParameters(scope constructs.Construct, id *string) error {
	return nil
}

