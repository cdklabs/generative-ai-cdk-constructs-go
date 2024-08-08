package amazonaurora

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraDefaultVectorStore",
		reflect.TypeOf((*AmazonAuroraDefaultVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdentifier", GoGetter: "ClusterIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsSecretArn", GoGetter: "CredentialsSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "embeddingsModelVectorDimension", GoGetter: "EmbeddingsModelVectorDimension"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyField", GoGetter: "PrimaryKeyField"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArn", GoGetter: "ResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonAuroraDefaultVectorStore{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraDefaultVectorStoreProps",
		reflect.TypeOf((*AmazonAuroraDefaultVectorStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		reflect.TypeOf((*AmazonAuroraVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "credentialsSecretArn", GoGetter: "CredentialsSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "metadataField", GoGetter: "MetadataField"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyField", GoGetter: "PrimaryKeyField"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArn", GoGetter: "ResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "textField", GoGetter: "TextField"},
			_jsii_.MemberProperty{JsiiProperty: "vectorField", GoGetter: "VectorField"},
		},
		func() interface{} {
			return &jsiiProxy_AmazonAuroraVectorStore{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStoreProps",
		reflect.TypeOf((*AmazonAuroraVectorStoreProps)(nil)).Elem(),
	)
}
