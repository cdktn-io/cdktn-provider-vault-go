// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package jwtauthbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JwtAuthBackendRoleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#role_name JwtAuthBackendRole#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// The claim to use to uniquely identify the user;
	//
	// this will be used as the name for the Identity entity alias created due to a successful login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#user_claim JwtAuthBackendRole#user_claim}
	UserClaim *string `field:"required" json:"userClaim" yaml:"userClaim"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#alias_metadata JwtAuthBackendRole#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// The list of allowed values for redirect_uri during OIDC logins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#allowed_redirect_uris JwtAuthBackendRole#allowed_redirect_uris}
	AllowedRedirectUris *[]*string `field:"optional" json:"allowedRedirectUris" yaml:"allowedRedirectUris"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#backend JwtAuthBackendRole#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// List of aud claims to match against. Any match is sufficient.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#bound_audiences JwtAuthBackendRole#bound_audiences}
	BoundAudiences *[]*string `field:"optional" json:"boundAudiences" yaml:"boundAudiences"`
	// Map of claims/values to match against. The expected value may be a single string or a comma-separated string list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#bound_claims JwtAuthBackendRole#bound_claims}
	BoundClaims *map[string]*string `field:"optional" json:"boundClaims" yaml:"boundClaims"`
	// How to interpret values in the claims/values map: can be either "string" (exact match) or "glob" (wildcard match).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#bound_claims_type JwtAuthBackendRole#bound_claims_type}
	BoundClaimsType *string `field:"optional" json:"boundClaimsType" yaml:"boundClaimsType"`
	// If set, requires that the sub claim matches this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#bound_subject JwtAuthBackendRole#bound_subject}
	BoundSubject *string `field:"optional" json:"boundSubject" yaml:"boundSubject"`
	// Map of claims (keys) to be copied to specified metadata fields (values).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#claim_mappings JwtAuthBackendRole#claim_mappings}
	ClaimMappings *map[string]*string `field:"optional" json:"claimMappings" yaml:"claimMappings"`
	// The amount of leeway to add to all claims to account for clock skew, in seconds.
	//
	// Defaults to 60 seconds if set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#clock_skew_leeway JwtAuthBackendRole#clock_skew_leeway}
	ClockSkewLeeway *float64 `field:"optional" json:"clockSkewLeeway" yaml:"clockSkewLeeway"`
	// Disable bound claim value parsing. Useful when values contain commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#disable_bound_claims_parsing JwtAuthBackendRole#disable_bound_claims_parsing}
	DisableBoundClaimsParsing interface{} `field:"optional" json:"disableBoundClaimsParsing" yaml:"disableBoundClaimsParsing"`
	// The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds.
	//
	// Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#expiration_leeway JwtAuthBackendRole#expiration_leeway}
	ExpirationLeeway *float64 `field:"optional" json:"expirationLeeway" yaml:"expirationLeeway"`
	// The claim to use to uniquely identify the set of groups to which the user belongs;
	//
	// this will be used as the names for the Identity group aliases created due to a successful login. The claim value must be a list of strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#groups_claim JwtAuthBackendRole#groups_claim}
	GroupsClaim *string `field:"optional" json:"groupsClaim" yaml:"groupsClaim"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#id JwtAuthBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#max_age JwtAuthBackendRole#max_age}
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#namespace JwtAuthBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds.
	//
	// Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Only applicable with 'jwt' roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#not_before_leeway JwtAuthBackendRole#not_before_leeway}
	NotBeforeLeeway *float64 `field:"optional" json:"notBeforeLeeway" yaml:"notBeforeLeeway"`
	// List of OIDC scopes to be used with an OIDC role.
	//
	// The standard scope "openid" is automatically included and need not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#oidc_scopes JwtAuthBackendRole#oidc_scopes}
	OidcScopes *[]*string `field:"optional" json:"oidcScopes" yaml:"oidcScopes"`
	// Type of role, either "oidc" (default) or "jwt".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#role_type JwtAuthBackendRole#role_type}
	RoleType *string `field:"optional" json:"roleType" yaml:"roleType"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_bound_cidrs JwtAuthBackendRole#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_explicit_max_ttl JwtAuthBackendRole#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_max_ttl JwtAuthBackendRole#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_no_default_policy JwtAuthBackendRole#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_num_uses JwtAuthBackendRole#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_period JwtAuthBackendRole#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_policies JwtAuthBackendRole#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_ttl JwtAuthBackendRole#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#token_type JwtAuthBackendRole#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// Specifies if the user_claim value uses JSON pointer syntax for referencing claims.
	//
	// By default, the user_claim value will not use JSON pointer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#user_claim_json_pointer JwtAuthBackendRole#user_claim_json_pointer}
	UserClaimJsonPointer interface{} `field:"optional" json:"userClaimJsonPointer" yaml:"userClaimJsonPointer"`
	// Log received OIDC tokens and claims when debug-level logging is active.
	//
	// Not recommended in production since sensitive information may be present in OIDC responses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/jwt_auth_backend_role#verbose_oidc_logging JwtAuthBackendRole#verbose_oidc_logging}
	VerboseOidcLogging interface{} `field:"optional" json:"verboseOidcLogging" yaml:"verboseOidcLogging"`
}

