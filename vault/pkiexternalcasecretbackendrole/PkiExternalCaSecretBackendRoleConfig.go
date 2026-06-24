// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiExternalCaSecretBackendRoleConfig struct {
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
	// The ACME account to use when validating certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#acme_account_name PkiExternalCaSecretBackendRole#acme_account_name}
	AcmeAccountName *string `field:"required" json:"acmeAccountName" yaml:"acmeAccountName"`
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#mount PkiExternalCaSecretBackendRole#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#name PkiExternalCaSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of challenge types that are allowed to be used.
	//
	// Valid values are: `http-01`, `dns-01`, `tls-alpn-01`. Defaults to all challenge types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#allowed_challenge_types PkiExternalCaSecretBackendRole#allowed_challenge_types}
	AllowedChallengeTypes *[]*string `field:"optional" json:"allowedChallengeTypes" yaml:"allowedChallengeTypes"`
	// A list of keyword options that influence how values within allowed_domains are interpreted against the requested set of identifiers from the client.
	//
	// Valid values are: `bare_domains`, `subdomains`, `wildcards`, `globs`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#allowed_domain_options PkiExternalCaSecretBackendRole#allowed_domain_options}
	AllowedDomainOptions *[]*string `field:"optional" json:"allowedDomainOptions" yaml:"allowedDomainOptions"`
	// A list of domains the role will accept certificates for. May contain templates, as with ACL Path Templating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#allowed_domains PkiExternalCaSecretBackendRole#allowed_domains}
	AllowedDomains *[]*string `field:"optional" json:"allowedDomains" yaml:"allowedDomains"`
	// The key type and size/parameters to use when generating a new key if running in the identifier workflow.
	//
	// Valid values are: `ec-256`, `ec-384`, `ec-521`, `rsa-2048`, `rsa-4096`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#csr_generate_key_type PkiExternalCaSecretBackendRole#csr_generate_key_type}
	CsrGenerateKeyType *string `field:"optional" json:"csrGenerateKeyType" yaml:"csrGenerateKeyType"`
	// The technique used to populate a CSR from the provided identifiers in the identifier workflow.
	//
	// Valid values are: `cn_first`, `sans_only`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#csr_identifier_population PkiExternalCaSecretBackendRole#csr_identifier_population}
	CsrIdentifierPopulation *string `field:"optional" json:"csrIdentifierPopulation" yaml:"csrIdentifierPopulation"`
	// Force deletion even when active orders exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#force PkiExternalCaSecretBackendRole#force}
	Force interface{} `field:"optional" json:"force" yaml:"force"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_external_ca_secret_backend_role#namespace PkiExternalCaSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

