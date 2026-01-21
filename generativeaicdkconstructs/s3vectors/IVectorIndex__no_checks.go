//go:build no_runtime_type_checking

package s3vectors

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVectorIndex) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

