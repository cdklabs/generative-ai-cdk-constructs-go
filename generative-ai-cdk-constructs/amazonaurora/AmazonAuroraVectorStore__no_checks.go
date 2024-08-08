//go:build no_runtime_type_checking

package amazonaurora

// Building without runtime type checking enabled, so all the below just return nil

func validateNewAmazonAuroraVectorStoreParameters(props *AmazonAuroraVectorStoreProps) error {
	return nil
}

