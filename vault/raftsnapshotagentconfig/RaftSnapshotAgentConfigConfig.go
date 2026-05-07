// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package raftsnapshotagentconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RaftSnapshotAgentConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Number of seconds between snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#interval_seconds RaftSnapshotAgentConfig#interval_seconds}
	IntervalSeconds *float64 `field:"required" json:"intervalSeconds" yaml:"intervalSeconds"`
	// Name of the snapshot agent configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#name RaftSnapshotAgentConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The directory or bucket prefix to to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#path_prefix RaftSnapshotAgentConfig#path_prefix}
	PathPrefix *string `field:"required" json:"pathPrefix" yaml:"pathPrefix"`
	// What storage service to send snapshots to. One of "local", "azure-blob", "aws-s3", or "google-gcs".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#storage_type RaftSnapshotAgentConfig#storage_type}
	StorageType *string `field:"required" json:"storageType" yaml:"storageType"`
	// Have Vault automatically load the latest snapshot after it is written.
	//
	// This will replace the previously loaded snapshot. Note that this does not mean the snapshot is automatically applied to the cluster, it is just loaded and available for recovery operations. Requires Vault Enterprise 1.21.0+. Not supported with storage_type = "local".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#autoload_enabled RaftSnapshotAgentConfig#autoload_enabled}
	AutoloadEnabled interface{} `field:"optional" json:"autoloadEnabled" yaml:"autoloadEnabled"`
	// AWS access key ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_access_key_id RaftSnapshotAgentConfig#aws_access_key_id}
	AwsAccessKeyId *string `field:"optional" json:"awsAccessKeyId" yaml:"awsAccessKeyId"`
	// S3 bucket to write snapshots to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_bucket RaftSnapshotAgentConfig#aws_s3_bucket}
	AwsS3Bucket *string `field:"optional" json:"awsS3Bucket" yaml:"awsS3Bucket"`
	// Disable TLS for the S3 endpoint. This should only be used for testing purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_disable_tls RaftSnapshotAgentConfig#aws_s3_disable_tls}
	AwsS3DisableTls interface{} `field:"optional" json:"awsS3DisableTls" yaml:"awsS3DisableTls"`
	// Use KMS to encrypt bucket contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_enable_kms RaftSnapshotAgentConfig#aws_s3_enable_kms}
	AwsS3EnableKms interface{} `field:"optional" json:"awsS3EnableKms" yaml:"awsS3EnableKms"`
	// AWS endpoint. This is typically only set when using a non-AWS S3 implementation like Minio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_endpoint RaftSnapshotAgentConfig#aws_s3_endpoint}
	AwsS3Endpoint *string `field:"optional" json:"awsS3Endpoint" yaml:"awsS3Endpoint"`
	// Use the endpoint/bucket URL style instead of bucket.endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_force_path_style RaftSnapshotAgentConfig#aws_s3_force_path_style}
	AwsS3ForcePathStyle interface{} `field:"optional" json:"awsS3ForcePathStyle" yaml:"awsS3ForcePathStyle"`
	// Use named KMS key, when aws_s3_enable_kms=true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_kms_key RaftSnapshotAgentConfig#aws_s3_kms_key}
	AwsS3KmsKey *string `field:"optional" json:"awsS3KmsKey" yaml:"awsS3KmsKey"`
	// AWS region bucket is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_region RaftSnapshotAgentConfig#aws_s3_region}
	AwsS3Region *string `field:"optional" json:"awsS3Region" yaml:"awsS3Region"`
	// Use AES256 to encrypt bucket contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_s3_server_side_encryption RaftSnapshotAgentConfig#aws_s3_server_side_encryption}
	AwsS3ServerSideEncryption interface{} `field:"optional" json:"awsS3ServerSideEncryption" yaml:"awsS3ServerSideEncryption"`
	// AWS secret access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_secret_access_key RaftSnapshotAgentConfig#aws_secret_access_key}
	AwsSecretAccessKey *string `field:"optional" json:"awsSecretAccessKey" yaml:"awsSecretAccessKey"`
	// AWS session token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#aws_session_token RaftSnapshotAgentConfig#aws_session_token}
	AwsSessionToken *string `field:"optional" json:"awsSessionToken" yaml:"awsSessionToken"`
	// Azure account key. Required when azure_auth_mode is 'shared'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_account_key RaftSnapshotAgentConfig#azure_account_key}
	AzureAccountKey *string `field:"optional" json:"azureAccountKey" yaml:"azureAccountKey"`
	// Azure account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_account_name RaftSnapshotAgentConfig#azure_account_name}
	AzureAccountName *string `field:"optional" json:"azureAccountName" yaml:"azureAccountName"`
	// Azure authentication mode. Required for azure-blob storage. Possible values are 'shared', 'managed', or 'environment'. Requires Vault Enterprise 1.18.0+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_auth_mode RaftSnapshotAgentConfig#azure_auth_mode}
	AzureAuthMode *string `field:"optional" json:"azureAuthMode" yaml:"azureAuthMode"`
	// Azure blob environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_blob_environment RaftSnapshotAgentConfig#azure_blob_environment}
	AzureBlobEnvironment *string `field:"optional" json:"azureBlobEnvironment" yaml:"azureBlobEnvironment"`
	// Azure client ID for authentication. Required when azure_auth_mode is 'managed'. Requires Vault Enterprise 1.18.0+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_client_id RaftSnapshotAgentConfig#azure_client_id}
	AzureClientId *string `field:"optional" json:"azureClientId" yaml:"azureClientId"`
	// Azure container name to write snapshots to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_container_name RaftSnapshotAgentConfig#azure_container_name}
	AzureContainerName *string `field:"optional" json:"azureContainerName" yaml:"azureContainerName"`
	// Azure blob storage endpoint. This is typically only set when using a non-Azure implementation like Azurite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#azure_endpoint RaftSnapshotAgentConfig#azure_endpoint}
	AzureEndpoint *string `field:"optional" json:"azureEndpoint" yaml:"azureEndpoint"`
	// The file or object name of snapshot files will start with this string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#file_prefix RaftSnapshotAgentConfig#file_prefix}
	FilePrefix *string `field:"optional" json:"filePrefix" yaml:"filePrefix"`
	// Disable TLS for the GCS endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#google_disable_tls RaftSnapshotAgentConfig#google_disable_tls}
	GoogleDisableTls interface{} `field:"optional" json:"googleDisableTls" yaml:"googleDisableTls"`
	// GCS endpoint. This is typically only set when using a non-Google GCS implementation like fake-gcs-server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#google_endpoint RaftSnapshotAgentConfig#google_endpoint}
	GoogleEndpoint *string `field:"optional" json:"googleEndpoint" yaml:"googleEndpoint"`
	// GCS bucket to write snapshots to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#google_gcs_bucket RaftSnapshotAgentConfig#google_gcs_bucket}
	GoogleGcsBucket *string `field:"optional" json:"googleGcsBucket" yaml:"googleGcsBucket"`
	// Google service account key in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#google_service_account_key RaftSnapshotAgentConfig#google_service_account_key}
	GoogleServiceAccountKey *string `field:"optional" json:"googleServiceAccountKey" yaml:"googleServiceAccountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#id RaftSnapshotAgentConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The maximum space, in bytes, to use for snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#local_max_space RaftSnapshotAgentConfig#local_max_space}
	LocalMaxSpace *float64 `field:"optional" json:"localMaxSpace" yaml:"localMaxSpace"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#namespace RaftSnapshotAgentConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// How many snapshots are to be kept.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/raft_snapshot_agent_config#retain RaftSnapshotAgentConfig#retain}
	Retain *float64 `field:"optional" json:"retain" yaml:"retain"`
}

