package generative-ai-cdk-constructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// The properties for the BedrockCwDashboardProps class.
// Experimental.
type BedrockCwDashboardProps struct {
	// Optional A name for the dashboard which will be created.
	//
	// If existingDashboard is defined, this value will be ignored.
	// If not provided, the construct will create a new dashboard named 'BedrockMetricsDashboard'.
	// Default: - none.
	//
	// Experimental.
	DashboardName *string `field:"optional" json:"dashboardName" yaml:"dashboardName"`
	// Optional An existing dashboard where metrics will be added to.
	//
	// If not provided, the construct will create a new dashboard.
	// Default: - none.
	//
	// Experimental.
	ExistingDashboard awscloudwatch.Dashboard `field:"optional" json:"existingDashboard" yaml:"existingDashboard"`
}

