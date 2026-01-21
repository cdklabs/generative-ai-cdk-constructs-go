//go:build no_runtime_type_checking

package s3vectors

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (v *jsiiProxy_VectorBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VectorBucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VectorBucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VectorBucket) validateGrantDeleteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_VectorBucket) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (v *jsiiProxy_VectorBucket) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateVectorBucket_FromCfnVectorBucketParameters(cfnVectorBucket awss3vectors.CfnVectorBucket) error {
	return nil
}

func validateVectorBucket_FromVectorBucketArnParameters(scope constructs.Construct, id *string, vectorBucketArn *string) error {
	return nil
}

func validateVectorBucket_FromVectorBucketAttributesParameters(scope constructs.Construct, id *string, attrs *VectorBucketAttributes) error {
	return nil
}

func validateVectorBucket_FromVectorBucketNameParameters(scope constructs.Construct, id *string, vectorBucketName *string) error {
	return nil
}

func validateVectorBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVectorBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVectorBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_VectorBucket) validateSetAutoCreatePolicyParameters(val *bool) error {
	return nil
}

func validateNewVectorBucketParameters(scope constructs.Construct, id *string, props *VectorBucketProps) error {
	return nil
}

