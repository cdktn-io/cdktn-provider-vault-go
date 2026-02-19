// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package passwordpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PasswordPolicyConfig struct {
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
	// Name of the password policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/password_policy#name PasswordPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The password policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/password_policy#policy PasswordPolicy#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Specifies an override to the default source of entropy (randomness) used to generate the passwords.
	//
	// Must be one of: '', 'platform', or 'seal'. Requires Vault 1.21+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/password_policy#entropy_source PasswordPolicy#entropy_source}
	EntropySource *string `field:"optional" json:"entropySource" yaml:"entropySource"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/password_policy#namespace PasswordPolicy#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

