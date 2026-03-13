// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitygroupmembergroupids

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityGroupMemberGroupIdsConfig struct {
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
	// ID of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_group_member_group_ids#group_id IdentityGroupMemberGroupIds#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// If set to true, allows the resource to manage member group ids exclusively.
	//
	// Beware of race conditions when disabling exclusive management
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_group_member_group_ids#exclusive IdentityGroupMemberGroupIds#exclusive}
	Exclusive interface{} `field:"optional" json:"exclusive" yaml:"exclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_group_member_group_ids#id IdentityGroupMemberGroupIds#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Group IDs to be assigned as group members.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_group_member_group_ids#member_group_ids IdentityGroupMemberGroupIds#member_group_ids}
	MemberGroupIds *[]*string `field:"optional" json:"memberGroupIds" yaml:"memberGroupIds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_group_member_group_ids#namespace IdentityGroupMemberGroupIds#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

