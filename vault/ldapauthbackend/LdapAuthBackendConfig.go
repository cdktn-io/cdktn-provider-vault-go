// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ldapauthbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LdapAuthBackendConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#url LdapAuthBackend#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The metadata to be tied to generated entity alias.
	//
	// This should be a list or map containing the metadata in key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#alias_metadata LdapAuthBackend#alias_metadata}
	AliasMetadata *map[string]*string `field:"optional" json:"aliasMetadata" yaml:"aliasMetadata"`
	// Allows anonymous group searches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#anonymous_group_search LdapAuthBackend#anonymous_group_search}
	AnonymousGroupSearch interface{} `field:"optional" json:"anonymousGroupSearch" yaml:"anonymousGroupSearch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#binddn LdapAuthBackend#binddn}.
	Binddn *string `field:"optional" json:"binddn" yaml:"binddn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#bindpass LdapAuthBackend#bindpass}.
	Bindpass *string `field:"optional" json:"bindpass" yaml:"bindpass"`
	// Write-only bind password to use for LDAP authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#bindpass_wo LdapAuthBackend#bindpass_wo}
	BindpassWo *string `field:"optional" json:"bindpassWo" yaml:"bindpassWo"`
	// Version counter for write-only bind password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#bindpass_wo_version LdapAuthBackend#bindpass_wo_version}
	BindpassWoVersion *float64 `field:"optional" json:"bindpassWoVersion" yaml:"bindpassWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#case_sensitive_names LdapAuthBackend#case_sensitive_names}.
	CaseSensitiveNames interface{} `field:"optional" json:"caseSensitiveNames" yaml:"caseSensitiveNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#certificate LdapAuthBackend#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#client_tls_cert LdapAuthBackend#client_tls_cert}.
	ClientTlsCert *string `field:"optional" json:"clientTlsCert" yaml:"clientTlsCert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#client_tls_key LdapAuthBackend#client_tls_key}.
	ClientTlsKey *string `field:"optional" json:"clientTlsKey" yaml:"clientTlsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#connection_timeout LdapAuthBackend#connection_timeout}.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#deny_null_bind LdapAuthBackend#deny_null_bind}.
	DenyNullBind interface{} `field:"optional" json:"denyNullBind" yaml:"denyNullBind"`
	// Specifies how aliases are dereferenced during LDAP searches. Valid values are 'never','searching','finding', and 'always'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#dereference_aliases LdapAuthBackend#dereference_aliases}
	DereferenceAliases *string `field:"optional" json:"dereferenceAliases" yaml:"dereferenceAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#description LdapAuthBackend#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#disable_automated_rotation LdapAuthBackend#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#disable_remount LdapAuthBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#discoverdn LdapAuthBackend#discoverdn}.
	Discoverdn interface{} `field:"optional" json:"discoverdn" yaml:"discoverdn"`
	// Enables login using the sAMAccountName attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#enable_samaccountname_login LdapAuthBackend#enable_samaccountname_login}
	EnableSamaccountnameLogin interface{} `field:"optional" json:"enableSamaccountnameLogin" yaml:"enableSamaccountnameLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#groupattr LdapAuthBackend#groupattr}.
	Groupattr *string `field:"optional" json:"groupattr" yaml:"groupattr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#groupdn LdapAuthBackend#groupdn}.
	Groupdn *string `field:"optional" json:"groupdn" yaml:"groupdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#groupfilter LdapAuthBackend#groupfilter}.
	Groupfilter *string `field:"optional" json:"groupfilter" yaml:"groupfilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#id LdapAuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#insecure_tls LdapAuthBackend#insecure_tls}.
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// Specifies if the auth method is local only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#local LdapAuthBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#max_page_size LdapAuthBackend#max_page_size}.
	MaxPageSize *float64 `field:"optional" json:"maxPageSize" yaml:"maxPageSize"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#namespace LdapAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#path LdapAuthBackend#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The timeout(in sec) for requests to the LDAP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#request_timeout LdapAuthBackend#request_timeout}
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#rotation_period LdapAuthBackend#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#rotation_schedule LdapAuthBackend#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#rotation_window LdapAuthBackend#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#starttls LdapAuthBackend#starttls}.
	Starttls interface{} `field:"optional" json:"starttls" yaml:"starttls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#tls_max_version LdapAuthBackend#tls_max_version}.
	TlsMaxVersion *string `field:"optional" json:"tlsMaxVersion" yaml:"tlsMaxVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#tls_min_version LdapAuthBackend#tls_min_version}.
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
	// Specifies the blocks of IP addresses which are allowed to use the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_bound_cidrs LdapAuthBackend#token_bound_cidrs}
	TokenBoundCidrs *[]*string `field:"optional" json:"tokenBoundCidrs" yaml:"tokenBoundCidrs"`
	// Generated Token's Explicit Maximum TTL in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_explicit_max_ttl LdapAuthBackend#token_explicit_max_ttl}
	TokenExplicitMaxTtl *float64 `field:"optional" json:"tokenExplicitMaxTtl" yaml:"tokenExplicitMaxTtl"`
	// The maximum lifetime of the generated token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_max_ttl LdapAuthBackend#token_max_ttl}
	TokenMaxTtl *float64 `field:"optional" json:"tokenMaxTtl" yaml:"tokenMaxTtl"`
	// If true, the 'default' policy will not automatically be added to generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_no_default_policy LdapAuthBackend#token_no_default_policy}
	TokenNoDefaultPolicy interface{} `field:"optional" json:"tokenNoDefaultPolicy" yaml:"tokenNoDefaultPolicy"`
	// The maximum number of times a token may be used, a value of zero means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_num_uses LdapAuthBackend#token_num_uses}
	TokenNumUses *float64 `field:"optional" json:"tokenNumUses" yaml:"tokenNumUses"`
	// Generated Token's Period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_period LdapAuthBackend#token_period}
	TokenPeriod *float64 `field:"optional" json:"tokenPeriod" yaml:"tokenPeriod"`
	// Generated Token's Policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_policies LdapAuthBackend#token_policies}
	TokenPolicies *[]*string `field:"optional" json:"tokenPolicies" yaml:"tokenPolicies"`
	// The initial ttl of the token to generate in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_ttl LdapAuthBackend#token_ttl}
	TokenTtl *float64 `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
	// The type of token to generate, service or batch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#token_type LdapAuthBackend#token_type}
	TokenType *string `field:"optional" json:"tokenType" yaml:"tokenType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#tune LdapAuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#upndomain LdapAuthBackend#upndomain}.
	Upndomain *string `field:"optional" json:"upndomain" yaml:"upndomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#userattr LdapAuthBackend#userattr}.
	Userattr *string `field:"optional" json:"userattr" yaml:"userattr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#userdn LdapAuthBackend#userdn}.
	Userdn *string `field:"optional" json:"userdn" yaml:"userdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#userfilter LdapAuthBackend#userfilter}.
	Userfilter *string `field:"optional" json:"userfilter" yaml:"userfilter"`
	// Force the auth method to use the username passed by the user as the alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#username_as_alias LdapAuthBackend#username_as_alias}
	UsernameAsAlias interface{} `field:"optional" json:"usernameAsAlias" yaml:"usernameAsAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/ldap_auth_backend#use_token_groups LdapAuthBackend#use_token_groups}.
	UseTokenGroups interface{} `field:"optional" json:"useTokenGroups" yaml:"useTokenGroups"`
}

