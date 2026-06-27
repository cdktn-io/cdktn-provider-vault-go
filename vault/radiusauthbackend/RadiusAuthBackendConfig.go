// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package radiusauthbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RadiusAuthBackendConfig struct {
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
	// The RADIUS server to connect to. Examples: `radius.myorg.com`, `127.0.0.1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#host RadiusAuthBackend#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Path of the enabled RADIUS auth backend mount to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#mount RadiusAuthBackend#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The RADIUS shared secret. This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#secret_wo RadiusAuthBackend#secret_wo}
	SecretWo *string `field:"required" json:"secretWo" yaml:"secretWo"`
	// Version counter for the write-only `secret_wo` field.
	//
	// Since write-only values are not stored in state, Terraform cannot detect when the secret changes. Increment this value whenever you update `secret_wo` so Terraform detects the change and applies an update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#secret_wo_version RadiusAuthBackend#secret_wo_version}
	SecretWoVersion *float64 `field:"required" json:"secretWoVersion" yaml:"secretWoVersion"`
	// A map of string to string that will be set as metadata on the identity alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#alias_metadata RadiusAuthBackend#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Number of seconds to wait for a backend connection before timing out.
	//
	// Defaults to `10`. If removed from configuration after being set, Vault retains the previously stored value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#dial_timeout RadiusAuthBackend#dial_timeout}
	DialTimeout *float64 `field:"optional" json:"dialTimeout" yaml:"dialTimeout"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#namespace RadiusAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The NAS-Port attribute of the RADIUS request.
	//
	// Defaults to `10`. If removed from configuration after being set, Vault retains the previously stored value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#nas_port RadiusAuthBackend#nas_port}
	NasPort *float64 `field:"optional" json:"nasPort" yaml:"nasPort"`
	// The UDP port where the RADIUS server is listening on. Defaults to `1812`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#port RadiusAuthBackend#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Number of seconds to wait for a response from the RADIUS server.
	//
	// Defaults to `10`. If removed from configuration after being set, Vault retains the previously stored value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#read_timeout RadiusAuthBackend#read_timeout}
	ReadTimeout *float64 `field:"optional" json:"readTimeout" yaml:"readTimeout"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_bound_cidrs RadiusAuthBackend#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_explicit_max_ttl RadiusAuthBackend#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_max_ttl RadiusAuthBackend#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_no_default_policy RadiusAuthBackend#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_num_uses RadiusAuthBackend#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_period RadiusAuthBackend#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_policies RadiusAuthBackend#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_ttl RadiusAuthBackend#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#token_type RadiusAuthBackend#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// A set of policies to be granted to unregistered users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/radius_auth_backend#unregistered_user_policies RadiusAuthBackend#unregistered_user_policies}
	UnregisteredUserPolicies *[]*string `field:"optional" json:"unregisteredUserPolicies" yaml:"unregisteredUserPolicies"`
}

