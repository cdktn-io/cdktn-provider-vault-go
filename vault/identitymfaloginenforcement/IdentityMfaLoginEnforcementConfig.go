// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitymfaloginenforcement

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityMfaLoginEnforcementConfig struct {
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
	// Set of MFA method UUIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#mfa_method_ids IdentityMfaLoginEnforcement#mfa_method_ids}
	MfaMethodIds *[]*string `field:"required" json:"mfaMethodIds" yaml:"mfaMethodIds"`
	// Login enforcement name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#name IdentityMfaLoginEnforcement#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Set of auth method accessor IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#auth_method_accessors IdentityMfaLoginEnforcement#auth_method_accessors}
	AuthMethodAccessors *[]*string `field:"optional" json:"authMethodAccessors" yaml:"authMethodAccessors"`
	// Set of auth method types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#auth_method_types IdentityMfaLoginEnforcement#auth_method_types}
	AuthMethodTypes *[]*string `field:"optional" json:"authMethodTypes" yaml:"authMethodTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#id IdentityMfaLoginEnforcement#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of identity entity IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#identity_entity_ids IdentityMfaLoginEnforcement#identity_entity_ids}
	IdentityEntityIds *[]*string `field:"optional" json:"identityEntityIds" yaml:"identityEntityIds"`
	// Set of identity group IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#identity_group_ids IdentityMfaLoginEnforcement#identity_group_ids}
	IdentityGroupIds *[]*string `field:"optional" json:"identityGroupIds" yaml:"identityGroupIds"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_login_enforcement#namespace IdentityMfaLoginEnforcement#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

