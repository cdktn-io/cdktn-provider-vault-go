// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keymgmtdistributekey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KeymgmtDistributeKeyConfig struct {
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
	// Specifies the name of the key to distribute to the given KMS provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/keymgmt_distribute_key#key_name KeymgmtDistributeKey#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// Specifies the name of the KMS provider to distribute the given key to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/keymgmt_distribute_key#kms_name KeymgmtDistributeKey#kms_name}
	KmsName *string `field:"required" json:"kmsName" yaml:"kmsName"`
	// Path of the Key Management secrets engine mount.
	//
	// Must match the `path` of a `vault_mount` resource with `type = "keymgmt"`. Use `vault_mount.keymgmt.path` here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/keymgmt_distribute_key#mount KeymgmtDistributeKey#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Specifies the purpose of the key.
	//
	// The purpose defines a set of cryptographic capabilities that the key will have in the KMS provider. A key must have at least one of the supported purposes. The following values are supported : encrypt, decrypt, sign, verify, wrap, unwrap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/keymgmt_distribute_key#purpose KeymgmtDistributeKey#purpose}
	Purpose *[]*string `field:"required" json:"purpose" yaml:"purpose"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/keymgmt_distribute_key#namespace KeymgmtDistributeKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Specifies the protection of the key.
	//
	// The protection defines where cryptographic operations are performed with the key in the KMS provider. The following values are supported: hsm, software. Defaults to `hsm`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/keymgmt_distribute_key#protection KeymgmtDistributeKey#protection}
	Protection *string `field:"optional" json:"protection" yaml:"protection"`
}

