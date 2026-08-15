// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultterraformtoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultTerraformTokenConfig struct {
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
	// Specifies the name of the role to create credentials against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/terraform_token#role_name EphemeralVaultTerraformToken#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Mount path for the Terraform engine in Vault. Default is `terraform`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/terraform_token#mount EphemeralVaultTerraformToken#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/terraform_token#mount_id EphemeralVaultTerraformToken#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/terraform_token#namespace EphemeralVaultTerraformToken#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

