package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type TextToSqlProps struct {
	// Database secret.
	//
	// DB credentials to connect to Database.
	// Experimental.
	DatabaseSecretARN *string `field:"required" json:"databaseSecretARN" yaml:"databaseSecretARN"`
	// Database name.
	//
	// This is the target database against which the query will be generated.
	// Experimental.
	DbName DbName `field:"required" json:"dbName" yaml:"dbName"`
	// The source of metatdata.
	//
	// This metadata is required to reduce the natual language ambiguity
	// in order to generate the correct sql query. A knowledge layer is used to map the natural language
	// to the database schema.
	// Two metatdata source are supported:
	// 1: config_file - A local json file containing the knowledge layer key value pair.
	// 2: Knowledge base - Semantic search is used to fetch the knowledge layer from AWS knowledge base.
	// Default: - config_file.
	//
	// Experimental.
	MetadataSource *string `field:"required" json:"metadataSource" yaml:"metadataSource"`
	// Optional.
	//
	// User provided props to override the default props for the S3 Bucket.
	// Default: - Default props are used.
	//
	// Experimental.
	ConfigAssetsBucketProps *awss3.BucketProps `field:"optional" json:"configAssetsBucketProps" yaml:"configAssetsBucketProps"`
	// Optional.
	//
	// Allows to provide custom lambda code for all pre steps required before generating the query.
	// If not provided, default code will be used.
	// Experimental.
	CustomQueryConfigurerLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customQueryConfigurerLambdaProps" yaml:"customQueryConfigurerLambdaProps"`
	// Optional.
	//
	// Allows to provide custom lambda code for executing the query.
	// If not provided, default code will be used.
	// Experimental.
	CustomQueryExecutorLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customQueryExecutorLambdaProps" yaml:"customQueryExecutorLambdaProps"`
	// Optional.
	//
	// Allows to provide custom lambda code for generating the query.
	// If not provided, default code will be used.
	// Experimental.
	CustomQueryGeneratorLambdaProps *DockerLambdaCustomProps `field:"optional" json:"customQueryGeneratorLambdaProps" yaml:"customQueryGeneratorLambdaProps"`
	// Optional.
	//
	// db port number.
	// Default: -1534.
	//
	// Experimental.
	DbPort *float64 `field:"optional" json:"dbPort" yaml:"dbPort"`
	// Optional user provided event bus props.
	// Default: - Default props are used.
	//
	// Experimental.
	EventBusProps *awsevents.EventBusProps `field:"optional" json:"eventBusProps" yaml:"eventBusProps"`
	// Optional.
	//
	// Existing s3 Bucket to store the config files.
	// Default: - None.
	//
	// Experimental.
	ExistingconfigAssetsBucketObj awss3.IBucket `field:"optional" json:"existingconfigAssetsBucketObj" yaml:"existingconfigAssetsBucketObj"`
	// Optional.
	//
	// Security group for the db instance which this construct will use.
	// If no exisiting security group is provided it will create one from the vpc.
	// Default: - none.
	//
	// Experimental.
	ExistingDBSecurityGroup awsec2.SecurityGroup `field:"optional" json:"existingDBSecurityGroup" yaml:"existingDBSecurityGroup"`
	// Optional.
	//
	// Existing instance of event bus, providing both this and `eventBusProps` will cause an error.
	// Default: - None.
	//
	// Experimental.
	ExistingEventBusInterface awsevents.IEventBus `field:"optional" json:"existingEventBusInterface" yaml:"existingEventBusInterface"`
	// Optional.
	//
	// Security group for the lambda function which this construct will use.
	// If no exisiting security group is provided it will create one from the vpc.
	// Default: - none.
	//
	// Experimental.
	ExistingLambdaSecurityGroup awsec2.SecurityGroup `field:"optional" json:"existingLambdaSecurityGroup" yaml:"existingLambdaSecurityGroup"`
	// Optional.
	//
	// An existing subnet group can be used to deploy the construct.
	// Default: - none.
	//
	// Experimental.
	ExistingSubnetGroup awsrds.SubnetGroup `field:"optional" json:"existingSubnetGroup" yaml:"existingSubnetGroup"`
	// Optional.
	//
	// An existing VPC can be used to deploy the construct.
	// Providing both this and vpcProps is an error.
	// Default: - none.
	//
	// Experimental.
	ExistingVpc awsec2.IVpc `field:"optional" json:"existingVpc" yaml:"existingVpc"`
	// Enable observability.
	//
	// Warning: associated cost with the services
	// used. Best practice to enable by default.
	// Default: - true.
	//
	// Experimental.
	Observability *bool `field:"optional" json:"observability" yaml:"observability"`
	// Value will be appended to resources name.
	// Default: - _dev.
	//
	// Experimental.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Optional.
	//
	// The construct creates a custom VPC based on vpcProps.
	// Providing both this and existingVpc is an error.
	// Default: - none.
	//
	// Experimental.
	VpcProps *awsec2.VpcProps `field:"optional" json:"vpcProps" yaml:"vpcProps"`
}

