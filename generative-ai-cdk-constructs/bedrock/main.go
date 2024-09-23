package bedrock

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ActionGroupExecutor",
		reflect.TypeOf((*ActionGroupExecutor)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AddAgentAliasProps",
		reflect.TypeOf((*AddAgentAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Agent",
		reflect.TypeOf((*Agent)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionGroups", GoGetter: "ActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addActionGroup", GoMethod: "AddActionGroup"},
			_jsii_.MemberMethod{JsiiMethod: "addActionGroups", GoMethod: "AddActionGroups"},
			_jsii_.MemberMethod{JsiiMethod: "addAlias", GoMethod: "AddAlias"},
			_jsii_.MemberMethod{JsiiMethod: "addGuardrail", GoMethod: "AddGuardrail"},
			_jsii_.MemberMethod{JsiiMethod: "addKnowledgeBase", GoMethod: "AddKnowledgeBase"},
			_jsii_.MemberMethod{JsiiMethod: "addKnowledgeBases", GoMethod: "AddKnowledgeBases"},
			_jsii_.MemberProperty{JsiiProperty: "agentArn", GoGetter: "AgentArn"},
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "agentInstance", GoGetter: "AgentInstance"},
			_jsii_.MemberProperty{JsiiProperty: "agentversion", GoGetter: "Agentversion"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBases", GoGetter: "KnowledgeBases"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Agent{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentActionGroup",
		reflect.TypeOf((*AgentActionGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionGroupExecutor", GoGetter: "ActionGroupExecutor"},
			_jsii_.MemberProperty{JsiiProperty: "actionGroupName", GoGetter: "ActionGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "actionGroupProperty", GoGetter: "ActionGroupProperty"},
			_jsii_.MemberProperty{JsiiProperty: "actionGroupState", GoGetter: "ActionGroupState"},
			_jsii_.MemberProperty{JsiiProperty: "apiSchema", GoGetter: "ApiSchema"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "functionSchema", GoGetter: "FunctionSchema"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "parentActionGroupSignature", GoGetter: "ParentActionGroupSignature"},
			_jsii_.MemberProperty{JsiiProperty: "skipResourceInUseCheckOnDelete", GoGetter: "SkipResourceInUseCheckOnDelete"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgentActionGroup{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
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
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
			_jsii_.MemberProperty{JsiiProperty: "aliasName", GoGetter: "AliasName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AgentAlias{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAgentAlias)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentAliasProps",
		reflect.TypeOf((*AgentAliasProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.AgentProps",
		reflect.TypeOf((*AgentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchema",
		reflect.TypeOf((*ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_ApiSchema{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ApiSchemaConfig",
		reflect.TypeOf((*ApiSchemaConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModel",
		reflect.TypeOf((*BedrockFoundationModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "asArn", GoMethod: "AsArn"},
			_jsii_.MemberMethod{JsiiMethod: "asIModel", GoMethod: "AsIModel"},
			_jsii_.MemberProperty{JsiiProperty: "modelId", GoGetter: "ModelId"},
			_jsii_.MemberProperty{JsiiProperty: "supportsAgents", GoGetter: "SupportsAgents"},
			_jsii_.MemberProperty{JsiiProperty: "supportsKnowledgeBase", GoGetter: "SupportsKnowledgeBase"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorDimensions", GoGetter: "VectorDimensions"},
		},
		func() interface{} {
			return &jsiiProxy_BedrockFoundationModel{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.BedrockFoundationModelProps",
		reflect.TypeOf((*BedrockFoundationModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CanadaSpecific",
		reflect.TypeOf((*CanadaSpecific)(nil)).Elem(),
		map[string]interface{}{
			"CA_HEALTH_NUMBER": CanadaSpecific_CA_HEALTH_NUMBER,
			"CA_SOCIAL_INSURANCE_NUMBER": CanadaSpecific_CA_SOCIAL_INSURANCE_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ChunkingStrategy",
		reflect.TypeOf((*ChunkingStrategy)(nil)).Elem(),
		map[string]interface{}{
			"FIXED_SIZE": ChunkingStrategy_FIXED_SIZE,
			"DEFAULT": ChunkingStrategy_DEFAULT,
			"NONE": ChunkingStrategy_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.CommonPromptVariantProps",
		reflect.TypeOf((*CommonPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentPolicyConfig",
		reflect.TypeOf((*ContentPolicyConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "contentPolicyConfigList", GoGetter: "ContentPolicyConfigList"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ContentPolicyConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContentPolicyConfigProps",
		reflect.TypeOf((*ContentPolicyConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextualGroundingFilterConfigType",
		reflect.TypeOf((*ContextualGroundingFilterConfigType)(nil)).Elem(),
		map[string]interface{}{
			"GROUNDING": ContextualGroundingFilterConfigType_GROUNDING,
			"RELEVANCE": ContextualGroundingFilterConfigType_RELEVANCE,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ContextualGroundingPolicyConfigProps",
		reflect.TypeOf((*ContextualGroundingPolicyConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.FiltersConfigStrength",
		reflect.TypeOf((*FiltersConfigStrength)(nil)).Elem(),
		map[string]interface{}{
			"NONE": FiltersConfigStrength_NONE,
			"LOW": FiltersConfigStrength_LOW,
			"MEDIUM": FiltersConfigStrength_MEDIUM,
			"HIGH": FiltersConfigStrength_HIGH,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.FiltersConfigType",
		reflect.TypeOf((*FiltersConfigType)(nil)).Elem(),
		map[string]interface{}{
			"VIOLENCE": FiltersConfigType_VIOLENCE,
			"HATE": FiltersConfigType_HATE,
			"INSULTS": FiltersConfigType_INSULTS,
			"MISCONDUCT": FiltersConfigType_MISCONDUCT,
			"PROMPT_ATTACK": FiltersConfigType_PROMPT_ATTACK,
			"SEXUAL": FiltersConfigType_SEXUAL,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Finance",
		reflect.TypeOf((*Finance)(nil)).Elem(),
		map[string]interface{}{
			"CREDIT_DEBIT_CARD_CVV": Finance_CREDIT_DEBIT_CARD_CVV,
			"CREDIT_DEBIT_CARD_EXPIRY": Finance_CREDIT_DEBIT_CARD_EXPIRY,
			"CREDIT_DEBIT_CARD_NUMBER": Finance_CREDIT_DEBIT_CARD_NUMBER,
			"PIN": Finance_PIN,
			"SWIFT_CODE": Finance_SWIFT_CODE,
			"INTERNATIONAL_BANK_ACCOUNT_NUMBER": Finance_INTERNATIONAL_BANK_ACCOUNT_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.General",
		reflect.TypeOf((*General)(nil)).Elem(),
		map[string]interface{}{
			"ADDRESS": General_ADDRESS,
			"AGE": General_AGE,
			"DRIVER_ID": General_DRIVER_ID,
			"EMAIL": General_EMAIL,
			"LICENSE_PLATE": General_LICENSE_PLATE,
			"NAME": General_NAME,
			"PASSWORD": General_PASSWORD,
			"PHONE": General_PHONE,
			"USERNAME": General_USERNAME,
			"VEHICLE_IDENTIFICATION_NUMBER": General_VEHICLE_IDENTIFICATION_NUMBER,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Guardrail",
		reflect.TypeOf((*Guardrail)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContextualGroundingPolicyConfig", GoMethod: "AddContextualGroundingPolicyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "addSensitiveInformationPolicyConfig", GoMethod: "AddSensitiveInformationPolicyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "addTags", GoMethod: "AddTags"},
			_jsii_.MemberMethod{JsiiMethod: "addTopicPolicyConfig", GoMethod: "AddTopicPolicyConfig"},
			_jsii_.MemberMethod{JsiiMethod: "addVersion", GoMethod: "AddVersion"},
			_jsii_.MemberMethod{JsiiMethod: "addWordPolicyConfig", GoMethod: "AddWordPolicyConfig"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailInstance", GoGetter: "GuardrailInstance"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyArn", GoGetter: "KmsKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "uploadWordPolicyFromFile", GoMethod: "UploadWordPolicyFromFile"},
		},
		func() interface{} {
			j := jsiiProxy_Guardrail{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailConfiguration",
		reflect.TypeOf((*GuardrailConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailProps",
		reflect.TypeOf((*GuardrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailVersion",
		reflect.TypeOf((*GuardrailVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersionInstance", GoGetter: "GuardrailVersionInstance"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GuardrailVersion{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IAgentAlias",
		reflect.TypeOf((*IAgentAlias)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "aliasArn", GoGetter: "AliasArn"},
			_jsii_.MemberProperty{JsiiProperty: "aliasId", GoGetter: "AliasId"},
		},
		func() interface{} {
			return &jsiiProxy_IAgentAlias{}
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IPrompt",
		reflect.TypeOf((*IPrompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
		},
		func() interface{} {
			return &jsiiProxy_IPrompt{}
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InferenceConfiguration",
		reflect.TypeOf((*InferenceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InformationTechnology",
		reflect.TypeOf((*InformationTechnology)(nil)).Elem(),
		map[string]interface{}{
			"URL": InformationTechnology_URL,
			"IP_ADDRESS": InformationTechnology_IP_ADDRESS,
			"MAC_ADDRESS": InformationTechnology_MAC_ADDRESS,
			"AWS_ACCESS_KEY": InformationTechnology_AWS_ACCESS_KEY,
			"AWS_SECRET_KEY": InformationTechnology_AWS_SECRET_KEY,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InlineApiSchema",
		reflect.TypeOf((*InlineApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_InlineApiSchema{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ApiSchema)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBase",
		reflect.TypeOf((*KnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "associateToAgent", GoMethod: "AssociateToAgent"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseInstance", GoGetter: "KnowledgeBaseInstance"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseState", GoGetter: "KnowledgeBaseState"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorStore", GoGetter: "VectorStore"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBaseProps",
		reflect.TypeOf((*KnowledgeBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParserMode",
		reflect.TypeOf((*ParserMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": ParserMode_DEFAULT,
			"OVERRIDDEN": ParserMode_OVERRIDDEN,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PiiEntitiesConfigAction",
		reflect.TypeOf((*PiiEntitiesConfigAction)(nil)).Elem(),
		map[string]interface{}{
			"BLOCK": PiiEntitiesConfigAction_BLOCK,
			"ANONYMIZE": PiiEntitiesConfigAction_ANONYMIZE,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Prompt",
		reflect.TypeOf((*Prompt)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addVariant", GoMethod: "AddVariant"},
			_jsii_.MemberMethod{JsiiMethod: "createVersion", GoMethod: "CreateVersion"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "promptArn", GoGetter: "PromptArn"},
			_jsii_.MemberProperty{JsiiProperty: "promptId", GoGetter: "PromptId"},
			_jsii_.MemberProperty{JsiiProperty: "promptName", GoGetter: "PromptName"},
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
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptConfiguration",
		reflect.TypeOf((*PromptConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptCreationMode",
		reflect.TypeOf((*PromptCreationMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": PromptCreationMode_DEFAULT,
			"OVERRIDDEN": PromptCreationMode_OVERRIDDEN,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptOverrideConfiguration",
		reflect.TypeOf((*PromptOverrideConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptProps",
		reflect.TypeOf((*PromptProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptState",
		reflect.TypeOf((*PromptState)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": PromptState_ENABLED,
			"DISABLED": PromptState_DISABLED,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptTemplateType",
		reflect.TypeOf((*PromptTemplateType)(nil)).Elem(),
		map[string]interface{}{
			"TEXT": PromptTemplateType_TEXT,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptType",
		reflect.TypeOf((*PromptType)(nil)).Elem(),
		map[string]interface{}{
			"PRE_PROCESSING": PromptType_PRE_PROCESSING,
			"ORCHESTRATION": PromptType_ORCHESTRATION,
			"POST_PROCESSING": PromptType_POST_PROCESSING,
			"KNOWLEDGE_BASE_RESPONSE_GENERATION": PromptType_KNOWLEDGE_BASE_RESPONSE_GENERATION,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PromptVariant",
		reflect.TypeOf((*PromptVariant)(nil)).Elem(),
		[]_jsii_.Member{
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
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3ApiSchema",
		reflect.TypeOf((*S3ApiSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
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
			_jsii_.MemberProperty{JsiiProperty: "dataSource", GoGetter: "DataSource"},
			_jsii_.MemberProperty{JsiiProperty: "dataSourceId", GoGetter: "DataSourceId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_S3DataSource{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3DataSourceProps",
		reflect.TypeOf((*S3DataSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3Identifier",
		reflect.TypeOf((*S3Identifier)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SensitiveInformationPolicyConfig",
		reflect.TypeOf((*SensitiveInformationPolicyConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "piiConfigList", GoGetter: "PiiConfigList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SensitiveInformationPolicyConfig{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.SensitiveInformationPolicyConfigProps",
		reflect.TypeOf((*SensitiveInformationPolicyConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.TextPromptVariantProps",
		reflect.TypeOf((*TextPromptVariantProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.Topic",
		reflect.TypeOf((*Topic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createTopic", GoMethod: "CreateTopic"},
			_jsii_.MemberMethod{JsiiMethod: "financialAdviceTopic", GoMethod: "FinancialAdviceTopic"},
			_jsii_.MemberMethod{JsiiMethod: "inappropriateContent", GoMethod: "InappropriateContent"},
			_jsii_.MemberMethod{JsiiMethod: "legalAdvice", GoMethod: "LegalAdvice"},
			_jsii_.MemberMethod{JsiiMethod: "medicalAdvice", GoMethod: "MedicalAdvice"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "politicalAdviceTopic", GoMethod: "PoliticalAdviceTopic"},
			_jsii_.MemberMethod{JsiiMethod: "topicConfigPropertyList", GoMethod: "TopicConfigPropertyList"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Topic{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.TopicProps",
		reflect.TypeOf((*TopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.UKSpecific",
		reflect.TypeOf((*UKSpecific)(nil)).Elem(),
		map[string]interface{}{
			"UK_NATIONAL_HEALTH_SERVICE_NUMBER": UKSpecific_UK_NATIONAL_HEALTH_SERVICE_NUMBER,
			"UK_NATIONAL_INSURANCE_NUMBER": UKSpecific_UK_NATIONAL_INSURANCE_NUMBER,
			"UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER": UKSpecific_UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.USASpecific",
		reflect.TypeOf((*USASpecific)(nil)).Elem(),
		map[string]interface{}{
			"US_BANK_ACCOUNT_NUMBER": USASpecific_US_BANK_ACCOUNT_NUMBER,
			"US_BANK_ROUTING_NUMBER": USASpecific_US_BANK_ROUTING_NUMBER,
			"US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER": USASpecific_US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER,
			"US_PASSPORT_NUMBER": USASpecific_US_PASSPORT_NUMBER,
			"US_SOCIAL_SECURITY_NUMBER": USASpecific_US_SOCIAL_SECURITY_NUMBER,
		},
	)
}
