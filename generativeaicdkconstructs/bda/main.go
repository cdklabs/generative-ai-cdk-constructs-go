package bda

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bda.BedrockDataAutomation",
		reflect.TypeOf((*BedrockDataAutomation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "bdaBlueprintLambdaFunction", GoGetter: "BdaBlueprintLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "bdaInputBucket", GoGetter: "BdaInputBucket"},
			_jsii_.MemberProperty{JsiiProperty: "bdaInvocationFunction", GoGetter: "BdaInvocationFunction"},
			_jsii_.MemberProperty{JsiiProperty: "bdaOutputBucket", GoGetter: "BdaOutputBucket"},
			_jsii_.MemberProperty{JsiiProperty: "bdaProjectFunction", GoGetter: "BdaProjectFunction"},
			_jsii_.MemberProperty{JsiiProperty: "bdaResultStatusFunction", GoGetter: "BdaResultStatusFunction"},
			_jsii_.MemberProperty{JsiiProperty: "boto3Layer", GoGetter: "Boto3Layer"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "powertoolsLayer", GoGetter: "PowertoolsLayer"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockDataAutomation{}
			_jsii_.InitJsiiProxy(&j.Type__generativeaicdkconstructsBaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bda.BedrockDataAutomationProps",
		reflect.TypeOf((*BedrockDataAutomationProps)(nil)).Elem(),
	)
}
