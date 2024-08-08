package amazonaurora

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/jsii"
)

// Class to define a AmazonAuroraVectorStore.
// Experimental.
type AmazonAuroraVectorStore interface {
	// Experimental.
	CredentialsSecretArn() *string
	// Experimental.
	DatabaseName() *string
	// Experimental.
	MetadataField() *string
	// Experimental.
	PrimaryKeyField() *string
	// Experimental.
	ResourceArn() *string
	// Experimental.
	TableName() *string
	// Experimental.
	TextField() *string
	// Experimental.
	VectorField() *string
}

// The jsii proxy struct for AmazonAuroraVectorStore
type jsiiProxy_AmazonAuroraVectorStore struct {
	_ byte // padding
}

func (j *jsiiProxy_AmazonAuroraVectorStore) CredentialsSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) MetadataField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) PrimaryKeyField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKeyField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) ResourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) TextField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AmazonAuroraVectorStore) VectorField() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorField",
		&returns,
	)
	return returns
}


// Experimental.
func NewAmazonAuroraVectorStore(props *AmazonAuroraVectorStoreProps) AmazonAuroraVectorStore {
	_init_.Initialize()

	if err := validateNewAmazonAuroraVectorStoreParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AmazonAuroraVectorStore{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAmazonAuroraVectorStore_Override(a AmazonAuroraVectorStore, props *AmazonAuroraVectorStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		[]interface{}{props},
		a,
	)
}

