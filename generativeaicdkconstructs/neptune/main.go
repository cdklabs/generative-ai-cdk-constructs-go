package neptune

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.neptune.INeptuneGraph",
		reflect.TypeOf((*INeptuneGraph)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createNotebook", GoMethod: "CreateNotebook"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantExportTask", GoMethod: "GrantExportTask"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadOnly", GoMethod: "GrantReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "graphArn", GoGetter: "GraphArn"},
			_jsii_.MemberProperty{JsiiProperty: "graphEndpoint", GoGetter: "GraphEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "graphId", GoGetter: "GraphId"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricGraphSizeBytes", GoMethod: "MetricGraphSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricGraphStorageUsagePercent", GoMethod: "MetricGraphStorageUsagePercent"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumEdgeProperties", GoMethod: "MetricNumEdgeProperties"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumEdges", GoMethod: "MetricNumEdges"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherClientErrorsPerSec", GoMethod: "MetricNumOpenCypherClientErrorsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherRequestsPerSec", GoMethod: "MetricNumOpenCypherRequestsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherServerErrorsPerSec", GoMethod: "MetricNumOpenCypherServerErrorsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumQueuedRequestsPerSec", GoMethod: "MetricNumQueuedRequestsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumVectors", GoMethod: "MetricNumVectors"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumVertexProperties", GoMethod: "MetricNumVertexProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_INeptuneGraph{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.neptune.NeptuneGraph",
		reflect.TypeOf((*NeptuneGraph)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createNotebook", GoMethod: "CreateNotebook"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantExportTask", GoMethod: "GrantExportTask"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadOnly", GoMethod: "GrantReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "graphArn", GoGetter: "GraphArn"},
			_jsii_.MemberProperty{JsiiProperty: "graphEndpoint", GoGetter: "GraphEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "graphId", GoGetter: "GraphId"},
			_jsii_.MemberProperty{JsiiProperty: "graphName", GoGetter: "GraphName"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricGraphSizeBytes", GoMethod: "MetricGraphSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricGraphStorageUsagePercent", GoMethod: "MetricGraphStorageUsagePercent"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumEdgeProperties", GoMethod: "MetricNumEdgeProperties"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumEdges", GoMethod: "MetricNumEdges"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherClientErrorsPerSec", GoMethod: "MetricNumOpenCypherClientErrorsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherRequestsPerSec", GoMethod: "MetricNumOpenCypherRequestsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherServerErrorsPerSec", GoMethod: "MetricNumOpenCypherServerErrorsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumQueuedRequestsPerSec", GoMethod: "MetricNumQueuedRequestsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumVectors", GoMethod: "MetricNumVectors"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumVertexProperties", GoMethod: "MetricNumVertexProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "provisionedMemory", GoGetter: "ProvisionedMemory"},
			_jsii_.MemberProperty{JsiiProperty: "publicConnectivity", GoGetter: "PublicConnectivity"},
			_jsii_.MemberProperty{JsiiProperty: "replicaCount", GoGetter: "ReplicaCount"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NeptuneGraph{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_NeptuneGraphBase)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INeptuneGraph)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.neptune.NeptuneGraphBase",
		reflect.TypeOf((*NeptuneGraphBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "createNotebook", GoMethod: "CreateNotebook"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberMethod{JsiiMethod: "grantExportTask", GoMethod: "GrantExportTask"},
			_jsii_.MemberMethod{JsiiMethod: "grantFullAccess", GoMethod: "GrantFullAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantQuery", GoMethod: "GrantQuery"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadOnly", GoMethod: "GrantReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "graphArn", GoGetter: "GraphArn"},
			_jsii_.MemberProperty{JsiiProperty: "graphEndpoint", GoGetter: "GraphEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "graphId", GoGetter: "GraphId"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricCPUUtilization", GoMethod: "MetricCPUUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "metricGraphSizeBytes", GoMethod: "MetricGraphSizeBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricGraphStorageUsagePercent", GoMethod: "MetricGraphStorageUsagePercent"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumEdgeProperties", GoMethod: "MetricNumEdgeProperties"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumEdges", GoMethod: "MetricNumEdges"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherClientErrorsPerSec", GoMethod: "MetricNumOpenCypherClientErrorsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherRequestsPerSec", GoMethod: "MetricNumOpenCypherRequestsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumOpenCypherServerErrorsPerSec", GoMethod: "MetricNumOpenCypherServerErrorsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumQueuedRequestsPerSec", GoMethod: "MetricNumQueuedRequestsPerSec"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumVectors", GoMethod: "MetricNumVectors"},
			_jsii_.MemberMethod{JsiiMethod: "metricNumVertexProperties", GoMethod: "MetricNumVertexProperties"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NeptuneGraphBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INeptuneGraph)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.neptune.NeptuneGraphNotebook",
		reflect.TypeOf((*NeptuneGraphNotebook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "graphExplorerEndpoint", GoGetter: "GraphExplorerEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "jupyterLabEndpoint", GoGetter: "JupyterLabEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfig", GoGetter: "LifecycleConfig"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSize", GoGetter: "VolumeSize"},
		},
		func() interface{} {
			j := jsiiProxy_NeptuneGraphNotebook{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.neptune.NeptuneGraphNotebookProps",
		reflect.TypeOf((*NeptuneGraphNotebookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.neptune.NeptuneGraphProps",
		reflect.TypeOf((*NeptuneGraphProps)(nil)).Elem(),
	)
}
