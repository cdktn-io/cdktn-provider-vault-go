// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesauthbackendconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KubernetesAuthBackendConfigConfig struct {
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
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#kubernetes_host KubernetesAuthBackendConfig#kubernetes_host}
	KubernetesHost *string `field:"required" json:"kubernetesHost" yaml:"kubernetesHost"`
	// Unique name of the kubernetes backend to configure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#backend KubernetesAuthBackendConfig#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Optional disable JWT issuer validation. Allows to skip ISS validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#disable_iss_validation KubernetesAuthBackendConfig#disable_iss_validation}
	DisableIssValidation interface{} `field:"optional" json:"disableIssValidation" yaml:"disableIssValidation"`
	// Optional disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#disable_local_ca_jwt KubernetesAuthBackendConfig#disable_local_ca_jwt}
	DisableLocalCaJwt interface{} `field:"optional" json:"disableLocalCaJwt" yaml:"disableLocalCaJwt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#id KubernetesAuthBackendConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional JWT issuer. If no issuer is specified, kubernetes.io/serviceaccount will be used as the default issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#issuer KubernetesAuthBackendConfig#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#kubernetes_ca_cert KubernetesAuthBackendConfig#kubernetes_ca_cert}
	KubernetesCaCert *string `field:"optional" json:"kubernetesCaCert" yaml:"kubernetesCaCert"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#namespace KubernetesAuthBackendConfig#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs.
	//
	// If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#pem_keys KubernetesAuthBackendConfig#pem_keys}
	PemKeys *[]*string `field:"optional" json:"pemKeys" yaml:"pemKeys"`
	// A service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login.
	//
	// If not set the JWT used for login will be used to access the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#token_reviewer_jwt KubernetesAuthBackendConfig#token_reviewer_jwt}
	TokenReviewerJwt *string `field:"optional" json:"tokenReviewerJwt" yaml:"tokenReviewerJwt"`
	// A write-only service account JWT (or other token) used as a bearer token to access the TokenReview API to validate other JWTs during login.
	//
	// If not set the JWT used for login will be used to access the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#token_reviewer_jwt_wo KubernetesAuthBackendConfig#token_reviewer_jwt_wo}
	TokenReviewerJwtWo *string `field:"optional" json:"tokenReviewerJwtWo" yaml:"tokenReviewerJwtWo"`
	// The version of token_reviewer_jwt_wo to use during write operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#token_reviewer_jwt_wo_version KubernetesAuthBackendConfig#token_reviewer_jwt_wo_version}
	TokenReviewerJwtWoVersion *float64 `field:"optional" json:"tokenReviewerJwtWoVersion" yaml:"tokenReviewerJwtWoVersion"`
	// Use annotations from the client token's associated service account as alias metadata for the Vault entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/kubernetes_auth_backend_config#use_annotations_as_alias_metadata KubernetesAuthBackendConfig#use_annotations_as_alias_metadata}
	UseAnnotationsAsAliasMetadata interface{} `field:"optional" json:"useAnnotationsAsAliasMetadata" yaml:"useAnnotationsAsAliasMetadata"`
}

