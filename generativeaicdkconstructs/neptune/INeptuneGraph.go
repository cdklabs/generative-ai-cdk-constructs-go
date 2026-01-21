package neptune

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generativeaicdkconstructs/neptune/internal"
)

// Interface representing a Neptune Graph database.
// Experimental.
type INeptuneGraph interface {
	awscdk.IResource
	// Creates a Neptune Graph Notebook for the graph.
	//
	// Defaults to a ml.t3.medium instance type.
	// **Note: Creating a notebook will incur additional AWS costs for the notebook instance.**
	// Experimental.
	CreateNotebook(params *NeptuneGraphNotebookProps) NeptuneGraphNotebook
	// Grant the given principal identity permissions to perform actions on this agent alias.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions *[]*string) awsiam.Grant
	// Grant the given identity the permissions to export the graph into columnar structured .csv and .parquet files.
	// Experimental.
	GrantExportTask(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity full access to the Graph.
	// Experimental.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity the permissions to query the Graph.
	// Experimental.
	GrantQuery(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity the permissions to read the Graph.
	// Experimental.
	GrantReadOnly(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this guardrail.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for CPU utilization.
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for graph size in bytes.
	// Experimental.
	MetricGraphSizeBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for graph storage usage percentage.
	// Experimental.
	MetricGraphStorageUsagePercent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of edge properties.
	// Experimental.
	MetricNumEdgeProperties(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of edges.
	// Experimental.
	MetricNumEdges(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of OpenCypher client errors per second.
	// Experimental.
	MetricNumOpenCypherClientErrorsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of OpenCypher requests per second.
	// Experimental.
	MetricNumOpenCypherRequestsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of OpenCypher server errors per second.
	// Experimental.
	MetricNumOpenCypherServerErrorsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of queued requests per second.
	// Experimental.
	MetricNumQueuedRequestsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of vectors.
	// Experimental.
	MetricNumVectors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns metric for number of vertex properties.
	// Experimental.
	MetricNumVertexProperties(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The Neptune Graph ARN.
	//
	// Example:
	//   "arn:aws:neptune-graph:us-east-1:111122223333:graph/g-12a3bcdef4"
	//
	// Experimental.
	GraphArn() *string
	// The Neptune Graph endpoint.
	//
	// Example:
	//   "g-12a3bcdef4.us-east-1.neptune-graph.amazonaws.com"
	//
	// Experimental.
	GraphEndpoint() *string
	// The Neptune Graph Identifier.
	//
	// Example:
	//   "g-12a3bcdef4"
	//
	// Experimental.
	GraphId() *string
}

// The jsii proxy for INeptuneGraph
type jsiiProxy_INeptuneGraph struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INeptuneGraph) CreateNotebook(params *NeptuneGraphNotebookProps) NeptuneGraphNotebook {
	if err := i.validateCreateNotebookParameters(params); err != nil {
		panic(err)
	}
	var returns NeptuneGraphNotebook

	_jsii_.Invoke(
		i,
		"createNotebook",
		[]interface{}{params},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) Grant(grantee awsiam.IGrantable, actions *[]*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee, actions); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		[]interface{}{grantee, actions},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) GrantExportTask(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantExportTaskParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantExportTask",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantFullAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) GrantQuery(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantQueryParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantQuery",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) GrantReadOnly(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadOnlyParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadOnly",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricGraphSizeBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGraphSizeBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGraphSizeBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricGraphStorageUsagePercent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGraphStorageUsagePercentParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGraphStorageUsagePercent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumEdgeProperties(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumEdgePropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumEdgeProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumEdges(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumEdgesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumEdges",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumOpenCypherClientErrorsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumOpenCypherClientErrorsPerSecParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumOpenCypherClientErrorsPerSec",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumOpenCypherRequestsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumOpenCypherRequestsPerSecParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumOpenCypherRequestsPerSec",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumOpenCypherServerErrorsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumOpenCypherServerErrorsPerSecParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumOpenCypherServerErrorsPerSec",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumQueuedRequestsPerSec(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumQueuedRequestsPerSecParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumQueuedRequestsPerSec",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumVectors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumVectorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumVectors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INeptuneGraph) MetricNumVertexProperties(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNumVertexPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNumVertexProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_INeptuneGraph) GraphArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INeptuneGraph) GraphEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INeptuneGraph) GraphId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"graphId",
		&returns,
	)
	return returns
}

