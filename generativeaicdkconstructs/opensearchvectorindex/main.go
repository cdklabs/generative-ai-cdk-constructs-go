package opensearchvectorindex

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearch_vectorindex.Analyzer",
		reflect.TypeOf((*Analyzer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearch_vectorindex.MetadataManagementFieldProps",
		reflect.TypeOf((*MetadataManagementFieldProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.opensearch_vectorindex.VectorIndex",
		reflect.TypeOf((*VectorIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "indexName", GoGetter: "IndexName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorDimensions", GoGetter: "VectorDimensions"},
			_jsii_.MemberProperty{JsiiProperty: "vectorField", GoGetter: "VectorField"},
		},
		func() interface{} {
			j := jsiiProxy_VectorIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearch_vectorindex.VectorIndexProps",
		reflect.TypeOf((*VectorIndexProps)(nil)).Elem(),
	)
}
