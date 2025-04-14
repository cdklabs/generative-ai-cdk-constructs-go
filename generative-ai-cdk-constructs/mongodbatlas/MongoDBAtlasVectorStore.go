package mongodbatlas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// Construct for MongoDB Atlas vector store.
// Experimental.
type MongoDBAtlasVectorStore interface {
	// The name of the collection.
	// Experimental.
	CollectionName() *string
	// The ARN of the secret containing MongoDB Atlas credentials.
	// Experimental.
	CredentialsSecretArn() *string
	// The name of the database.
	// Experimental.
	DatabaseName() *string
	// The endpoint URL for MongoDB Atlas.
	// Experimental.
	Endpoint() *string
	// The name of the endpoint service.
	// Experimental.
	EndpointServiceName() *string
	// The field mapping for MongoDB Atlas.
	// Experimental.
	FieldMapping() *MongoDbAtlasFieldMapping
	// The name of the vector index.
	// Experimental.
	VectorIndexName() *string
}

// The jsii proxy struct for MongoDBAtlasVectorStore
type jsiiProxy_MongoDBAtlasVectorStore struct {
	_ byte // padding
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) CollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) CredentialsSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) EndpointServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) FieldMapping() *MongoDbAtlasFieldMapping {
	var returns *MongoDbAtlasFieldMapping
	_jsii_.Get(
		j,
		"fieldMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MongoDBAtlasVectorStore) VectorIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexName",
		&returns,
	)
	return returns
}


// Creates a new instance of the MongoDBAtlas class.
// Experimental.
func NewMongoDBAtlasVectorStore(props *MongoDBAtlasVectorStoreProps) MongoDBAtlasVectorStore {
	_init_.Initialize()

	if err := validateNewMongoDBAtlasVectorStoreParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MongoDBAtlasVectorStore{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.mongodbAtlas.MongoDBAtlasVectorStore",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Creates a new instance of the MongoDBAtlas class.
// Experimental.
func NewMongoDBAtlasVectorStore_Override(m MongoDBAtlasVectorStore, props *MongoDBAtlasVectorStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.mongodbAtlas.MongoDBAtlasVectorStore",
		[]interface{}{props},
		m,
	)
}

