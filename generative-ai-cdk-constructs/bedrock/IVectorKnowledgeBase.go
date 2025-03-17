package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Represents a Knowledge Base, either created with CDK or imported.
// Experimental.
type IVectorKnowledgeBase interface {
	IKnowledgeBase
	// Add a Confluence data source to the knowledge base.
	// Experimental.
	AddConfluenceDataSource(props *ConfluenceDataSourceAssociationProps) ConfluenceDataSource
	// Add a Custom data source to the knowledge base.
	// Experimental.
	AddCustomDataSource(props *CustomDataSourceAssociationProps) CustomDataSource
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
	// Grant the given identity permissions to retrieve content from the knowledge base.
	// Experimental.
	GrantRetrieve(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to retrieve content from the knowledge base.
	// Experimental.
	GrantRetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant
}

// The jsii proxy for IVectorKnowledgeBase
type jsiiProxy_IVectorKnowledgeBase struct {
	jsiiProxy_IKnowledgeBase
}

func (i *jsiiProxy_IVectorKnowledgeBase) AddConfluenceDataSource(props *ConfluenceDataSourceAssociationProps) ConfluenceDataSource {
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

func (i *jsiiProxy_IVectorKnowledgeBase) AddCustomDataSource(props *CustomDataSourceAssociationProps) CustomDataSource {
	if err := i.validateAddCustomDataSourceParameters(props); err != nil {
		panic(err)
	}
	var returns CustomDataSource

	_jsii_.Invoke(
		i,
		"addCustomDataSource",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IVectorKnowledgeBase) AddS3DataSource(props *S3DataSourceAssociationProps) S3DataSource {
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

func (i *jsiiProxy_IVectorKnowledgeBase) AddSalesforceDataSource(props *SalesforceDataSourceAssociationProps) SalesforceDataSource {
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

func (i *jsiiProxy_IVectorKnowledgeBase) AddSharePointDataSource(props *SharePointDataSourceAssociationProps) SharePointDataSource {
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

func (i *jsiiProxy_IVectorKnowledgeBase) AddWebCrawlerDataSource(props *WebCrawlerDataSourceAssociationProps) WebCrawlerDataSource {
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

func (i *jsiiProxy_IVectorKnowledgeBase) GrantRetrieve(grantee awsiam.IGrantable) awsiam.Grant {
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

func (i *jsiiProxy_IVectorKnowledgeBase) GrantRetrieveAndGenerate(grantee awsiam.IGrantable) awsiam.Grant {
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

