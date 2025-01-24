//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAgentAlias) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAgentAlias) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAgentAlias) validateOnCloudTrailEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

