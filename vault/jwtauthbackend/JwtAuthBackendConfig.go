// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package jwtauthbackend

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type JwtAuthBackendConfig struct {
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
	// The value against which to match the iss claim in a JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#bound_issuer JwtAuthBackend#bound_issuer}
	BoundIssuer *string `field:"optional" json:"boundIssuer" yaml:"boundIssuer"`
	// The default role to use if none is provided during login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#default_role JwtAuthBackend#default_role}
	DefaultRole *string `field:"optional" json:"defaultRole" yaml:"defaultRole"`
	// The description of the auth backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#description JwtAuthBackend#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set, opts out of mount migration on path updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#disable_remount JwtAuthBackend#disable_remount}
	DisableRemount interface{} `field:"optional" json:"disableRemount" yaml:"disableRemount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#id JwtAuthBackend#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL.
	//
	// If not set, system certificates are used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#jwks_ca_pem JwtAuthBackend#jwks_ca_pem}
	JwksCaPem *string `field:"optional" json:"jwksCaPem" yaml:"jwksCaPem"`
	// List of JWKS URL and optional CA certificate pairs. Cannot be used with 'jwks_url' or 'jwks_ca_pem'. Requires Vault 1.16+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#jwks_pairs JwtAuthBackend#jwks_pairs}
	JwksPairs interface{} `field:"optional" json:"jwksPairs" yaml:"jwksPairs"`
	// JWKS URL to use to authenticate signatures. Cannot be used with 'oidc_discovery_url' or 'jwt_validation_pubkeys'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#jwks_url JwtAuthBackend#jwks_url}
	JwksUrl *string `field:"optional" json:"jwksUrl" yaml:"jwksUrl"`
	// A list of supported signing algorithms. Defaults to [RS256].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#jwt_supported_algs JwtAuthBackend#jwt_supported_algs}
	JwtSupportedAlgs *[]*string `field:"optional" json:"jwtSupportedAlgs" yaml:"jwtSupportedAlgs"`
	// A list of PEM-encoded public keys to use to authenticate signatures locally.
	//
	// Cannot be used with 'jwks_url' or 'oidc_discovery_url'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#jwt_validation_pubkeys JwtAuthBackend#jwt_validation_pubkeys}
	JwtValidationPubkeys *[]*string `field:"optional" json:"jwtValidationPubkeys" yaml:"jwtValidationPubkeys"`
	// Specifies if the auth method is local only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#local JwtAuthBackend#local}
	Local interface{} `field:"optional" json:"local" yaml:"local"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#namespace JwtAuthBackend#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Pass namespace in the OIDC state parameter instead of as a separate query parameter.
	//
	// With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#namespace_in_state JwtAuthBackend#namespace_in_state}
	NamespaceInState interface{} `field:"optional" json:"namespaceInState" yaml:"namespaceInState"`
	// Client ID used for OIDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_client_id JwtAuthBackend#oidc_client_id}
	OidcClientId *string `field:"optional" json:"oidcClientId" yaml:"oidcClientId"`
	// Client Secret used for OIDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_client_secret JwtAuthBackend#oidc_client_secret}
	OidcClientSecret *string `field:"optional" json:"oidcClientSecret" yaml:"oidcClientSecret"`
	// Write-only Client Secret used for OIDC. This field is recommended over oidc_client_secret for enhanced security.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_client_secret_wo JwtAuthBackend#oidc_client_secret_wo}
	OidcClientSecretWo *string `field:"optional" json:"oidcClientSecretWo" yaml:"oidcClientSecretWo"`
	// Version counter for write-only oidc_client_secret field. Increment this value to force update of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_client_secret_wo_version JwtAuthBackend#oidc_client_secret_wo_version}
	OidcClientSecretWoVersion *float64 `field:"optional" json:"oidcClientSecretWoVersion" yaml:"oidcClientSecretWoVersion"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL.
	//
	// If not set, system certificates are used
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_discovery_ca_pem JwtAuthBackend#oidc_discovery_ca_pem}
	OidcDiscoveryCaPem *string `field:"optional" json:"oidcDiscoveryCaPem" yaml:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used with 'jwks_url' or 'jwt_validation_pubkeys'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_discovery_url JwtAuthBackend#oidc_discovery_url}
	OidcDiscoveryUrl *string `field:"optional" json:"oidcDiscoveryUrl" yaml:"oidcDiscoveryUrl"`
	// The response mode to be used in the OAuth2 request.
	//
	// Allowed values are 'query' and 'form_post'. Defaults to 'query'. If using Vault namespaces, and oidc_response_mode is 'form_post', then 'namespace_in_state' should be set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_response_mode JwtAuthBackend#oidc_response_mode}
	OidcResponseMode *string `field:"optional" json:"oidcResponseMode" yaml:"oidcResponseMode"`
	// The response types to request.
	//
	// Allowed values are 'code' and 'id_token'. Defaults to 'code'. Note: 'id_token' may only be used if 'oidc_response_mode' is set to 'form_post'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#oidc_response_types JwtAuthBackend#oidc_response_types}
	OidcResponseTypes *[]*string `field:"optional" json:"oidcResponseTypes" yaml:"oidcResponseTypes"`
	// path to mount the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#path JwtAuthBackend#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Provider specific handling configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#provider_config JwtAuthBackend#provider_config}
	ProviderConfig *map[string]*string `field:"optional" json:"providerConfig" yaml:"providerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#tune JwtAuthBackend#tune}.
	Tune interface{} `field:"optional" json:"tune" yaml:"tune"`
	// Type of backend. Can be either 'jwt' or 'oidc'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/jwt_auth_backend#type JwtAuthBackend#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

