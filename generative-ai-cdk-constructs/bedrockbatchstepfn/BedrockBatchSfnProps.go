package bedrockbatchstepfn

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type BedrockBatchSfnProps struct {
	// The S3 bucket where the Bedrock Batch Inference Job gets the input manifests.
	// Experimental.
	BedrockBatchInputBucket awss3.IBucket `field:"required" json:"bedrockBatchInputBucket" yaml:"bedrockBatchInputBucket"`
	// The S3 bucket where the Bedrock Batch Inference Job stores the output.
	// Experimental.
	BedrockBatchOutputBucket awss3.IBucket `field:"required" json:"bedrockBatchOutputBucket" yaml:"bedrockBatchOutputBucket"`
	// IAM policy used for Bedrock batch processing.
	//
	// The policy must have the following permissions for the models and inference profiles you plan to use:
	// - bedrock:InvokeModel
	// - bedrock:CreateModelInvocationJob.
	// Default: const bedrockBatchPolicy = new iam.ManagedPolicy(this, 'BedrockBatchPolicy', {
	//         statements: [
	//           new iam.PolicyStatement({
	//             sid: 'Inference',
	//             actions: ['bedrock:InvokeModel', 'bedrock:CreateModelInvocationJob'],
	//             resources: [
	//               'arn:aws:bedrock:*::foundation-model/*',
	//               Stack.of(this).formatArn({
	//                 service: 'bedrock',
	//                 resource: 'inference-profile',
	//                 resourceName: '*',
	//               }),
	//             ],
	//           }),
	//         ],
	// });.
	//
	// Experimental.
	BedrockBatchPolicy awsiam.IManagedPolicy `field:"optional" json:"bedrockBatchPolicy" yaml:"bedrockBatchPolicy"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath. DISCARD, which will cause the effective input to be the empty object {}.
	//
	// Input schema:
	// ```
	// {
	//   "job_name": string,        // Required. Name of the batch inference job
	//   "manifest_keys": string[],    // Required. List of S3 keys of the input manifest files
	//   "model_id": string        // Required. Model ID to use.
	// }
	// ```.
	// Default: The entire task input (JSON path '$').
	//
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// JSONPath expression to indicate where to inject the state's output May also be the special value JsonPath.
	//
	// DISCARD, which will cause the state's input to become its output.
	//
	// Output schema:
	// ```
	// {
	//   "status": string,        // Required. Status of the batch job. One of: "Completed" or "PartiallyCompleted"
	//   "bucket": string,        // Required. S3 bucket where the output is stored
	//   "keys": string[]         // Required. Array of S3 keys for the output files
	// }
	// ```.
	// Default: Replaces the entire input with the result (JSON path '$').
	//
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The timeout duration for the batch inference job.
	//
	// Must be between 24 hours and 1 week (168 hours).
	// Default: Duration.hours(48)
	//
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

