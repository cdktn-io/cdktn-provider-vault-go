// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityoidcclient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityOidcClientConfig struct {
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
	// The name of the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#name IdentityOidcClient#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time-to-live for access tokens obtained by the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#access_token_ttl IdentityOidcClient#access_token_ttl}
	AccessTokenTtl *float64 `field:"optional" json:"accessTokenTtl" yaml:"accessTokenTtl"`
	// A list of assignment resources associated with the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#assignments IdentityOidcClient#assignments}
	Assignments *[]*string `field:"optional" json:"assignments" yaml:"assignments"`
	// The client type based on its ability to maintain confidentiality of credentials.Defaults to 'confidential'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#client_type IdentityOidcClient#client_type}
	ClientType *string `field:"optional" json:"clientType" yaml:"clientType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#id IdentityOidcClient#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The time-to-live for ID tokens obtained by the client.
	//
	// The value should be less than the verification_ttl on the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#id_token_ttl IdentityOidcClient#id_token_ttl}
	IdTokenTtl *float64 `field:"optional" json:"idTokenTtl" yaml:"idTokenTtl"`
	// A reference to a named key resource in Vault. This cannot be modified after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#key IdentityOidcClient#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#namespace IdentityOidcClient#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Redirection URI values used by the client.
	//
	// One of these values must exactly match the redirect_uri parameter value used in each authentication request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_client#redirect_uris IdentityOidcClient#redirect_uris}
	RedirectUris *[]*string `field:"optional" json:"redirectUris" yaml:"redirectUris"`
}

