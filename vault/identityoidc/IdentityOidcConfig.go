// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityoidc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityOidcConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_oidc#id IdentityOidc#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Issuer URL to be used in the iss claim of the token.
	//
	// If not set, Vault's api_addr will be used. The issuer is a case sensitive URL using the https scheme that contains scheme, host, and optionally, port number and path components, but no query or fragment components.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_oidc#issuer IdentityOidc#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/identity_oidc#namespace IdentityOidc#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

