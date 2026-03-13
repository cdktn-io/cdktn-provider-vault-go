// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitymfapingid

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityMfaPingidConfig struct {
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
	// A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_pingid#settings_file_base64 IdentityMfaPingid#settings_file_base64}
	SettingsFileBase64 *string `field:"required" json:"settingsFileBase64" yaml:"settingsFileBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_pingid#id IdentityMfaPingid#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_pingid#namespace IdentityMfaPingid#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A template string for mapping Identity names to MFA methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_pingid#username_format IdentityMfaPingid#username_format}
	UsernameFormat *string `field:"optional" json:"usernameFormat" yaml:"usernameFormat"`
}

