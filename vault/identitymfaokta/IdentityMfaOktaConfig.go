// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitymfaokta

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityMfaOktaConfig struct {
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
	// Okta API token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#api_token IdentityMfaOkta#api_token}
	ApiToken *string `field:"required" json:"apiToken" yaml:"apiToken"`
	// Name of the organization to be used in the Okta API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#org_name IdentityMfaOkta#org_name}
	OrgName *string `field:"required" json:"orgName" yaml:"orgName"`
	// The base domain to use for API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#base_url IdentityMfaOkta#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#id IdentityMfaOkta#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#namespace IdentityMfaOkta#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Only match the primary email for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#primary_email IdentityMfaOkta#primary_email}
	PrimaryEmail interface{} `field:"optional" json:"primaryEmail" yaml:"primaryEmail"`
	// A template string for mapping Identity names to MFA methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_mfa_okta#username_format IdentityMfaOkta#username_format}
	UsernameFormat *string `field:"optional" json:"usernameFormat" yaml:"usernameFormat"`
}

