//go:build no_runtime_type_checking

package opensearchserverless

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VectorCollection) validateGrantDataAccessParameters(grantee awsiam.IRole) error {
	return nil
}

func validateVectorCollection_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VectorCollection) validateSetAossPolicyParameters(val awsiam.ManagedPolicy) error {
	return nil
}

func (j *jsiiProxy_VectorCollection) validateSetCollectionArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VectorCollection) validateSetCollectionIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VectorCollection) validateSetCollectionNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_VectorCollection) validateSetDataAccessPolicyParameters(val awsopensearchserverless.CfnAccessPolicy) error {
	return nil
}

func (j *jsiiProxy_VectorCollection) validateSetStandbyReplicasParameters(val VectorCollectionStandbyReplicas) error {
	return nil
}

func validateNewVectorCollectionParameters(scope constructs.Construct, id *string, props *VectorCollectionProps) error {
	return nil
}

