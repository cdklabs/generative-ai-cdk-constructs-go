//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3ApiSchema) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateS3ApiSchema_FromAssetParameters(path *string) error {
	return nil
}

func validateS3ApiSchema_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateS3ApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateNewS3ApiSchemaParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

