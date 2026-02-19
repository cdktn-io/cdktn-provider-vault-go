// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mfaduo

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MfaDuoConfig struct {
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
	// API hostname for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#api_hostname MfaDuo#api_hostname}
	ApiHostname *string `field:"required" json:"apiHostname" yaml:"apiHostname"`
	// Integration key for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#integration_key MfaDuo#integration_key}
	IntegrationKey *string `field:"required" json:"integrationKey" yaml:"integrationKey"`
	// The mount to tie this method to for use in automatic mappings.
	//
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#mount_accessor MfaDuo#mount_accessor}
	MountAccessor *string `field:"required" json:"mountAccessor" yaml:"mountAccessor"`
	// Name of the MFA method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#name MfaDuo#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Secret key for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#secret_key MfaDuo#secret_key}
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#id MfaDuo#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#namespace MfaDuo#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Push information for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#push_info MfaDuo#push_info}
	PushInfo *string `field:"optional" json:"pushInfo" yaml:"pushInfo"`
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_duo#username_format MfaDuo#username_format}
	UsernameFormat *string `field:"optional" json:"usernameFormat" yaml:"usernameFormat"`
}

