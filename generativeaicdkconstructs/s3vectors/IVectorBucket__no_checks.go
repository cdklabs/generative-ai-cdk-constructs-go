//go:build no_runtime_type_checking

package s3vectors

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVectorBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IVectorBucket) validateGrantDeleteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVectorBucket) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IVectorBucket) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

