// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaulttransitdecrypt

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultTransitDecryptConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#backend DataVaultTransitDecrypt#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Transit encrypted cipher text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#ciphertext DataVaultTransitDecrypt#ciphertext}
	Ciphertext *string `field:"required" json:"ciphertext" yaml:"ciphertext"`
	// Name of the decryption key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#key DataVaultTransitDecrypt#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Specifies the context for key derivation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#context DataVaultTransitDecrypt#context}
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Specifies the hash algorithm to use for RSA key decryption.
	//
	// Only applies to RSA key types; ignored for all other key types. Supported values are: sha1, sha2-224, sha2-256, sha2-384, sha2-512, sha3-224, sha3-256, sha3-384, sha3-512. If not set, Vault defaults to sha2-256 for RSA keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#hash_algorithm DataVaultTransitDecrypt#hash_algorithm}
	HashAlgorithm *string `field:"optional" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#id DataVaultTransitDecrypt#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#namespace DataVaultTransitDecrypt#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the RSA padding scheme to use for decryption.
	//
	// Only applies to RSA key types; ignored for all other key types. Supported values are: oaep, pkcs1v15. If not set, Vault defaults to oaep for RSA keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/data-sources/transit_decrypt#padding_scheme DataVaultTransitDecrypt#padding_scheme}
	PaddingScheme *string `field:"optional" json:"paddingScheme" yaml:"paddingScheme"`
}

