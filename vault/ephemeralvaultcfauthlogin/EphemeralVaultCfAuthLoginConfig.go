// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultcfauthlogin

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultCfAuthLoginConfig struct {
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
	// The full body of the file available at the path denoted by `CF_INSTANCE_CERT`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#cf_instance_cert EphemeralVaultCfAuthLogin#cf_instance_cert}
	CfInstanceCert *string `field:"required" json:"cfInstanceCert" yaml:"cfInstanceCert"`
	// Name of the CF auth role to log in with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#role EphemeralVaultCfAuthLogin#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The RSA-PSS/SHA256 signature generated using `CF_INSTANCE_KEY` over the concatenation of signing_time, cf_instance_cert, and role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#signature EphemeralVaultCfAuthLogin#signature}
	Signature *string `field:"required" json:"signature" yaml:"signature"`
	// The date and time used to construct the signature (e.g. `2006-01-02T15:04:05Z`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#signing_time EphemeralVaultCfAuthLogin#signing_time}
	SigningTime *string `field:"required" json:"signingTime" yaml:"signingTime"`
	// Mount path for the CF auth engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#mount EphemeralVaultCfAuthLogin#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#mount_id EphemeralVaultCfAuthLogin#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/cf_auth_login#namespace EphemeralVaultCfAuthLogin#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

