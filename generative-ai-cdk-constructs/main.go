// AWS Generative AI CDK Constructs is a library for well-architected generative AI patterns.
package generative-ai-cdk-constructs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
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
		"@cdklabs/generative-ai-cdk-constructs.BedrockCwDashboard",
		reflect.TypeOf((*BedrockCwDashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAllModelsMonitoring", GoMethod: "AddAllModelsMonitoring"},
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
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.ConstructName",
		reflect.TypeOf((*ConstructName)(nil)).Elem(),
		map[string]interface{}{
			"AWSRAGAPPSYNCSTEPFNOPENSEARCH": ConstructName_AWSRAGAPPSYNCSTEPFNOPENSEARCH,
			"AWSQAAPPSYNCOPENSEARCH": ConstructName_AWSQAAPPSYNCOPENSEARCH,
			"AWSSUMMARIZATIONAPPSYNCSTEPFN": ConstructName_AWSSUMMARIZATIONAPPSYNCSTEPFN,
			"AWSMODELDEPLOYMENTSAGEMAKER": ConstructName_AWSMODELDEPLOYMENTSAGEMAKER,
			"CUSTOMSAGEMAKERENDPOINT": ConstructName_CUSTOMSAGEMAKERENDPOINT,
			"HUGGINGFACESAGEMAKERENDPOINT": ConstructName_HUGGINGFACESAGEMAKERENDPOINT,
			"JUMPSTARTSAGEMAKERENDPOINT": ConstructName_JUMPSTARTSAGEMAKERENDPOINT,
			"AWSCONTENTGENAPPSYNCLAMBDA": ConstructName_AWSCONTENTGENAPPSYNCLAMBDA,
			"AWSRAGAPPSYNCSTEPFNKENDRA": ConstructName_AWSRAGAPPSYNCSTEPFNKENDRA,
			"AWSWEBCRAWLER": ConstructName_AWSWEBCRAWLER,
			"AWSTEXTTOSQL": ConstructName_AWSTEXTTOSQL,
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
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambda",
		reflect.TypeOf((*ContentGenerationAppSyncLambda)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "cgLambdaFunction", GoGetter: "CgLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "generatedImageBus", GoGetter: "GeneratedImageBus"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlApi", GoGetter: "GraphqlApi"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "s3GenerateAssetsBucket", GoGetter: "S3GenerateAssetsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3GenerateAssetsBucketInterface", GoGetter: "S3GenerateAssetsBucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ContentGenerationAppSyncLambda{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.ContentGenerationAppSyncLambdaProps",
		reflect.TypeOf((*ContentGenerationAppSyncLambdaProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.CrawlerTarget",
		reflect.TypeOf((*CrawlerTarget)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.CrawlerTargetType",
		reflect.TypeOf((*CrawlerTargetType)(nil)).Elem(),
		map[string]interface{}{
			"WEBSITE": CrawlerTargetType_WEBSITE,
			"RSS_FEED": CrawlerTargetType_RSS_FEED,
		},
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
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.DbName",
		reflect.TypeOf((*DbName)(nil)).Elem(),
		map[string]interface{}{
			"MYSQL": DbName_MYSQL,
			"POSTGRESQL": DbName_POSTGRESQL,
		},
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
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.DockerLambdaCustomProps",
		reflect.TypeOf((*DockerLambdaCustomProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.LangchainCommonDepsLayer",
		reflect.TypeOf((*LangchainCommonDepsLayer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "layer", GoGetter: "Layer"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LangchainCommonDepsLayer{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.LangchainLayerProps",
		reflect.TypeOf((*LangchainLayerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.LangchainProps",
		reflect.TypeOf((*LangchainProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.MetatdataSource",
		reflect.TypeOf((*MetatdataSource)(nil)).Elem(),
		map[string]interface{}{
			"CONFIG_FILE": MetatdataSource_CONFIG_FILE,
			"KNOWLEDGE_BASE": MetatdataSource_KNOWLEDGE_BASE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.ModelMonitoringProps",
		reflect.TypeOf((*ModelMonitoringProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearch",
		reflect.TypeOf((*QaAppsyncOpensearch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlApi", GoGetter: "GraphqlApi"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "qaBus", GoGetter: "QaBus"},
			_jsii_.MemberProperty{JsiiProperty: "qaLambdaFunction", GoGetter: "QaLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputAssetsBucket", GoGetter: "S3InputAssetsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputAssetsBucketInterface", GoGetter: "S3InputAssetsBucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_QaAppsyncOpensearch{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.QaAppsyncOpensearchProps",
		reflect.TypeOf((*QaAppsyncOpensearchProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearch",
		reflect.TypeOf((*RagAppsyncStepfnOpensearch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "embeddingsLambdaFunction", GoGetter: "EmbeddingsLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "fileTransformerLambdaFunction", GoGetter: "FileTransformerLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlApi", GoGetter: "GraphqlApi"},
			_jsii_.MemberProperty{JsiiProperty: "ingestionBus", GoGetter: "IngestionBus"},
			_jsii_.MemberProperty{JsiiProperty: "inputValidationLambdaFunction", GoGetter: "InputValidationLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputAssetsBucket", GoGetter: "S3InputAssetsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3InputAssetsBucketInterface", GoGetter: "S3InputAssetsBucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "s3ProcessedAssetsBucket", GoGetter: "S3ProcessedAssetsBucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3ProcessedAssetsBucketInterface", GoGetter: "S3ProcessedAssetsBucketInterface"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_RagAppsyncStepfnOpensearch{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.RagAppsyncStepfnOpensearchProps",
		reflect.TypeOf((*RagAppsyncStepfnOpensearchProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfn",
		reflect.TypeOf((*SummarizationAppsyncStepfn)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "documentReaderLambdaFunction", GoGetter: "DocumentReaderLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridgeBus", GoGetter: "EventBridgeBus"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlApi", GoGetter: "GraphqlApi"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlApiId", GoGetter: "GraphqlApiId"},
			_jsii_.MemberProperty{JsiiProperty: "graphqlUrl", GoGetter: "GraphqlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "inputAssetBucket", GoGetter: "InputAssetBucket"},
			_jsii_.MemberProperty{JsiiProperty: "inputValidationLambdaFunction", GoGetter: "InputValidationLambdaFunction"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "processedAssetBucket", GoGetter: "ProcessedAssetBucket"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroup", GoGetter: "SecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "stateMachine", GoGetter: "StateMachine"},
			_jsii_.MemberProperty{JsiiProperty: "summaryGeneratorLambdaFunction", GoGetter: "SummaryGeneratorLambdaFunction"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_SummarizationAppsyncStepfn{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.SummarizationAppsyncStepfnProps",
		reflect.TypeOf((*SummarizationAppsyncStepfnProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.TextToSql",
		reflect.TypeOf((*TextToSql)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "configAssetBucket", GoGetter: "ConfigAssetBucket"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "dbSecurityGroup", GoGetter: "DbSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "eventBus", GoGetter: "EventBus"},
			_jsii_.MemberProperty{JsiiProperty: "eventsRule", GoGetter: "EventsRule"},
			_jsii_.MemberProperty{JsiiProperty: "feedbackQueue", GoGetter: "FeedbackQueue"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaSecurityGroup", GoGetter: "LambdaSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputQueue", GoGetter: "OutputQueue"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "stepFunction", GoGetter: "StepFunction"},
			_jsii_.MemberProperty{JsiiProperty: "subnetGroup", GoGetter: "SubnetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_TextToSql{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.TextToSqlProps",
		reflect.TypeOf((*TextToSqlProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawler",
		reflect.TypeOf((*WebCrawler)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addObservabilityToConstruct", GoMethod: "AddObservabilityToConstruct"},
			_jsii_.MemberProperty{JsiiProperty: "constructUsageMetric", GoGetter: "ConstructUsageMetric"},
			_jsii_.MemberProperty{JsiiProperty: "dataBucket", GoGetter: "DataBucket"},
			_jsii_.MemberProperty{JsiiProperty: "enablexray", GoGetter: "Enablexray"},
			_jsii_.MemberProperty{JsiiProperty: "fieldLogLevel", GoGetter: "FieldLogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "jobQueue", GoGetter: "JobQueue"},
			_jsii_.MemberProperty{JsiiProperty: "jobsTable", GoGetter: "JobsTable"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaCrawler", GoGetter: "LambdaCrawler"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaCrawlerApiSchemaPath", GoGetter: "LambdaCrawlerApiSchemaPath"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaTracing", GoGetter: "LambdaTracing"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "retention", GoGetter: "Retention"},
			_jsii_.MemberProperty{JsiiProperty: "snsTopic", GoGetter: "SnsTopic"},
			_jsii_.MemberProperty{JsiiProperty: "stage", GoGetter: "Stage"},
			_jsii_.MemberProperty{JsiiProperty: "targetsTable", GoGetter: "TargetsTable"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "updateConstructUsageMetricCode", GoMethod: "UpdateConstructUsageMetricCode"},
			_jsii_.MemberMethod{JsiiMethod: "updateEnvSuffix", GoMethod: "UpdateEnvSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
			_jsii_.MemberProperty{JsiiProperty: "webCrawlerJobDefinition", GoGetter: "WebCrawlerJobDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_WebCrawler{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseClass)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.WebCrawlerProps",
		reflect.TypeOf((*WebCrawlerProps)(nil)).Elem(),
	)
}
