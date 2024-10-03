package amazonaurora

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Interface representing the resources required for a database cluster.
// Experimental.
type DatabaseClusterResources struct {
	// The security group associated with the Aurora cluster.
	// Experimental.
	AuroraSecurityGroup awsec2.SecurityGroup `field:"required" json:"auroraSecurityGroup" yaml:"auroraSecurityGroup"`
	// The unique cluster identifier of the Aurora RDS cluster.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The ARN of your existing Amazon Aurora DB cluster.
	// Experimental.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The secret containing the database credentials.
	//
	// The secret must contain `username` and `password` values.
	// Experimental.
	Secret awssecretsmanager.ISecret `field:"required" json:"secret" yaml:"secret"`
	// The VPC in which the database cluster is located.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The Amazon Aurora RDS cluster.
	// Experimental.
	AuroraCluster awsrds.DatabaseCluster `field:"optional" json:"auroraCluster" yaml:"auroraCluster"`
}

