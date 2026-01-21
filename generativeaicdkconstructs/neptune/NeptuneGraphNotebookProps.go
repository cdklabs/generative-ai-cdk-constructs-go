package neptune

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties for creating a new Neptune Graph Notebook.
// Experimental.
type NeptuneGraphNotebookProps struct {
	// The Neptune Analytics Graph this notebook will be connected to.
	// Experimental.
	Graph INeptuneGraph `field:"required" json:"graph" yaml:"graph"`
	// The instance type of the notebook instance.
	// Default: "ec2.InstanceType.of(ec2.InstanceClass.T3, ec2.InstanceSize.MEDIUM)"
	//
	// Experimental.
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The size of the EBS volume.
	// Default: - 5 GiB.
	//
	// Experimental.
	VolumeSize awscdk.Size `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

