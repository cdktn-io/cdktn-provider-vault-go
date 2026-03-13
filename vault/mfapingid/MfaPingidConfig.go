// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mfapingid

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MfaPingidConfig struct {
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
	// The mount to tie this method to for use in automatic mappings.
	//
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/mfa_pingid#mount_accessor MfaPingid#mount_accessor}
	MountAccessor *string `field:"required" json:"mountAccessor" yaml:"mountAccessor"`
	// Name of the MFA method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/mfa_pingid#name MfaPingid#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A base64-encoded third-party settings file retrieved from PingID's configuration page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/mfa_pingid#settings_file_base64 MfaPingid#settings_file_base64}
	SettingsFileBase64 *string `field:"required" json:"settingsFileBase64" yaml:"settingsFileBase64"`
	// ID computed by Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/mfa_pingid#id MfaPingid#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/mfa_pingid#namespace MfaPingid#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/mfa_pingid#username_format MfaPingid#username_format}
	UsernameFormat *string `field:"optional" json:"usernameFormat" yaml:"usernameFormat"`
}

