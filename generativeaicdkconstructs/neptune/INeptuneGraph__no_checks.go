//go:build no_runtime_type_checking

package neptune

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_INeptuneGraph) validateCreateNotebookParameters(params *NeptuneGraphNotebookProps) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateGrantParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateGrantExportTaskParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateGrantReadOnlyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricGraphSizeBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricGraphStorageUsagePercentParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumEdgePropertiesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumEdgesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumOpenCypherClientErrorsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumOpenCypherRequestsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumOpenCypherServerErrorsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumQueuedRequestsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumVectorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_INeptuneGraph) validateMetricNumVertexPropertiesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

