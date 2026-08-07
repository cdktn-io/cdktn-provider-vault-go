// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultdatabasesecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultDatabaseSecretConfig struct {
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
	// Mount path for the DB engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/database_secret#mount EphemeralVaultDatabaseSecret#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Specifies the name of the role to create credentials against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/database_secret#name EphemeralVaultDatabaseSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/database_secret#mount_id EphemeralVaultDatabaseSecret#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/database_secret#namespace EphemeralVaultDatabaseSecret#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

