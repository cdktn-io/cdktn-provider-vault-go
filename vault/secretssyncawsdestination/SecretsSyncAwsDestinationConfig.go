// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secretssyncawsdestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecretsSyncAwsDestinationConfig struct {
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
	// Unique name of the AWS destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#name SecretsSyncAwsDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Access key id to authenticate against the AWS secrets manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#access_key_id SecretsSyncAwsDestination#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Allowed IPv4 addresses for outbound connections from Vault to AWS Secrets Manager.
	//
	// Can also be set via an IP address range using CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#allowed_ipv4_addresses SecretsSyncAwsDestination#allowed_ipv4_addresses}
	AllowedIpv4Addresses *[]*string `field:"optional" json:"allowedIpv4Addresses" yaml:"allowedIpv4Addresses"`
	// Allowed IPv6 addresses for outbound connections from Vault to AWS Secrets Manager.
	//
	// Can also be set via an IP address range using CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#allowed_ipv6_addresses SecretsSyncAwsDestination#allowed_ipv6_addresses}
	AllowedIpv6Addresses *[]*string `field:"optional" json:"allowedIpv6Addresses" yaml:"allowedIpv6Addresses"`
	// Allowed ports for outbound connections from Vault to AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#allowed_ports SecretsSyncAwsDestination#allowed_ports}
	AllowedPorts *[]*float64 `field:"optional" json:"allowedPorts" yaml:"allowedPorts"`
	// Custom tags to set on the secret managed at the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#custom_tags SecretsSyncAwsDestination#custom_tags}
	CustomTags *map[string]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Disable strict networking mode. When set to true, Vault will not enforce allowed IP addresses and ports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#disable_strict_networking SecretsSyncAwsDestination#disable_strict_networking}
	DisableStrictNetworking interface{} `field:"optional" json:"disableStrictNetworking" yaml:"disableStrictNetworking"`
	// Extra protection that must match the trust policy granting access to the AWS IAM role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#external_id SecretsSyncAwsDestination#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Determines what level of information is synced as a distinct resource at the destination. Can be 'secret-path' or 'secret-key'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#granularity SecretsSyncAwsDestination#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#id SecretsSyncAwsDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value for identity tokens.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#identity_token_audience_wo SecretsSyncAwsDestination#identity_token_audience_wo}
	IdentityTokenAudienceWo *string `field:"optional" json:"identityTokenAudienceWo" yaml:"identityTokenAudienceWo"`
	// A version counter for the write-only identity_token_audience_wo field. Incrementing this value will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#identity_token_audience_wo_version SecretsSyncAwsDestination#identity_token_audience_wo_version}
	IdentityTokenAudienceWoVersion *float64 `field:"optional" json:"identityTokenAudienceWoVersion" yaml:"identityTokenAudienceWoVersion"`
	// The key to use for signing identity tokens.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#identity_token_key_wo SecretsSyncAwsDestination#identity_token_key_wo}
	IdentityTokenKeyWo *string `field:"optional" json:"identityTokenKeyWo" yaml:"identityTokenKeyWo"`
	// A version counter for the write-only identity_token_key_wo field. Incrementing this value will trigger an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#identity_token_key_wo_version SecretsSyncAwsDestination#identity_token_key_wo_version}
	IdentityTokenKeyWoVersion *float64 `field:"optional" json:"identityTokenKeyWoVersion" yaml:"identityTokenKeyWoVersion"`
	// The TTL of generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#identity_token_ttl SecretsSyncAwsDestination#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#namespace SecretsSyncAwsDestination#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Region where to manage the secrets manager entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#region SecretsSyncAwsDestination#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Specifies a role to assume when connecting to AWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#role_arn SecretsSyncAwsDestination#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Secret access key to authenticate against the AWS secrets manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#secret_access_key SecretsSyncAwsDestination#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
	// Template describing how to generate external secret names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/secrets_sync_aws_destination#secret_name_template SecretsSyncAwsDestination#secret_name_template}
	SecretNameTemplate *string `field:"optional" json:"secretNameTemplate" yaml:"secretNameTemplate"`
}

