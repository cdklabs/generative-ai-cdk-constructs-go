package mongodbatlas


// Interface for MongoDB Atlas field mapping.
// Experimental.
type MongoDbAtlasFieldMapping struct {
	// The field name for the metadata field.
	// Experimental.
	MetadataField *string `field:"required" json:"metadataField" yaml:"metadataField"`
	// The field name for the text field.
	// Experimental.
	TextField *string `field:"required" json:"textField" yaml:"textField"`
	// The field name for the vector field.
	// Experimental.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
}

