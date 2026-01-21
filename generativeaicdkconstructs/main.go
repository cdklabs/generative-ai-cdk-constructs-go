// AWS Generative AI CDK Constructs is a library for well-architected generative AI patterns.
package generativeaicdkconstructs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.BaseClass",
		reflect.TypeOf((*BaseClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_BaseClass{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.BaseClassProps",
		reflect.TypeOf((*BaseClassProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.ConstructName",
		reflect.TypeOf((*ConstructName)(nil)).Elem(),
		map[string]interface{}{
			"AWSMODELDEPLOYMENTSAGEMAKER": ConstructName_AWSMODELDEPLOYMENTSAGEMAKER,
			"CUSTOMSAGEMAKERENDPOINT": ConstructName_CUSTOMSAGEMAKERENDPOINT,
			"HUGGINGFACESAGEMAKERENDPOINT": ConstructName_HUGGINGFACESAGEMAKERENDPOINT,
			"JUMPSTARTSAGEMAKERENDPOINT": ConstructName_JUMPSTARTSAGEMAKERENDPOINT,
		},
	)
}
