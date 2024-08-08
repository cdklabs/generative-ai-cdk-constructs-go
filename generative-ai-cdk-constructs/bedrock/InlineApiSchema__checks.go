//go:build !no_runtime_type_checking

package bedrock

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_InlineApiSchema) validateBindParameters(_scope constructs.Construct) error {
	if _scope == nil {
		return fmt.Errorf("parameter _scope is required, but nil was provided")
	}

	return nil
}

func validateInlineApiSchema_FromAssetParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func validateInlineApiSchema_FromBucketParameters(bucket awss3.IBucket, key *string) error {
	if bucket == nil {
		return fmt.Errorf("parameter bucket is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateInlineApiSchema_FromInlineParameters(schema *string) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

func validateNewInlineApiSchemaParameters(schema *string) error {
	if schema == nil {
		return fmt.Errorf("parameter schema is required, but nil was provided")
	}

	return nil
}

