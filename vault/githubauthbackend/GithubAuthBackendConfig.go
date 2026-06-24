// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package githubauthbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GithubAuthBackendConfig struct {
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
	// The organization users must be part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#organization GithubAuthBackend#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#alias_metadata GithubAuthBackend#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// The API endpoint to use. Useful if you are running GitHub Enterprise or an API-compatible authentication server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#base_url GithubAuthBackend#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Specifies the description of the mount. This overrides the current stored value, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#description GithubAuthBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#disable_remount GithubAuthBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#id GithubAuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#namespace GithubAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The ID of the organization users must be part of.
	//
	// Vault will attempt to fetch and set this value if it is not provided (vault-1.10+)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#organization_id GithubAuthBackend#organization_id}
	OrganizationId *float64 `field:"optional" json:"organizationId" yaml:"organizationId"`
	// Path where the auth backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#path GithubAuthBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_bound_cidrs GithubAuthBackend#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_explicit_max_ttl GithubAuthBackend#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_max_ttl GithubAuthBackend#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_no_default_policy GithubAuthBackend#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_num_uses GithubAuthBackend#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_period GithubAuthBackend#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_policies GithubAuthBackend#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_ttl GithubAuthBackend#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#token_type GithubAuthBackend#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/github_auth_backend#tune GithubAuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
}

