package kendra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties needed for importing an existing Kendra Index.
// Experimental.
type KendraGenAiIndexAttributes struct {
	// The Id of the index.
	// Experimental.
	IndexId *string `field:"required" json:"indexId" yaml:"indexId"`
	// An IAM role that gives your Amazon Kendra index permissions.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

