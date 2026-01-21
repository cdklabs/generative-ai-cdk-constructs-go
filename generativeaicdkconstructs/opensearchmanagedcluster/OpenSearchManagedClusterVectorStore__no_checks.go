//go:build no_runtime_type_checking

package opensearchmanagedcluster

// Building without runtime type checking enabled, so all the below just return nil

func validateNewOpenSearchManagedClusterVectorStoreParameters(props *OpenSearchManagedClusterVectorStoreProps) error {
	return nil
}

