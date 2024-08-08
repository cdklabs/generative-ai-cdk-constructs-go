package generative-ai-cdk-constructs


// Experimental.
type ContainerImageConfig struct {
	// The image name. Images in Amazon ECR repositories can be specified by either using the full registry/repository:tag or registry/repository@digest.
	//
	// For example, `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>:latest` or
	// `012345678910.dkr.ecr.<region-name>.amazonaws.com/<repository-name>@sha256:94afd1f2e64d908bc90dbca0035a5b567EXAMPLE`.
	// Experimental.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
}

