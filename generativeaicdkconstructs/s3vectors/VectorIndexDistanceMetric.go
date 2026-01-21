package s3vectors


// The distance metric to be used for similarity search.
// Experimental.
type VectorIndexDistanceMetric string

const (
	// Measures the straight-line distance between two points in multi-dimensional space.
	//
	// Lower values indicate greater similarity.
	// Experimental.
	VectorIndexDistanceMetric_EUCLIDEAN VectorIndexDistanceMetric = "EUCLIDEAN"
	// Measures the cosine of the angle between two vectors.
	// Experimental.
	VectorIndexDistanceMetric_COSINE VectorIndexDistanceMetric = "COSINE"
)

