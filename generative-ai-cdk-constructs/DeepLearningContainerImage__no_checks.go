//go:build no_runtime_type_checking

package generative-ai-cdk-constructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeepLearningContainerImage) validateBindParameters(scope constructs.Construct, grantable awsiam.IGrantable) error {
	return nil
}

func validateDeepLearningContainerImage_FromAssetParameters(directory *string, options *awsecrassets.DockerImageAssetOptions) error {
	return nil
}

func validateDeepLearningContainerImage_FromDeepLearningContainerImageParameters(repositoryName *string, tag *string) error {
	return nil
}

func validateDeepLearningContainerImage_FromEcrRepositoryParameters(repository awsecr.IRepository) error {
	return nil
}

func validateNewDeepLearningContainerImageParameters(repositoryName *string, tag *string) error {
	return nil
}

