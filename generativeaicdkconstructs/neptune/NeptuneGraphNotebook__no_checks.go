//go:build no_runtime_type_checking

package neptune

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NeptuneGraphNotebook) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphNotebook) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NeptuneGraphNotebook) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNeptuneGraphNotebook_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNeptuneGraphNotebook_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNeptuneGraphNotebook_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNeptuneGraphNotebookParameters(scope constructs.Construct, id *string, props *NeptuneGraphNotebookProps) error {
	return nil
}

