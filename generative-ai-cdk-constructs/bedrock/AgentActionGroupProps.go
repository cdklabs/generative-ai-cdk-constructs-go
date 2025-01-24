package bedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock"
)

// ****************************************************************************                        PROPS - Action Group Class ***************************************************************************.
// Experimental.
type AgentActionGroupProps struct {
	// The name of the action group.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API Schema.
	// Default: - No API Schema.
	//
	// Experimental.
	ApiSchema ApiSchema `field:"optional" json:"apiSchema" yaml:"apiSchema"`
	// A description of the action group.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the action group is available for the agent to invoke or not when sending an InvokeAgent request.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The action group executor.
	// Default: - No executor.
	//
	// Experimental.
	Executor ActionGroupExecutor `field:"optional" json:"executor" yaml:"executor"`
	// Specifies whether to delete the resource even if it's in use.
	// Default: false.
	//
	// Experimental.
	ForceDelete *bool `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Defines functions that each define parameters that the agent needs to invoke from the user.
	//
	// NO L2 yet as this doesn't make much sense IMHO.
	// Experimental.
	FunctionSchema *awsbedrock.CfnAgent_FunctionSchemaProperty `field:"optional" json:"functionSchema" yaml:"functionSchema"`
	// The AWS Defined signature for enabling certain capabilities in your agent.
	//
	// When this property is specified, you must leave the description, apiSchema,
	// and actionGroupExecutor fields blank for this action group.
	// Experimental.
	ParentActionGroupSignature ParentActionGroupSignature `field:"optional" json:"parentActionGroupSignature" yaml:"parentActionGroupSignature"`
}

