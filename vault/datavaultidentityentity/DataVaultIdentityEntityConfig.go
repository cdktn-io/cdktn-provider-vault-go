// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datavaultidentityentity

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataVaultIdentityEntityConfig struct {
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
	// ID of the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#alias_id DataVaultIdentityEntity#alias_id}
	AliasId *string `field:"optional" json:"aliasId" yaml:"aliasId"`
	// Accessor of the mount to which the alias belongs to. This should be supplied in conjunction with `alias_name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#alias_mount_accessor DataVaultIdentityEntity#alias_mount_accessor}
	AliasMountAccessor *string `field:"optional" json:"aliasMountAccessor" yaml:"aliasMountAccessor"`
	// Name of the alias. This should be supplied in conjunction with `alias_mount_accessor`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#alias_name DataVaultIdentityEntity#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
	// ID of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#entity_id DataVaultIdentityEntity#entity_id}
	EntityId *string `field:"optional" json:"entityId" yaml:"entityId"`
	// Name of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#entity_name DataVaultIdentityEntity#entity_name}
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#id DataVaultIdentityEntity#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/data-sources/identity_entity#namespace DataVaultIdentityEntity#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

