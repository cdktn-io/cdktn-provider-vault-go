// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralvaultapproleauthbackendrolesecretid

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EphemeralVaultApproleAuthBackendRoleSecretIdConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#role_name EphemeralVaultApproleAuthBackendRoleSecretId#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Unique name of the auth backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#backend EphemeralVaultApproleAuthBackendRoleSecretId#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// List of CIDR blocks that can log in using the SecretID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#cidr_list EphemeralVaultApproleAuthBackendRoleSecretId#cidr_list}
	CidrList *[]*string `field:"optional" json:"cidrList" yaml:"cidrList"`
	// JSON-encoded secret data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#metadata EphemeralVaultApproleAuthBackendRoleSecretId#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Terraform ID of the mount resource. Used to defer the provisioning of the ephemeral resource till the apply stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#mount_id EphemeralVaultApproleAuthBackendRoleSecretId#mount_id}
	MountId *string `field:"optional" json:"mountId" yaml:"mountId"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#namespace EphemeralVaultApproleAuthBackendRoleSecretId#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The number of uses for the secret-id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#num_uses EphemeralVaultApproleAuthBackendRoleSecretId#num_uses}
	NumUses *float64 `field:"optional" json:"numUses" yaml:"numUses"`
	// The TTL duration of the SecretID in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/ephemeral-resources/approle_auth_backend_role_secret_id#ttl EphemeralVaultApproleAuthBackendRoleSecretId#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

