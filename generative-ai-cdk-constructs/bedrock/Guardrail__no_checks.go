//go:build no_runtime_type_checking

package bedrock

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Guardrail) validateAddContentFilterParameters(filter *ContentFilter) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddContextualGroundingFilterParameters(filter *ContextualGroundingFilter) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddDeniedTopicFilterParameters(filter Topic) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddManagedWordListFilterParameters(filter ManagedWordFilterType) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddPIIFilterParameters(filter *PIIFilter) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddRegexFilterParameters(filter *RegexFilter) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddWordFilterParameters(filter *string) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateAddWordFilterFromFileParameters(filePath *string) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (g *jsiiProxy_Guardrail) validateGrantApplyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateGuardrail_FromCfnGuardrailParameters(cfnGuardrail awsbedrock.CfnGuardrail) error {
	return nil
}

func validateGuardrail_FromGuardrailAttributesParameters(scope constructs.Construct, id *string, attrs *GuardrailAttributes) error {
	return nil
}

func validateGuardrail_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGuardrail_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateGuardrail_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_Guardrail) validateSetGuardrailVersionParameters(val *string) error {
	return nil
}

func validateNewGuardrailParameters(scope constructs.Construct, id *string, props *GuardrailProps) error {
	return nil
}

