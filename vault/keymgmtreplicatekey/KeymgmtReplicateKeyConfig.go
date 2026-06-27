// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keymgmtreplicatekey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KeymgmtReplicateKeyConfig struct {
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
	// Specifies the name of the key to replicate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_replicate_key#key_name KeymgmtReplicateKey#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Specifies the name of the AWS KMS provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_replicate_key#kms_name KeymgmtReplicateKey#kms_name}
	KmsName *string `field:"required" json:"kmsName" yaml:"kmsName"`
	// Path of the Key Management secrets engine mount.
	//
	// Must match the `path` of a `vault_mount` resource with `type = "keymgmt"`. Use `vault_mount.keymgmt.path` here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_replicate_key#mount KeymgmtReplicateKey#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_replicate_key#namespace KeymgmtReplicateKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

