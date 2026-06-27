// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package identityoidckey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IdentityOidcKeyConfig struct {
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
	// Name of the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#name IdentityOidcKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Signing algorithm to use. Signing algorithm to use. Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#algorithm IdentityOidcKey#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Array of role client ids allowed to use this key for signing.
	//
	// If empty, no roles are allowed. If "*", all roles are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#allowed_client_ids IdentityOidcKey#allowed_client_ids}
	AllowedClientIds *[]*string `field:"optional" json:"allowedClientIds" yaml:"allowedClientIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#id IdentityOidcKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#namespace IdentityOidcKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// How often to generate a new signing key in number of seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#rotation_period IdentityOidcKey#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// Controls how long the public portion of a signing key will be available for verification after being rotated in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/identity_oidc_key#verification_ttl IdentityOidcKey#verification_ttl}
	VerificationTtl *float64 `field:"optional" json:"verificationTtl" yaml:"verificationTtl"`
}

