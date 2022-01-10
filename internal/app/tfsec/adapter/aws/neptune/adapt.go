package neptune

import (
	"github.com/aquasecurity/defsec/provider/aws/neptune"
	"github.com/aquasecurity/defsec/types"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
)

func Adapt(modules []block.Module) neptune.Neptune {
	return neptune.Neptune{
		Clusters: adaptClusters(modules),
	}
}

func adaptClusters(modules []block.Module) []neptune.Cluster {
	var clusters []neptune.Cluster
	for _, module := range modules {
		for _, resource := range module.GetResourcesByType("aws_neptune_cluster") {
			clusters = append(clusters, adaptCluster(resource))
		}
	}
	return clusters
}

func adaptCluster(resource block.Block) neptune.Cluster {
	enableLogExportsAttr := resource.GetAttribute("enable_cloudwatch_logs_exports")
	auditVal := types.Bool(false, *resource.GetMetadata())
	if enableLogExportsAttr.Contains("audit") {
		auditVal = types.Bool(true, *resource.GetMetadata())
	}

	storageEncryptedAttr := resource.GetAttribute("storage_encrypted")
	storageEncryptedVal := storageEncryptedAttr.AsBoolValueOrDefault(false, resource)

	KMSKeyAttr := resource.GetAttribute("kms_key_arn")
	KMSKeyVal := KMSKeyAttr.AsStringValueOrDefault("", resource)

	return neptune.Cluster{
		Metadata: *resource.GetMetadata(),
		Logging: neptune.Logging{
			Audit: auditVal,
		},
		StorageEncrypted: storageEncryptedVal,
		KMSKeyID:         KMSKeyVal,
	}
}
