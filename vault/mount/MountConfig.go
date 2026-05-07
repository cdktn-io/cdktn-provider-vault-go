// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MountConfig struct {
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
	// Where the secret backend will be mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#path Mount#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Type of the backend, such as 'aws'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#type Mount#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#allowed_managed_keys Mount#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#allowed_response_headers Mount#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#audit_non_hmac_request_keys Mount#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#audit_non_hmac_response_keys Mount#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// Default lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#default_lease_ttl_seconds Mount#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#delegated_auth_accessors Mount#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#description Mount#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#external_entropy_access Mount#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#force_no_cache Mount#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#id Mount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The key to use for signing plugin workload identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#identity_token_key Mount#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#listing_visibility Mount#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#local Mount#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#max_lease_ttl_seconds Mount#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#namespace Mount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#options Mount#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#passthrough_request_headers Mount#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#plugin_version Mount#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/mount#seal_wrap Mount#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
}

