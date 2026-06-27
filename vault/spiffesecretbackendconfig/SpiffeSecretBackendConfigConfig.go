// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spiffesecretbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SpiffeSecretBackendConfigConfig struct {
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
	// Mount path for the SPIFFE secrets engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#mount SpiffeSecretBackendConfig#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The SPIFFE trust domain for this backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#trust_domain SpiffeSecretBackendConfig#trust_domain}
	TrustDomain *string `field:"required" json:"trustDomain" yaml:"trustDomain"`
	// Refresh hint to use in trust bundles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#bundle_refresh_hint SpiffeSecretBackendConfig#bundle_refresh_hint}
	BundleRefreshHint *string `field:"optional" json:"bundleRefreshHint" yaml:"bundleRefreshHint"`
	// Base URL to use for JWT iss claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#jwt_issuer_url SpiffeSecretBackendConfig#jwt_issuer_url}
	JwtIssuerUrl *string `field:"optional" json:"jwtIssuerUrl" yaml:"jwtIssuerUrl"`
	// If true, SPIFFE IDs in JWT SVIDs must not exceed 255 bytes, the limit for the sub claim in OIDC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#jwt_oidc_compatibility_mode SpiffeSecretBackendConfig#jwt_oidc_compatibility_mode}
	JwtOidcCompatibilityMode interface{} `field:"optional" json:"jwtOidcCompatibilityMode" yaml:"jwtOidcCompatibilityMode"`
	// Signing algorithm to use for JWTs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#jwt_signing_algorithm SpiffeSecretBackendConfig#jwt_signing_algorithm}
	JwtSigningAlgorithm *string `field:"optional" json:"jwtSigningAlgorithm" yaml:"jwtSigningAlgorithm"`
	// How long a signing key will live for once it starts being used to sign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#key_lifetime SpiffeSecretBackendConfig#key_lifetime}
	KeyLifetime *string `field:"optional" json:"keyLifetime" yaml:"keyLifetime"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/spiffe_secret_backend_config#namespace SpiffeSecretBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

