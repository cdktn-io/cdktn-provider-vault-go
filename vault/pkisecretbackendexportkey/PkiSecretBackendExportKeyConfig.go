// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendexportkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendExportKeyConfig struct {
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
	// Algorithm for the wrapping keypair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_export_key#key_type PkiSecretBackendExportKey#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Path of the PKI secrets engine mount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_export_key#mount PkiSecretBackendExportKey#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Human-readable name for this export key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_export_key#name PkiSecretBackendExportKey#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_secret_backend_export_key#namespace PkiSecretBackendExportKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

