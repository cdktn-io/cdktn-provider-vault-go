// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package awssecretbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AwsSecretBackendConfig struct {
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
	// The AWS Access Key ID to use when generating new credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#access_key AwsSecretBackend#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// List of managed key registry entry names that the mount in question is allowed to access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#allowed_managed_keys AwsSecretBackend#allowed_managed_keys}
	AllowedManagedKeys *[]*string `field:"optional" json:"allowedManagedKeys" yaml:"allowedManagedKeys"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#allowed_response_headers AwsSecretBackend#allowed_response_headers}
	AllowedResponseHeaders *[]*string `field:"optional" json:"allowedResponseHeaders" yaml:"allowedResponseHeaders"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#audit_non_hmac_request_keys AwsSecretBackend#audit_non_hmac_request_keys}
	AuditNonHmacRequestKeys *[]*string `field:"optional" json:"auditNonHmacRequestKeys" yaml:"auditNonHmacRequestKeys"`
	// Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#audit_non_hmac_response_keys AwsSecretBackend#audit_non_hmac_response_keys}
	AuditNonHmacResponseKeys *[]*string `field:"optional" json:"auditNonHmacResponseKeys" yaml:"auditNonHmacResponseKeys"`
	// Default lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#default_lease_ttl_seconds AwsSecretBackend#default_lease_ttl_seconds}
	DefaultLeaseTtlSeconds *float64 `field:"optional" json:"defaultLeaseTtlSeconds" yaml:"defaultLeaseTtlSeconds"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#delegated_auth_accessors AwsSecretBackend#delegated_auth_accessors}
	DelegatedAuthAccessors *[]*string `field:"optional" json:"delegatedAuthAccessors" yaml:"delegatedAuthAccessors"`
	// Human-friendly description of the mount for the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#description AwsSecretBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#disable_automated_rotation AwsSecretBackend#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#disable_remount AwsSecretBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Enable the secrets engine to access Vault's external entropy source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#external_entropy_access AwsSecretBackend#external_entropy_access}
	ExternalEntropyAccess interface{} `field:"optional" json:"externalEntropyAccess" yaml:"externalEntropyAccess"`
	// If set to true, disables caching.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#force_no_cache AwsSecretBackend#force_no_cache}
	ForceNoCache interface{} `field:"optional" json:"forceNoCache" yaml:"forceNoCache"`
	// Specifies a custom HTTP IAM endpoint to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#iam_endpoint AwsSecretBackend#iam_endpoint}
	IamEndpoint *string `field:"optional" json:"iamEndpoint" yaml:"iamEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#id AwsSecretBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The audience claim value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#identity_token_audience AwsSecretBackend#identity_token_audience}
	IdentityTokenAudience *string `field:"optional" json:"identityTokenAudience" yaml:"identityTokenAudience"`
	// The key to use for signing identity tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#identity_token_key AwsSecretBackend#identity_token_key}
	IdentityTokenKey *string `field:"optional" json:"identityTokenKey" yaml:"identityTokenKey"`
	// The TTL of generated identity tokens in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#identity_token_ttl AwsSecretBackend#identity_token_ttl}
	IdentityTokenTtl *float64 `field:"optional" json:"identityTokenTtl" yaml:"identityTokenTtl"`
	// Specifies whether to show this mount in the UI-specific listing endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#listing_visibility AwsSecretBackend#listing_visibility}
	ListingVisibility *string `field:"optional" json:"listingVisibility" yaml:"listingVisibility"`
	// Specifies if the secret backend is local only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#local AwsSecretBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Maximum possible lease duration for secrets in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#max_lease_ttl_seconds AwsSecretBackend#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Number of max retries the client should use for recoverable errors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#max_retries AwsSecretBackend#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#namespace AwsSecretBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies mount type specific options that are passed to the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#options AwsSecretBackend#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// List of headers to allow and pass from the request to the plugin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#passthrough_request_headers AwsSecretBackend#passthrough_request_headers}
	PassthroughRequestHeaders *[]*string `field:"optional" json:"passthroughRequestHeaders" yaml:"passthroughRequestHeaders"`
	// Path to mount the backend at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#path AwsSecretBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Specifies the semantic version of the plugin to use, e.g. 'v1.0.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#plugin_version AwsSecretBackend#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// The AWS region to make API calls against. Defaults to us-east-1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#region AwsSecretBackend#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Role ARN to assume for plugin identity token federation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#role_arn AwsSecretBackend#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#rotation_period AwsSecretBackend#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#rotation_schedule AwsSecretBackend#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#rotation_window AwsSecretBackend#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#seal_wrap AwsSecretBackend#seal_wrap}
	SealWrap interface{} `field:"optional" json:"sealWrap" yaml:"sealWrap"`
	// The AWS Secret Access Key to use when generating new credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#secret_key AwsSecretBackend#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// The AWS Secret Access Key to use when generating new credentials.
	//
	// This is a write-only field and will not be read back from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#secret_key_wo AwsSecretBackend#secret_key_wo}
	SecretKeyWo *string `field:"optional" json:"secretKeyWo" yaml:"secretKeyWo"`
	// A version counter for the write-only secret_key_wo field. Incrementing this value will trigger an update to the secret_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#secret_key_wo_version AwsSecretBackend#secret_key_wo_version}
	SecretKeyWoVersion *float64 `field:"optional" json:"secretKeyWoVersion" yaml:"secretKeyWoVersion"`
	// Specifies a custom HTTP STS endpoint to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#sts_endpoint AwsSecretBackend#sts_endpoint}
	StsEndpoint *string `field:"optional" json:"stsEndpoint" yaml:"stsEndpoint"`
	// Specifies a list of custom STS fallback endpoints to use (in order).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#sts_fallback_endpoints AwsSecretBackend#sts_fallback_endpoints}
	StsFallbackEndpoints *[]*string `field:"optional" json:"stsFallbackEndpoints" yaml:"stsFallbackEndpoints"`
	// Specifies a list of custom STS fallback regions to use (in order).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#sts_fallback_regions AwsSecretBackend#sts_fallback_regions}
	StsFallbackRegions *[]*string `field:"optional" json:"stsFallbackRegions" yaml:"stsFallbackRegions"`
	// Specifies a custom STS region to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#sts_region AwsSecretBackend#sts_region}
	StsRegion *string `field:"optional" json:"stsRegion" yaml:"stsRegion"`
	// Template describing how dynamic usernames are generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/aws_secret_backend#username_template AwsSecretBackend#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

