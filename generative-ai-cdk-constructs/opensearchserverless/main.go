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
			_jsii_.MemberProperty{JsiiProperty: "collectionArn", GoGetter: "CollectionArn"},
			_jsii_.MemberProperty{JsiiProperty: "collectionId", GoGetter: "CollectionId"},
			_jsii_.MemberProperty{JsiiProperty: "collectionName", GoGetter: "CollectionName"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccessPolicy", GoGetter: "DataAccessPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "grantDataAccess", GoMethod: "GrantDataAccess"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "standbyReplicas", GoGetter: "StandbyReplicas"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_VectorCollection{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
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
}
