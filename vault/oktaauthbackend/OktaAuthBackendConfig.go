// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oktaauthbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OktaAuthBackendConfig struct {
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
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#alias_metadata OktaAuthBackend#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// The Okta API token.
	//
	// This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#api_token OktaAuthBackend#api_token}
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Write-only Okta API token.
	//
	// This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#api_token_wo OktaAuthBackend#api_token_wo}
	ApiTokenWo *string `field:"optional" json:"apiTokenWo" yaml:"apiTokenWo"`
	// Version counter for write-only api_token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#api_token_wo_version OktaAuthBackend#api_token_wo_version}
	ApiTokenWoVersion *float64 `field:"optional" json:"apiTokenWoVersion" yaml:"apiTokenWoVersion"`
	// The Okta url. Examples: oktapreview.com, okta.com (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#base_url OktaAuthBackend#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// When true, requests by Okta for a MFA check will be bypassed.
	//
	// This also disallows certain status checks on the account, such as whether the password is expired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#bypass_okta_mfa OktaAuthBackend#bypass_okta_mfa}
	BypassOktaMfa interface{} `field:"optional" json:"bypassOktaMfa" yaml:"bypassOktaMfa"`
	// The description of the auth backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#description OktaAuthBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#disable_remount OktaAuthBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#group OktaAuthBackend#group}.
	Group interface{} `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#id OktaAuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#namespace OktaAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The Okta organization. This will be the first part of the url https://XXX.okta.com. Use org_name instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#organization OktaAuthBackend#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The Okta organization. This will be the first part of the url https://XXX.okta.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#org_name OktaAuthBackend#org_name}
	OrgName *string `field:"optional" json:"orgName" yaml:"orgName"`
	// path to mount the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#path OktaAuthBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The Okta API token.
	//
	// This is required to query Okta for user group membership. If this is not supplied only locally configured groups will be enabled. Use api_token instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token OktaAuthBackend#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_bound_cidrs OktaAuthBackend#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_explicit_max_ttl OktaAuthBackend#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_max_ttl OktaAuthBackend#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_no_default_policy OktaAuthBackend#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_num_uses OktaAuthBackend#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_period OktaAuthBackend#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_policies OktaAuthBackend#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_ttl OktaAuthBackend#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#token_type OktaAuthBackend#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#tune OktaAuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/okta_auth_backend#user OktaAuthBackend#user}.
	User interface{} `field:"optional" json:"user" yaml:"user"`
}

