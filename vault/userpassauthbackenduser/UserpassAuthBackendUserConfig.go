// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package userpassauthbackenduser

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type UserpassAuthBackendUserConfig struct {
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
	// Mount path for the Userpass auth engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#mount UserpassAuthBackendUser#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Username for this Userpass user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#username UserpassAuthBackendUser#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// A map of string to string that will be set as metadata on the identity alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#alias_metadata UserpassAuthBackendUser#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#namespace UserpassAuthBackendUser#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Pre-hashed password for this user in bcrypt format. Available in Vault 1.17 and later. Mutually exclusive with password_wo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#password_hash_wo UserpassAuthBackendUser#password_hash_wo}
	PasswordHashWo *string `field:"optional" json:"passwordHashWo" yaml:"passwordHashWo"`
	// Version counter for the write-only `password_hash_wo` field.
	//
	// Since write-only values are not stored in state, Terraform cannot detect when the password hash changes. Increment this value whenever you update `password_hash_wo` to ensure the new password hash is sent to Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#password_hash_wo_version UserpassAuthBackendUser#password_hash_wo_version}
	PasswordHashWoVersion *float64 `field:"optional" json:"passwordHashWoVersion" yaml:"passwordHashWoVersion"`
	// Password for this user. This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#password_wo UserpassAuthBackendUser#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Version counter for the write-only `password_wo` field.
	//
	// Since write-only values are not stored in state, Terraform cannot detect when the password changes. Increment this value whenever you update `password_wo` to ensure the new password is sent to Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#password_wo_version UserpassAuthBackendUser#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_bound_cidrs UserpassAuthBackendUser#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_explicit_max_ttl UserpassAuthBackendUser#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_max_ttl UserpassAuthBackendUser#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_no_default_policy UserpassAuthBackendUser#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_num_uses UserpassAuthBackendUser#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_period UserpassAuthBackendUser#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_policies UserpassAuthBackendUser#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_ttl UserpassAuthBackendUser#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/userpass_auth_backend_user#token_type UserpassAuthBackendUser#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
}

