// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tpmauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TpmAuthBackendRoleConfig struct {
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
	// TPM auth backend mount path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#mount TpmAuthBackendRole#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the TPM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#name TpmAuthBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A map of string to string that will be set as metadata on the identity alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#alias_metadata TpmAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Certificate TTL for the TPM role, in seconds. Zero means inherit the global `default_cert_ttl` from the backend config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#cert_ttl TpmAuthBackendRole#cert_ttl}
	CertTtl *float64 `field:"optional" json:"certTtl" yaml:"certTtl"`
	// Display name for the role. Defaults to the role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#display_name TpmAuthBackendRole#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#namespace TpmAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_bound_cidrs TpmAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_explicit_max_ttl TpmAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_max_ttl TpmAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_no_default_policy TpmAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_num_uses TpmAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_period TpmAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_policies TpmAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_ttl TpmAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#token_type TpmAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// Set of TPM group IDs authorized to authenticate with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#tpmgroup_ids TpmAuthBackendRole#tpmgroup_ids}
	TpmgroupIds *[]*string `field:"optional" json:"tpmgroupIds" yaml:"tpmgroupIds"`
	// Set of TPM record IDs authorized to authenticate with this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/tpm_auth_backend_role#tpm_ids TpmAuthBackendRole#tpm_ids}
	TpmIds *[]*string `field:"optional" json:"tpmIds" yaml:"tpmIds"`
}

