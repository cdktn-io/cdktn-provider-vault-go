// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityentityalias

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityEntityAliasConfig struct {
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
	// ID of the entity to which this is an alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity_alias#canonical_id IdentityEntityAlias#canonical_id}
	CanonicalId *string `field:"required" json:"canonicalId" yaml:"canonicalId"`
	// Mount accessor to which this alias belongs toMount accessor to which this alias belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity_alias#mount_accessor IdentityEntityAlias#mount_accessor}
	MountAccessor *string `field:"required" json:"mountAccessor" yaml:"mountAccessor"`
	// Name of the entity alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity_alias#name IdentityEntityAlias#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Custom metadata to be associated with this alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity_alias#custom_metadata IdentityEntityAlias#custom_metadata}
	CustomMetadata *map[string]*string `field:"optional" json:"customMetadata" yaml:"customMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity_alias#id IdentityEntityAlias#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_entity_alias#namespace IdentityEntityAlias#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

