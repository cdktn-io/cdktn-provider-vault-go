// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultsshsecretbackendsign

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultSshSecretBackendSignConfig struct {
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
	// Specifies the name of the role to sign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#name DataVaultSshSecretBackendSign#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Full path where SSH backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#path DataVaultSshSecretBackendSign#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Specifies the SSH public key that should be signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#public_key DataVaultSshSecretBackendSign#public_key}
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Specifies the type of certificate to be created; either "user" or "host".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#cert_type DataVaultSshSecretBackendSign#cert_type}
	CertType *string `field:"optional" json:"certType" yaml:"certType"`
	// Specifies a map of the critical options that the certificate should be signed for. Defaults to none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#critical_options DataVaultSshSecretBackendSign#critical_options}
	CriticalOptions *map[string]*string `field:"optional" json:"criticalOptions" yaml:"criticalOptions"`
	// Specifies a map of the extensions that the certificate should be signed for. Defaults to none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#extensions DataVaultSshSecretBackendSign#extensions}
	Extensions *map[string]*string `field:"optional" json:"extensions" yaml:"extensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#id DataVaultSshSecretBackendSign#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the key id that the created certificate should have.
	//
	// If not specified, the display name of the token will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#key_id DataVaultSshSecretBackendSign#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#namespace DataVaultSshSecretBackendSign#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the Requested Time To Live.
	//
	// Cannot be greater than the role's max_ttl value. If not provided, the role's ttl value will be used. Note that the role values default to system values if not explicitly set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#ttl DataVaultSshSecretBackendSign#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Specifies valid principals, either usernames or hostnames, that the certificate should be signed for.
	//
	// Required unless the role has specified allow_empty_principals or a value has been set for either the default_user or default_user_template role parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/ssh_secret_backend_sign#valid_principals DataVaultSshSecretBackendSign#valid_principals}
	ValidPrincipals *string `field:"optional" json:"validPrincipals" yaml:"validPrincipals"`
}

