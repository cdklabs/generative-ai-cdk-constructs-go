//go:build no_runtime_type_checking

package s3vectors

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorBucketPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VectorBucketPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VectorBucketPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVectorBucketPolicy_FromCfnVectorBucketPolicyParameters(cfnVectorBucketPolicy awss3vectors.CfnVectorBucketPolicy) error {
	return nil
}

func validateVectorBucketPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVectorBucketPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVectorBucketPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVectorBucketPolicyParameters(scope constructs.Construct, id *string, props *VectorBucketPolicyProps) error {
	return nil
}

