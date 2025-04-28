package opensearchmanagedcluster

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearchmanagedcluster.OpenSearchFieldMapping",
		reflect.TypeOf((*OpenSearchFieldMapping)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.opensearchmanagedcluster.OpenSearchManagedClusterVectorStore",
		reflect.TypeOf((*OpenSearchManagedClusterVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "domainArn", GoGetter: "DomainArn"},
			_jsii_.MemberProperty{JsiiProperty: "domainEndpoint", GoGetter: "DomainEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "fieldMapping", GoGetter: "FieldMapping"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexName", GoGetter: "VectorIndexName"},
		},
		func() interface{} {
			return &jsiiProxy_OpenSearchManagedClusterVectorStore{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearchmanagedcluster.OpenSearchManagedClusterVectorStoreProps",
		reflect.TypeOf((*OpenSearchManagedClusterVectorStoreProps)(nil)).Elem(),
	)
}
