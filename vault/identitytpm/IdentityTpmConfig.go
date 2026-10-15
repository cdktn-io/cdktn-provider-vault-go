// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitytpm

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityTpmConfig struct {
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
	// PEM-encoded TPM Endorsement Key (EK) public key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm#tpm_ek_public_key IdentityTpm#tpm_ek_public_key}
	TpmEkPublicKey *string `field:"required" json:"tpmEkPublicKey" yaml:"tpmEkPublicKey"`
	// Whether the TPM is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm#disabled IdentityTpm#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Metadata to associate with the TPM record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm#metadata IdentityTpm#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Name of the TPM record. Vault generates a name if one is not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm#name IdentityTpm#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/identity_tpm#namespace IdentityTpm#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

