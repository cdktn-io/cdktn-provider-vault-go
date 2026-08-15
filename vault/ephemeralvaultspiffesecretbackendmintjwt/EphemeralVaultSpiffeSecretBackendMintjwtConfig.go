// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultspiffesecretbackendmintjwt

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultSpiffeSecretBackendMintjwtConfig struct {
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
	// The audience claim to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/spiffe_secret_backend_mintjwt#audience EphemeralVaultSpiffeSecretBackendMintjwt#audience}
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// Mount path for the SPIFFE secrets engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/spiffe_secret_backend_mintjwt#mount EphemeralVaultSpiffeSecretBackendMintjwt#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the SPIFFE role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/spiffe_secret_backend_mintjwt#name EphemeralVaultSpiffeSecretBackendMintjwt#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/spiffe_secret_backend_mintjwt#mount_id EphemeralVaultSpiffeSecretBackendMintjwt#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/spiffe_secret_backend_mintjwt#namespace EphemeralVaultSpiffeSecretBackendMintjwt#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

