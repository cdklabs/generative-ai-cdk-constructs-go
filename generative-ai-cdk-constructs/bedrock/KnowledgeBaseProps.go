package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/opensearchvectorindex"
)

// Properties for a knowledge base.
// Experimental.
type KnowledgeBaseProps struct {
	// The embeddings model for the knowledge base.
	// Experimental.
	EmbeddingsModel BedrockFoundationModel `field:"required" json:"embeddingsModel" yaml:"embeddingsModel"`
	// The description of the knowledge base.
	// Default: - No description provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Existing IAM role with a policy statement granting permission to invoke the specific embeddings model.
	//
	// Any entity (e.g., an AWS service or application) that assumes
	// this role will be able to invoke or use the
	// specified embeddings model within the Bedrock service.
	// Experimental.
	ExistingRole awsiam.IRole `field:"optional" json:"existingRole" yaml:"existingRole"`
	// The name of the vector index.
	//
	// If vectorStore is not of type `VectorCollection`,
	// do not include this property as it will throw error.
	// Default: - 'bedrock-knowledge-base-default-index'.
	//
	// Experimental.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// A narrative description of the knowledge base.
	//
	// A Bedrock Agent can use this instruction to determine if it should
	// query this Knowledge Base.
	// Default: - No description provided.
	//
	// Experimental.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// Specifies whether to use the knowledge base or not when sending an InvokeAgent request.
	// Experimental.
	KnowledgeBaseState *string `field:"optional" json:"knowledgeBaseState" yaml:"knowledgeBaseState"`
	// The name of the knowledge base.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// OPTIONAL: Tag (KEY-VALUE) bedrock agent resource.
	// Default: - false.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The name of the field in the vector index.
	//
	// If vectorStore is not of type `VectorCollection`,
	// do not include this property as it will throw error.
	// Default: - 'bedrock-knowledge-base-default-vector'.
	//
	// Experimental.
	VectorField *string `field:"optional" json:"vectorField" yaml:"vectorField"`
	// The vector index for the OpenSearch Serverless backed knowledge base.
	//
	// If vectorStore is not of type `VectorCollection`, do not include
	// this property as it will throw error.
	// Default: - A new vector index is created on the Vector Collection
	// if vector store is of `VectorCollection` type.
	//
	// Experimental.
	VectorIndex opensearchvectorindex.VectorIndex `field:"optional" json:"vectorIndex" yaml:"vectorIndex"`
	// The vector store for the knowledge base.
	//
	// Must be either of
	// type `VectorCollection`, `RedisEnterpriseVectorStore`,
	// `PineconeVectorStore`, `AmazonAuroraVectorStore` or
	// `AmazonAuroraDefaultVectorStore`.
	// Default: - A new OpenSearch Serverless vector collection is created.
	//
	// Experimental.
	VectorStore interface{} `field:"optional" json:"vectorStore" yaml:"vectorStore"`
}

