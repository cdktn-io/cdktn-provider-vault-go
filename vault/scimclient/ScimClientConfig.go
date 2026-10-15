// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scimclient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ScimClientConfig struct {
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
	// Entity ID authorized to call SCIM protocol endpoints on behalf of this client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#access_grant_principal ScimClient#access_grant_principal}
	AccessGrantPrincipal *string `field:"required" json:"accessGrantPrincipal" yaml:"accessGrantPrincipal"`
	// The name of the SCIM client.
	//
	// Must be lowercase, because Vault lowercases client names when it stores them. May contain only letters, digits, '_', '-' and '.', and must start and end with a letter, digit or '_'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#client_name ScimClient#client_name}
	ClientName *string `field:"required" json:"clientName" yaml:"clientName"`
	// Auth mount accessor used to create entity aliases for SCIM users. Immutable after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#alias_mount_accessor ScimClient#alias_mount_accessor}
	AliasMountAccessor *string `field:"optional" json:"aliasMountAccessor" yaml:"aliasMountAccessor"`
	// Additional mount accessors whose aliases may be managed by this SCIM client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#allowed_extra_alias_mount_accessors ScimClient#allowed_extra_alias_mount_accessors}
	AllowedExtraAliasMountAccessors *[]*string `field:"optional" json:"allowedExtraAliasMountAccessors" yaml:"allowedExtraAliasMountAccessors"`
	// Whether to allow adoption of existing Vault groups by SCIM-provisioned groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#allow_group_adoption ScimClient#allow_group_adoption}
	AllowGroupAdoption interface{} `field:"optional" json:"allowGroupAdoption" yaml:"allowGroupAdoption"`
	// Whether to allow adoption of existing Vault entities by SCIM users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#allow_user_adoption ScimClient#allow_user_adoption}
	AllowUserAdoption interface{} `field:"optional" json:"allowUserAdoption" yaml:"allowUserAdoption"`
	// The Vault SCIM extension schema version used for this client.
	//
	// Must be "2.0" or "2.2". Leave unset to use Vault's default for new clients ("2.2").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#default_schema_version ScimClient#default_schema_version}
	DefaultSchemaVersion *string `field:"optional" json:"defaultSchemaVersion" yaml:"defaultSchemaVersion"`
	// Controls how linked entities/groups/aliases are handled when this resource is destroyed.
	//
	// `orphan_child_resources` detaches managed resources without deleting them; `delete_child_resources` deletes all managed resources along with the client. Leave unset for a plain delete (fails if the client still owns linked resources). Only consulted at destroy time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#deletion_policy ScimClient#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#id ScimClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum number of active tokens for this SCIM client.
	//
	// Must be a positive non-zero value. Leave unset to use Vault's default (2).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#max_active_tokens ScimClient#max_active_tokens}
	MaxActiveTokens *float64 `field:"optional" json:"maxActiveTokens" yaml:"maxActiveTokens"`
	// Maximum TTL (in seconds) for tokens issued to this SCIM client. Zero means the cluster default is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#max_token_ttl ScimClient#max_token_ttl}
	MaxTokenTtl *float64 `field:"optional" json:"maxTokenTtl" yaml:"maxTokenTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#namespace ScimClient#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#timeouts ScimClient#timeouts}
	Timeouts *ScimClientTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

