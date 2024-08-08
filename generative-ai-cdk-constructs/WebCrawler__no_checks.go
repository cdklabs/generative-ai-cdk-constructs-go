//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WebCrawler) validateAddObservabilityToConstructParameters(props *BaseClassProps) error {
	return nil
}

func (w *jsiiProxy_WebCrawler) validateUpdateConstructUsageMetricCodeParameters(props *BaseClassProps, scope constructs.Construct, lambdaFunctions *[]awslambda.DockerImageFunction) error {
	return nil
}

func (w *jsiiProxy_WebCrawler) validateUpdateEnvSuffixParameters(props *BaseClassProps) error {
	return nil
}

func validateWebCrawler_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_WebCrawler) validateSetEnablexrayParameters(val *bool) error {
	return nil
}

func (j *jsiiProxy_WebCrawler) validateSetFieldLogLevelParameters(val awsappsync.FieldLogLevel) error {
	return nil
}

func (j *jsiiProxy_WebCrawler) validateSetLambdaTracingParameters(val awslambda.Tracing) error {
	return nil
}

func (j *jsiiProxy_WebCrawler) validateSetRetentionParameters(val awslogs.RetentionDays) error {
	return nil
}

func (j *jsiiProxy_WebCrawler) validateSetStageParameters(val *string) error {
	return nil
}

func validateWebCrawler_SetUsageMetricMapParameters(val *map[string]*float64) error {
	return nil
}

func validateNewWebCrawlerParameters(scope constructs.Construct, id *string, props *WebCrawlerProps) error {
	return nil
}

