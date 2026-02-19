// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitygroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityGroupConfig struct {
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
	// Manage member entities externally through `vault_identity_group_member_entity_ids`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#external_member_entity_ids IdentityGroup#external_member_entity_ids}
	ExternalMemberEntityIds interface{} `field:"optional" json:"externalMemberEntityIds" yaml:"externalMemberEntityIds"`
	// Manage member groups externally through `vault_identity_group_member_group_ids`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#external_member_group_ids IdentityGroup#external_member_group_ids}
	ExternalMemberGroupIds interface{} `field:"optional" json:"externalMemberGroupIds" yaml:"externalMemberGroupIds"`
	// Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#external_policies IdentityGroup#external_policies}
	ExternalPolicies interface{} `field:"optional" json:"externalPolicies" yaml:"externalPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#id IdentityGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Entity IDs to be assigned as group members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#member_entity_ids IdentityGroup#member_entity_ids}
	MemberEntityIds *[]*string `field:"optional" json:"memberEntityIds" yaml:"memberEntityIds"`
	// Group IDs to be assigned as group members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#member_group_ids IdentityGroup#member_group_ids}
	MemberGroupIds *[]*string `field:"optional" json:"memberGroupIds" yaml:"memberGroupIds"`
	// Metadata to be associated with the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#metadata IdentityGroup#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Name of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#name IdentityGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#namespace IdentityGroup#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Policies to be tied to the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#policies IdentityGroup#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Type of the group, internal or external. Defaults to internal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/identity_group#type IdentityGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

