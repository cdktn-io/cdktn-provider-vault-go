// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkiexternalcasecretbackenddnsprovidergcp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiExternalCaSecretBackendDnsProviderGcpConfig struct {
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
	// List of domain identifiers this provider handles. Supports wildcard patterns with leftmost `*` (e.g. `*.example.com`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#identifiers PkiExternalCaSecretBackendDnsProviderGcp#identifiers}
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
	// The path where the PKI External CA secret backend is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#mount PkiExternalCaSecretBackendDnsProviderGcp#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the DNS provider configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#name PkiExternalCaSecretBackendDnsProviderGcp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// GCP service account credentials as JSON content. Write-only — not returned by Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#credentials_wo PkiExternalCaSecretBackendDnsProviderGcp#credentials_wo}
	CredentialsWo *string `field:"optional" json:"credentialsWo" yaml:"credentialsWo"`
	// Version counter for the write-only `credentials` field. Increment this value to trigger an update to the credentials in Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#credentials_wo_version PkiExternalCaSecretBackendDnsProviderGcp#credentials_wo_version}
	CredentialsWoVersion *float64 `field:"optional" json:"credentialsWoVersion" yaml:"credentialsWoVersion"`
	// Service account email to impersonate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#impersonate_service_account PkiExternalCaSecretBackendDnsProviderGcp#impersonate_service_account}
	ImpersonateServiceAccount *string `field:"optional" json:"impersonateServiceAccount" yaml:"impersonateServiceAccount"`
	// Address of a DNS nameserver (`host` or `host:port`) to use when verifying DNS-01 challenge propagation instead of the domain's primary nameserver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#nameserver PkiExternalCaSecretBackendDnsProviderGcp#nameserver}
	Nameserver *string `field:"optional" json:"nameserver" yaml:"nameserver"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#namespace PkiExternalCaSecretBackendDnsProviderGcp#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// GCP project name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#project PkiExternalCaSecretBackendDnsProviderGcp#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// TTL for DNS TXT records used in DNS-01 challenges in seconds. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#ttl PkiExternalCaSecretBackendDnsProviderGcp#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// GCP Cloud DNS zone name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/pki_external_ca_secret_backend_dns_provider_gcp#zone_name PkiExternalCaSecretBackendDnsProviderGcp#zone_name}
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}

