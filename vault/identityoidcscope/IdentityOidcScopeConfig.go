// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityoidcscope

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityOidcScopeConfig struct {
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
	// The name of the scope. The openid scope name is reserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_oidc_scope#name IdentityOidcScope#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The scope's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_oidc_scope#description IdentityOidcScope#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_oidc_scope#id IdentityOidcScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_oidc_scope#namespace IdentityOidcScope#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/identity_oidc_scope#template IdentityOidcScope#template}
	Template *string `field:"optional" json:"template" yaml:"template"`
}

