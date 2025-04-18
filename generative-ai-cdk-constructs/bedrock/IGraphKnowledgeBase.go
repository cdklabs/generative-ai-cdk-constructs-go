package bedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs/neptune"
)

// ****************************************************************************                            COMMON INTERFACES ***************************************************************************.
// Experimental.
type IGraphKnowledgeBase interface {
	IKnowledgeBase
	// The Neptune Analytics vector store.
	// Experimental.
	Graph() neptune.INeptuneGraph
}

// The jsii proxy for IGraphKnowledgeBase
type jsiiProxy_IGraphKnowledgeBase struct {
	jsiiProxy_IKnowledgeBase
}

func (j *jsiiProxy_IGraphKnowledgeBase) Graph() neptune.INeptuneGraph {
	var returns neptune.INeptuneGraph
	_jsii_.Get(
		j,
		"graph",
		&returns,
	)
	return returns
}

