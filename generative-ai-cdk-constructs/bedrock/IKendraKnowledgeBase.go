package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/kendra"
)

// ****************************************************************************                            COMMON INTERFACES ***************************************************************************.
// Experimental.
type IKendraKnowledgeBase interface {
	IKnowledgeBase
	// The GenAI Kendra Index.
	// Experimental.
	KendraIndex() kendra.IKendraGenAiIndex
}

// The jsii proxy for IKendraKnowledgeBase
type jsiiProxy_IKendraKnowledgeBase struct {
	jsiiProxy_IKnowledgeBase
}

func (j *jsiiProxy_IKendraKnowledgeBase) KendraIndex() kendra.IKendraGenAiIndex {
	var returns kendra.IKendraGenAiIndex
	_jsii_.Get(
		j,
		"kendraIndex",
		&returns,
	)
	return returns
}

