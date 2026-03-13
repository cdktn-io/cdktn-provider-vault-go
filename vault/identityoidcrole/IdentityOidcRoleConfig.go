// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityoidcrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityOidcRoleConfig struct {
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
	// A configured named key, the key must already exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#key IdentityOidcRole#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#name IdentityOidcRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value that will be included in the `aud` field of all the OIDC identity tokens issued by this role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#client_id IdentityOidcRole#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#id IdentityOidcRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#namespace IdentityOidcRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The template string to use for generating tokens. This may be in string-ified JSON or base64 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#template IdentityOidcRole#template}
	Template *string `field:"optional" json:"template" yaml:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_role#ttl IdentityOidcRole#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

