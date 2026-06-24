// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package consulsecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConsulSecretBackendConfig struct {
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
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#address ConsulSecretBackend#address}
	Address *string `field:"required" json:"address" yaml:"address"`
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#allowed_managed_keys ConsulSecretBackend#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#allowed_response_headers ConsulSecretBackend#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#audit_non_hmac_request_keys ConsulSecretBackend#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#audit_non_hmac_response_keys ConsulSecretBackend#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// Denotes a backend resource that is used to bootstrap the Consul ACL system.
	//
	// Only one resource may be used to bootstrap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#bootstrap ConsulSecretBackend#bootstrap}
	Bootstrap interface{} `field:"optional" json:"bootstrap" yaml:"bootstrap"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#ca_cert ConsulSecretBackend#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#client_cert ConsulSecretBackend#client_cert}
	ClientCert *string `field:"optional" json:"clientCert" yaml:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
	//
	// Mutually exclusive with 'client_key_wo'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#client_key ConsulSecretBackend#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded.
	//
	// This field is write-only and will never be stored in state. Mutually exclusive with 'client_key'. Requires 'client_key_wo_version' to trigger updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#client_key_wo ConsulSecretBackend#client_key_wo}
	ClientKeyWo *string `field:"optional" json:"clientKeyWo" yaml:"clientKeyWo"`
	// Version counter for the write-only client key.
	//
	// Increment this value to trigger rotation of the client key. Required when using 'client_key_wo'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#client_key_wo_version ConsulSecretBackend#client_key_wo_version}
	ClientKeyWoVersion *float64 `field:"optional" json:"clientKeyWoVersion" yaml:"clientKeyWoVersion"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#default_lease_ttl_seconds ConsulSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#delegated_auth_accessors ConsulSecretBackend#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#description ConsulSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#disable_remount ConsulSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#external_entropy_access ConsulSecretBackend#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#force_no_cache ConsulSecretBackend#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#id ConsulSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The key to use for signing plugin workload identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#identity_token_key ConsulSecretBackend#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#listing_visibility ConsulSecretBackend#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Specifies if the secret backend is local only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#local ConsulSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#max_lease_ttl_seconds ConsulSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#namespace ConsulSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#options ConsulSecretBackend#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#passthrough_request_headers ConsulSecretBackend#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Unique name of the Vault Consul mount to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#path ConsulSecretBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#plugin_version ConsulSecretBackend#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// Specifies the URL scheme to use. Defaults to "http".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#scheme ConsulSecretBackend#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#seal_wrap ConsulSecretBackend#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
	// Specifies the Consul token to use when managing or issuing new tokens. Mutually exclusive with 'token_wo'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#token ConsulSecretBackend#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Specifies the Consul token to use when managing or issuing new tokens.
	//
	// This field is write-only and will never be stored in state. Mutually exclusive with 'token'. Requires 'token_wo_version' to trigger updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#token_wo ConsulSecretBackend#token_wo}
	TokenWo *string `field:"optional" json:"tokenWo" yaml:"tokenWo"`
	// Version counter for the write-only token. Increment this value to trigger rotation of the token. Required when using 'token_wo'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/consul_secret_backend#token_wo_version ConsulSecretBackend#token_wo_version}
	TokenWoVersion *float64 `field:"optional" json:"tokenWoVersion" yaml:"tokenWoVersion"`
}

