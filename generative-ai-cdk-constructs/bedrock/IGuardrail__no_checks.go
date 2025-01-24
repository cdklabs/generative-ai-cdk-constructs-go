//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGuardrail) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGuardrail) validateGrantApplyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (j *jsiiProxy_IGuardrail) validateSetGuardrailVersionParameters(val *string) error {
	return nil
}

