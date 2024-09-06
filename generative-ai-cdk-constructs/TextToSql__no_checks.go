//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TextToSql) validateAddObservabilityToConstructParameters(props *BaseClassProps) error {
	return nil
}

func (t *jsiiProxy_TextToSql) validateUpdateConstructUsageMetricCodeParameters(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) error {
	return nil
}

func (t *jsiiProxy_TextToSql) validateUpdateEnvSuffixParameters(props *BaseClassProps) error {
	return nil
}

func validateTextToSql_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_TextToSql) validateSetEnablexrayParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_TextToSql) validateSetFieldLogLevelParameters(val awsappsync.FieldLogLevel) error {
	return nil
}

func (j *jsiiProxy_TextToSql) validateSetLambdaTracingParameters(val awslambda.Tracing) error {
	return nil
}

func (j *jsiiProxy_TextToSql) validateSetRetentionParameters(val awslogs.RetentionDays) error {
	return nil
}

func (j *jsiiProxy_TextToSql) validateSetStageParameters(val *string) error {
	return nil
}

func validateTextToSql_SetUsageMetricMapParameters(val *map[string]*float64) error {
	return nil
}

func validateNewTextToSqlParameters(scope constructs.Construct, id *string, props *TextToSqlProps) error {
	return nil
}

