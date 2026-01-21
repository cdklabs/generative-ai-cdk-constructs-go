//go:build !no_runtime_type_checking

package s3vectors

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_IVectorBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVectorBucket) validateGrantDeleteParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVectorBucket) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IVectorBucket) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

