// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultkvsecretv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultKvSecretV2Config struct {
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
	// Mount path for the KVV2 engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/kv_secret_v2#mount EphemeralVaultKvSecretV2#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Full name of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/kv_secret_v2#name EphemeralVaultKvSecretV2#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/kv_secret_v2#mount_id EphemeralVaultKvSecretV2#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/kv_secret_v2#namespace EphemeralVaultKvSecretV2#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Version of the secret to retrieve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.11.0/docs/ephemeral-resources/kv_secret_v2#version EphemeralVaultKvSecretV2#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

