// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendwrappedkeyimport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendWrappedKeyImportConfig struct {
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
	// HMAC fingerprint of the wrapping key used to encrypt the blob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_wrapped_key_import#export_key_hmac PkiSecretBackendWrappedKeyImport#export_key_hmac}
	ExportKeyHmac *string `field:"required" json:"exportKeyHmac" yaml:"exportKeyHmac"`
	// Path of the destination PKI secrets engine mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_wrapped_key_import#mount PkiSecretBackendWrappedKeyImport#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Base64-encoded wrapped CA private key blob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_wrapped_key_import#wrapped_key PkiSecretBackendWrappedKeyImport#wrapped_key}
	WrappedKey *string `field:"required" json:"wrappedKey" yaml:"wrappedKey"`
	// Human-readable name for the imported key in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_wrapped_key_import#key_name PkiSecretBackendWrappedKeyImport#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_wrapped_key_import#namespace PkiSecretBackendWrappedKeyImport#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

