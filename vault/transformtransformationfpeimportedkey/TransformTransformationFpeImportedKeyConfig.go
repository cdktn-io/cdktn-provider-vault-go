// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transformtransformationfpeimportedkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransformTransformationFpeImportedKeyConfig struct {
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
	// Path where the Transform secret backend is mounted. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#mount TransformTransformationFpeImportedKey#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the transformation to create. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#name TransformTransformationFpeImportedKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the template to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#template TransformTransformationFpeImportedKey#template}
	Template *string `field:"required" json:"template" yaml:"template"`
	// The list of roles allowed to perform this transformation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#allowed_roles TransformTransformationFpeImportedKey#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// A base64-encoded string containing the wrapped ephemeral AES key and the encrypted 256-bit AES key material (write-only).
	//
	// Required with `ciphertext_wo_version`. Required to create the transformation, but may be omitted for a transformation adopted with `terraform import`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#ciphertext_wo TransformTransformationFpeImportedKey#ciphertext_wo}
	CiphertextWo *string `field:"optional" json:"ciphertextWo" yaml:"ciphertextWo"`
	// Version identifier for `ciphertext_wo`.
	//
	// Vault imports ciphertext on create; changing this value replaces the transformation. Removing it does not replace the resource. Required with `ciphertext_wo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#ciphertext_wo_version TransformTransformationFpeImportedKey#ciphertext_wo_version}
	CiphertextWoVersion *float64 `field:"optional" json:"ciphertextWoVersion" yaml:"ciphertextWoVersion"`
	// If true, this transformation can be deleted.
	//
	// Deletion is blocked while this is false, so it must be applied as `true` before the resource can be destroyed. Vault's import endpoint does not accept this parameter, so it is applied in an update immediately after the key is imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#deletion_allowed TransformTransformationFpeImportedKey#deletion_allowed}
	DeletionAllowed interface{} `field:"optional" json:"deletionAllowed" yaml:"deletionAllowed"`
	// The hash function used for the RSA-OAEP step of creating the ciphertext (`SHA1`, `SHA224`, `SHA256`, `SHA384`, `SHA512`).
	//
	// Vault uses `SHA256` if not set. Vault does not return this value, so it is not verified against the transformation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#hash_function TransformTransformationFpeImportedKey#hash_function}
	HashFunction *string `field:"optional" json:"hashFunction" yaml:"hashFunction"`
	// The maximum tweak length in bytes (0-32).
	//
	// If the value is 0, the tweak length can be any value from 0 to 32 bytes. Vault does not allow this field to be changed after creation, so changing it returns an error from Vault. Requires Vault 2.2.0 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#max_tweak_len TransformTransformationFpeImportedKey#max_tweak_len}
	MaxTweakLen *float64 `field:"optional" json:"maxTweakLen" yaml:"maxTweakLen"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#namespace TransformTransformationFpeImportedKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The source of the tweak value (`supplied`, `generated` or `internal`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transform_transformation_fpe_imported_key#tweak_source TransformTransformationFpeImportedKey#tweak_source}
	TweakSource *string `field:"optional" json:"tweakSource" yaml:"tweakSource"`
}

