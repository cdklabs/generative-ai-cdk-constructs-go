// AWS Generative AI CDK Constructs is a library for well-architected generative AI patterns.
package generative-ai-cdk-constructs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.AossCwDashboard",
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
		"@cdklabs/generative-ai-cdk-constructs.AossCwDashboardProps",
		reflect.TypeOf((*AossCwDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.AsyncInferenceConfig",
		reflect.TypeOf((*AsyncInferenceConfig)(nil)).Elem(),
	)
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
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.BedrockBatchSfn",
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
		"@cdklabs/generative-ai-cdk-constructs.BedrockBatchSfnProps",
		reflect.TypeOf((*BedrockBatchSfnProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.BedrockCwDashboard",
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
		"@cdklabs/generative-ai-cdk-constructs.BedrockCwDashboardProps",
		reflect.TypeOf((*BedrockCwDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.BedrockDataAutomation",
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
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.BedrockDataAutomationProps",
		reflect.TypeOf((*BedrockDataAutomationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.CollectionMonitoringProps",
		reflect.TypeOf((*CollectionMonitoringProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.ContainerImage",
		reflect.TypeOf((*ContainerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ContainerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.ContainerImageConfig",
		reflect.TypeOf((*ContainerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpoint",
		reflect.TypeOf((*CustomSageMakerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpoint", GoGetter: "CfnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpointConfig", GoGetter: "CfnEndpointConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cfnModel", GoGetter: "CfnModel"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberMethod{JsiiMethod: "createSageMakerRole", GoMethod: "CreateSageMakerRole"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "errorTopic", GoGetter: "ErrorTopic"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "modelDataDownloadTimeoutInSeconds", GoGetter: "ModelDataDownloadTimeoutInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "modelDataUrl", GoGetter: "ModelDataUrl"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicy", GoGetter: "ScalingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "successTopic", GoGetter: "SuccessTopic"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_CustomSageMakerEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SageMakerEndpointBase)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.CustomSageMakerEndpointProps",
		reflect.TypeOf((*CustomSageMakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.DeepLearningContainerImage",
		reflect.TypeOf((*DeepLearningContainerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_DeepLearningContainerImage{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ContainerImage)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.HuggingFaceSageMakerEndpoint",
		reflect.TypeOf((*HuggingFaceSageMakerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpoint", GoGetter: "CfnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpointConfig", GoGetter: "CfnEndpointConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cfnModel", GoGetter: "CfnModel"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberMethod{JsiiMethod: "createSageMakerRole", GoMethod: "CreateSageMakerRole"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_HuggingFaceSageMakerEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SageMakerEndpointBase)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.HuggingFaceSageMakerEndpointProps",
		reflect.TypeOf((*HuggingFaceSageMakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.IInstanceAliase",
		reflect.TypeOf((*IInstanceAliase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "aliases", GoGetter: "Aliases"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
		},
		func() interface{} {
			return &jsiiProxy_IInstanceAliase{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.IInstanceValiant",
		reflect.TypeOf((*IInstanceValiant)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "imageUri", GoGetter: "ImageUri"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
		},
		func() interface{} {
			return &jsiiProxy_IInstanceValiant{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.IJumpStartModelSpec",
		reflect.TypeOf((*IJumpStartModelSpec)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "artifactKey", GoGetter: "ArtifactKey"},
			_jsii_.MemberProperty{JsiiProperty: "defaultInstanceType", GoGetter: "DefaultInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "gatedBucket", GoGetter: "GatedBucket"},
			_jsii_.MemberProperty{JsiiProperty: "instanceAliases", GoGetter: "InstanceAliases"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "instanceVariants", GoGetter: "InstanceVariants"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "modelPackageArns", GoGetter: "ModelPackageArns"},
			_jsii_.MemberProperty{JsiiProperty: "prepackedArtifactKey", GoGetter: "PrepackedArtifactKey"},
			_jsii_.MemberProperty{JsiiProperty: "requiresEula", GoGetter: "RequiresEula"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_IJumpStartModelSpec{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.IndexMonitoringProps",
		reflect.TypeOf((*IndexMonitoringProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartModel",
		reflect.TypeOf((*JumpStartModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_JumpStartModel{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartSageMakerEndpoint",
		reflect.TypeOf((*JumpStartSageMakerEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberMethod{JsiiMethod: "addToRolePolicy", GoMethod: "AddToRolePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpoint", GoGetter: "CfnEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "cfnEndpointConfig", GoGetter: "CfnEndpointConfig"},
			_jsii_.MemberProperty{JsiiProperty: "cfnModel", GoGetter: "CfnModel"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberMethod{JsiiMethod: "createSageMakerRole", GoMethod: "CreateSageMakerRole"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "endpointArn", GoGetter: "EndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "instanceCount", GoGetter: "InstanceCount"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "model", GoGetter: "Model"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
		},
		func() interface{} {
			j := jsiiProxy_JumpStartSageMakerEndpoint{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_SageMakerEndpointBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.JumpStartSageMakerEndpointProps",
		reflect.TypeOf((*JumpStartSageMakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.ModelMonitoringProps",
		reflect.TypeOf((*ModelMonitoringProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.SageMakerEndpointBase",
		reflect.TypeOf((*SageMakerEndpointBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberMethod{JsiiMethod: "createSageMakerRole", GoMethod: "CreateSageMakerRole"},
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
			j := jsiiProxy_SageMakerEndpointBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.SageMakerInstanceType",
		reflect.TypeOf((*SageMakerInstanceType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_SageMakerInstanceType{}
		},
	)
}
