// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keymgmtkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KeymgmtKeyConfig struct {
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
	// Path of the Key Management secrets engine mount.
	//
	// Must match the `path` of a `vault_mount` resource with `type = "keymgmt"`. Use `vault_mount.keymgmt.path` here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#mount KeymgmtKey#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Specifies the name of the key to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#name KeymgmtKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies if the key is allowed to be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#deletion_allowed KeymgmtKey#deletion_allowed}
	DeletionAllowed interface{} `field:"optional" json:"deletionAllowed" yaml:"deletionAllowed"`
	// Specifies the minimum enabled version of the key.
	//
	// All versions of the key less than the specified version will be disabled for cryptographic operations in the KMS provider that the key has been distributed to. Setting this value to 0 means that all versions will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#min_enabled_version KeymgmtKey#min_enabled_version}
	MinEnabledVersion *float64 `field:"optional" json:"minEnabledVersion" yaml:"minEnabledVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#namespace KeymgmtKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the regions in which the key should be replicated. Supported only for AWS KMS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#replica_regions KeymgmtKey#replica_regions}
	ReplicaRegions *[]*string `field:"optional" json:"replicaRegions" yaml:"replicaRegions"`
	// Specifies the type of cryptographic key to create.
	//
	// aes256-gcm96, rsa-2048, rsa-3072, rsa-4096, ecdsa-p256, ecdsa-p384, ecdsa-p521 key types are supported. Defaults to `rsa-2048`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key#type KeymgmtKey#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

