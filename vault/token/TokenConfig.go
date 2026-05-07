// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package token

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TokenConfig struct {
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
	// The display name of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#display_name Token#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The explicit max TTL of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#explicit_max_ttl Token#explicit_max_ttl}
	ExplicitMaxTtl *string `field:"optional" json:"explicitMaxTtl" yaml:"explicitMaxTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#id Token#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Metadata to be associated with the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#metadata Token#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#namespace Token#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Flag to disable the default policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#no_default_policy Token#no_default_policy}
	NoDefaultPolicy interface{} `field:"optional" json:"noDefaultPolicy" yaml:"noDefaultPolicy"`
	// Flag to create a token without parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#no_parent Token#no_parent}
	NoParent interface{} `field:"optional" json:"noParent" yaml:"noParent"`
	// The number of allowed uses of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#num_uses Token#num_uses}
	NumUses *float64 `field:"optional" json:"numUses" yaml:"numUses"`
	// The period of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#period Token#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
	// List of policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#policies Token#policies}
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Flag to allow the token to be renewed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#renewable Token#renewable}
	Renewable interface{} `field:"optional" json:"renewable" yaml:"renewable"`
	// The renew increment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#renew_increment Token#renew_increment}
	RenewIncrement *float64 `field:"optional" json:"renewIncrement" yaml:"renewIncrement"`
	// The minimum lease to renew token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#renew_min_lease Token#renew_min_lease}
	RenewMinLease *float64 `field:"optional" json:"renewMinLease" yaml:"renewMinLease"`
	// The token role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#role_name Token#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// The TTL period of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#ttl Token#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// The TTL period of the wrapped token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/token#wrapping_ttl Token#wrapping_ttl}
	WrappingTtl *string `field:"optional" json:"wrappingTtl" yaml:"wrappingTtl"`
}

