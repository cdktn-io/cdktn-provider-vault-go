// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package azuresecretbackendstaticrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AzureSecretBackendStaticRoleConfig struct {
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
	// Application object ID for an existing service principal that is managed by the static role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#application_object_id AzureSecretBackendStaticRole#application_object_id}
	ApplicationObjectId *string `field:"required" json:"applicationObjectId" yaml:"applicationObjectId"`
	// The path where the Azure secrets backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#backend AzureSecretBackendStaticRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the static role to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#role AzureSecretBackendStaticRole#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The plaintext secret value of the credential you want to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#client_secret AzureSecretBackendStaticRole#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// If true, the initial creation of credentials will be deferred until first static-creds read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#defer_initial_creds AzureSecretBackendStaticRole#defer_initial_creds}
	DeferInitialCreds interface{} `field:"optional" json:"deferInitialCreds" yaml:"deferInitialCreds"`
	// A future expiration time for the imported credential, in RFC3339 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#expiration AzureSecretBackendStaticRole#expiration}
	Expiration *string `field:"optional" json:"expiration" yaml:"expiration"`
	// A map of string key/value pairs that will be stored as metadata on the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#metadata AzureSecretBackendStaticRole#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#namespace AzureSecretBackendStaticRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The secret ID of the Azure password credential you want to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#secret_id AzureSecretBackendStaticRole#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// If true, skip rotation of the client secret on import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#skip_import_rotation AzureSecretBackendStaticRole#skip_import_rotation}
	SkipImportRotation interface{} `field:"optional" json:"skipImportRotation" yaml:"skipImportRotation"`
	// Timespan of 1 month or more during which the role credentials are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/azure_secret_backend_static_role#ttl AzureSecretBackendStaticRole#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

