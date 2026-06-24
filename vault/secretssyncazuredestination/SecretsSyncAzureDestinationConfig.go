// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncazuredestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncAzureDestinationConfig struct {
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
	// Unique name of the Azure destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#name SecretsSyncAzureDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Set of allowed IPv4 addresses in CIDR notation (e.g., 192.168.1.1/32) for outbound connections from Vault to the destination. If not set, all IPv4 addresses are allowed. Requires Vault 1.19+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#allowed_ipv4_addresses SecretsSyncAzureDestination#allowed_ipv4_addresses}
	AllowedIpv4Addresses *[]*string `field:"optional" json:"allowedIpv4Addresses" yaml:"allowedIpv4Addresses"`
	// Set of allowed IPv6 addresses in CIDR notation (e.g., 2001:db8::1/128) for outbound connections from Vault to the destination. If not set, all IPv6 addresses are allowed. Requires Vault 1.19+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#allowed_ipv6_addresses SecretsSyncAzureDestination#allowed_ipv6_addresses}
	AllowedIpv6Addresses *[]*string `field:"optional" json:"allowedIpv6Addresses" yaml:"allowedIpv6Addresses"`
	// Set of allowed ports for outbound connections from Vault to the destination.
	//
	// If not set, all ports are allowed. Requires Vault 1.19+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#allowed_ports SecretsSyncAzureDestination#allowed_ports}
	AllowedPorts *[]*float64 `field:"optional" json:"allowedPorts" yaml:"allowedPorts"`
	// Client ID of an Azure app registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#client_id SecretsSyncAzureDestination#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client Secret of an Azure app registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#client_secret SecretsSyncAzureDestination#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Specifies a cloud for the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#cloud SecretsSyncAzureDestination#cloud}
	Cloud *string `field:"optional" json:"cloud" yaml:"cloud"`
	// Custom tags to set on the secret managed at the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#custom_tags SecretsSyncAzureDestination#custom_tags}
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// If set to true, disables strict networking enforcement for this destination.
	//
	// When disabled, Vault will not enforce allowed IP addresses and ports. Requires Vault 1.19+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#disable_strict_networking SecretsSyncAzureDestination#disable_strict_networking}
	DisableStrictNetworking interface{} `field:"optional" json:"disableStrictNetworking" yaml:"disableStrictNetworking"`
	// Determines what level of information is synced as a distinct resource at the destination. Can be 'secret-path' or 'secret-key'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#granularity SecretsSyncAzureDestination#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#id SecretsSyncAzureDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value for identity tokens.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#identity_token_audience_wo SecretsSyncAzureDestination#identity_token_audience_wo}
	IdentityTokenAudienceWo *string `field:"optional" json:"identityTokenAudienceWo" yaml:"identityTokenAudienceWo"`
	// A version counter for the write-only identity_token_audience_wo field. Incrementing this value will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#identity_token_audience_wo_version SecretsSyncAzureDestination#identity_token_audience_wo_version}
	IdentityTokenAudienceWoVersion *float64 `field:"optional" json:"identityTokenAudienceWoVersion" yaml:"identityTokenAudienceWoVersion"`
	// The key to use for signing identity tokens.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#identity_token_key_wo SecretsSyncAzureDestination#identity_token_key_wo}
	IdentityTokenKeyWo *string `field:"optional" json:"identityTokenKeyWo" yaml:"identityTokenKeyWo"`
	// A version counter for the write-only identity_token_key_wo field. Incrementing this value will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#identity_token_key_wo_version SecretsSyncAzureDestination#identity_token_key_wo_version}
	IdentityTokenKeyWoVersion *float64 `field:"optional" json:"identityTokenKeyWoVersion" yaml:"identityTokenKeyWoVersion"`
	// The TTL of generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#identity_token_ttl SecretsSyncAzureDestination#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// URI of an existing Azure Key Vault instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#key_vault_uri SecretsSyncAzureDestination#key_vault_uri}
	KeyVaultUri *string `field:"optional" json:"keyVaultUri" yaml:"keyVaultUri"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#namespace SecretsSyncAzureDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#secret_name_template SecretsSyncAzureDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
	// ID of the target Azure tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/secrets_sync_azure_destination#tenant_id SecretsSyncAzureDestination#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

