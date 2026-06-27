// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthresourceserverconfigprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OauthResourceServerConfigProfileConfig struct {
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
	// The issuer ID (iss claim) to validate against in incoming JWTs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#issuer_id OauthResourceServerConfigProfile#issuer_id}
	IssuerId *string `field:"required" json:"issuerId" yaml:"issuerId"`
	// The name of the OAuth Resource Server Configuration profile. Must be unique within the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#profile_name OauthResourceServerConfigProfile#profile_name}
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// List of allowed audiences (aud claim) to validate in JWTs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#audiences OauthResourceServerConfigProfile#audiences}
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
	// Leeway for clock skew in seconds when validating time-based claims. Defaults to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#clock_skew_leeway OauthResourceServerConfigProfile#clock_skew_leeway}
	ClockSkewLeeway *float64 `field:"optional" json:"clockSkewLeeway" yaml:"clockSkewLeeway"`
	// Whether this profile is enabled for JWT validation. Disabled profiles are ignored. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#enabled OauthResourceServerConfigProfile#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Optional CA certificate (PEM format) for JWKS URI TLS validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#jwks_ca_pem OauthResourceServerConfigProfile#jwks_ca_pem}
	JwksCaPem *string `field:"optional" json:"jwksCaPem" yaml:"jwksCaPem"`
	// The JWKS URI to fetch public keys from. Required when use_jwks=true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#jwks_uri OauthResourceServerConfigProfile#jwks_uri}
	JwksUri *string `field:"optional" json:"jwksUri" yaml:"jwksUri"`
	// The JWT type: 'access_token' or 'transaction_token'. Defaults to 'access_token'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#jwt_type OauthResourceServerConfigProfile#jwt_type}
	JwtType *string `field:"optional" json:"jwtType" yaml:"jwtType"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#namespace OauthResourceServerConfigProfile#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// If true, JWT-authenticated tokens omit the default policy unless added elsewhere. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#no_default_policy OauthResourceServerConfigProfile#no_default_policy}
	NoDefaultPolicy interface{} `field:"optional" json:"noDefaultPolicy" yaml:"noDefaultPolicy"`
	// When false, RAR (Rich Authorization Requests) is mandatory and authorization_details must be present in the token.
	//
	// When set to true, authorization_details in the JWT token are optional. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#optional_authorization_details OauthResourceServerConfigProfile#optional_authorization_details}
	OptionalAuthorizationDetails interface{} `field:"optional" json:"optionalAuthorizationDetails" yaml:"optionalAuthorizationDetails"`
	// public_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#public_keys OauthResourceServerConfigProfile#public_keys}
	PublicKeys interface{} `field:"optional" json:"publicKeys" yaml:"publicKeys"`
	// List of supported signing algorithms (e.g., RS256, ES256). Defaults to all supported algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#supported_algorithms OauthResourceServerConfigProfile#supported_algorithms}
	SupportedAlgorithms *[]*string `field:"optional" json:"supportedAlgorithms" yaml:"supportedAlgorithms"`
	// If true, use JWKS URI for key validation; if false, use static public keys. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#use_jwks OauthResourceServerConfigProfile#use_jwks}
	UseJwks interface{} `field:"optional" json:"useJwks" yaml:"useJwks"`
	// The claim to use as the user identifier. Defaults to 'sub'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/oauth_resource_server_config_profile#user_claim OauthResourceServerConfigProfile#user_claim}
	UserClaim *string `field:"optional" json:"userClaim" yaml:"userClaim"`
}

