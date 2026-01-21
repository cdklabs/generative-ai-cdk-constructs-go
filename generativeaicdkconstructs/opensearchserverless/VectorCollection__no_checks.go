//go:build no_runtime_type_checking

package opensearchserverless

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorCollection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateGrantDataAccessParameters(grantee awsiam.IRole) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateMetricIndexRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateMetricSearchLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateMetricSearchLatencyP90Parameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VectorCollection) validateMetricSearchRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateVectorCollection_FromCollectionAttributesParameters(constructScope constructs.Construct, constructId *string, attrs *VectorCollectionAttributes) error {
	return nil
}

func validateVectorCollection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVectorCollection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVectorCollection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVectorCollection_MetricAllParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateVectorCollection_MetricAllIndexRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateVectorCollection_MetricAllSearchLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateVectorCollection_MetricAllSearchRequestCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateNewVectorCollectionParameters(scope constructs.Construct, id *string, props *VectorCollectionProps) error {
	return nil
}

