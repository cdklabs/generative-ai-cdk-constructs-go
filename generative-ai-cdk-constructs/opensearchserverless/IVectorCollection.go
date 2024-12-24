package opensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/opensearchserverless/internal"
)

// Interface representing a vector collection.
// Experimental.
type IVectorCollection interface {
	awscdk.IResource
	// Return the given named metric for this VectorCollection.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of index requests.
	// Experimental.
	MetricIndexRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the search latency.
	// Experimental.
	MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the 90th percentile search latency.
	// Experimental.
	MetricSearchLatencyP90(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of search requests.
	// Experimental.
	MetricSearchRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// An IAM policy that allows API access to the collection.
	// Experimental.
	AossPolicy() awsiam.ManagedPolicy
	// The ARN of the collection.
	// Experimental.
	CollectionArn() *string
	// The ID of the collection.
	// Experimental.
	CollectionId() *string
	// The name of the collection.
	// Experimental.
	CollectionName() *string
	// Type of collection.
	// Experimental.
	CollectionType() VectorCollectionType
	// An OpenSearch Access Policy that allows access to the index.
	// Experimental.
	DataAccessPolicy() awsopensearchserverless.CfnAccessPolicy
	// Indicates whether standby replicas are enabled.
	// Experimental.
	StandbyReplicas() VectorCollectionStandbyReplicas
}

// The jsii proxy for IVectorCollection
type jsiiProxy_IVectorCollection struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IVectorCollection) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IVectorCollection) MetricIndexRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIndexRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIndexRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorCollection) MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSearchLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSearchLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorCollection) MetricSearchLatencyP90(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSearchLatencyP90Parameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSearchLatencyP90",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorCollection) MetricSearchRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSearchRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSearchRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IVectorCollection) AossPolicy() awsiam.ManagedPolicy {
	var returns awsiam.ManagedPolicy
	_jsii_.Get(
		j,
		"aossPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorCollection) CollectionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorCollection) CollectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorCollection) CollectionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorCollection) CollectionType() VectorCollectionType {
	var returns VectorCollectionType
	_jsii_.Get(
		j,
		"collectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorCollection) DataAccessPolicy() awsopensearchserverless.CfnAccessPolicy {
	var returns awsopensearchserverless.CfnAccessPolicy
	_jsii_.Get(
		j,
		"dataAccessPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVectorCollection) StandbyReplicas() VectorCollectionStandbyReplicas {
	var returns VectorCollectionStandbyReplicas
	_jsii_.Get(
		j,
		"standbyReplicas",
		&returns,
	)
	return returns
}

