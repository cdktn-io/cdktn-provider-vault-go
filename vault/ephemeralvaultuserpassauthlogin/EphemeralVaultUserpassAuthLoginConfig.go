// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultuserpassauthlogin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultUserpassAuthLoginConfig struct {
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
	// Mount path for the Userpass auth engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/userpass_auth_login#mount EphemeralVaultUserpassAuthLogin#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Password to log in with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/userpass_auth_login#password EphemeralVaultUserpassAuthLogin#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Username to log in with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/userpass_auth_login#username EphemeralVaultUserpassAuthLogin#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/userpass_auth_login#mount_id EphemeralVaultUserpassAuthLogin#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/ephemeral-resources/userpass_auth_login#namespace EphemeralVaultUserpassAuthLogin#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

