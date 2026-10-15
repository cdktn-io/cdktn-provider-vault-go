// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultgcpkmsreencrypt

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultGcpkmsReencryptConfig struct {
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformEphemeralResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Base64-encoded ciphertext to re-encrypt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#ciphertext EphemeralVaultGcpkmsReencrypt#ciphertext}
	Ciphertext *string `field:"required" json:"ciphertext" yaml:"ciphertext"`
	// Name of the Vault key to use for re-encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#key_name EphemeralVaultGcpkmsReencrypt#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Path where the GCP KMS secrets engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#mount EphemeralVaultGcpkmsReencrypt#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Base64-encoded additional authenticated data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#additional_authenticated_data EphemeralVaultGcpkmsReencrypt#additional_authenticated_data}
	AdditionalAuthenticatedData *string `field:"optional" json:"additionalAuthenticatedData" yaml:"additionalAuthenticatedData"`
	// Version of the key to use for re-encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#key_version EphemeralVaultGcpkmsReencrypt#key_version}
	KeyVersion *float64 `field:"optional" json:"keyVersion" yaml:"keyVersion"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#mount_id EphemeralVaultGcpkmsReencrypt#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/gcpkms_reencrypt#namespace EphemeralVaultGcpkmsReencrypt#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

