//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

func (i *jsiiProxy_IInvokable) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	if grantee == nil {
		return fmt.Errorf("parameter grantee is required, but nil was provided")
	}

	return nil
}

