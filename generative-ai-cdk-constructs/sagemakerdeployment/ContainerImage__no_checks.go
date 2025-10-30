//go:build no_runtime_type_checking

package sagemakerdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerImage) validateBindParameters(scope constructs.Construct, grantable awsiam.IGrantable) error {
	return nil
}

func validateContainerImage_FromAssetParameters(directory *string, options *awsecrassets.DockerImageAssetOptions) error {
	return nil
}

func validateContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

