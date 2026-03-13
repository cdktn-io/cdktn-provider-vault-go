// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identitymfaduo

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityMfaDuoConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#api_hostname IdentityMfaDuo#api_hostname}
	ApiHostname *string `field:"required" json:"apiHostname" yaml:"apiHostname"`
	// Integration key for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#integration_key IdentityMfaDuo#integration_key}
	IntegrationKey *string `field:"required" json:"integrationKey" yaml:"integrationKey"`
	// Secret key for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#secret_key IdentityMfaDuo#secret_key}
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#id IdentityMfaDuo#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#namespace IdentityMfaDuo#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Push information for Duo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#push_info IdentityMfaDuo#push_info}
	PushInfo *string `field:"optional" json:"pushInfo" yaml:"pushInfo"`
	// Require passcode upon MFA validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#use_passcode IdentityMfaDuo#use_passcode}
	UsePasscode interface{} `field:"optional" json:"usePasscode" yaml:"usePasscode"`
	// A template string for mapping Identity names to MFA methods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_mfa_duo#username_format IdentityMfaDuo#username_format}
	UsernameFormat *string `field:"optional" json:"usernameFormat" yaml:"usernameFormat"`
}

