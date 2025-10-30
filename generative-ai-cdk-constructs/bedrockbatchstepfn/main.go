package bedrockbatchstepfn

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock_batch_stepfn.BedrockBatchSfn",
		reflect.TypeOf((*BedrockBatchSfn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "prefixStates", GoMethod: "PrefixStates"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberMethod{JsiiMethod: "toSingleState", GoMethod: "ToSingleState"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockBatchSfn{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsStateMachineFragment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock_batch_stepfn.BedrockBatchSfnProps",
		reflect.TypeOf((*BedrockBatchSfnProps)(nil)).Elem(),
	)
}
