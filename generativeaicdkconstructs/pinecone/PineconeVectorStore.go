package pinecone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"
)

// Class to define a PineconeVectorStore.
// Experimental.
type PineconeVectorStore interface {
	// Experimental.
	ConnectionString() *string
	// Experimental.
	CredentialsSecretArn() *string
	// Experimental.
	KmsKey() *string
	// Experimental.
	MetadataField() *string
	// Experimental.
	Namespace() *string
	// Experimental.
	TextField() *string
}

// The jsii proxy struct for PineconeVectorStore
type jsiiProxy_PineconeVectorStore struct {
	_ byte // padding
}

func (j *jsiiProxy_PineconeVectorStore) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PineconeVectorStore) CredentialsSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PineconeVectorStore) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PineconeVectorStore) MetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PineconeVectorStore) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PineconeVectorStore) TextField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textField",
		&returns,
	)
	return returns
}


// Experimental.
func NewPineconeVectorStore(props *PineconeVectorStoreProps) PineconeVectorStore {
	_init_.Initialize()

	if err := validateNewPineconeVectorStoreParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_PineconeVectorStore{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.pinecone.PineconeVectorStore",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewPineconeVectorStore_Override(p PineconeVectorStore, props *PineconeVectorStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.pinecone.PineconeVectorStore",
		[]interface{}{props},
		p,
	)
}

