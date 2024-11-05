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
		},
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
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "hash", GoGetter: "Hash"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "managedWordListFilters", GoGetter: "ManagedWordListFilters"},
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
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
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
		"@cdklabs/generative-ai-cdk-constructs.bedrock.GuardrailConfiguration",
		reflect.TypeOf((*GuardrailConfiguration)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "grantApply", GoMethod: "GrantApply"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailArn", GoGetter: "GuardrailArn"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailId", GoGetter: "GuardrailId"},
			_jsii_.MemberProperty{JsiiProperty: "guardrailVersion", GoGetter: "GuardrailVersion"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
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
		"@cdklabs/generative-ai-cdk-constructs.bedrock.IKnowledgeBase",
		reflect.TypeOf((*IKnowledgeBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addConfluenceDataSource", GoMethod: "AddConfluenceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addS3DataSource", GoMethod: "AddS3DataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSalesforceDataSource", GoMethod: "AddSalesforceDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addSharePointDataSource", GoMethod: "AddSharePointDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "addWebCrawlerDataSource", GoMethod: "AddWebCrawlerDataSource"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
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
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.InferenceConfiguration",
		reflect.TypeOf((*InferenceConfiguration)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "addConfluenceDataSource", GoMethod: "AddConfluenceDataSource"},
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
			_jsii_.MemberProperty{JsiiProperty: "instruction", GoGetter: "Instruction"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseArn", GoGetter: "KnowledgeBaseArn"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseId", GoGetter: "KnowledgeBaseId"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseInstance", GoGetter: "KnowledgeBaseInstance"},
			_jsii_.MemberProperty{JsiiProperty: "knowledgeBaseState", GoGetter: "KnowledgeBaseState"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorStore", GoGetter: "VectorStore"},
		},
		func() interface{} {
			j := jsiiProxy_KnowledgeBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IKnowledgeBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBaseAttributes",
		reflect.TypeOf((*KnowledgeBaseAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.KnowledgeBaseProps",
		reflect.TypeOf((*KnowledgeBaseProps)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.PIIFilter",
		reflect.TypeOf((*PIIFilter)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.ParserMode",
		reflect.TypeOf((*ParserMode)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": ParserMode_DEFAULT,
			"OVERRIDDEN": ParserMode_OVERRIDDEN,
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
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.bedrock.RegexFilter",
		reflect.TypeOf((*RegexFilter)(nil)).Elem(),
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
		"@cdklabs/generative-ai-cdk-constructs.bedrock.S3Identifier",
		reflect.TypeOf((*S3Identifier)(nil)).Elem(),
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
