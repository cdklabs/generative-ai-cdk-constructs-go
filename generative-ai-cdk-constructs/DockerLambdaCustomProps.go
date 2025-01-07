package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// DockerLambdaCustomProps.
// Experimental.
type DockerLambdaCustomProps struct {
	// The source code of your Lambda function.
	//
	// You can point to a file in an
	// Amazon Simple Storage Service (Amazon S3) bucket or specify your source
	// code as inline text.
	// Experimental.
	Code awslambda.DockerImageCode `field:"required" json:"code" yaml:"code"`
	// Specify the configuration of AWS Distro for OpenTelemetry (ADOT) instrumentation.
	// Default: - No ADOT instrumentation.
	//
	// Experimental.
	AdotInstrumentation *awslambda.AdotInstrumentationConfig `field:"optional" json:"adotInstrumentation" yaml:"adotInstrumentation"`
	// Whether to allow the Lambda to send all ipv6 network traffic.
	//
	// If set to true, there will only be a single egress rule which allows all
	// outbound ipv6 traffic. If set to false, you must individually add traffic rules to allow the
	// Lambda to connect to network targets using ipv6.
	//
	// Do not specify this property if the `securityGroups` or `securityGroup` property is set.
	// Instead, configure `allowAllIpv6Outbound` directly on the security group.
	// Default: false.
	//
	// Experimental.
	AllowAllIpv6Outbound *bool `field:"optional" json:"allowAllIpv6Outbound" yaml:"allowAllIpv6Outbound"`
	// Sets the application log level for the function.
	// Default: ApplicationLogLevel.INFO
	//
	// Experimental.
	ApplicationLogLevelV2 awslambda.ApplicationLogLevel `field:"optional" json:"applicationLogLevelV2" yaml:"applicationLogLevelV2"`
	// The system architectures compatible with this lambda function.
	// Default: Architecture.X86_64
	//
	// Experimental.
	Architecture awslambda.Architecture `field:"optional" json:"architecture" yaml:"architecture"`
	// Code signing config associated with this function.
	// Default: - Not Sign the Code.
	//
	// Experimental.
	CodeSigningConfig awslambda.ICodeSigningConfig `field:"optional" json:"codeSigningConfig" yaml:"codeSigningConfig"`
	// Options for the `lambda.Version` resource automatically created by the `fn.currentVersion` method.
	// Default: - default options as described in `VersionOptions`.
	//
	// Experimental.
	CurrentVersionOptions *awslambda.VersionOptions `field:"optional" json:"currentVersionOptions" yaml:"currentVersionOptions"`
	// The SQS queue to use if DLQ is enabled.
	//
	// If SNS topic is desired, specify `deadLetterTopic` property instead.
	// Default: - SQS queue with 14 day retention period if `deadLetterQueueEnabled` is `true`.
	//
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Enabled DLQ.
	//
	// If `deadLetterQueue` is undefined,
	// an SQS queue with default options will be defined for your Function.
	// Default: - false unless `deadLetterQueue` is set, which implies DLQ is enabled.
	//
	// Experimental.
	DeadLetterQueueEnabled *bool `field:"optional" json:"deadLetterQueueEnabled" yaml:"deadLetterQueueEnabled"`
	// The SNS topic to use as a DLQ.
	//
	// Note that if `deadLetterQueueEnabled` is set to `true`, an SQS queue will be created
	// rather than an SNS topic. Using an SNS topic as a DLQ requires this property to be set explicitly.
	// Default: - no SNS topic.
	//
	// Experimental.
	DeadLetterTopic awssns.ITopic `field:"optional" json:"deadLetterTopic" yaml:"deadLetterTopic"`
	// Key-value pairs that Lambda caches and makes available for your Lambda functions.
	//
	// Use environment variables to apply configuration changes, such
	// as test and production environment configurations, without changing your
	// Lambda function source code.
	// Default: - No environment variables.
	//
	// Experimental.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The AWS KMS key that's used to encrypt your function's environment variables.
	// Default: - AWS Lambda creates and uses an AWS managed customer master key (CMK).
	//
	// Experimental.
	EnvironmentEncryption awskms.IKey `field:"optional" json:"environmentEncryption" yaml:"environmentEncryption"`
	// The size of the function’s /tmp directory in MiB.
	// Default: 512 MiB.
	//
	// Experimental.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// Event sources for this function.
	//
	// You can also add event sources using `addEventSource`.
	// Default: - No event sources.
	//
	// Experimental.
	Events *[]awslambda.IEventSource `field:"optional" json:"events" yaml:"events"`
	// The filesystem configuration for the lambda function.
	// Default: - will not mount any filesystem.
	//
	// Experimental.
	Filesystem awslambda.FileSystem `field:"optional" json:"filesystem" yaml:"filesystem"`
	// Initial policy statements to add to the created Lambda Role.
	//
	// You can call `addToRolePolicy` to the created lambda to add statements post creation.
	// Default: - No policy statements are added to the created Lambda role.
	//
	// Experimental.
	InitialPolicy *[]awsiam.PolicyStatement `field:"optional" json:"initialPolicy" yaml:"initialPolicy"`
	// Specify the version of CloudWatch Lambda insights to use for monitoring.
	// Default: - No Lambda Insights.
	//
	// Experimental.
	InsightsVersion awslambda.LambdaInsightsVersion `field:"optional" json:"insightsVersion" yaml:"insightsVersion"`
	// Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack subnets.
	//
	// Only used if 'vpc' is supplied.
	// Default: false.
	//
	// Experimental.
	Ipv6AllowedForDualStack *bool `field:"optional" json:"ipv6AllowedForDualStack" yaml:"ipv6AllowedForDualStack"`
	// Sets the loggingFormat for the function.
	// Default: LoggingFormat.TEXT
	//
	// Experimental.
	LoggingFormat awslambda.LoggingFormat `field:"optional" json:"loggingFormat" yaml:"loggingFormat"`
	// The log group the function sends logs to.
	//
	// By default, Lambda functions send logs to an automatically created default log group named /aws/lambda/\<function name\>.
	// However you cannot change the properties of this auto-created log group using the AWS CDK, e.g. you cannot set a different log retention.
	//
	// Use the `logGroup` property to create a fully customizable LogGroup ahead of time, and instruct the Lambda function to send logs to it.
	//
	// Providing a user-controlled log group was rolled out to commercial regions on 2023-11-16.
	// If you are deploying to another type of region, please check regional availability first.
	// Default: `/aws/lambda/${this.functionName}` - default log group created by Lambda
	//
	// Experimental.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `INFINITE`.
	//
	// This is a legacy API and we strongly recommend you move away from it if you can.
	// Instead create a fully customizable log group with `logs.LogGroup` and use the `logGroup` property
	// to instruct the Lambda function to send logs to it.
	// Migrating from `logRetention` to `logGroup` will cause the name of the log group to change.
	// Users and code and referencing the name verbatim will have to adjust.
	//
	// In AWS CDK code, you can access the log group name directly from the LogGroup construct:
	// ```ts
	// import * as logs from 'aws-cdk-lib/aws-logs';
	//
	// declare const myLogGroup: logs.LogGroup;
	// myLogGroup.logGroupName;
	// ```.
	// Default: logs.RetentionDays.INFINITE
	//
	// Experimental.
	LogRetention awslogs.RetentionDays `field:"optional" json:"logRetention" yaml:"logRetention"`
	// When log retention is specified, a custom resource attempts to create the CloudWatch log group.
	//
	// These options control the retry policy when interacting with CloudWatch APIs.
	//
	// This is a legacy API and we strongly recommend you migrate to `logGroup` if you can.
	// `logGroup` allows you to create a fully customizable log group and instruct the Lambda function to send logs to it.
	// Default: - Default AWS SDK retry options.
	//
	// Experimental.
	LogRetentionRetryOptions *awslambda.LogRetentionRetryOptions `field:"optional" json:"logRetentionRetryOptions" yaml:"logRetentionRetryOptions"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	//
	// This is a legacy API and we strongly recommend you migrate to `logGroup` if you can.
	// `logGroup` allows you to create a fully customizable log group and instruct the Lambda function to send logs to it.
	// Default: - A new role is created.
	//
	// Experimental.
	LogRetentionRole awsiam.IRole `field:"optional" json:"logRetentionRole" yaml:"logRetentionRole"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum: 60 seconds
	// Maximum: 6 hours.
	// Default: Duration.hours(6)
	//
	// Experimental.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The amount of memory, in MB, that is allocated to your Lambda function.
	//
	// Lambda uses this value to proportionally allocate the amount of CPU
	// power. For more information, see Resource Model in the AWS Lambda
	// Developer Guide.
	// Default: 128.
	//
	// Experimental.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// The destination for failed invocations.
	// Default: - no destination.
	//
	// Experimental.
	OnFailure awslambda.IDestination `field:"optional" json:"onFailure" yaml:"onFailure"`
	// The destination for successful invocations.
	// Default: - no destination.
	//
	// Experimental.
	OnSuccess awslambda.IDestination `field:"optional" json:"onSuccess" yaml:"onSuccess"`
	// Specify the configuration of Parameters and Secrets Extension.
	// Default: - No Parameters and Secrets Extension.
	//
	// Experimental.
	ParamsAndSecrets awslambda.ParamsAndSecretsLayerVersion `field:"optional" json:"paramsAndSecrets" yaml:"paramsAndSecrets"`
	// Enable profiling.
	// Default: - No profiling.
	//
	// Experimental.
	Profiling *bool `field:"optional" json:"profiling" yaml:"profiling"`
	// Profiling Group.
	// Default: - A new profiling group will be created if `profiling` is set.
	//
	// Experimental.
	ProfilingGroup awscodeguruprofiler.IProfilingGroup `field:"optional" json:"profilingGroup" yaml:"profilingGroup"`
	// Sets the Recursive Loop Protection for Lambda Function.
	//
	// It lets Lambda detect and terminate unintended recursive loops.
	// Default: RecursiveLoop.Terminate
	//
	// Experimental.
	RecursiveLoop awslambda.RecursiveLoop `field:"optional" json:"recursiveLoop" yaml:"recursiveLoop"`
	// The maximum of concurrent executions you want to reserve for the function.
	// Default: - No specific limit - account limit.
	//
	// Experimental.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum: 0
	// Maximum: 2.
	// Default: 2.
	//
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Sets the runtime management configuration for a function's version.
	// Default: Auto.
	//
	// Experimental.
	RuntimeManagementMode awslambda.RuntimeManagementMode `field:"optional" json:"runtimeManagementMode" yaml:"runtimeManagementMode"`
	// Enable SnapStart for Lambda Function.
	//
	// SnapStart is currently supported for Java 11, Java 17, Python 3.12, Python 3.13, and .NET 8 runtime
	// Default: - No snapstart.
	//
	// Experimental.
	SnapStart awslambda.SnapStartConf `field:"optional" json:"snapStart" yaml:"snapStart"`
	// Sets the system log level for the function.
	// Default: SystemLogLevel.INFO
	//
	// Experimental.
	SystemLogLevelV2 awslambda.SystemLogLevel `field:"optional" json:"systemLogLevelV2" yaml:"systemLogLevelV2"`
	// The function execution time (in seconds) after which Lambda terminates the function.
	//
	// Because the execution time affects cost, set this value
	// based on the function's expected execution time.
	// Default: Duration.seconds(3)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

