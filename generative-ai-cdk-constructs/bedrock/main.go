package bedrock

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ActionGroupExecutor",
		reflect.TypeOf((*ActionGroupExecutor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "customControl", GoGetter: "CustomControl"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
		},
		func() interface{} {
			return &jsiiProxy_ActionGroupExecutor{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Agent",
		reflect.TypeOf((*Agent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionGroups", GoGetter: "ActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addActionGroup", GoMethod: "AddActionGroup"},
			_jsii_.MemberMethod{JsiiMethod: "addActionGroups", GoMethod: "AddActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addAgentCollaborator", GoMethod: "AddAgentCollaborator"},
			_jsii_.MemberMethod{JsiiMethod: "addGuardrail", GoMethod: "AddGuardrail"},
			_jsii_.MemberMethod{JsiiMethod: "addKnowledgeBase", GoMethod: "AddKnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentCollaboration", GoGetter: "AgentCollaboration"},
			_jsii_.MemberProperty{JsiiProperty: "agentCollaborators", GoGetter: "AgentCollaborators"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "agentVersion", GoGetter: "AgentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "codeInterpreterEnabled", GoGetter: "CodeInterpreterEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "customOrchestration", GoGetter: "CustomOrchestration"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "forceDelete", GoGetter: "ForceDelete"},
			_jsii_.MemberProperty{JsiiProperty: "foundationModel", GoGetter: "FoundationModel"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "guardrail", GoGetter: "Guardrail"},
			_jsii_.MemberProperty{JsiiProperty: "idleSessionTTL", GoGetter: "IdleSessionTTL"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBases", GoGetter: "KnowledgeBases"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "orchestrationType", GoGetter: "OrchestrationType"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "promptOverrideConfiguration", GoGetter: "PromptOverrideConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "shouldPrepareAgent", GoGetter: "ShouldPrepareAgent"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "testAlias", GoGetter: "TestAlias"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userInputEnabled", GoGetter: "UserInputEnabled"},
		},
		func() interface{} {
			j := jsiiProxy_Agent{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AgentBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		reflect.TypeOf((*AgentActionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiSchema", GoGetter: "ApiSchema"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "executor", GoGetter: "Executor"},
			_jsii_.MemberProperty{JsiiProperty: "forceDelete", GoGetter: "ForceDelete"},
			_jsii_.MemberProperty{JsiiProperty: "functionSchema", GoGetter: "FunctionSchema"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentActionGroupSignature", GoGetter: "ParentActionGroupSignature"},
		},
		func() interface{} {
			return &jsiiProxy_AgentActionGroup{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroupProps",
		reflect.TypeOf((*AgentActionGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAlias",
		reflect.TypeOf((*AgentAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgentAlias{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AgentAliasBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAliasAttributes",
		reflect.TypeOf((*AgentAliasAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAliasBase",
		reflect.TypeOf((*AgentAliasBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgentAliasBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAgentAlias)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAliasProps",
		reflect.TypeOf((*AgentAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAttributes",
		reflect.TypeOf((*AgentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentBase",
		reflect.TypeOf((*AgentBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "agentVersion", GoGetter: "AgentVersion"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgentBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAgent)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentCollaborator",
		reflect.TypeOf((*AgentCollaborator)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentAlias", GoGetter: "AgentAlias"},
			_jsii_.MemberProperty{JsiiProperty: "collaborationInstruction", GoGetter: "CollaborationInstruction"},
			_jsii_.MemberProperty{JsiiProperty: "collaboratorName", GoGetter: "CollaboratorName"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "relayConversationHistory", GoGetter: "RelayConversationHistory"},
		},
		func() interface{} {
			return &jsiiProxy_AgentCollaborator{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentCollaboratorProps",
		reflect.TypeOf((*AgentCollaboratorProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentCollaboratorType",
		reflect.TypeOf((*AgentCollaboratorType)(nil)).Elem(),
		map[string]interface{}{
			"SUPERVISOR": AgentCollaboratorType_SUPERVISOR,
			"DISABLED": AgentCollaboratorType_DISABLED,
			"SUPERVISOR_ROUTER": AgentCollaboratorType_SUPERVISOR_ROUTER,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentPromptVariantProps",
		reflect.TypeOf((*AgentPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentProps",
		reflect.TypeOf((*AgentProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentStepType",
		reflect.TypeOf((*AgentStepType)(nil)).Elem(),
		map[string]interface{}{
			"PRE_PROCESSING": AgentStepType_PRE_PROCESSING,
			"ORCHESTRATION": AgentStepType_ORCHESTRATION,
			"POST_PROCESSING": AgentStepType_POST_PROCESSING,
			"ROUTING_CLASSIFIER": AgentStepType_ROUTING_CLASSIFIER,
			"MEMORY_SUMMARIZATION": AgentStepType_MEMORY_SUMMARIZATION,
			"KNOWLEDGE_BASE_RESPONSE_GENERATION": AgentStepType_KNOWLEDGE_BASE_RESPONSE_GENERATION,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		reflect.TypeOf((*ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			return &jsiiProxy_ApiSchema{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApplicationInferenceProfile",
		reflect.TypeOf((*ApplicationInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileModel", GoGetter: "InferenceProfileModel"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileName", GoGetter: "InferenceProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationInferenceProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_InferenceProfileBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInvokable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApplicationInferenceProfileAttributes",
		reflect.TypeOf((*ApplicationInferenceProfileAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApplicationInferenceProfileProps",
		reflect.TypeOf((*ApplicationInferenceProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		reflect.TypeOf((*BedrockFoundationModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asArn", GoMethod: "AsArn"},
			_jsii_.MemberMethod{JsiiMethod: "asIModel", GoMethod: "AsIModel"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvokeAllRegions", GoMethod: "GrantInvokeAllRegions"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelArn", GoGetter: "ModelArn"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "supportedVectorType", GoGetter: "SupportedVectorType"},
			_jsii_.MemberProperty{JsiiProperty: "supportsAgents", GoGetter: "SupportsAgents"},
			_jsii_.MemberProperty{JsiiProperty: "supportsCrossRegion", GoGetter: "SupportsCrossRegion"},
			_jsii_.MemberProperty{JsiiProperty: "supportsKnowledgeBase", GoGetter: "SupportsKnowledgeBase"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorDimensions", GoGetter: "VectorDimensions"},
		},
		func() interface{} {
			j := jsiiProxy_BedrockFoundationModel{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInvokable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModelProps",
		reflect.TypeOf((*BedrockFoundationModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatMessage",
		reflect.TypeOf((*ChatMessage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
		},
		func() interface{} {
			return &jsiiProxy_ChatMessage{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatMessageRole",
		reflect.TypeOf((*ChatMessageRole)(nil)).Elem(),
		map[string]interface{}{
			"USER": ChatMessageRole_USER,
			"ASSISTANT": ChatMessageRole_ASSISTANT,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChatPromptVariantProps",
		reflect.TypeOf((*ChatPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		reflect.TypeOf((*ChunkingStrategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
		},
		func() interface{} {
			return &jsiiProxy_ChunkingStrategy{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CommonKnowledgeBaseAttributes",
		reflect.TypeOf((*CommonKnowledgeBaseAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CommonKnowledgeBaseProps",
		reflect.TypeOf((*CommonKnowledgeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CommonPromptVariantProps",
		reflect.TypeOf((*CommonPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ConfluenceCrawlingFilters",
		reflect.TypeOf((*ConfluenceCrawlingFilters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ConfluenceDataSource",
		reflect.TypeOf((*ConfluenceDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authSecret", GoGetter: "AuthSecret"},
			_jsii_.MemberProperty{JsiiProperty: "confluenceUrl", GoGetter: "ConfluenceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ConfluenceDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceNew)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ConfluenceDataSourceAssociationProps",
		reflect.TypeOf((*ConfluenceDataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ConfluenceDataSourceAuthType",
		reflect.TypeOf((*ConfluenceDataSourceAuthType)(nil)).Elem(),
		map[string]interface{}{
			"OAUTH2_CLIENT_CREDENTIALS": ConfluenceDataSourceAuthType_OAUTH2_CLIENT_CREDENTIALS,
			"BASIC": ConfluenceDataSourceAuthType_BASIC,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ConfluenceDataSourceProps",
		reflect.TypeOf((*ConfluenceDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ConfluenceObjectType",
		reflect.TypeOf((*ConfluenceObjectType)(nil)).Elem(),
		map[string]interface{}{
			"SPACE": ConfluenceObjectType_SPACE,
			"PAGE": ConfluenceObjectType_PAGE,
			"BLOG": ConfluenceObjectType_BLOG,
			"COMMENT": ConfluenceObjectType_COMMENT,
			"ATTACHMENT": ConfluenceObjectType_ATTACHMENT,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentFilter",
		reflect.TypeOf((*ContentFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentFilterStrength",
		reflect.TypeOf((*ContentFilterStrength)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ContentFilterStrength_NONE,
			"LOW": ContentFilterStrength_LOW,
			"MEDIUM": ContentFilterStrength_MEDIUM,
			"HIGH": ContentFilterStrength_HIGH,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentFilterType",
		reflect.TypeOf((*ContentFilterType)(nil)).Elem(),
		map[string]interface{}{
			"SEXUAL": ContentFilterType_SEXUAL,
			"VIOLENCE": ContentFilterType_VIOLENCE,
			"HATE": ContentFilterType_HATE,
			"INSULTS": ContentFilterType_INSULTS,
			"MISCONDUCT": ContentFilterType_MISCONDUCT,
			"PROMPT_ATTACK": ContentFilterType_PROMPT_ATTACK,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextEnrichment",
		reflect.TypeOf((*ContextEnrichment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "generatePolicyStatements", GoMethod: "GeneratePolicyStatements"},
		},
		func() interface{} {
			return &jsiiProxy_ContextEnrichment{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextEnrichmentType",
		reflect.TypeOf((*ContextEnrichmentType)(nil)).Elem(),
		map[string]interface{}{
			"BEDROCK_FOUNDATION_MODEL": ContextEnrichmentType_BEDROCK_FOUNDATION_MODEL,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextualGroundingFilter",
		reflect.TypeOf((*ContextualGroundingFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextualGroundingFilterType",
		reflect.TypeOf((*ContextualGroundingFilterType)(nil)).Elem(),
		map[string]interface{}{
			"GROUNDING": ContextualGroundingFilterType_GROUNDING,
			"RELEVANCE": ContextualGroundingFilterType_RELEVANCE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CrawlingFilters",
		reflect.TypeOf((*CrawlingFilters)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CrawlingScope",
		reflect.TypeOf((*CrawlingScope)(nil)).Elem(),
		map[string]interface{}{
			"HOST_ONLY": CrawlingScope_HOST_ONLY,
			"SUBDOMAINS": CrawlingScope_SUBDOMAINS,
			"DEFAULT": CrawlingScope_DEFAULT,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CrossRegionInferenceProfile",
		reflect.TypeOf((*CrossRegionInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileModel", GoGetter: "InferenceProfileModel"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_CrossRegionInferenceProfile{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInferenceProfile)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInvokable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CrossRegionInferenceProfileProps",
		reflect.TypeOf((*CrossRegionInferenceProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CrossRegionInferenceProfileRegion",
		reflect.TypeOf((*CrossRegionInferenceProfileRegion)(nil)).Elem(),
		map[string]interface{}{
			"EU": CrossRegionInferenceProfileRegion_EU,
			"US": CrossRegionInferenceProfileRegion_US,
			"APAC": CrossRegionInferenceProfileRegion_APAC,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomDataSource",
		reflect.TypeOf((*CustomDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CustomDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceNew)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomDataSourceAssociationProps",
		reflect.TypeOf((*CustomDataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomDataSourceProps",
		reflect.TypeOf((*CustomDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomOrchestration",
		reflect.TypeOf((*CustomOrchestration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomParserProps",
		reflect.TypeOf((*CustomParserProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomTopicProps",
		reflect.TypeOf((*CustomTopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CustomTransformation",
		reflect.TypeOf((*CustomTransformation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "generatePolicyStatements", GoMethod: "GeneratePolicyStatements"},
		},
		func() interface{} {
			return &jsiiProxy_CustomTransformation{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DataDeletionPolicy",
		reflect.TypeOf((*DataDeletionPolicy)(nil)).Elem(),
		map[string]interface{}{
			"DELETE": DataDeletionPolicy_DELETE,
			"RETAIN": DataDeletionPolicy_RETAIN,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DataSourceAssociationProps",
		reflect.TypeOf((*DataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DataSourceBase",
		reflect.TypeOf((*DataSourceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataSourceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IDataSource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DataSourceNew",
		reflect.TypeOf((*DataSourceNew)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataSourceNew{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceBase)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DataSourceType",
		reflect.TypeOf((*DataSourceType)(nil)).Elem(),
		map[string]interface{}{
			"S3": DataSourceType_S3,
			"CONFLUENCE": DataSourceType_CONFLUENCE,
			"SALESFORCE": DataSourceType_SALESFORCE,
			"SHAREPOINT": DataSourceType_SHAREPOINT,
			"WEB_CRAWLER": DataSourceType_WEB_CRAWLER,
			"CUSTOM": DataSourceType_CUSTOM,
			"REDSHIFT_METADATA": DataSourceType_REDSHIFT_METADATA,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.DefaultPromptRouterIdentifier",
		reflect.TypeOf((*DefaultPromptRouterIdentifier)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "promptRouterId", GoGetter: "PromptRouterId"},
			_jsii_.MemberProperty{JsiiProperty: "routingModels", GoGetter: "RoutingModels"},
		},
		func() interface{} {
			return &jsiiProxy_DefaultPromptRouterIdentifier{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.EnrichmentStrategyConfigurationType",
		reflect.TypeOf((*EnrichmentStrategyConfigurationType)(nil)).Elem(),
		map[string]interface{}{
			"CHUNK_ENTITY_EXTRACTION": EnrichmentStrategyConfigurationType_CHUNK_ENTITY_EXTRACTION,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.FoundationModelContextEnrichmentProps",
		reflect.TypeOf((*FoundationModelContextEnrichmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.FoundationModelParsingStategyProps",
		reflect.TypeOf((*FoundationModelParsingStategyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Guardrail",
		reflect.TypeOf((*Guardrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContentFilter", GoMethod: "AddContentFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addContextualGroundingFilter", GoMethod: "AddContextualGroundingFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addDeniedTopicFilter", GoMethod: "AddDeniedTopicFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addManagedWordListFilter", GoMethod: "AddManagedWordListFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addPIIFilter", GoMethod: "AddPIIFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addRegexFilter", GoMethod: "AddRegexFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addWordFilter", GoMethod: "AddWordFilter"},
			_jsii_.MemberMethod{JsiiMethod: "addWordFilterFromFile", GoMethod: "AddWordFilterFromFile"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "contentFilters", GoGetter: "ContentFilters"},
			_jsii_.MemberProperty{JsiiProperty: "contextualGroundingFilters", GoGetter: "ContextualGroundingFilters"},
			_jsii_.MemberMethod{JsiiMethod: "createVersion", GoMethod: "CreateVersion"},
			_jsii_.MemberProperty{JsiiProperty: "deniedTopics", GoGetter: "DeniedTopics"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "hash", GoGetter: "Hash"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "managedWordListFilters", GoGetter: "ManagedWordListFilters"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationClientErrors", GoMethod: "MetricInvocationClientErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationLatency", GoMethod: "MetricInvocationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationServerErrors", GoMethod: "MetricInvocationServerErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsIntervened", GoMethod: "MetricInvocationsIntervened"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationThrottles", GoMethod: "MetricInvocationThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTextUnitCount", GoMethod: "MetricTextUnitCount"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "piiFilters", GoGetter: "PiiFilters"},
			_jsii_.MemberProperty{JsiiProperty: "regexFilters", GoGetter: "RegexFilters"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wordFilters", GoGetter: "WordFilters"},
		},
		func() interface{} {
			j := jsiiProxy_Guardrail{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_GuardrailBase)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailAction",
		reflect.TypeOf((*GuardrailAction)(nil)).Elem(),
		map[string]interface{}{
			"BLOCK": GuardrailAction_BLOCK,
			"ANONYMIZE": GuardrailAction_ANONYMIZE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailAttributes",
		reflect.TypeOf((*GuardrailAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailBase",
		reflect.TypeOf((*GuardrailBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationClientErrors", GoMethod: "MetricInvocationClientErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationLatency", GoMethod: "MetricInvocationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationServerErrors", GoMethod: "MetricInvocationServerErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsIntervened", GoMethod: "MetricInvocationsIntervened"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationThrottles", GoMethod: "MetricInvocationThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTextUnitCount", GoMethod: "MetricTextUnitCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GuardrailBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGuardrail)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailProps",
		reflect.TypeOf((*GuardrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.HierarchicalChunkingProps",
		reflect.TypeOf((*HierarchicalChunkingProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IAgent",
		reflect.TypeOf((*IAgent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAgent{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IAgentAlias",
		reflect.TypeOf((*IAgentAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agent", GoGetter: "Agent"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IAgentAlias{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IDataSource",
		reflect.TypeOf((*IDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IDataSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IGuardrail",
		reflect.TypeOf((*IGuardrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "lastUpdated", GoGetter: "LastUpdated"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationClientErrors", GoMethod: "MetricInvocationClientErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationLatency", GoMethod: "MetricInvocationLatency"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocations", GoMethod: "MetricInvocations"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationServerErrors", GoMethod: "MetricInvocationServerErrors"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationsIntervened", GoMethod: "MetricInvocationsIntervened"},
			_jsii_.MemberMethod{JsiiMethod: "metricInvocationThrottles", GoMethod: "MetricInvocationThrottles"},
			_jsii_.MemberMethod{JsiiMethod: "metricTextUnitCount", GoMethod: "MetricTextUnitCount"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IGuardrail{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IInferenceProfile",
		reflect.TypeOf((*IInferenceProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			return &jsiiProxy_IInferenceProfile{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IInvokable",
		reflect.TypeOf((*IInvokable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
		},
		func() interface{} {
			return &jsiiProxy_IInvokable{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IKendraKnowledgeBase",
		reflect.TypeOf((*IKendraKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "kendraIndex", GoGetter: "KendraIndex"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_IKendraKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKnowledgeBase)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IKnowledgeBase",
		reflect.TypeOf((*IKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_IKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IPrompt",
		reflect.TypeOf((*IPrompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptVersion", GoGetter: "PromptVersion"},
		},
		func() interface{} {
			return &jsiiProxy_IPrompt{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IPromptRouter",
		reflect.TypeOf((*IPromptRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "promptRouterArn", GoGetter: "PromptRouterArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptRouterId", GoGetter: "PromptRouterId"},
			_jsii_.MemberProperty{JsiiProperty: "routingEndpoints", GoGetter: "RoutingEndpoints"},
		},
		func() interface{} {
			return &jsiiProxy_IPromptRouter{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IVectorKnowledgeBase",
		reflect.TypeOf((*IVectorKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfluenceDataSource", GoMethod: "AddConfluenceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addCustomDataSource", GoMethod: "AddCustomDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addS3DataSource", GoMethod: "AddS3DataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSalesforceDataSource", GoMethod: "AddSalesforceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSharePointDataSource", GoMethod: "AddSharePointDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addWebCrawlerDataSource", GoMethod: "AddWebCrawlerDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieve", GoMethod: "GrantRetrieve"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieveAndGenerate", GoMethod: "GrantRetrieveAndGenerate"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "vectorStoreType", GoGetter: "VectorStoreType"},
		},
		func() interface{} {
			j := jsiiProxy_IVectorKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKnowledgeBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InferenceConfiguration",
		reflect.TypeOf((*InferenceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InferenceProfileBase",
		reflect.TypeOf((*InferenceProfileBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantProfileUsage", GoMethod: "GrantProfileUsage"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileArn", GoGetter: "InferenceProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceProfileId", GoGetter: "InferenceProfileId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_InferenceProfileBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInferenceProfile)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InferenceProfileType",
		reflect.TypeOf((*InferenceProfileType)(nil)).Elem(),
		map[string]interface{}{
			"SYSTEM_DEFINED": InferenceProfileType_SYSTEM_DEFINED,
			"APPLICATION": InferenceProfileType_APPLICATION,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		reflect.TypeOf((*InlineApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_InlineApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBase",
		reflect.TypeOf((*KendraKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieve", GoMethod: "GrantRetrieve"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieveAndGenerate", GoMethod: "GrantRetrieveAndGenerate"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "kendraIndex", GoGetter: "KendraIndex"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_KendraKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_KendraKnowledgeBaseBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseAttributes",
		reflect.TypeOf((*KendraKnowledgeBaseAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseBase",
		reflect.TypeOf((*KendraKnowledgeBaseBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieve", GoMethod: "GrantRetrieve"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieveAndGenerate", GoMethod: "GrantRetrieveAndGenerate"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "kendraIndex", GoGetter: "KendraIndex"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_KendraKnowledgeBaseBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_KnowledgeBaseBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KendraKnowledgeBaseProps",
		reflect.TypeOf((*KendraKnowledgeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBaseBase",
		reflect.TypeOf((*KnowledgeBaseBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieve", GoMethod: "GrantRetrieve"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieveAndGenerate", GoMethod: "GrantRetrieveAndGenerate"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBaseBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKnowledgeBase)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBaseType",
		reflect.TypeOf((*KnowledgeBaseType)(nil)).Elem(),
		map[string]interface{}{
			"VECTOR": KnowledgeBaseType_VECTOR,
			"KENDRA": KnowledgeBaseType_KENDRA,
			"SQL": KnowledgeBaseType_SQL,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.LambdaCustomTransformationProps",
		reflect.TypeOf((*LambdaCustomTransformationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ManagedWordFilterType",
		reflect.TypeOf((*ManagedWordFilterType)(nil)).Elem(),
		map[string]interface{}{
			"PROFANITY": ManagedWordFilterType_PROFANITY,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Memory",
		reflect.TypeOf((*Memory)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Memory{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.OrchestrationExecutor",
		reflect.TypeOf((*OrchestrationExecutor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunction", GoGetter: "LambdaFunction"},
		},
		func() interface{} {
			return &jsiiProxy_OrchestrationExecutor{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.OrchestrationType",
		reflect.TypeOf((*OrchestrationType)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": OrchestrationType_DEFAULT,
			"CUSTOM_ORCHESTRATION": OrchestrationType_CUSTOM_ORCHESTRATION,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIFilter",
		reflect.TypeOf((*PIIFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParentActionGroupSignature",
		reflect.TypeOf((*ParentActionGroupSignature)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ParentActionGroupSignature{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParsingModality",
		reflect.TypeOf((*ParsingModality)(nil)).Elem(),
		map[string]interface{}{
			"MULTIMODAL": ParsingModality_MULTIMODAL,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParsingStategy",
		reflect.TypeOf((*ParsingStategy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberMethod{JsiiMethod: "generatePolicyStatements", GoMethod: "GeneratePolicyStatements"},
		},
		func() interface{} {
			return &jsiiProxy_ParsingStategy{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParsingStategyType",
		reflect.TypeOf((*ParsingStategyType)(nil)).Elem(),
		map[string]interface{}{
			"FOUNDATION_MODEL": ParsingStategyType_FOUNDATION_MODEL,
			"DATA_AUTOMATION": ParsingStategyType_DATA_AUTOMATION,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Prompt",
		reflect.TypeOf((*Prompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVariant", GoMethod: "AddVariant"},
			_jsii_.MemberMethod{JsiiMethod: "createVersion", GoMethod: "CreateVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptName", GoGetter: "PromptName"},
			_jsii_.MemberProperty{JsiiProperty: "promptVersion", GoGetter: "PromptVersion"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "variants", GoGetter: "Variants"},
		},
		func() interface{} {
			j := jsiiProxy_Prompt{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrompt)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptAttributes",
		reflect.TypeOf((*PromptAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptBase",
		reflect.TypeOf((*PromptBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantGet", GoMethod: "GrantGet"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptVersion", GoGetter: "PromptVersion"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PromptBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPrompt)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptOverrideConfiguration",
		reflect.TypeOf((*PromptOverrideConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "parser", GoGetter: "Parser"},
			_jsii_.MemberProperty{JsiiProperty: "steps", GoGetter: "Steps"},
		},
		func() interface{} {
			return &jsiiProxy_PromptOverrideConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptProps",
		reflect.TypeOf((*PromptProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptRouter",
		reflect.TypeOf((*PromptRouter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "grantInvoke", GoMethod: "GrantInvoke"},
			_jsii_.MemberProperty{JsiiProperty: "invokableArn", GoGetter: "InvokableArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptRouterArn", GoGetter: "PromptRouterArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptRouterId", GoGetter: "PromptRouterId"},
			_jsii_.MemberProperty{JsiiProperty: "routingEndpoints", GoGetter: "RoutingEndpoints"},
		},
		func() interface{} {
			j := jsiiProxy_PromptRouter{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInvokable)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IPromptRouter)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptRouterProps",
		reflect.TypeOf((*PromptRouterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptStepConfiguration",
		reflect.TypeOf((*PromptStepConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptStepConfigurationCustomParser",
		reflect.TypeOf((*PromptStepConfigurationCustomParser)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptTemplateType",
		reflect.TypeOf((*PromptTemplateType)(nil)).Elem(),
		map[string]interface{}{
			"TEXT": PromptTemplateType_TEXT,
			"CHAT": PromptTemplateType_CHAT,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVariant",
		reflect.TypeOf((*PromptVariant)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "genAiResource", GoGetter: "GenAiResource"},
			_jsii_.MemberProperty{JsiiProperty: "inferenceConfiguration", GoGetter: "InferenceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "templateConfiguration", GoGetter: "TemplateConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "templateType", GoGetter: "TemplateType"},
		},
		func() interface{} {
			return &jsiiProxy_PromptVariant{}
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVersion",
		reflect.TypeOf((*PromptVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "prompt", GoGetter: "Prompt"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionArn", GoGetter: "VersionArn"},
		},
		func() interface{} {
			j := jsiiProxy_PromptVersion{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVersionProps",
		reflect.TypeOf((*PromptVersionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.RegexFilter",
		reflect.TypeOf((*RegexFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.RelayConversationHistoryType",
		reflect.TypeOf((*RelayConversationHistoryType)(nil)).Elem(),
		map[string]interface{}{
			"TO_COLLABORATOR": RelayConversationHistoryType_TO_COLLABORATOR,
			"DISABLED": RelayConversationHistoryType_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		reflect.TypeOf((*S3ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inlineSchema", GoGetter: "InlineSchema"},
			_jsii_.MemberProperty{JsiiProperty: "s3File", GoGetter: "S3File"},
		},
		func() interface{} {
			j := jsiiProxy_S3ApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSource",
		reflect.TypeOf((*S3DataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_S3DataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceNew)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSourceAssociationProps",
		reflect.TypeOf((*S3DataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSourceProps",
		reflect.TypeOf((*S3DataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SalesforceCrawlingFilters",
		reflect.TypeOf((*SalesforceCrawlingFilters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SalesforceDataSource",
		reflect.TypeOf((*SalesforceDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authSecret", GoGetter: "AuthSecret"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "endpoint", GoGetter: "Endpoint"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SalesforceDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceNew)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SalesforceDataSourceAssociationProps",
		reflect.TypeOf((*SalesforceDataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SalesforceDataSourceAuthType",
		reflect.TypeOf((*SalesforceDataSourceAuthType)(nil)).Elem(),
		map[string]interface{}{
			"OAUTH2_CLIENT_CREDENTIALS": SalesforceDataSourceAuthType_OAUTH2_CLIENT_CREDENTIALS,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SalesforceDataSourceProps",
		reflect.TypeOf((*SalesforceDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SalesforceObjectType",
		reflect.TypeOf((*SalesforceObjectType)(nil)).Elem(),
		map[string]interface{}{
			"ACCOUNT": SalesforceObjectType_ACCOUNT,
			"ATTACHMENT": SalesforceObjectType_ATTACHMENT,
			"CAMPAIGN": SalesforceObjectType_CAMPAIGN,
			"CONTENT_VERSION": SalesforceObjectType_CONTENT_VERSION,
			"PARTNER": SalesforceObjectType_PARTNER,
			"PRICEBOOK_2": SalesforceObjectType_PRICEBOOK_2,
			"CASE": SalesforceObjectType_CASE,
			"CONTACT": SalesforceObjectType_CONTACT,
			"CONTRACT": SalesforceObjectType_CONTRACT,
			"DOCUMENT": SalesforceObjectType_DOCUMENT,
			"IDEA": SalesforceObjectType_IDEA,
			"LEAD": SalesforceObjectType_LEAD,
			"OPPORTUNITY": SalesforceObjectType_OPPORTUNITY,
			"PRODUCT_2": SalesforceObjectType_PRODUCT_2,
			"SOLUTION": SalesforceObjectType_SOLUTION,
			"TASK": SalesforceObjectType_TASK,
			"FEED_ITEM": SalesforceObjectType_FEED_ITEM,
			"FEED_COMMENT": SalesforceObjectType_FEED_COMMENT,
			"KNOWLEDGE_KAV": SalesforceObjectType_KNOWLEDGE_KAV,
			"USER": SalesforceObjectType_USER,
			"COLLABORATION_GROUP": SalesforceObjectType_COLLABORATION_GROUP,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SessionSummaryMemoryProps",
		reflect.TypeOf((*SessionSummaryMemoryProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SharePointCrawlingFilters",
		reflect.TypeOf((*SharePointCrawlingFilters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SharePointDataSource",
		reflect.TypeOf((*SharePointDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authSecret", GoGetter: "AuthSecret"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "siteUrls", GoGetter: "SiteUrls"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SharePointDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceNew)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SharePointDataSourceAssociationProps",
		reflect.TypeOf((*SharePointDataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SharePointDataSourceAuthType",
		reflect.TypeOf((*SharePointDataSourceAuthType)(nil)).Elem(),
		map[string]interface{}{
			"OAUTH2_CLIENT_CREDENTIALS": SharePointDataSourceAuthType_OAUTH2_CLIENT_CREDENTIALS,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SharePointDataSourceProps",
		reflect.TypeOf((*SharePointDataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SharePointObjectType",
		reflect.TypeOf((*SharePointObjectType)(nil)).Elem(),
		map[string]interface{}{
			"PAGE": SharePointObjectType_PAGE,
			"EVENT": SharePointObjectType_EVENT,
			"FILE": SharePointObjectType_FILE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.TextPromptVariantProps",
		reflect.TypeOf((*TextPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolChoice",
		reflect.TypeOf((*ToolChoice)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "any", GoGetter: "Any"},
			_jsii_.MemberProperty{JsiiProperty: "auto", GoGetter: "Auto"},
			_jsii_.MemberProperty{JsiiProperty: "tool", GoGetter: "Tool"},
		},
		func() interface{} {
			return &jsiiProxy_ToolChoice{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ToolConfiguration",
		reflect.TypeOf((*ToolConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		reflect.TypeOf((*Topic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "definition", GoGetter: "Definition"},
			_jsii_.MemberProperty{JsiiProperty: "examples", GoGetter: "Examples"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
		},
		func() interface{} {
			return &jsiiProxy_Topic{}
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.TransformationStep",
		reflect.TypeOf((*TransformationStep)(nil)).Elem(),
		map[string]interface{}{
			"POST_CHUNKING": TransformationStep_POST_CHUNKING,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.VectorKnowledgeBase",
		reflect.TypeOf((*VectorKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfluenceDataSource", GoMethod: "AddConfluenceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addCustomDataSource", GoMethod: "AddCustomDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addS3DataSource", GoMethod: "AddS3DataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSalesforceDataSource", GoMethod: "AddSalesforceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSharePointDataSource", GoMethod: "AddSharePointDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addWebCrawlerDataSource", GoMethod: "AddWebCrawlerDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "associateToAgent", GoMethod: "AssociateToAgent"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieve", GoMethod: "GrantRetrieve"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieveAndGenerate", GoMethod: "GrantRetrieveAndGenerate"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseInstance", GoGetter: "KnowledgeBaseInstance"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "vectorStore", GoGetter: "VectorStore"},
			_jsii_.MemberProperty{JsiiProperty: "vectorStoreType", GoGetter: "VectorStoreType"},
		},
		func() interface{} {
			j := jsiiProxy_VectorKnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VectorKnowledgeBaseBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.VectorKnowledgeBaseAttributes",
		reflect.TypeOf((*VectorKnowledgeBaseAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.VectorKnowledgeBaseBase",
		reflect.TypeOf((*VectorKnowledgeBaseBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfluenceDataSource", GoMethod: "AddConfluenceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addCustomDataSource", GoMethod: "AddCustomDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addS3DataSource", GoMethod: "AddS3DataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSalesforceDataSource", GoMethod: "AddSalesforceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSharePointDataSource", GoMethod: "AddSharePointDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addWebCrawlerDataSource", GoMethod: "AddWebCrawlerDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieve", GoMethod: "GrantRetrieve"},
			_jsii_.MemberMethod{JsiiMethod: "grantRetrieveAndGenerate", GoMethod: "GrantRetrieveAndGenerate"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "vectorStoreType", GoGetter: "VectorStoreType"},
		},
		func() interface{} {
			j := jsiiProxy_VectorKnowledgeBaseBase{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_KnowledgeBaseBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVectorKnowledgeBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.VectorKnowledgeBaseProps",
		reflect.TypeOf((*VectorKnowledgeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.VectorStoreType",
		reflect.TypeOf((*VectorStoreType)(nil)).Elem(),
		map[string]interface{}{
			"OPENSEARCH_SERVERLESS": VectorStoreType_OPENSEARCH_SERVERLESS,
			"PINECONE": VectorStoreType_PINECONE,
			"AMAZON_AURORA": VectorStoreType_AMAZON_AURORA,
			"MONGO_DB_ATLAS": VectorStoreType_MONGO_DB_ATLAS,
			"NEPTUNE_ANALYTICS": VectorStoreType_NEPTUNE_ANALYTICS,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.VectorType",
		reflect.TypeOf((*VectorType)(nil)).Elem(),
		map[string]interface{}{
			"FLOATING_POINT": VectorType_FLOATING_POINT,
			"BINARY": VectorType_BINARY,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.WebCrawlerDataSource",
		reflect.TypeOf((*WebCrawlerDataSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "crawlingRate", GoGetter: "CrawlingRate"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceName", GoGetter: "DataSourceName"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceType", GoGetter: "DataSourceType"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "formatAsCfnProps", GoMethod: "FormatAsCfnProps"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "handleCommonPermissions", GoMethod: "HandleCommonPermissions"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBase", GoGetter: "KnowledgeBase"},
			_jsii_.MemberProperty{JsiiProperty: "maxPages", GoGetter: "MaxPages"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "siteUrls", GoGetter: "SiteUrls"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WebCrawlerDataSource{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_DataSourceNew)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.WebCrawlerDataSourceAssociationProps",
		reflect.TypeOf((*WebCrawlerDataSourceAssociationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.WebCrawlerDataSourceProps",
		reflect.TypeOf((*WebCrawlerDataSourceProps)(nil)).Elem(),
	)
}
