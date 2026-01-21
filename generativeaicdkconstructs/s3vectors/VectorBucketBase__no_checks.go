//go:build no_runtime_type_checking

package s3vectors

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorBucketBase) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (v *jsiiProxy_VectorBucketBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VectorBucketBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VectorBucketBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VectorBucketBase) validateGrantDeleteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_VectorBucketBase) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_VectorBucketBase) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateVectorBucketBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVectorBucketBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVectorBucketBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_VectorBucketBase) validateSetAutoCreatePolicyParameters(val *bool) error {
	return nil
}

func validateNewVectorBucketBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

