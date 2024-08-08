package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type WebCrawlerProps struct {
	// Optional user provided props to override the default props for the S3 Bucket.
	//
	// Providing both this and `existingOutputBucketObj` will cause an error.
	// Default: - Default props are used.
	//
	// Experimental.
	BucketOutputProps *awss3.BucketProps `field:"optional" json:"bucketOutputProps" yaml:"bucketOutputProps"`
	// Deploy Lambda crawler.
	// Default: - false.
	//
	// Experimental.
	EnableLambdaCrawler *bool `field:"optional" json:"enableLambdaCrawler" yaml:"enableLambdaCrawler"`
	// Existing instance of S3 Bucket object, providing both this and `bucketOutputProps` will cause an error.
	// Default: - None.
	//
	// Experimental.
	ExistingOutputBucketObj awss3.IBucket `field:"optional" json:"existingOutputBucketObj" yaml:"existingOutputBucketObj"`
	// Optional An existing VPC in which to deploy the construct.
	//
	// Providing both this and
	// vpcProps is an error.
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
	// Targets to be crawled.
	// Default: - none.
	//
	// Experimental.
	Targets *[]*CrawlerTarget `field:"optional" json:"targets" yaml:"targets"`
	// Optional custom properties for a VPC the construct will create.
	//
	// This VPC will
	// be used by the compute resources the construct creates. Providing
	// both this and existingVpc is an error.
	// Default: - none.
	//
	// Experimental.
	VpcProps *awsec2.VpcProps `field:"optional" json:"vpcProps" yaml:"vpcProps"`
}

