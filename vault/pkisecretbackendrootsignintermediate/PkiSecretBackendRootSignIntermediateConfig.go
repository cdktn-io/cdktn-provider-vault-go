// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrootsignintermediate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendRootSignIntermediateConfig struct {
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
	// The PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#backend PkiSecretBackendRootSignIntermediate#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of intermediate to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#common_name PkiSecretBackendRootSignIntermediate#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// The CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#csr PkiSecretBackendRootSignIntermediate#csr}
	Csr *string `field:"required" json:"csr" yaml:"csr"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#alt_names PkiSecretBackendRootSignIntermediate#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// The country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#country PkiSecretBackendRootSignIntermediate#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#exclude_cn_from_sans PkiSecretBackendRootSignIntermediate#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// List of domains for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#excluded_dns_domains PkiSecretBackendRootSignIntermediate#excluded_dns_domains}
	ExcludedDnsDomains *[]*string `field:"optional" json:"excludedDnsDomains" yaml:"excludedDnsDomains"`
	// List of email addresses for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#excluded_email_addresses PkiSecretBackendRootSignIntermediate#excluded_email_addresses}
	ExcludedEmailAddresses *[]*string `field:"optional" json:"excludedEmailAddresses" yaml:"excludedEmailAddresses"`
	// List of IP ranges for which certificates are NOT allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#excluded_ip_ranges PkiSecretBackendRootSignIntermediate#excluded_ip_ranges}
	ExcludedIpRanges *[]*string `field:"optional" json:"excludedIpRanges" yaml:"excludedIpRanges"`
	// List of URI domains for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#excluded_uri_domains PkiSecretBackendRootSignIntermediate#excluded_uri_domains}
	ExcludedUriDomains *[]*string `field:"optional" json:"excludedUriDomains" yaml:"excludedUriDomains"`
	// The format of data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#format PkiSecretBackendRootSignIntermediate#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#id PkiSecretBackendRootSignIntermediate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#ip_sans PkiSecretBackendRootSignIntermediate#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// Specifies the default issuer of this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#issuer_ref PkiSecretBackendRootSignIntermediate#issuer_ref}
	IssuerRef *string `field:"optional" json:"issuerRef" yaml:"issuerRef"`
	// Specify the key usages to be added to the existing set of key usages ("CRL", "CertSign") on the generated certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#key_usage PkiSecretBackendRootSignIntermediate#key_usage}
	KeyUsage *[]*string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The locality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#locality PkiSecretBackendRootSignIntermediate#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The maximum path length to encode in the generated certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#max_path_length PkiSecretBackendRootSignIntermediate#max_path_length}
	MaxPathLength *float64 `field:"optional" json:"maxPathLength" yaml:"maxPathLength"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#namespace PkiSecretBackendRootSignIntermediate#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set the Not After field of the certificate with specified date value.
	//
	// The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#not_after PkiSecretBackendRootSignIntermediate#not_after}
	NotAfter *string `field:"optional" json:"notAfter" yaml:"notAfter"`
	// Specifies the duration by which to backdate the NotBefore property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#not_before_duration PkiSecretBackendRootSignIntermediate#not_before_duration}
	NotBeforeDuration *string `field:"optional" json:"notBeforeDuration" yaml:"notBeforeDuration"`
	// The organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#organization PkiSecretBackendRootSignIntermediate#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#other_sans PkiSecretBackendRootSignIntermediate#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// The organization unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#ou PkiSecretBackendRootSignIntermediate#ou}
	Ou *string `field:"optional" json:"ou" yaml:"ou"`
	// List of domains for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#permitted_dns_domains PkiSecretBackendRootSignIntermediate#permitted_dns_domains}
	PermittedDnsDomains *[]*string `field:"optional" json:"permittedDnsDomains" yaml:"permittedDnsDomains"`
	// List of email addresses for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#permitted_email_addresses PkiSecretBackendRootSignIntermediate#permitted_email_addresses}
	PermittedEmailAddresses *[]*string `field:"optional" json:"permittedEmailAddresses" yaml:"permittedEmailAddresses"`
	// List of IP ranges for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#permitted_ip_ranges PkiSecretBackendRootSignIntermediate#permitted_ip_ranges}
	PermittedIpRanges *[]*string `field:"optional" json:"permittedIpRanges" yaml:"permittedIpRanges"`
	// List of URI domains for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#permitted_uri_domains PkiSecretBackendRootSignIntermediate#permitted_uri_domains}
	PermittedUriDomains *[]*string `field:"optional" json:"permittedUriDomains" yaml:"permittedUriDomains"`
	// The postal code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#postal_code PkiSecretBackendRootSignIntermediate#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The province.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#province PkiSecretBackendRootSignIntermediate#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// Revoke the certificate upon resource destruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#revoke PkiSecretBackendRootSignIntermediate#revoke}
	Revoke interface{} `field:"optional" json:"revoke" yaml:"revoke"`
	// The number of bits to use in the signature algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#signature_bits PkiSecretBackendRootSignIntermediate#signature_bits}
	SignatureBits *float64 `field:"optional" json:"signatureBits" yaml:"signatureBits"`
	// Value for the Subject Key Identifier field   (RFC 5280 Section 4.2.1.2). Specified as a string in hex format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#skid PkiSecretBackendRootSignIntermediate#skid}
	Skid *string `field:"optional" json:"skid" yaml:"skid"`
	// The street address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#street_address PkiSecretBackendRootSignIntermediate#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#ttl PkiSecretBackendRootSignIntermediate#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#uri_sans PkiSecretBackendRootSignIntermediate#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
	// Preserve CSR values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#use_csr_values PkiSecretBackendRootSignIntermediate#use_csr_values}
	UseCsrValues interface{} `field:"optional" json:"useCsrValues" yaml:"useCsrValues"`
	// Specifies whether or not to use PSS signatures   over PKCS#1v1.5 signatures when a RSA-type issuer is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_sign_intermediate#use_pss PkiSecretBackendRootSignIntermediate#use_pss}
	UsePss interface{} `field:"optional" json:"usePss" yaml:"usePss"`
}

