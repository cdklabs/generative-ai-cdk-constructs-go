//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InlineApiSchema) validateBindParameters(_scope constructs.Construct) error {
	return nil
}

func validateInlineApiSchema_FromAssetParameters(path *string) error {
	return nil
}

func validateInlineApiSchema_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	return nil
}

func validateInlineApiSchema_FromInlineParameters(schema *string) error {
	return nil
}

func validateNewInlineApiSchemaParameters(schema *string) error {
	return nil
}

