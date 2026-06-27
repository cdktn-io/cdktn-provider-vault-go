// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackendorder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiExternalCaSecretBackendOrderConfig struct {
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
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_order#mount PkiExternalCaSecretBackendOrder#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the role to create the order for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_order#role_name PkiExternalCaSecretBackendOrder#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// PEM-encoded Certificate Signing Request containing identifiers. Required if `identifiers` is not provided. Mutually exclusive with `identifiers`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_order#csr PkiExternalCaSecretBackendOrder#csr}
	Csr *string `field:"optional" json:"csr" yaml:"csr"`
	// List of identifiers (domain names) for the certificate order. Required if `csr` is not provided. Mutually exclusive with `csr`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_order#identifiers PkiExternalCaSecretBackendOrder#identifiers}
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_external_ca_secret_backend_order#namespace PkiExternalCaSecretBackendOrder#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

