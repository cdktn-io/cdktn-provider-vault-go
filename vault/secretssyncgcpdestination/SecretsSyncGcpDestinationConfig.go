// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncgcpdestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncGcpDestinationConfig struct {
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
	// Unique name of the GCP destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#name SecretsSyncGcpDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allowed IPv4 addresses for outbound network connectivity in CIDR notation. If not set, all IPv4 addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#allowed_ipv4_addresses SecretsSyncGcpDestination#allowed_ipv4_addresses}
	AllowedIpv4Addresses *[]*string `field:"optional" json:"allowedIpv4Addresses" yaml:"allowedIpv4Addresses"`
	// Allowed IPv6 addresses for outbound network connectivity in CIDR notation. If not set, all IPv6 addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#allowed_ipv6_addresses SecretsSyncGcpDestination#allowed_ipv6_addresses}
	AllowedIpv6Addresses *[]*string `field:"optional" json:"allowedIpv6Addresses" yaml:"allowedIpv6Addresses"`
	// Allowed ports for outbound network connectivity. If not set, all ports are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#allowed_ports SecretsSyncGcpDestination#allowed_ports}
	AllowedPorts *[]*float64 `field:"optional" json:"allowedPorts" yaml:"allowedPorts"`
	// JSON-encoded credentials to use to connect to GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#credentials SecretsSyncGcpDestination#credentials}
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Custom tags to set on the secret managed at the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#custom_tags SecretsSyncGcpDestination#custom_tags}
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Disable strict networking requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#disable_strict_networking SecretsSyncGcpDestination#disable_strict_networking}
	DisableStrictNetworking interface{} `field:"optional" json:"disableStrictNetworking" yaml:"disableStrictNetworking"`
	// Global KMS key for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#global_kms_key SecretsSyncGcpDestination#global_kms_key}
	GlobalKmsKey *string `field:"optional" json:"globalKmsKey" yaml:"globalKmsKey"`
	// Determines what level of information is synced as a distinct resource at the destination. Can be 'secret-path' or 'secret-key'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#granularity SecretsSyncGcpDestination#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#id SecretsSyncGcpDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value for identity tokens.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#identity_token_audience_wo SecretsSyncGcpDestination#identity_token_audience_wo}
	IdentityTokenAudienceWo *string `field:"optional" json:"identityTokenAudienceWo" yaml:"identityTokenAudienceWo"`
	// A version counter for the write-only identity_token_audience_wo field. Incrementing this value will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#identity_token_audience_wo_version SecretsSyncGcpDestination#identity_token_audience_wo_version}
	IdentityTokenAudienceWoVersion *float64 `field:"optional" json:"identityTokenAudienceWoVersion" yaml:"identityTokenAudienceWoVersion"`
	// The key to use for signing identity tokens.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#identity_token_key_wo SecretsSyncGcpDestination#identity_token_key_wo}
	IdentityTokenKeyWo *string `field:"optional" json:"identityTokenKeyWo" yaml:"identityTokenKeyWo"`
	// A version counter for the write-only identity_token_key_wo field. Incrementing this value will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#identity_token_key_wo_version SecretsSyncGcpDestination#identity_token_key_wo_version}
	IdentityTokenKeyWoVersion *float64 `field:"optional" json:"identityTokenKeyWoVersion" yaml:"identityTokenKeyWoVersion"`
	// The TTL of generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#identity_token_ttl SecretsSyncGcpDestination#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Locational KMS keys for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#locational_kms_keys SecretsSyncGcpDestination#locational_kms_keys}
	LocationalKmsKeys *map[string]*string `field:"optional" json:"locationalKmsKeys" yaml:"locationalKmsKeys"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#namespace SecretsSyncGcpDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The target project to manage secrets in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#project_id SecretsSyncGcpDestination#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Replication locations for secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#replication_locations SecretsSyncGcpDestination#replication_locations}
	ReplicationLocations *[]*string `field:"optional" json:"replicationLocations" yaml:"replicationLocations"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#secret_name_template SecretsSyncGcpDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
	// Service Account to impersonate for workload identity federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/secrets_sync_gcp_destination#service_account_email SecretsSyncGcpDestination#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

