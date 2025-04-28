package mongodbatlas

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.mongodbAtlas.MongoDBAtlasVectorStore",
		reflect.TypeOf((*MongoDBAtlasVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collectionName", GoGetter: "CollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsSecretArn", GoGetter: "CredentialsSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "endpointServiceName", GoGetter: "EndpointServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "fieldMapping", GoGetter: "FieldMapping"},
			_jsii_.MemberProperty{JsiiProperty: "textIndexName", GoGetter: "TextIndexName"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexName", GoGetter: "VectorIndexName"},
		},
		func() interface{} {
			return &jsiiProxy_MongoDBAtlasVectorStore{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.mongodbAtlas.MongoDBAtlasVectorStoreProps",
		reflect.TypeOf((*MongoDBAtlasVectorStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.mongodbAtlas.MongoDbAtlasFieldMapping",
		reflect.TypeOf((*MongoDbAtlasFieldMapping)(nil)).Elem(),
	)
}
