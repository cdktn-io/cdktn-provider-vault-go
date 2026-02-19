// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package adsecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AdSecretBackendConfig struct {
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
	// Distinguished name of object to bind when performing user and group search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#binddn AdSecretBackend#binddn}
	Binddn *string `field:"required" json:"binddn" yaml:"binddn"`
	// LDAP password for searching for the user DN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#bindpass AdSecretBackend#bindpass}
	Bindpass *string `field:"required" json:"bindpass" yaml:"bindpass"`
	// Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#anonymous_group_search AdSecretBackend#anonymous_group_search}
	AnonymousGroupSearch interface{} `field:"optional" json:"anonymousGroupSearch" yaml:"anonymousGroupSearch"`
	// The mount path for a backend, for example, the path given in "$ vault auth enable -path=my-ad ad".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#backend AdSecretBackend#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// If true, case sensitivity will be used when comparing usernames and groups for matching policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#case_sensitive_names AdSecretBackend#case_sensitive_names}
	CaseSensitiveNames interface{} `field:"optional" json:"caseSensitiveNames" yaml:"caseSensitiveNames"`
	// CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#certificate AdSecretBackend#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Client certificate to provide to the LDAP server, must be x509 PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#client_tls_cert AdSecretBackend#client_tls_cert}
	ClientTlsCert *string `field:"optional" json:"clientTlsCert" yaml:"clientTlsCert"`
	// Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#client_tls_key AdSecretBackend#client_tls_key}
	ClientTlsKey *string `field:"optional" json:"clientTlsKey" yaml:"clientTlsKey"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#default_lease_ttl_seconds AdSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// Denies an unauthenticated LDAP bind request if the user's password is empty; defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#deny_null_bind AdSecretBackend#deny_null_bind}
	DenyNullBind interface{} `field:"optional" json:"denyNullBind" yaml:"denyNullBind"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#description AdSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#disable_remount AdSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Use anonymous bind to discover the bind DN of a user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#discoverdn AdSecretBackend#discoverdn}
	Discoverdn interface{} `field:"optional" json:"discoverdn" yaml:"discoverdn"`
	// LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate user group membership.
	//
	// Examples: "cn" or "memberOf", etc. Default: cn
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#groupattr AdSecretBackend#groupattr}
	Groupattr *string `field:"optional" json:"groupattr" yaml:"groupattr"`
	// LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#groupdn AdSecretBackend#groupdn}
	Groupdn *string `field:"optional" json:"groupdn" yaml:"groupdn"`
	// Go template for querying group membership of user.
	//
	// The template can access the following context variables: UserDN, Username Example: (&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}})) Default: (|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#groupfilter AdSecretBackend#groupfilter}
	Groupfilter *string `field:"optional" json:"groupfilter" yaml:"groupfilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#id AdSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Skip LDAP server SSL Certificate verification - insecure and not recommended for production use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#insecure_tls AdSecretBackend#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// The number of seconds after a Vault rotation where, if Active Directory shows a later rotation, it should be considered out-of-band.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#last_rotation_tolerance AdSecretBackend#last_rotation_tolerance}
	LastRotationTolerance *float64 `field:"optional" json:"lastRotationTolerance" yaml:"lastRotationTolerance"`
	// Mark the secrets engine as local-only.
	//
	// Local engines are not replicated or removed by replication.Tolerance duration to use when checking the last rotation time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#local AdSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#max_lease_ttl_seconds AdSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// In seconds, the maximum password time-to-live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#max_ttl AdSecretBackend#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#namespace AdSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Name of the password policy to use to generate passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#password_policy AdSecretBackend#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Timeout, in seconds, for the connection when making requests against the server before returning back an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#request_timeout AdSecretBackend#request_timeout}
	RequestTimeout *float64 `field:"optional" json:"requestTimeout" yaml:"requestTimeout"`
	// Issue a StartTLS command after establishing unencrypted connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#starttls AdSecretBackend#starttls}
	Starttls interface{} `field:"optional" json:"starttls" yaml:"starttls"`
	// Maximum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#tls_max_version AdSecretBackend#tls_max_version}
	TlsMaxVersion *string `field:"optional" json:"tlsMaxVersion" yaml:"tlsMaxVersion"`
	// Minimum TLS version to use. Accepted values are 'tls10', 'tls11', 'tls12' or 'tls13'. Defaults to 'tls12'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#tls_min_version AdSecretBackend#tls_min_version}
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
	// In seconds, the default password time-to-live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#ttl AdSecretBackend#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Enables userPrincipalDomain login with [username]@UPNDomain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#upndomain AdSecretBackend#upndomain}
	Upndomain *string `field:"optional" json:"upndomain" yaml:"upndomain"`
	// LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#url AdSecretBackend#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#use_pre111_group_cn_behavior AdSecretBackend#use_pre111_group_cn_behavior}
	UsePre111GroupCnBehavior interface{} `field:"optional" json:"usePre111GroupCnBehavior" yaml:"usePre111GroupCnBehavior"`
	// Attribute used for users (default: cn).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#userattr AdSecretBackend#userattr}
	Userattr *string `field:"optional" json:"userattr" yaml:"userattr"`
	// LDAP domain to use for users (eg: ou=People,dc=example,dc=org).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#userdn AdSecretBackend#userdn}
	Userdn *string `field:"optional" json:"userdn" yaml:"userdn"`
	// If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships.
	//
	// This will find all security groups including nested ones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/ad_secret_backend#use_token_groups AdSecretBackend#use_token_groups}
	UseTokenGroups interface{} `field:"optional" json:"useTokenGroups" yaml:"useTokenGroups"`
}

