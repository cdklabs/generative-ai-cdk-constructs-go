package kendra

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.kendra.Kendra.Edition",
		reflect.TypeOf((*Edition)(nil)).Elem(),
		map[string]interface{}{
			"DEVELOPER_EDITION": Edition_DEVELOPER_EDITION,
			"ENTERPRISE_EDITION": Edition_ENTERPRISE_EDITION,
			"GEN_AI_ENTERPRISE_EDITION": Edition_GEN_AI_ENTERPRISE_EDITION,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.kendra.Kendra.IndexFieldTypes",
		reflect.TypeOf((*IndexFieldTypes)(nil)).Elem(),
		map[string]interface{}{
			"STRING": IndexFieldTypes_STRING,
			"STRING_LIST": IndexFieldTypes_STRING_LIST,
			"LONG": IndexFieldTypes_LONG,
			"DATE": IndexFieldTypes_DATE,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.kendra.Kendra.UserContextPolicy",
		reflect.TypeOf((*UserContextPolicy)(nil)).Elem(),
		map[string]interface{}{
			"ATTRIBUTE_FILTER": UserContextPolicy_ATTRIBUTE_FILTER,
			"USER_TOKEN": UserContextPolicy_USER_TOKEN,
		},
	)
}
