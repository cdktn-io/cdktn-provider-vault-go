// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keymgmtkeyrotate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KeymgmtKeyRotateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key_rotate#mount KeymgmtKeyRotate#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Specifies the name of the key to rotate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key_rotate#name KeymgmtKeyRotate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/keymgmt_key_rotate#namespace KeymgmtKeyRotate#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

