// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentRegistrationConfig struct {
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
	// Human-readable name for the agent registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#display_name AgentRegistration#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Entity ID representing this agent. Each entity can only have one registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#entity_id AgentRegistration#entity_id}
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// List of policy identifiers that define the authorization ceiling for this agent.
	//
	// Cannot include 'root' policy. Note: Vault automatically adds default policies (['default', 'default-ceiling']) unless no_default_ceiling_policy is true, but these are filtered out when reading the state to match your configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#ceiling_policies AgentRegistration#ceiling_policies}
	CeilingPolicies *[]*string `field:"optional" json:"ceilingPolicies" yaml:"ceilingPolicies"`
	// Detailed description of the agent's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#description AgentRegistration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#namespace AgentRegistration#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// If true, opts out of automatically adding the default-ceiling policy to this agent registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#no_default_ceiling_policy AgentRegistration#no_default_ceiling_policy}
	NoDefaultCeilingPolicy interface{} `field:"optional" json:"noDefaultCeilingPolicy" yaml:"noDefaultCeilingPolicy"`
	// When false, RAR (Rich Authorization Requests) is mandatory and authorization_details must be present in the token.
	//
	// When set to true, authorization_details in the JWT token are optional for this agent. This setting works in conjunction with the OAuth Resource Server profile's optional_authorization_details setting - RAR is optional if EITHER is true. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#optional_authorization_details AgentRegistration#optional_authorization_details}
	OptionalAuthorizationDetails interface{} `field:"optional" json:"optionalAuthorizationDetails" yaml:"optionalAuthorizationDetails"`
	// Owner of the agent registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/agent_registration#owner AgentRegistration#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}

