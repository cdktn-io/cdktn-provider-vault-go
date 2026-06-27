// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackendacmeaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiExternalCaSecretBackendAcmeAccountConfig struct {
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
	// ACME Directory URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#directory_url PkiExternalCaSecretBackendAcmeAccount#directory_url}
	DirectoryUrl *string `field:"required" json:"directoryUrl" yaml:"directoryUrl"`
	// Email addresses for the ACME account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#email_contacts PkiExternalCaSecretBackendAcmeAccount#email_contacts}
	EmailContacts *[]*string `field:"required" json:"emailContacts" yaml:"emailContacts"`
	// The path where the PKI secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#mount PkiExternalCaSecretBackendAcmeAccount#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the ACME account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#name PkiExternalCaSecretBackendAcmeAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An url base64 encoded external binding token to create the initial account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#eab_key PkiExternalCaSecretBackendAcmeAccount#eab_key}
	EabKey *string `field:"optional" json:"eabKey" yaml:"eabKey"`
	// The external binding key ID to create the initial account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#eab_kid PkiExternalCaSecretBackendAcmeAccount#eab_kid}
	EabKid *string `field:"optional" json:"eabKid" yaml:"eabKid"`
	// Key type to generate for the account key. Valid values are `ec-256`, `ec-384`, `ec-521`, `rsa-2048`, `rsa-4096`, `rsa-8192`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#key_type PkiExternalCaSecretBackendAcmeAccount#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#namespace PkiExternalCaSecretBackendAcmeAccount#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Trusted CA certificates for the ACME server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_acme_account#trusted_ca PkiExternalCaSecretBackendAcmeAccount#trusted_ca}
	TrustedCa *string `field:"optional" json:"trustedCa" yaml:"trustedCa"`
}

