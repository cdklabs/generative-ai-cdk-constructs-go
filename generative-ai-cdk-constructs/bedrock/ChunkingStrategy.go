package bedrock


// Knowledge base can split your source data into chunks.
//
// A chunk refers to an
// excerpt from a data source that is returned when the knowledge base that it
// belongs to is queried. You have the following options for chunking your
// data. If you opt for NONE, then you may want to pre-process your files by
// splitting them up such that each file corresponds to a chunk.
// Experimental.
type ChunkingStrategy string

const (
	// Amazon Bedrock splits your source data into chunks of the approximate size that you set in the `fixedSizeChunkingConfiguration`.
	// Experimental.
	ChunkingStrategy_FIXED_SIZE ChunkingStrategy = "FIXED_SIZE"
	// `FIXED_SIZE` with the default chunk size of 300 tokens and 20% overlap.
	//
	// If default is selected, chunk size and overlap set by the user will be
	// ignored.
	// Experimental.
	ChunkingStrategy_DEFAULT ChunkingStrategy = "DEFAULT"
	// Amazon Bedrock treats each file as one chunk.
	//
	// If you choose this option,
	// you may want to pre-process your documents by splitting them into separate
	// files.
	// Experimental.
	ChunkingStrategy_NONE ChunkingStrategy = "NONE"
)

