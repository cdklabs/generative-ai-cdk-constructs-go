package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/bedrock/internal"
)

// Represents a Knowledge Base, either created with CDK or imported.
// Experimental.
type IKnowledgeBase interface {
	awscdk.IResource
	// Add a Confluence data source to the knowledge base.
	// Experimental.
	AddConfluenceDataSource(props *ConfluenceDataSourceAssociationProps) ConfluenceDataSource
	// Add an S3 data source to the knowledge base.
	// Experimental.
	AddS3DataSource(props *S3DataSourceAssociationProps) S3DataSource
	// Add a Salesforce data source to the knowledge base.
	// Experimental.
	AddSalesforceDataSource(props *SalesforceDataSourceAssociationProps) SalesforceDataSource
	// Add a SharePoint data source to the knowledge base.
	// Experimental.
	AddSharePointDataSource(props *SharePointDataSourceAssociationProps) SharePointDataSource
	// Add a web crawler data source to the knowledge base.
	// Experimental.
	AddWebCrawlerDataSource(props *WebCrawlerDataSourceAssociationProps) WebCrawlerDataSource
	// Grant the given principal identity permissions to perform actions on this knowledge base.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant the given identity permissions to query the knowledge base.
	//
	// This contains:
	// - Retrieve
	// - RetrieveAndGenerate.
	// Experimental.
	GrantQuery(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to retrieve content from the knowledge base.
	// Experimental.
	GrantRetrieve(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to retrieve content from the knowledge base.
	// Experimental.
	GrantRetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant
	// A description of the knowledge base.
	// Experimental.
	Description() *string
	// Instructions for agents based on the design and type of information of the Knowledge Base.
	//
	// This will impact how Agents interact with the Knowledge Base.
	// Experimental.
	Instruction() *string
	// The ARN of the knowledge base.
	//
	// Example:
	//   "arn:aws:bedrock:us-east-1:123456789012:knowledge-base/KB12345678"
	//
	// Experimental.
	KnowledgeBaseArn() *string
	// The ID of the knowledge base.
	//
	// Example:
	//   "KB12345678"
	//
	// Experimental.
	KnowledgeBaseId() *string
	// The role associated with the knowledge base.
	// Experimental.
	Role() awsiam.IRole
}

// The jsii proxy for IKnowledgeBase
type jsiiProxy_IKnowledgeBase struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IKnowledgeBase) AddConfluenceDataSource(props *ConfluenceDataSourceAssociationProps) ConfluenceDataSource {
	if err := i.validateAddConfluenceDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns ConfluenceDataSource

	_jsii_.Invoke(
		i,
		"addConfluenceDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) AddS3DataSource(props *S3DataSourceAssociationProps) S3DataSource {
	if err := i.validateAddS3DataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns S3DataSource

	_jsii_.Invoke(
		i,
		"addS3DataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) AddSalesforceDataSource(props *SalesforceDataSourceAssociationProps) SalesforceDataSource {
	if err := i.validateAddSalesforceDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns SalesforceDataSource

	_jsii_.Invoke(
		i,
		"addSalesforceDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) AddSharePointDataSource(props *SharePointDataSourceAssociationProps) SharePointDataSource {
	if err := i.validateAddSharePointDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns SharePointDataSource

	_jsii_.Invoke(
		i,
		"addSharePointDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) AddWebCrawlerDataSource(props *WebCrawlerDataSourceAssociationProps) WebCrawlerDataSource {
	if err := i.validateAddWebCrawlerDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns WebCrawlerDataSource

	_jsii_.Invoke(
		i,
		"addWebCrawlerDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) GrantQuery(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IKnowledgeBase) GrantRetrieve(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantRetrieveParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRetrieve",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IKnowledgeBase) GrantRetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantRetrieveAndGenerateParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRetrieveAndGenerate",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IKnowledgeBase) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

