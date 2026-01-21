//go:build no_runtime_type_checking

package pinecone

// Building without runtime type checking enabled, so all the below just return nil

func validateNewPineconeVectorStoreParameters(props *PineconeVectorStoreProps) error {
	return nil
}

