package opensearchvectorindex


// Metadata field definitions.
// Experimental.
type MetadataManagementFieldProps struct {
	// The data type of the field.
	// Experimental.
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Whether the field is filterable.
	// Experimental.
	Filterable *bool `field:"required" json:"filterable" yaml:"filterable"`
	// The name of the field.
	// Experimental.
	MappingField *string `field:"required" json:"mappingField" yaml:"mappingField"`
}

