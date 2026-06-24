// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetessecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesSecretBackendConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#path KubernetesSecretBackend#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#allowed_managed_keys KubernetesSecretBackend#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#allowed_response_headers KubernetesSecretBackend#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#audit_non_hmac_request_keys KubernetesSecretBackend#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#audit_non_hmac_response_keys KubernetesSecretBackend#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// Default lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#default_lease_ttl_seconds KubernetesSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#delegated_auth_accessors KubernetesSecretBackend#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#description KubernetesSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Disable defaulting to the local CA certificate and service account JWT when running in a Kubernetes pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#disable_local_ca_jwt KubernetesSecretBackend#disable_local_ca_jwt}
	DisableLocalCaJwt interface{} `field:"optional" json:"disableLocalCaJwt" yaml:"disableLocalCaJwt"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#external_entropy_access KubernetesSecretBackend#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#force_no_cache KubernetesSecretBackend#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#id KubernetesSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The key to use for signing plugin workload identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#identity_token_key KubernetesSecretBackend#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// A PEM-encoded CA certificate used by the secret engine to verify the Kubernetes API server certificate.
	//
	// Defaults to the local pod’s CA if found, or otherwise the host's root CA set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#kubernetes_ca_cert KubernetesSecretBackend#kubernetes_ca_cert}
	KubernetesCaCert *string `field:"optional" json:"kubernetesCaCert" yaml:"kubernetesCaCert"`
	// The Kubernetes API URL to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#kubernetes_host KubernetesSecretBackend#kubernetes_host}
	KubernetesHost *string `field:"optional" json:"kubernetesHost" yaml:"kubernetesHost"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#listing_visibility KubernetesSecretBackend#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Local mount flag that can be explicitly set to true to enforce local mount in HA environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#local KubernetesSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#max_lease_ttl_seconds KubernetesSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#namespace KubernetesSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#options KubernetesSecretBackend#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#passthrough_request_headers KubernetesSecretBackend#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#plugin_version KubernetesSecretBackend#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#seal_wrap KubernetesSecretBackend#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
	// The JSON web token of the service account used by the secrets engine to manage Kubernetes credentials.
	//
	// Defaults to the local pod’s JWT if found.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#service_account_jwt KubernetesSecretBackend#service_account_jwt}
	ServiceAccountJwt *string `field:"optional" json:"serviceAccountJwt" yaml:"serviceAccountJwt"`
	// Write-only JSON web token of the service account used by the secrets engine to manage Kubernetes credentials.
	//
	// This value will not be stored in state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#service_account_jwt_wo KubernetesSecretBackend#service_account_jwt_wo}
	ServiceAccountJwtWo *string `field:"optional" json:"serviceAccountJwtWo" yaml:"serviceAccountJwtWo"`
	// Version counter for write-only service account JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/kubernetes_secret_backend#service_account_jwt_wo_version KubernetesSecretBackend#service_account_jwt_wo_version}
	ServiceAccountJwtWoVersion *float64 `field:"optional" json:"serviceAccountJwtWoVersion" yaml:"serviceAccountJwtWoVersion"`
}

