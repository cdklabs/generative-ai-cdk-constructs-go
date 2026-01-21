package s3vectors


// The data type of the vectors to be inserted into the vector index.
// Experimental.
type VectorIndexDataType string

const (
	// 32-bit floating-point numbers.
	// Experimental.
	VectorIndexDataType_FLOAT_32 VectorIndexDataType = "FLOAT_32"
)

