package aosscwdashboard

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The properties for the CollectionMonitoringProps class.
// Experimental.
type CollectionMonitoringProps struct {
	// Experimental.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

