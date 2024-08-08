package amazonaurora

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Experimental.
type AmazonAuroraDefaultVectorStoreProps struct {
	// The embeddings model vector dimension for the knowledge base.
	//
	// Must be identical as in the KnowledgeBase construct.
	// This is due to the factor that the embeddings models
	// have different vector dimensions and this construct
	// needs to know the vector dimensions to create the vector
	// index of appropriate dimensions in the Aurora database.
	// Experimental.
	EmbeddingsModelVectorDimension *float64 `field:"required" json:"embeddingsModelVectorDimension" yaml:"embeddingsModelVectorDimension"`
	// The VPC where the Aurora Vector Store will be deployed in.
	//
	// The provided VPC must have at least one subnet of type
	// `ec2.SubnetType.PUBLIC` and at least one subnet of type
	// `ec2.SubnetType.PRIVATE_WITH_EGRESS`. If no subnets of these
	// types are available, the deployment will fail.
	// If not provided, a new VPC with the required subnet
	// configuration will be created automatically.
	// Default: - "A new VPC will be created."
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

