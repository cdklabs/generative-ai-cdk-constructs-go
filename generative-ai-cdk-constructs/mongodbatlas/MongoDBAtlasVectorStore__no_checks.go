//go:build no_runtime_type_checking

package mongodbatlas

// Building without runtime type checking enabled, so all the below just return nil

func validateNewMongoDBAtlasVectorStoreParameters(props *MongoDBAtlasVectorStoreProps) error {
	return nil
}

