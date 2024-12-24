package opensearchserverless

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.CharacterFilterType",
		reflect.TypeOf((*CharacterFilterType)(nil)).Elem(),
		map[string]interface{}{
			"ICU_NORMALIZER": CharacterFilterType_ICU_NORMALIZER,
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.IVectorCollection",
		reflect.TypeOf((*IVectorCollection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aossPolicy", GoGetter: "AossPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "collectionArn", GoGetter: "CollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "collectionId", GoGetter: "CollectionId"},
			_jsii_.MemberProperty{JsiiProperty: "collectionName", GoGetter: "CollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "collectionType", GoGetter: "CollectionType"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccessPolicy", GoGetter: "DataAccessPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricIndexRequestCount", GoMethod: "MetricIndexRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchLatency", GoMethod: "MetricSearchLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchLatencyP90", GoMethod: "MetricSearchLatencyP90"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchRequestCount", GoMethod: "MetricSearchRequestCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "standbyReplicas", GoGetter: "StandbyReplicas"},
		},
		func() interface{} {
			j := jsiiProxy_IVectorCollection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.TokenFilterType",
		reflect.TypeOf((*TokenFilterType)(nil)).Elem(),
		map[string]interface{}{
			"KUROMOJI_BASEFORM": TokenFilterType_KUROMOJI_BASEFORM,
			"KUROMOJI_PART_OF_SPEECH": TokenFilterType_KUROMOJI_PART_OF_SPEECH,
			"KUROMOJI_STEMMER": TokenFilterType_KUROMOJI_STEMMER,
			"CJK_WIDTH": TokenFilterType_CJK_WIDTH,
			"JA_STOP": TokenFilterType_JA_STOP,
			"LOWERCASE": TokenFilterType_LOWERCASE,
			"ICU_FOLDING": TokenFilterType_ICU_FOLDING,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.TokenizerType",
		reflect.TypeOf((*TokenizerType)(nil)).Elem(),
		map[string]interface{}{
			"KUROMOJI_TOKENIZER": TokenizerType_KUROMOJI_TOKENIZER,
			"ICU_TOKENIZER": TokenizerType_ICU_TOKENIZER,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollection",
		reflect.TypeOf((*VectorCollection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aossPolicy", GoGetter: "AossPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "collectionArn", GoGetter: "CollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "collectionEndpoint", GoGetter: "CollectionEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "collectionId", GoGetter: "CollectionId"},
			_jsii_.MemberProperty{JsiiProperty: "collectionName", GoGetter: "CollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "collectionType", GoGetter: "CollectionType"},
			_jsii_.MemberProperty{JsiiProperty: "dashboardEndpoint", GoGetter: "DashboardEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccessPolicy", GoGetter: "DataAccessPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDataAccess", GoMethod: "GrantDataAccess"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricIndexRequestCount", GoMethod: "MetricIndexRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchLatency", GoMethod: "MetricSearchLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchLatencyP90", GoMethod: "MetricSearchLatencyP90"},
			_jsii_.MemberMethod{JsiiMethod: "metricSearchRequestCount", GoMethod: "MetricSearchRequestCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "standbyReplicas", GoGetter: "StandbyReplicas"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VectorCollection{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVectorCollection)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollectionAttributes",
		reflect.TypeOf((*VectorCollectionAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollectionProps",
		reflect.TypeOf((*VectorCollectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollectionStandbyReplicas",
		reflect.TypeOf((*VectorCollectionStandbyReplicas)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": VectorCollectionStandbyReplicas_ENABLED,
			"DISABLED": VectorCollectionStandbyReplicas_DISABLED,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.opensearchserverless.VectorCollectionType",
		reflect.TypeOf((*VectorCollectionType)(nil)).Elem(),
		map[string]interface{}{
			"SEARCH": VectorCollectionType_SEARCH,
			"TIMESERIES": VectorCollectionType_TIMESERIES,
			"VECTORSEARCH": VectorCollectionType_VECTORSEARCH,
		},
	)
}
