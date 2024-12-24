//go:build no_runtime_type_checking

package opensearchserverless

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVectorCollection) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVectorCollection) validateMetricIndexRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVectorCollection) validateMetricSearchLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVectorCollection) validateMetricSearchLatencyP90Parameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IVectorCollection) validateMetricSearchRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

