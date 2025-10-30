package aosscwdashboard

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.aosscwdashboard.AossCwDashboard",
		reflect.TypeOf((*AossCwDashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCollectionMonitoringbyAttributes", GoMethod: "AddCollectionMonitoringbyAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "addCollectionMonitoringByCollection", GoMethod: "AddCollectionMonitoringByCollection"},
			_jsii_.MemberMethod{JsiiMethod: "addIndexMonitoringByAtributes", GoMethod: "AddIndexMonitoringByAtributes"},
			_jsii_.MemberProperty{JsiiProperty: "dashboard", GoGetter: "Dashboard"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AossCwDashboard{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.aosscwdashboard.AossCwDashboardProps",
		reflect.TypeOf((*AossCwDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.aosscwdashboard.CollectionMonitoringProps",
		reflect.TypeOf((*CollectionMonitoringProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.aosscwdashboard.IndexMonitoringProps",
		reflect.TypeOf((*IndexMonitoringProps)(nil)).Elem(),
	)
}
