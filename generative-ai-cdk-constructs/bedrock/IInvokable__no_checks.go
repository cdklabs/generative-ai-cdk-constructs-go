//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IInvokable) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}
