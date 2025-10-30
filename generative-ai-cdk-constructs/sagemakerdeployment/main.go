package sagemakerdeployment

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.AsyncInferenceConfig",
		reflect.TypeOf((*AsyncInferenceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.ContainerImage",
		reflect.TypeOf((*ContainerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ContainerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.ContainerImageConfig",
		reflect.TypeOf((*ContainerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.CustomSageMakerEndpoint",
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
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.CustomSageMakerEndpointProps",
		reflect.TypeOf((*CustomSageMakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.DeepLearningContainerImage",
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
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpoint",
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
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.HuggingFaceSageMakerEndpointProps",
		reflect.TypeOf((*HuggingFaceSageMakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.IInstanceAliase",
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
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.IInstanceValiant",
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
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.IJumpStartModelSpec",
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
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartModel",
		reflect.TypeOf((*JumpStartModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_JumpStartModel{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpoint",
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
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.JumpStartSageMakerEndpointProps",
		reflect.TypeOf((*JumpStartSageMakerEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerEndpointBase",
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
			_jsii_.InitJsiiProxy(&j.Type__generativeaicdkconstructsBaseClass)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.sagemaker_deployment.SageMakerInstanceType",
		reflect.TypeOf((*SageMakerInstanceType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_SageMakerInstanceType{}
		},
	)
}
