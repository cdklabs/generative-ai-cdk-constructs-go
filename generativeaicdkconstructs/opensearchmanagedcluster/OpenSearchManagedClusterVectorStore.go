package opensearchmanagedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/jsii"
)

// Class to define an OpenSearchManagedClusterVectorStore.
// Experimental.
type OpenSearchManagedClusterVectorStore interface {
	// Experimental.
	DomainArn() *string
	// Experimental.
	DomainEndpoint() *string
	// Experimental.
	FieldMapping() *OpenSearchFieldMapping
	// Experimental.
	VectorIndexName() *string
}

// The jsii proxy struct for OpenSearchManagedClusterVectorStore
type jsiiProxy_OpenSearchManagedClusterVectorStore struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchManagedClusterVectorStore) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchManagedClusterVectorStore) DomainEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchManagedClusterVectorStore) FieldMapping() *OpenSearchFieldMapping {
	var returns *OpenSearchFieldMapping
	_jsii_.Get(
		j,
		"fieldMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchManagedClusterVectorStore) VectorIndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vectorIndexName",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenSearchManagedClusterVectorStore(props *OpenSearchManagedClusterVectorStoreProps) OpenSearchManagedClusterVectorStore {
	_init_.Initialize()

	if err := validateNewOpenSearchManagedClusterVectorStoreParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchManagedClusterVectorStore{}

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.opensearchmanagedcluster.OpenSearchManagedClusterVectorStore",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchManagedClusterVectorStore_Override(o OpenSearchManagedClusterVectorStore, props *OpenSearchManagedClusterVectorStoreProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdklabs/generative-ai-cdk-constructs.opensearchmanagedcluster.OpenSearchManagedClusterVectorStore",
		[]interface{}{props},
		o,
	)
}

