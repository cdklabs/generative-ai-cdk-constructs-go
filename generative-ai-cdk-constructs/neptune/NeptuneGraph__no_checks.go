//go:build no_runtime_type_checking

package neptune

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NeptuneGraph) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateCreateNotebookParameters(params *NeptuneGraphNotebookProps) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGrantParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGrantExportTaskParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateGrantReadOnlyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricGraphSizeBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricGraphStorageUsagePercentParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumEdgePropertiesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumEdgesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumOpenCypherClientErrorsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumOpenCypherRequestsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumOpenCypherServerErrorsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumQueuedRequestsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumVectorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraph) validateMetricNumVertexPropertiesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNeptuneGraph_FromGraphIdParameters(scope constructs.Construct, id *string, graphId *string) error {
	return nil
}

func validateNeptuneGraph_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNeptuneGraph_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNeptuneGraph_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNeptuneGraphParameters(scope constructs.Construct, id *string, props *NeptuneGraphProps) error {
	return nil
}

