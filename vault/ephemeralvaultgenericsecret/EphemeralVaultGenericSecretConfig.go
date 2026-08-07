// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultgenericsecret

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultGenericSecretConfig struct {
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
	// Full path from which a secret will be read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/generic_secret#path EphemeralVaultGenericSecret#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/generic_secret#mount_id EphemeralVaultGenericSecret#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/generic_secret#namespace EphemeralVaultGenericSecret#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Version of the secret to retrieve. Use -1 for latest version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/generic_secret#version EphemeralVaultGenericSecret#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
	// If set to true, stores 'lease_start_time' in the result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/generic_secret#with_lease_start_time EphemeralVaultGenericSecret#with_lease_start_time}
	WithLeaseStartTime interface{} `field:"optional" json:"withLeaseStartTime" yaml:"withLeaseStartTime"`
}

