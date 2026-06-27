// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitencrypt

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultTransitEncryptConfig struct {
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
	// The Transit secret backend the key belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#backend DataVaultTransitEncrypt#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the encryption key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#key DataVaultTransitEncrypt#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Map of strings read from Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#plaintext DataVaultTransitEncrypt#plaintext}
	Plaintext *string `field:"required" json:"plaintext" yaml:"plaintext"`
	// Specifies the context for key derivation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#context DataVaultTransitEncrypt#context}
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#id DataVaultTransitEncrypt#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The version of the key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#key_version DataVaultTransitEncrypt#key_version}
	KeyVersion *float64 `field:"optional" json:"keyVersion" yaml:"keyVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/data-sources/transit_encrypt#namespace DataVaultTransitEncrypt#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

