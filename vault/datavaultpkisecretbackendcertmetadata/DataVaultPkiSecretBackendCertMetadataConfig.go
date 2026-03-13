// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultpkisecretbackendcertmetadata

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultPkiSecretBackendCertMetadataConfig struct {
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
	// Full path where PKI backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_cert_metadata#path DataVaultPkiSecretBackendCertMetadata#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Specifies the serial of the certificate whose metadata to read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_cert_metadata#serial DataVaultPkiSecretBackendCertMetadata#serial}
	Serial *string `field:"required" json:"serial" yaml:"serial"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_cert_metadata#id DataVaultPkiSecretBackendCertMetadata#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/data-sources/pki_secret_backend_cert_metadata#namespace DataVaultPkiSecretBackendCertMetadata#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

