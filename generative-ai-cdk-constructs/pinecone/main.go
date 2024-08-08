package pinecone

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.pinecone.PineconeVectorStore",
		reflect.TypeOf((*PineconeVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "connectionString", GoGetter: "ConnectionString"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsSecretArn", GoGetter: "CredentialsSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "metadataField", GoGetter: "MetadataField"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "textField", GoGetter: "TextField"},
		},
		func() interface{} {
			return &jsiiProxy_PineconeVectorStore{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.pinecone.PineconeVectorStoreProps",
		reflect.TypeOf((*PineconeVectorStoreProps)(nil)).Elem(),
	)
}
