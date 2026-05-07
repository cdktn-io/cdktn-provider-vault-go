// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package azuresecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AzureSecretBackendConfig struct {
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
	// The subscription id for the Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#subscription_id AzureSecretBackend#subscription_id}
	SubscriptionId *string `field:"required" json:"subscriptionId" yaml:"subscriptionId"`
	// The tenant id for the Azure Active Directory organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#tenant_id AzureSecretBackend#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#allowed_managed_keys AzureSecretBackend#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#allowed_response_headers AzureSecretBackend#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#audit_non_hmac_request_keys AzureSecretBackend#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#audit_non_hmac_response_keys AzureSecretBackend#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// The client id for credentials to query the Azure APIs. Currently read permissions to query compute resources are required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#client_id AzureSecretBackend#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret for credentials to query the Azure APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#client_secret AzureSecretBackend#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The client secret for credentials to query the Azure APIs.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#client_secret_wo AzureSecretBackend#client_secret_wo}
	ClientSecretWo *string `field:"optional" json:"clientSecretWo" yaml:"clientSecretWo"`
	// A version counter for the write-only client_secret_wo field. Incrementing this value will trigger an update to the client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#client_secret_wo_version AzureSecretBackend#client_secret_wo_version}
	ClientSecretWoVersion *float64 `field:"optional" json:"clientSecretWoVersion" yaml:"clientSecretWoVersion"`
	// Default lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#default_lease_ttl_seconds AzureSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#delegated_auth_accessors AzureSecretBackend#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#description AzureSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#disable_automated_rotation AzureSecretBackend#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#disable_remount AzureSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// The Azure cloud environment. Valid values: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#environment AzureSecretBackend#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#external_entropy_access AzureSecretBackend#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#force_no_cache AzureSecretBackend#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#id AzureSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#identity_token_audience AzureSecretBackend#identity_token_audience}
	IdentityTokenAudience *string `field:"optional" json:"identityTokenAudience" yaml:"identityTokenAudience"`
	// The key to use for signing identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#identity_token_key AzureSecretBackend#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// The TTL of generated identity tokens in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#identity_token_ttl AzureSecretBackend#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#listing_visibility AzureSecretBackend#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#local AzureSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#max_lease_ttl_seconds AzureSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#namespace AzureSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#options AzureSecretBackend#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#passthrough_request_headers AzureSecretBackend#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Path to mount the backend at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#path AzureSecretBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#plugin_version AzureSecretBackend#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// The TTL in seconds of the root password in Azure when rotate-root generates a new client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#root_password_ttl AzureSecretBackend#root_password_ttl}
	RootPasswordTtl *float64 `field:"optional" json:"rootPasswordTtl" yaml:"rootPasswordTtl"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#rotation_period AzureSecretBackend#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#rotation_schedule AzureSecretBackend#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#rotation_window AzureSecretBackend#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/azure_secret_backend#seal_wrap AzureSecretBackend#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
}

