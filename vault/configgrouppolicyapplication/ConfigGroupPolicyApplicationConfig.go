// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configgrouppolicyapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigGroupPolicyApplicationConfig struct {
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
	// Mode for group policy application.
	//
	// Must be either "within_namespace_hierarchy" or "any". "within_namespace_hierarchy" means policies only apply when the token authorizing a request was created in the same namespace as the group, or a descendant namespace. "any" means group policies apply to all members of a group, regardless of what namespace the request token came from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_group_policy_application#group_policy_application_mode ConfigGroupPolicyApplication#group_policy_application_mode}
	GroupPolicyApplicationMode *string `field:"required" json:"groupPolicyApplicationMode" yaml:"groupPolicyApplicationMode"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/config_group_policy_application#namespace ConfigGroupPolicyApplication#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

