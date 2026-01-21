# S3 Vectors

<!--BEGIN STABILITY BANNER-->---


![Stability: Experimental](https://img.shields.io/badge/stability-Experimental-important.svg?style=for-the-badge)

> All classes are under active development and subject to non-backward compatible changes or removal in any
> future version. These are not subject to the [Semantic Versioning](https://semver.org/) model.
> This means that while you may use them, you may need to update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

| **Language**                                                                                   | **Package**                                                                                            |
| :---------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------- |
| ![Typescript Logo](https://docs.aws.amazon.com/cdk/api/latest/img/typescript32.png) TypeScript | `@cdklabs/generative-ai-cdk-constructs`                                                               |
| ![Python Logo](https://docs.aws.amazon.com/cdk/api/latest/img/python32.png) Python             | `cdklabs.generative_ai_cdk_constructs`                                                                |
| ![.Net](https://docs.aws.amazon.com/cdk/api/latest/img/dotnet32.png) .Net                     | `CdkLabs.GenerativeAICdkConstructs`                                                                    |
| ![Go](https://docs.aws.amazon.com/cdk/api/latest/img/go32.png) Go                             | `github.com/cdklabs/generative-ai-cdk-constructs-go/generative-ai-cdk-constructs`                      |

Amazon S3 Vectors delivers purpose-built, cost-optimized vector storage for your semantic search and AI applications. With Amazon S3 level elasticity and durability for storing vector datasets with sub-second query performance, S3 Vectors is ideal for applications that need to build and grow vector indexes. You get a dedicated set of API operations to store, access, and perform similarity queries on vector data without provisioning any infrastructure. S3 Vectors consists of several key components that work together:

* **Vector buckets** – A new bucket type that's purpose-built to store and query vectors.
* **Vector indexes** – Within a vector bucket, you can organize your vector data within vector indexes. You perform similarity queries on your vector data within vector indexes.
* **Vectors** – You store vectors in your vector index. For similarity search and AI applications, vectors are created as vector embeddings which are numerical representations that preserve semantic relationships between content (such as text, images, or audio) so similar items are positioned closer together. S3 Vectors can perform similarity searches based on semantic meaning rather than exact matching through comparing how close vectors are to each other mathematically. When adding vector data to a vector index, you can also attach metadata for future filtering queries based on a set of conditions (for example, timestamps, categories, and user preferences).

This construct library provides L2 constructs to manage S3 vectors resources.

## Table of contents

* [Vector Bucket](#vector-bucket)
* [Vector Index](#vector-index)
* [Vector Bucket Policy](#vector-bucket-policy)

## Vector bucket

Vector buckets are a type of Amazon S3 bucket designed specifically for storing and querying vector data. Vector buckets use dedicated APIs to manage vector data efficiently and reduce costs of upload, storing, and querying vector embeddings. Vector buckets provide the foundation for organizing your vector data into indexes, enabling you to perform similarity searches across large datasets while benefiting from the availability, durability, scalability, and cost-effectiveness of Amazon S3.

Vector buckets are optimized for long-term vector storage with sub-second search times. You can perform similarity queries on your vector data and optionally attach metadata to filter queries based on specific conditions such as dates, categories, or user preferences.

All data stored in vector buckets are always encrypted at rest. By default, vector buckets use SSE-S3 to encrypt vector data. You can choose to configure buckets to use server-side encryption with AWS Key Management Service (AWS KMS) keys (SSE-KMS) instead. The bucket encryption settings can't be changed after a vector bucket is created, so it's important to choose the appropriate encryption method based on your security requirements and compliance needs.

### Creating a vector bucket

Create a vector bucket with default settings:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
```

Create a vector bucket with a custom name:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"), &VectorBucketProps{
	VectorBucketName: jsii.String("my-custom-bucket-name"),
})
```

### Encryption

You can configure the encryption type for your vector bucket. By default, vector buckets use SSE-S3 (server-side encryption with Amazon S3 managed keys). You can choose to use SSE-KMS (server-side encryption with AWS KMS keys) for enhanced control and auditability.

#### SSE-S3 for buckets

SSE-S3 provides a simple and effective encryption solution where AWS manages all aspects of the encryption process:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"), &VectorBucketProps{
	Encryption: s3vectors.VectorBucketEncryption_S3_MANAGED,
})
```

#### SSE-KMS for buckets

SSE-KMS provides enhanced control over encryption keys and enables detailed audit logging of key usage. When you use KMS encryption, the construct automatically grants the S3 Vectors service principal (`indexing.s3vectors.amazonaws.com`) the necessary permissions to use the key for background operations.

**Using an auto-created KMS key:**

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"), &VectorBucketProps{
	Encryption: s3vectors.VectorBucketEncryption_KMS,
})
```

**Using your own KMS key:**

```go
myKmsKey := kms.NewKey(this, jsii.String("MyKey"), &KeyProps{
	Description: jsii.String("KMS key for S3 Vectors bucket"),
	EnableKeyRotation: jsii.Boolean(true),
})

vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"), &VectorBucketProps{
	Encryption: s3vectors.VectorBucketEncryption_KMS,
	EncryptionKey: myKmsKey,
})
```

**Important:** When using KMS encryption with customer-managed keys, the construct automatically grants the S3 Vectors service principal (`indexing.s3vectors.amazonaws.com`) the `kms:Decrypt` permission on your KMS key. This is required for the service to maintain and optimize indexes in background operations. The key policy includes appropriate conditions to ensure the service can only access keys for resources in your account.

For more information about KMS key requirements and service principal permissions, see [Data protection and encryption in S3 Vectors](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-data-encryption.html).

### Bucket permissions

#### Resource-based policies

A bucket policy will be automatically created for the bucket upon the first call to `addToResourcePolicy()`:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
result := vectorBucket.AddToResourcePolicy(
iam.NewPolicyStatement(&PolicyStatementProps{
	Actions: []*string{
		jsii.String("s3vectors:GetVectorBucket"),
		jsii.String("s3vectors:ListIndexes"),
	},
	Resources: []*string{
		vectorBucket.VectorBucketArn,
	},
	Principals: []IPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

The bucket policy can be directly accessed after creation to add statements or adjust the removal policy:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorBucket.Policy.ApplyRemovalPolicy(cdk.RemovalPolicy_RETAIN)
```

#### Grant methods

Most of the time, you won't have to manipulate the bucket policy directly. Instead, buckets have "grant" methods that give prepackaged sets of permissions to other resources.

**Grant read permissions:**

```go
var myLambda Function


vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
// Grant read access to all indexes
vectorBucket.GrantRead(myLambda)

// Grant read access to specific indexes
vectorBucket.GrantRead(myLambda, []interface{}{
	jsii.String("index-1"),
	jsii.String("index-2"),
})
```

**Grant write permissions:**

```go
var myLambda Function

vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))

// Grant write access to all indexes
vectorBucket.GrantWrite(myLambda)

// Grant write access to specific indexes
vectorBucket.GrantWrite(myLambda, []interface{}{
	jsii.String("index-1"),
	jsii.String("index-2"),
})
```

**Grant delete permissions:**

```go
var myLambda Function

vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))

// Grant delete access to all indexes
vectorBucket.GrantDelete(myLambda)

// Grant delete access to specific indexes
vectorBucket.GrantDelete(myLambda, []interface{}{
	jsii.String("index-1"),
	jsii.String("index-2"),
})
```

When using KMS encryption, the grant methods automatically include the necessary KMS permissions (`kms:Decrypt` and `kms:DescribeKey`) for the granted principal.

### Bucket deletion

When a bucket is removed from a stack (or the stack is deleted), the S3 vector bucket will be removed according to its removal policy (which by default will simply orphan the bucket and leave it in your AWS account). If the removal policy is set to `RemovalPolicy.DESTROY`, the bucket will be deleted as long as it does not contain any indexes.

To override this and force all indexes to get deleted during bucket deletion, enable the `autoDeleteObjects` option:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"), &VectorBucketProps{
	AutoDeleteObjects: jsii.Boolean(true),
	RemovalPolicy: cdk.RemovalPolicy_DESTROY,
})
```

**Note:** The `autoDeleteObjects` feature uses a custom resource to delete all indexes in the bucket before the bucket is deleted. This ensures that the bucket can be successfully removed even if it contains indexes.

### Importing existing vector buckets

To import an existing vector bucket into your CDK application, use one of the static factory methods:

**Import by ARN:**

```go
importedBucket := s3vectors.VectorBucket_FromVectorBucketArn(this, jsii.String("ImportedBucket"), jsii.String("arn:aws:s3vectors:us-east-1:123456789012:bucket/my-bucket-name"))
```

**Import by name:**

```go
importedBucket := s3vectors.VectorBucket_FromVectorBucketName(this, jsii.String("ImportedBucket"), jsii.String("my-bucket-name"))
```

**Import with attributes:**

```go
importedBucket := s3vectors.VectorBucket_FromVectorBucketAttributes(this, jsii.String("ImportedBucket"), &VectorBucketAttributes{
	VectorBucketArn: jsii.String("arn:aws:s3vectors:us-east-1:123456789012:bucket/my-bucket-name"),
	CreationTime: jsii.String("2024-01-01T00:00:00Z"),
})
```

### Sharing buckets between stacks

To use a bucket in a different stack in the same CDK application, pass the object to the other stack:

```go
// In Stack B
var stackB Stack
var someOtherConstruct interface{}
// In Stack A
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
// Pass the bucket reference
// Pass the bucket reference
NewSomeOtherConstruct(stackB, jsii.String("OtherConstruct"), map[string]VectorBucket{
	"bucket": vectorBucket,
})
```

## Vector index

Vector indexes are resources within vector buckets that store and organize vector data for efficient similarity search operations. When you create a vector index, you specify the distance metric (Cosine or Euclidean), the number of dimensions that a vector should have, and optionally a list of metadata fields that you want to exclude from filtering during similarity queries.

For more information about vector index limits per bucket, vector limits per index, and dimension limits per vector, see [Limitations and restrictions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-limitations.html).

### Creating a vector index

Create a vector index with minimal required properties:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
})
```

### Dimension

Dimension is a numeric value between 1 and 4096 that determines how many numbers will be in each vector that's generated by your vector embedding model. Embedding models are specialized machine learning (ML) models that convert data (such as text or images) into numerical vectors. Embedding models typically produce outputs between 500-2000 dimensions, with each dimension being a floating-point number.

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(1024),
})
```

### Distance metric

You can configure the distance metric to define how similarity between vectors is calculated during queries:

* **Cosine**: Measures the cosine of the angle between two vectors. This is the default.
* **Euclidean**: Measures the straight-line distance between two points in multi-dimensional space. Lower values indicate greater similarity.

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))

// Using cosine similarity (default)
cosineIndex := s3vectors.NewVectorIndex(this, jsii.String("CosineIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	DistanceMetric: s3vectors.VectorIndexDistanceMetric_COSINE,
})

// Using euclidean distance
euclideanIndex := s3vectors.NewVectorIndex(this, jsii.String("EuclideanIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	DistanceMetric: s3vectors.VectorIndexDistanceMetric_EUCLIDEAN,
})
```

### Data type

Currently, S3 Vectors supports 32-bit floating-point numbers (`float32`) for vector data:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	DataType: s3vectors.VectorIndexDataType_FLOAT_32,
})
```

### Non-filterable metadata keys

Metadata keys allow you to attach additional information to your vectors as key-value pairs during storage and retrieval. By default, all metadata is filterable, so you can use it to filter query results. However, you can designate specific metadata keys as non-filterable when you want to store information with vectors without using it for filtering.

Unlike default metadata keys, these keys can't be used as query filters. Non-filterable metadata keys can be retrieved but can't be searched, queried, or filtered. You can only access them after finding the vectors.

Non-filterable metadata keys allow you to enrich vectors with additional context that you want to retrieve with search results but don't need for filtering. A common example of a non-filterable metadata key is when you embed text into vectors and want to include the original text itself as non-filterable metadata. This allows you to return the source text alongside vector search results without increasing your filterable metadata size limits.

**Specify non-filterable metadata keys:**

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	NonFilterableMetadataKeys: []*string{
		jsii.String("originalText"),
		jsii.String("sourceUrl"),
		jsii.String("timestamp"),
	},
})
```

**Requirements:**

* You can specify 1 to 10 non-filterable metadata keys per index
* Each key must be 1 to 63 characters long
* Keys must follow the same naming rules as index names (lowercase letters, numbers, dots, and hyphens)

### Index encryption

You can configure the encryption type for your vector index. By default, if you don't specify encryption, the index will use SSE-S3 (server-side encryption with Amazon S3 managed keys). You can optionally override the bucket-level encryption settings and provide a dedicated encryption configuration at the index level.

**Important:** Encryption settings for a vector index can't be changed after the index is created.

#### SSE-S3 for indexes

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	Encryption: s3vectors.VectorIndexEncryption_S3_MANAGED,
})
```

#### SSE-KMS for indexes

When you use KMS encryption at the index level, the construct automatically grants the S3 Vectors service principal (`indexing.s3vectors.amazonaws.com`) the necessary permissions to use the key for background operations.

**Using an auto-created KMS key:**

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	Encryption: s3vectors.VectorIndexEncryption_KMS,
})
```

**Using your own KMS key:**

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
myKmsKey := kms.NewKey(this, jsii.String("MyIndexKey"), &KeyProps{
	Description: jsii.String("KMS key for S3 Vectors index"),
	EnableKeyRotation: jsii.Boolean(true),
})

vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
	Encryption: s3vectors.VectorIndexEncryption_KMS,
	EncryptionKey: myKmsKey,
})
```

**Important:** When using KMS encryption with customer-managed keys, the construct automatically grants the S3 Vectors service principal (`indexing.s3vectors.amazonaws.com`) the `kms:Decrypt` permission on your KMS key. This is required for the service to maintain and optimize indexes in background operations. The key policy includes appropriate conditions to ensure the service can only access keys for resources in your account.

### Index permissions

#### Granting permissions to indexes

Vector indexes provide a `grant()` method to grant IAM permissions to principals:

```go
var myLambda Function


vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
vectorIndex := s3vectors.NewVectorIndex(this, jsii.String("MyVectorIndex"), &VectorIndexProps{
	VectorBucket: vectorBucket,
	Dimension: jsii.Number(128),
})

// Grant specific actions
vectorIndex.Grant(myLambda, jsii.String("s3vectors:GetIndex"), jsii.String("s3vectors:QueryVectors"), jsii.String("s3vectors:PutVectors"))
```

When using KMS encryption, the grant method automatically includes the necessary KMS permissions for the granted principal.

### Importing existing vector indexes

To import an existing vector index into your CDK application, use one of the static factory methods:

**Import by ARN:**

```go
importedIndex := s3vectors.VectorIndex_FromVectorIndexArn(this, jsii.String("ImportedIndex"), jsii.String("arn:aws:s3vectors:us-east-1:123456789012:bucket/my-bucket/index/my-index"))
```

**Import by name:**

```go
importedIndex := s3vectors.VectorIndex_FromVectorIndexName(this, jsii.String("ImportedIndex"), jsii.String("my-bucket-name"), jsii.String("my-index-name"))
```

**Import with attributes:**

```go
importedIndex := s3vectors.VectorIndex_FromVectorIndexAttributes(this, jsii.String("ImportedIndex"), &VectorIndexAttributes{
	VectorIndexArn: jsii.String("arn:aws:s3vectors:us-east-1:123456789012:bucket/my-bucket/index/my-index"),
	CreationTime: jsii.String("2024-01-01T00:00:00Z"),
})
```

### Index name validation

Vector index names must follow specific rules:

* Must be 3 to 63 characters long
* Can only contain lowercase letters, numbers, dots (.), and hyphens (-)
* Must begin and end with a letter or number

If you don't specify a name, CloudFormation will automatically generate one for you.

## Vector bucket policy

Vector bucket policies are resource-based policies that you attach directly to vector buckets to control access to the bucket and its contents. Bucket policies for vector buckets can grant permissions to principals from other AWS accounts, making them useful for cross-account access scenarios.

### Basic usage

Create a vector bucket policy:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
policy := s3vectors.NewVectorBucketPolicy(this, jsii.String("MyBucketPolicy"), &VectorBucketPolicyProps{
	Bucket: vectorBucket,
})

// Add statements to the policy
policy.Document.AddStatements(
iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Actions: []*string{
		jsii.String("s3vectors:GetVectorBucket"),
		jsii.String("s3vectors:ListIndexes"),
	},
	Resources: []*string{
		vectorBucket.VectorBucketArn,
	},
	Principals: []IPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

### Providing a policy document

You can provide a complete policy document when creating the policy:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
policyDoc := iam.NewPolicyDocument(&PolicyDocumentProps{
	Statements: []PolicyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Effect: iam.Effect_ALLOW,
			Actions: []*string{
				jsii.String("s3vectors:GetVectorBucket"),
			},
			Resources: []*string{
				vectorBucket.VectorBucketArn,
			},
			Principals: []IPrincipal{
				iam.NewAnyPrincipal(),
			},
		}),
	},
})

policy := s3vectors.NewVectorBucketPolicy(this, jsii.String("MyBucketPolicy"), &VectorBucketPolicyProps{
	Bucket: vectorBucket,
	Document: policyDoc,
})
```

### Removal policy

You can specify a removal policy for the bucket policy:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
policyDoc := iam.NewPolicyDocument(&PolicyDocumentProps{
	Statements: []PolicyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Effect: iam.Effect_ALLOW,
			Actions: []*string{
				jsii.String("s3vectors:GetVectorBucket"),
			},
			Resources: []*string{
				vectorBucket.VectorBucketArn,
			},
			Principals: []IPrincipal{
				iam.NewAnyPrincipal(),
			},
		}),
	},
})

policy := s3vectors.NewVectorBucketPolicy(this, jsii.String("MyBucketPolicy"), &VectorBucketPolicyProps{
	Bucket: vectorBucket,
	Document: policyDoc,
	RemovalPolicy: cdk.RemovalPolicy_RETAIN,
})
```

Or apply it after creation:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
policyDoc := iam.NewPolicyDocument(&PolicyDocumentProps{
	Statements: []PolicyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Effect: iam.Effect_ALLOW,
			Actions: []*string{
				jsii.String("s3vectors:GetVectorBucket"),
			},
			Resources: []*string{
				vectorBucket.VectorBucketArn,
			},
			Principals: []IPrincipal{
				iam.NewAnyPrincipal(),
			},
		}),
	},
})

policy := s3vectors.NewVectorBucketPolicy(this, jsii.String("MyBucketPolicy"), &VectorBucketPolicyProps{
	Bucket: vectorBucket,
	Document: policyDoc,
})
policy.ApplyRemovalPolicy(cdk.RemovalPolicy_RETAIN)
```

### Policy resources

Vector bucket policies can reference both bucket and index resources:

```go
vectorBucket := s3vectors.NewVectorBucket(this, jsii.String("MyVectorBucket"))
policy := s3vectors.NewVectorBucketPolicy(this, jsii.String("MyBucketPolicy"), &VectorBucketPolicyProps{
	Bucket: vectorBucket,
})

// Grant access to the bucket
policy.Document.AddStatements(
iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Actions: []*string{
		jsii.String("s3vectors:GetVectorBucket"),
	},
	Resources: []*string{
		vectorBucket.VectorBucketArn,
	},
	Principals: []IPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))

// Grant access to all indexes in the bucket
policy.Document.AddStatements(
iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: iam.Effect_ALLOW,
	Actions: []*string{
		jsii.String("s3vectors:GetIndex"),
		jsii.String("s3vectors:QueryVectors"),
	},
	Resources: []*string{
		fmt.Sprintf("%v/index/*", vectorBucket.*VectorBucketArn),
	},
	Principals: []IPrincipal{
		iam.NewAccountRootPrincipal(),
	},
}))
```

## Additional resources

* [Amazon S3 Vectors User Guide](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors.html)
* [Data protection and encryption in S3 Vectors](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-data-encryption.html)
* [S3 Vectors access management](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-access-management.html)
* [S3 Vectors limitations and restrictions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-limitations.html)
