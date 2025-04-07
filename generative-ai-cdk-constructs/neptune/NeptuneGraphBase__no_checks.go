//go:build no_runtime_type_checking

package neptune

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NeptuneGraphBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateCreateNotebookParameters(params *NeptuneGraphNotebookProps) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGrantParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGrantExportTaskParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGrantQueryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateGrantReadOnlyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricGraphSizeBytesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricGraphStorageUsagePercentParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumEdgePropertiesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumEdgesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumOpenCypherClientErrorsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumOpenCypherRequestsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumOpenCypherServerErrorsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumQueuedRequestsPerSecParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumVectorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphBase) validateMetricNumVertexPropertiesParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNeptuneGraphBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNeptuneGraphBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNeptuneGraphBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNeptuneGraphBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

