package s3vectors

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.IVectorBucket",
		reflect.TypeOf((*IVectorBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "creationTime", GoGetter: "CreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketArn", GoGetter: "VectorBucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketName", GoGetter: "VectorBucketName"},
		},
		func() interface{} {
			j := jsiiProxy_IVectorBucket{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.IVectorIndex",
		reflect.TypeOf((*IVectorIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "creationTime", GoGetter: "CreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexArn", GoGetter: "VectorIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexName", GoGetter: "VectorIndexName"},
		},
		func() interface{} {
			j := jsiiProxy_IVectorIndex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucket",
		reflect.TypeOf((*VectorBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "creationTime", GoGetter: "CreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketArn", GoGetter: "VectorBucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketName", GoGetter: "VectorBucketName"},
		},
		func() interface{} {
			j := jsiiProxy_VectorBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VectorBucketBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucketAttributes",
		reflect.TypeOf((*VectorBucketAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucketBase",
		reflect.TypeOf((*VectorBucketBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "creationTime", GoGetter: "CreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketArn", GoGetter: "VectorBucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketName", GoGetter: "VectorBucketName"},
		},
		func() interface{} {
			j := jsiiProxy_VectorBucketBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVectorBucket)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucketEncryption",
		reflect.TypeOf((*VectorBucketEncryption)(nil)).Elem(),
		map[string]interface{}{
			"S3_MANAGED": VectorBucketEncryption_S3_MANAGED,
			"KMS": VectorBucketEncryption_KMS,
		},
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucketPolicy",
		reflect.TypeOf((*VectorBucketPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "document", GoGetter: "Document"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucketPolicyRef", GoGetter: "VectorBucketPolicyRef"},
		},
		func() interface{} {
			j := jsiiProxy_VectorBucketPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawss3vectorsIVectorBucketPolicyRef)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucketPolicyProps",
		reflect.TypeOf((*VectorBucketPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorBucketProps",
		reflect.TypeOf((*VectorBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndex",
		reflect.TypeOf((*VectorIndex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "creationTime", GoGetter: "CreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dimension", GoGetter: "Dimension"},
			_jsii_.MemberProperty{JsiiProperty: "distanceMetric", GoGetter: "DistanceMetric"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorBucket", GoGetter: "VectorBucket"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexArn", GoGetter: "VectorIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexName", GoGetter: "VectorIndexName"},
		},
		func() interface{} {
			j := jsiiProxy_VectorIndex{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_VectorIndexBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndexAttributes",
		reflect.TypeOf((*VectorIndexAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndexBase",
		reflect.TypeOf((*VectorIndexBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "creationTime", GoGetter: "CreationTime"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexArn", GoGetter: "VectorIndexArn"},
			_jsii_.MemberProperty{JsiiProperty: "vectorIndexName", GoGetter: "VectorIndexName"},
		},
		func() interface{} {
			j := jsiiProxy_VectorIndexBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IVectorIndex)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndexDataType",
		reflect.TypeOf((*VectorIndexDataType)(nil)).Elem(),
		map[string]interface{}{
			"FLOAT_32": VectorIndexDataType_FLOAT_32,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndexDistanceMetric",
		reflect.TypeOf((*VectorIndexDistanceMetric)(nil)).Elem(),
		map[string]interface{}{
			"EUCLIDEAN": VectorIndexDistanceMetric_EUCLIDEAN,
			"COSINE": VectorIndexDistanceMetric_COSINE,
		},
	)
	_jsii_.RegisterEnum(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndexEncryption",
		reflect.TypeOf((*VectorIndexEncryption)(nil)).Elem(),
		map[string]interface{}{
			"S3_MANAGED": VectorIndexEncryption_S3_MANAGED,
			"KMS": VectorIndexEncryption_KMS,
		},
	)
	_jsii_.RegisterStruct(
		"@cdklabs/generative-ai-cdk-constructs.s3vectors.VectorIndexProps",
		reflect.TypeOf((*VectorIndexProps)(nil)).Elem(),
	)
}
