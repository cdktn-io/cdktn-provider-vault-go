// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpsecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GcpSecretBackendConfig struct {
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
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#allowed_managed_keys GcpSecretBackend#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#allowed_response_headers GcpSecretBackend#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#audit_non_hmac_request_keys GcpSecretBackend#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#audit_non_hmac_response_keys GcpSecretBackend#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// JSON-encoded credentials to use to connect to GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#credentials GcpSecretBackend#credentials}
	Credentials *string `field:"optional" json:"credentials" yaml:"credentials"`
	// Write-only JSON-encoded credentials to use to connect to GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#credentials_wo GcpSecretBackend#credentials_wo}
	CredentialsWo *string `field:"optional" json:"credentialsWo" yaml:"credentialsWo"`
	// Version counter for write-only JSON-encoded credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#credentials_wo_version GcpSecretBackend#credentials_wo_version}
	CredentialsWoVersion *float64 `field:"optional" json:"credentialsWoVersion" yaml:"credentialsWoVersion"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#default_lease_ttl_seconds GcpSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#delegated_auth_accessors GcpSecretBackend#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#description GcpSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#disable_automated_rotation GcpSecretBackend#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#disable_remount GcpSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#external_entropy_access GcpSecretBackend#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#force_no_cache GcpSecretBackend#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#id GcpSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value for plugin identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#identity_token_audience GcpSecretBackend#identity_token_audience}
	IdentityTokenAudience *string `field:"optional" json:"identityTokenAudience" yaml:"identityTokenAudience"`
	// The key to use for signing identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#identity_token_key GcpSecretBackend#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// The TTL of generated tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#identity_token_ttl GcpSecretBackend#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#listing_visibility GcpSecretBackend#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#local GcpSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#max_lease_ttl_seconds GcpSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// The maximum TTL for long-lived credentials (i.e. service account keys).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#max_ttl GcpSecretBackend#max_ttl}
	MaxTtl *float64 `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#namespace GcpSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#options GcpSecretBackend#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#passthrough_request_headers GcpSecretBackend#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Path to mount the backend at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#path GcpSecretBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#plugin_version GcpSecretBackend#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#rotation_period GcpSecretBackend#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#rotation_schedule GcpSecretBackend#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#rotation_window GcpSecretBackend#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#seal_wrap GcpSecretBackend#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
	// Service Account to impersonate for plugin workload identity federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#service_account_email GcpSecretBackend#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The default TTL for long-lived credentials (i.e. service account keys).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/gcp_secret_backend#ttl GcpSecretBackend#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

