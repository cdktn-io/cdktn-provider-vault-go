// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncghdestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncGhDestinationConfig struct {
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
	// Unique name of the github destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#name SecretsSyncGhDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Fine-grained or personal access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#access_token SecretsSyncGhDestination#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// List of allowed IPv4 addresses in CIDR notation (e.g., 192.168.1.1/32) for outbound connections from Vault to the destination. If not set, all IPv4 addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#allowed_ipv4_addresses SecretsSyncGhDestination#allowed_ipv4_addresses}
	AllowedIpv4Addresses *[]*string `field:"optional" json:"allowedIpv4Addresses" yaml:"allowedIpv4Addresses"`
	// List of allowed IPv6 addresses in CIDR notation (e.g., 2001:db8::1/128) for outbound connections from Vault to the destination. If not set, all IPv6 addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#allowed_ipv6_addresses SecretsSyncGhDestination#allowed_ipv6_addresses}
	AllowedIpv6Addresses *[]*string `field:"optional" json:"allowedIpv6Addresses" yaml:"allowedIpv6Addresses"`
	// List of allowed ports for outbound connections from Vault to the destination. If not set, all ports are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#allowed_ports SecretsSyncGhDestination#allowed_ports}
	AllowedPorts *[]*float64 `field:"optional" json:"allowedPorts" yaml:"allowedPorts"`
	// The user-defined name of the GitHub App configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#app_name SecretsSyncGhDestination#app_name}
	AppName *string `field:"optional" json:"appName" yaml:"appName"`
	// If set to true, disables strict networking enforcement for this destination.
	//
	// When disabled, Vault will not enforce allowed IP addresses and ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#disable_strict_networking SecretsSyncGhDestination#disable_strict_networking}
	DisableStrictNetworking interface{} `field:"optional" json:"disableStrictNetworking" yaml:"disableStrictNetworking"`
	// GitHub environment name where secrets will be synced. Required when secrets_location is set to 'environment'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#environment_name SecretsSyncGhDestination#environment_name}
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// Determines what level of information is synced as a distinct resource at the destination. Can be 'secret-path' or 'secret-key'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#granularity SecretsSyncGhDestination#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#id SecretsSyncGhDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The ID of the installation generated by GitHub when the app referenced by the app_name was installed in the user's GitHub account.
	//
	// Necessary if the app_name field is also provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#installation_id SecretsSyncGhDestination#installation_id}
	InstallationId *float64 `field:"optional" json:"installationId" yaml:"installationId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#namespace SecretsSyncGhDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#repository_name SecretsSyncGhDestination#repository_name}
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
	// GitHub organization or username that owns the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#repository_owner SecretsSyncGhDestination#repository_owner}
	RepositoryOwner *string `field:"optional" json:"repositoryOwner" yaml:"repositoryOwner"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#secret_name_template SecretsSyncGhDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
	// Determines where secrets will be stored in GitHub. Valid values are 'repository' or 'environment'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/secrets_sync_gh_destination#secrets_location SecretsSyncGhDestination#secrets_location}
	SecretsLocation *string `field:"optional" json:"secretsLocation" yaml:"secretsLocation"`
}

