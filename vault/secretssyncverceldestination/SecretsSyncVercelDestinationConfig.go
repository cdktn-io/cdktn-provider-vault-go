// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncverceldestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncVercelDestinationConfig struct {
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
	// Vercel API access token with the permissions to manage environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#access_token SecretsSyncVercelDestination#access_token}
	AccessToken *string `field:"required" json:"accessToken" yaml:"accessToken"`
	// Deployment environments where the environment variables are available. Accepts 'development', 'preview' & 'production'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#deployment_environments SecretsSyncVercelDestination#deployment_environments}
	DeploymentEnvironments *[]*string `field:"required" json:"deploymentEnvironments" yaml:"deploymentEnvironments"`
	// Unique name of the Vercel destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#name SecretsSyncVercelDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Project ID where to manage environment variables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#project_id SecretsSyncVercelDestination#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Set of allowed IPv4 addresses in CIDR notation (e.g., 192.168.1.1/32) for outbound connections from Vault to the destination. If not set, all IPv4 addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#allowed_ipv4_addresses SecretsSyncVercelDestination#allowed_ipv4_addresses}
	AllowedIpv4Addresses *[]*string `field:"optional" json:"allowedIpv4Addresses" yaml:"allowedIpv4Addresses"`
	// Set of allowed IPv6 addresses in CIDR notation (e.g., 2001:db8::1/128) for outbound connections from Vault to the destination. If not set, all IPv6 addresses are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#allowed_ipv6_addresses SecretsSyncVercelDestination#allowed_ipv6_addresses}
	AllowedIpv6Addresses *[]*string `field:"optional" json:"allowedIpv6Addresses" yaml:"allowedIpv6Addresses"`
	// Set of allowed ports for outbound connections from Vault to the destination. If not set, all ports are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#allowed_ports SecretsSyncVercelDestination#allowed_ports}
	AllowedPorts *[]*float64 `field:"optional" json:"allowedPorts" yaml:"allowedPorts"`
	// If set to true, disables strict networking enforcement for this destination.
	//
	// When disabled, Vault will not enforce allowed IP addresses and ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#disable_strict_networking SecretsSyncVercelDestination#disable_strict_networking}
	DisableStrictNetworking interface{} `field:"optional" json:"disableStrictNetworking" yaml:"disableStrictNetworking"`
	// Determines what level of information is synced as a distinct resource at the destination. Can be 'secret-path' or 'secret-key'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#granularity SecretsSyncVercelDestination#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#id SecretsSyncVercelDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#namespace SecretsSyncVercelDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#secret_name_template SecretsSyncVercelDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
	// Team ID the project belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_vercel_destination#team_id SecretsSyncVercelDestination#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
}

