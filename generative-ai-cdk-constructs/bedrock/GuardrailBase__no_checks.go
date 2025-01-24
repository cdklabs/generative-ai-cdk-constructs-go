//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GuardrailBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_GuardrailBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_GuardrailBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_GuardrailBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (g *jsiiProxy_GuardrailBase) validateGrantApplyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateGuardrailBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGuardrailBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGuardrailBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_GuardrailBase) validateSetGuardrailVersionParameters(val *string) error {
	return nil
}

func validateNewGuardrailBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

