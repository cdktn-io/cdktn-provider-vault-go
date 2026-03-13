// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityoidcprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityOidcProviderConfig struct {
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
	// The name of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#name IdentityOidcProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The client IDs that are permitted to use the provider.
	//
	// If empty, no clients are allowed. If "*", all clients are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#allowed_client_ids IdentityOidcProvider#allowed_client_ids}
	AllowedClientIds *[]*string `field:"optional" json:"allowedClientIds" yaml:"allowedClientIds"`
	// Set to true if the issuer endpoint uses HTTPS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#https_enabled IdentityOidcProvider#https_enabled}
	HttpsEnabled interface{} `field:"optional" json:"httpsEnabled" yaml:"httpsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#id IdentityOidcProvider#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The host for the issuer. Can be either host or host:port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#issuer_host IdentityOidcProvider#issuer_host}
	IssuerHost *string `field:"optional" json:"issuerHost" yaml:"issuerHost"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#namespace IdentityOidcProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The scopes available for requesting on the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/identity_oidc_provider#scopes_supported IdentityOidcProvider#scopes_supported}
	ScopesSupported *[]*string `field:"optional" json:"scopesSupported" yaml:"scopesSupported"`
}

