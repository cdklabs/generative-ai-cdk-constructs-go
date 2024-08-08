//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiSchema) validateBindParameters(scope constructs.Construct) error {
	return nil
}

func validateApiSchema_FromAssetParameters(path *string) error {
	return nil
}

func validateApiSchema_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

