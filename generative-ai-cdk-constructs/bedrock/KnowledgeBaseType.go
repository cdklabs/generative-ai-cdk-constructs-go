package bedrock


// Types of possible knowledge bases supported by Amazon Bedrock Knowledge Bases.
// Experimental.
type KnowledgeBaseType string

const (
	// Vector database with emebeddings vectors.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-how-it-works.html
	//
	// Experimental.
	KnowledgeBaseType_VECTOR KnowledgeBaseType = "VECTOR"
	// Kendra GenAI Index.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-kendra-genai-index.html
	//
	// Experimental.
	KnowledgeBaseType_KENDRA KnowledgeBaseType = "KENDRA"
	// Structured data store (e.g. REDSHIFT).
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-build-structured.html
	//
	// Experimental.
	KnowledgeBaseType_SQL KnowledgeBaseType = "SQL"
)

