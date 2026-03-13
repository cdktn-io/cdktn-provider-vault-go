// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendexternalcaorder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendExternalCaOrderConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order#mount PkiSecretBackendExternalCaOrder#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the role to create the order for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order#role_name PkiSecretBackendExternalCaOrder#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// PEM-encoded Certificate Signing Request containing identifiers. Required if `identifiers` is not provided. Mutually exclusive with `identifiers`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order#csr PkiSecretBackendExternalCaOrder#csr}
	Csr *string `field:"optional" json:"csr" yaml:"csr"`
	// List of identifiers (domain names) for the certificate order. Required if `csr` is not provided. Mutually exclusive with `csr`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order#identifiers PkiSecretBackendExternalCaOrder#identifiers}
	Identifiers *[]*string `field:"optional" json:"identifiers" yaml:"identifiers"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/pki_secret_backend_external_ca_order#namespace PkiSecretBackendExternalCaOrder#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

