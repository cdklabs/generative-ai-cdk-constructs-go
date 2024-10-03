package amazonaurora

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStore",
		reflect.TypeOf((*AmazonAuroraVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIngressRuleToAuroraSecurityGroup", GoMethod: "AddIngressRuleToAuroraSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "createAuroraPgCRPolicy", GoMethod: "CreateAuroraPgCRPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createLambdaSecurityGroup", GoMethod: "CreateLambdaSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsSecretArn", GoGetter: "CredentialsSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "embeddingsModelVectorDimension", GoGetter: "EmbeddingsModelVectorDimension"},
			_jsii_.MemberMethod{JsiiMethod: "generateResourceArn", GoMethod: "GenerateResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "metadataField", GoGetter: "MetadataField"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyField", GoGetter: "PrimaryKeyField"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArn", GoGetter: "ResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "schemaName", GoGetter: "SchemaName"},
			_jsii_.MemberMethod{JsiiMethod: "setupCustomResource", GoMethod: "SetupCustomResource"},
			_jsii_.MemberMethod{JsiiMethod: "setupDatabaseClusterResources", GoMethod: "SetupDatabaseClusterResources"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "textField", GoGetter: "TextField"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorField", GoGetter: "VectorField"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_AmazonAuroraVectorStore{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.AmazonAuroraVectorStoreProps",
		reflect.TypeOf((*AmazonAuroraVectorStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.BaseAuroraVectorStoreProps",
		reflect.TypeOf((*BaseAuroraVectorStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.DatabaseClusterResources",
		reflect.TypeOf((*DatabaseClusterResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.ExistingAmazonAuroraVectorStore",
		reflect.TypeOf((*ExistingAmazonAuroraVectorStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addIngressRuleToAuroraSecurityGroup", GoMethod: "AddIngressRuleToAuroraSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "createAuroraPgCRPolicy", GoMethod: "CreateAuroraPgCRPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createLambdaSecurityGroup", GoMethod: "CreateLambdaSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsSecretArn", GoGetter: "CredentialsSecretArn"},
			_jsii_.MemberProperty{JsiiProperty: "databaseName", GoGetter: "DatabaseName"},
			_jsii_.MemberProperty{JsiiProperty: "embeddingsModelVectorDimension", GoGetter: "EmbeddingsModelVectorDimension"},
			_jsii_.MemberMethod{JsiiMethod: "generateResourceArn", GoMethod: "GenerateResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "metadataField", GoGetter: "MetadataField"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyField", GoGetter: "PrimaryKeyField"},
			_jsii_.MemberProperty{JsiiProperty: "resourceArn", GoGetter: "ResourceArn"},
			_jsii_.MemberProperty{JsiiProperty: "schemaName", GoGetter: "SchemaName"},
			_jsii_.MemberMethod{JsiiMethod: "setupCustomResource", GoMethod: "SetupCustomResource"},
			_jsii_.MemberMethod{JsiiMethod: "setupDatabaseClusterResources", GoMethod: "SetupDatabaseClusterResources"},
			_jsii_.MemberProperty{JsiiProperty: "tableName", GoGetter: "TableName"},
			_jsii_.MemberProperty{JsiiProperty: "textField", GoGetter: "TextField"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorField", GoGetter: "VectorField"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ExistingAmazonAuroraVectorStore{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.amazonaurora.ExistingAmazonAuroraVectorStoreProps",
		reflect.TypeOf((*ExistingAmazonAuroraVectorStoreProps)(nil)).Elem(),
	)
}
