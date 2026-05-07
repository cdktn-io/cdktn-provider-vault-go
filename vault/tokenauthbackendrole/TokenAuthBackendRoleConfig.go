// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package tokenauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TokenAuthBackendRoleConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#role_name TokenAuthBackendRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#alias_metadata TokenAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Set of allowed entity aliases for this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#allowed_entity_aliases TokenAuthBackendRole#allowed_entity_aliases}
	AllowedEntityAliases *[]*string `field:"optional" json:"allowedEntityAliases" yaml:"allowedEntityAliases"`
	// List of allowed policies for given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#allowed_policies TokenAuthBackendRole#allowed_policies}
	AllowedPolicies *[]*string `field:"optional" json:"allowedPolicies" yaml:"allowedPolicies"`
	// Set of allowed policies with glob match for given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#allowed_policies_glob TokenAuthBackendRole#allowed_policies_glob}
	AllowedPoliciesGlob *[]*string `field:"optional" json:"allowedPoliciesGlob" yaml:"allowedPoliciesGlob"`
	// List of disallowed policies for given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#disallowed_policies TokenAuthBackendRole#disallowed_policies}
	DisallowedPolicies *[]*string `field:"optional" json:"disallowedPolicies" yaml:"disallowedPolicies"`
	// Set of disallowed policies with glob match for given role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#disallowed_policies_glob TokenAuthBackendRole#disallowed_policies_glob}
	DisallowedPoliciesGlob *[]*string `field:"optional" json:"disallowedPoliciesGlob" yaml:"disallowedPoliciesGlob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#id TokenAuthBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#namespace TokenAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// If true, tokens created against this policy will be orphan tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#orphan TokenAuthBackendRole#orphan}
	Orphan interface{} `field:"optional" json:"orphan" yaml:"orphan"`
	// Tokens created against this role will have the given suffix as part of their path in addition to the role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#path_suffix TokenAuthBackendRole#path_suffix}
	PathSuffix *string `field:"optional" json:"pathSuffix" yaml:"pathSuffix"`
	// Whether to disable the ability of the token to be renewed past its initial TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#renewable TokenAuthBackendRole#renewable}
	Renewable interface{} `field:"optional" json:"renewable" yaml:"renewable"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_bound_cidrs TokenAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_explicit_max_ttl TokenAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_max_ttl TokenAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_no_default_policy TokenAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_num_uses TokenAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_period TokenAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_policies TokenAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_ttl TokenAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token_auth_backend_role#token_type TokenAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

