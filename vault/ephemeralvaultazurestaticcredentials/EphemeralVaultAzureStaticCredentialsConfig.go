// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultazurestaticcredentials

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultAzureStaticCredentialsConfig struct {
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
	// Azure Secret Backend to read credentials from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/azure_static_credentials#backend EphemeralVaultAzureStaticCredentials#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Static role name to fetch credentials for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/azure_static_credentials#role EphemeralVaultAzureStaticCredentials#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/azure_static_credentials#mount_id EphemeralVaultAzureStaticCredentials#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/azure_static_credentials#namespace EphemeralVaultAzureStaticCredentials#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Input metadata to send with the request to Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/azure_static_credentials#request_metadata EphemeralVaultAzureStaticCredentials#request_metadata}
	RequestMetadata *map[string]*string `field:"optional" json:"requestMetadata" yaml:"requestMetadata"`
}

