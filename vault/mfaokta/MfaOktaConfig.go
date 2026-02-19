// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mfaokta

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MfaOktaConfig struct {
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
	// Okta API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#api_token MfaOkta#api_token}
	ApiToken *string `field:"required" json:"apiToken" yaml:"apiToken"`
	// The mount to tie this method to for use in automatic mappings.
	//
	// The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#mount_accessor MfaOkta#mount_accessor}
	MountAccessor *string `field:"required" json:"mountAccessor" yaml:"mountAccessor"`
	// Name of the MFA method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#name MfaOkta#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the organization to be used in the Okta API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#org_name MfaOkta#org_name}
	OrgName *string `field:"required" json:"orgName" yaml:"orgName"`
	// If set, will be used as the base domain for API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#base_url MfaOkta#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// ID computed by Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#id MfaOkta#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#namespace MfaOkta#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// If set to true, the username will only match the primary email for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#primary_email MfaOkta#primary_email}
	PrimaryEmail interface{} `field:"optional" json:"primaryEmail" yaml:"primaryEmail"`
	// A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/mfa_okta#username_format MfaOkta#username_format}
	UsernameFormat *string `field:"optional" json:"usernameFormat" yaml:"usernameFormat"`
}

