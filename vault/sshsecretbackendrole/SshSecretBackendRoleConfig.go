// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sshsecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SshSecretBackendRoleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#backend SshSecretBackendRole#backend}.
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#key_type SshSecretBackendRole#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Unique name for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#name SshSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#algorithm_signer SshSecretBackendRole#algorithm_signer}.
	AlgorithmSigner *string `field:"optional" json:"algorithmSigner" yaml:"algorithmSigner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allow_bare_domains SshSecretBackendRole#allow_bare_domains}.
	AllowBareDomains interface{} `field:"optional" json:"allowBareDomains" yaml:"allowBareDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_critical_options SshSecretBackendRole#allowed_critical_options}.
	AllowedCriticalOptions *string `field:"optional" json:"allowedCriticalOptions" yaml:"allowedCriticalOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_domains SshSecretBackendRole#allowed_domains}.
	AllowedDomains *string `field:"optional" json:"allowedDomains" yaml:"allowedDomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_domains_template SshSecretBackendRole#allowed_domains_template}.
	AllowedDomainsTemplate interface{} `field:"optional" json:"allowedDomainsTemplate" yaml:"allowedDomainsTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_extensions SshSecretBackendRole#allowed_extensions}.
	AllowedExtensions *string `field:"optional" json:"allowedExtensions" yaml:"allowedExtensions"`
	// allowed_user_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_user_key_config SshSecretBackendRole#allowed_user_key_config}
	AllowedUserKeyConfig interface{} `field:"optional" json:"allowedUserKeyConfig" yaml:"allowedUserKeyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_users SshSecretBackendRole#allowed_users}.
	AllowedUsers *string `field:"optional" json:"allowedUsers" yaml:"allowedUsers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allowed_users_template SshSecretBackendRole#allowed_users_template}.
	AllowedUsersTemplate interface{} `field:"optional" json:"allowedUsersTemplate" yaml:"allowedUsersTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allow_empty_principals SshSecretBackendRole#allow_empty_principals}.
	AllowEmptyPrincipals interface{} `field:"optional" json:"allowEmptyPrincipals" yaml:"allowEmptyPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allow_host_certificates SshSecretBackendRole#allow_host_certificates}.
	AllowHostCertificates interface{} `field:"optional" json:"allowHostCertificates" yaml:"allowHostCertificates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allow_subdomains SshSecretBackendRole#allow_subdomains}.
	AllowSubdomains interface{} `field:"optional" json:"allowSubdomains" yaml:"allowSubdomains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allow_user_certificates SshSecretBackendRole#allow_user_certificates}.
	AllowUserCertificates interface{} `field:"optional" json:"allowUserCertificates" yaml:"allowUserCertificates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#allow_user_key_ids SshSecretBackendRole#allow_user_key_ids}.
	AllowUserKeyIds interface{} `field:"optional" json:"allowUserKeyIds" yaml:"allowUserKeyIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#cidr_list SshSecretBackendRole#cidr_list}.
	CidrList *string `field:"optional" json:"cidrList" yaml:"cidrList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#default_critical_options SshSecretBackendRole#default_critical_options}.
	DefaultCriticalOptions *map[string]*string `field:"optional" json:"defaultCriticalOptions" yaml:"defaultCriticalOptions"`
	// Default extensions to include in SSH certificates. Only applicable for CA key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#default_extensions SshSecretBackendRole#default_extensions}
	DefaultExtensions *map[string]*string `field:"optional" json:"defaultExtensions" yaml:"defaultExtensions"`
	// Specifies if the default_extensions field supports templating. Only applicable for CA key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#default_extensions_template SshSecretBackendRole#default_extensions_template}
	DefaultExtensionsTemplate interface{} `field:"optional" json:"defaultExtensionsTemplate" yaml:"defaultExtensionsTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#default_user SshSecretBackendRole#default_user}.
	DefaultUser *string `field:"optional" json:"defaultUser" yaml:"defaultUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#default_user_template SshSecretBackendRole#default_user_template}.
	DefaultUserTemplate interface{} `field:"optional" json:"defaultUserTemplate" yaml:"defaultUserTemplate"`
	// List of CIDR blocks for which credentials cannot be created. Applicable for OTP and dynamic key types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#exclude_cidr_list SshSecretBackendRole#exclude_cidr_list}
	ExcludeCidrList *[]*string `field:"optional" json:"excludeCidrList" yaml:"excludeCidrList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#id SshSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#key_id_format SshSecretBackendRole#key_id_format}.
	KeyIdFormat *string `field:"optional" json:"keyIdFormat" yaml:"keyIdFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#max_ttl SshSecretBackendRole#max_ttl}.
	MaxTtl *string `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#namespace SshSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#not_before_duration SshSecretBackendRole#not_before_duration}
	NotBeforeDuration *string `field:"optional" json:"notBeforeDuration" yaml:"notBeforeDuration"`
	// Specifies the port number for SSH connections (default 22). Applicable for OTP and dynamic key types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#port SshSecretBackendRole#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ssh_secret_backend_role#ttl SshSecretBackendRole#ttl}.
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
}

