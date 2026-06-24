// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package spiffeauthbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SpiffeAuthBackendConfigConfig struct {
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
	// Mount path for the SPIFFE auth engine in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#mount SpiffeAuthBackendConfig#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// The mechanism to fetch or embed the trust bundle to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#profile SpiffeAuthBackendConfig#profile}
	Profile *string `field:"required" json:"profile" yaml:"profile"`
	// The SPIFFE trust domain for this backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#trust_domain SpiffeAuthBackendConfig#trust_domain}
	TrustDomain *string `field:"required" json:"trustDomain" yaml:"trustDomain"`
	// A list of audience values allowed to match claims in JWT-SVIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#audience SpiffeAuthBackendConfig#audience}
	Audience *[]*string `field:"optional" json:"audience" yaml:"audience"`
	// When profile is 'https_spiffe_bundle', the bootstrapping bundle in SPIFFE format;
	//
	// when profile is 'static', either a bundle in SPIFFE format or PEM-encoded CA certificate(s)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#bundle SpiffeAuthBackendConfig#bundle}
	Bundle *string `field:"optional" json:"bundle" yaml:"bundle"`
	// Don't attempt to fetch a bundle immediately; only applies when profile != static.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#defer_bundle_fetch SpiffeAuthBackendConfig#defer_bundle_fetch}
	DeferBundleFetch interface{} `field:"optional" json:"deferBundleFetch" yaml:"deferBundleFetch"`
	// PEM-encoded CA certificate(s) to validate the server reached by 'endpoint_url', if set this will override the default TLS trust store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#endpoint_root_ca_truststore_pem SpiffeAuthBackendConfig#endpoint_root_ca_truststore_pem}
	EndpointRootCaTruststorePem *string `field:"optional" json:"endpointRootCaTruststorePem" yaml:"endpointRootCaTruststorePem"`
	// The server's SPIFFE ID to validate when profile is 'https_spiffe_bundle'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#endpoint_spiffe_id SpiffeAuthBackendConfig#endpoint_spiffe_id}
	EndpointSpiffeId *string `field:"optional" json:"endpointSpiffeId" yaml:"endpointSpiffeId"`
	// The URI to be used when profile is 'https_web_bundle' or 'https_spiffe_bundle'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#endpoint_url SpiffeAuthBackendConfig#endpoint_url}
	EndpointUrl *string `field:"optional" json:"endpointUrl" yaml:"endpointUrl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/spiffe_auth_backend_config#namespace SpiffeAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

