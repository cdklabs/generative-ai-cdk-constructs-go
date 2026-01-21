package bedrockcwdashboard

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrockcwdashboard.BedrockCwDashboard",
		reflect.TypeOf((*BedrockCwDashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAllGuardrailsMonitoring", GoMethod: "AddAllGuardrailsMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "addAllModelsMonitoring", GoMethod: "AddAllModelsMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "addGuardrailMonitoring", GoMethod: "AddGuardrailMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "addModelMonitoring", GoMethod: "AddModelMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "dashboard", GoGetter: "Dashboard"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockCwDashboard{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrockcwdashboard.BedrockCwDashboardProps",
		reflect.TypeOf((*BedrockCwDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrockcwdashboard.ModelMonitoringProps",
		reflect.TypeOf((*ModelMonitoringProps)(nil)).Elem(),
	)
}
