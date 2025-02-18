package opensearchvectorindex

import (
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/opensearchserverless"
)

// Properties for the VectorIndex.
// Experimental.
type VectorIndexProps struct {
	// The OpenSearch Vector Collection.
	// Experimental.
	Collection opensearchserverless.VectorCollection `field:"required" json:"collection" yaml:"collection"`
	// Experimental.
	DistanceType *string `field:"required" json:"distanceType" yaml:"distanceType"`
	// The name of the index.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The metadata management fields.
	// Experimental.
	Mappings *[]*MetadataManagementFieldProps `field:"required" json:"mappings" yaml:"mappings"`
	// Experimental.
	Precision *string `field:"required" json:"precision" yaml:"precision"`
	// The number of dimensions in the vector.
	// Experimental.
	VectorDimensions *float64 `field:"required" json:"vectorDimensions" yaml:"vectorDimensions"`
	// The name of the vector field.
	// Experimental.
	VectorField *string `field:"required" json:"vectorField" yaml:"vectorField"`
	// The analyzer to use.
	// Default: - No analyzer.
	//
	// Experimental.
	Analyzer *Analyzer `field:"optional" json:"analyzer" yaml:"analyzer"`
}

