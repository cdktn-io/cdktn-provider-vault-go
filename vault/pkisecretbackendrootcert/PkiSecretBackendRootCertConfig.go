// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrootcert

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendRootCertConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#backend PkiSecretBackendRootCert#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of root to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#common_name PkiSecretBackendRootCert#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Type of root to create. Must be either "existing", "exported", "internal" or "kms".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#type PkiSecretBackendRootCert#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#alt_names PkiSecretBackendRootCert#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// The country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#country PkiSecretBackendRootCert#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#exclude_cn_from_sans PkiSecretBackendRootCert#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// List of domains for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#excluded_dns_domains PkiSecretBackendRootCert#excluded_dns_domains}
	ExcludedDnsDomains *[]*string `field:"optional" json:"excludedDnsDomains" yaml:"excludedDnsDomains"`
	// List of email addresses for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#excluded_email_addresses PkiSecretBackendRootCert#excluded_email_addresses}
	ExcludedEmailAddresses *[]*string `field:"optional" json:"excludedEmailAddresses" yaml:"excludedEmailAddresses"`
	// List of IP ranges for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#excluded_ip_ranges PkiSecretBackendRootCert#excluded_ip_ranges}
	ExcludedIpRanges *[]*string `field:"optional" json:"excludedIpRanges" yaml:"excludedIpRanges"`
	// List of URI domains for which certificates are not allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#excluded_uri_domains PkiSecretBackendRootCert#excluded_uri_domains}
	ExcludedUriDomains *[]*string `field:"optional" json:"excludedUriDomains" yaml:"excludedUriDomains"`
	// The format of data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#format PkiSecretBackendRootCert#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#id PkiSecretBackendRootCert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#ip_sans PkiSecretBackendRootCert#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// Provides a name to the specified issuer.
	//
	// The name must be unique across all issuers and not be the reserved value 'default'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#issuer_name PkiSecretBackendRootCert#issuer_name}
	IssuerName *string `field:"optional" json:"issuerName" yaml:"issuerName"`
	// The number of bits to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#key_bits PkiSecretBackendRootCert#key_bits}
	KeyBits *float64 `field:"optional" json:"keyBits" yaml:"keyBits"`
	// When a new key is created with this request, optionally specifies the name for this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#key_name PkiSecretBackendRootCert#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Specifies the key to use for generating this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#key_ref PkiSecretBackendRootCert#key_ref}
	KeyRef *string `field:"optional" json:"keyRef" yaml:"keyRef"`
	// The desired key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#key_type PkiSecretBackendRootCert#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Specify the allowed key usage constraint on issued certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#key_usage PkiSecretBackendRootCert#key_usage}
	KeyUsage *[]*string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The locality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#locality PkiSecretBackendRootCert#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The ID of the previously configured managed key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#managed_key_id PkiSecretBackendRootCert#managed_key_id}
	ManagedKeyId *string `field:"optional" json:"managedKeyId" yaml:"managedKeyId"`
	// The name of the previously configured managed key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#managed_key_name PkiSecretBackendRootCert#managed_key_name}
	ManagedKeyName *string `field:"optional" json:"managedKeyName" yaml:"managedKeyName"`
	// The maximum path length to encode in the generated certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#max_path_length PkiSecretBackendRootCert#max_path_length}
	MaxPathLength *float64 `field:"optional" json:"maxPathLength" yaml:"maxPathLength"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#namespace PkiSecretBackendRootCert#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set the Not After field of the certificate with specified date value.
	//
	// The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#not_after PkiSecretBackendRootCert#not_after}
	NotAfter *string `field:"optional" json:"notAfter" yaml:"notAfter"`
	// Specifies the duration by which to backdate the NotBefore property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#not_before_duration PkiSecretBackendRootCert#not_before_duration}
	NotBeforeDuration *string `field:"optional" json:"notBeforeDuration" yaml:"notBeforeDuration"`
	// The organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#organization PkiSecretBackendRootCert#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#other_sans PkiSecretBackendRootCert#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// The organization unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#ou PkiSecretBackendRootCert#ou}
	Ou *string `field:"optional" json:"ou" yaml:"ou"`
	// List of domains for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#permitted_dns_domains PkiSecretBackendRootCert#permitted_dns_domains}
	PermittedDnsDomains *[]*string `field:"optional" json:"permittedDnsDomains" yaml:"permittedDnsDomains"`
	// List of email addresses for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#permitted_email_addresses PkiSecretBackendRootCert#permitted_email_addresses}
	PermittedEmailAddresses *[]*string `field:"optional" json:"permittedEmailAddresses" yaml:"permittedEmailAddresses"`
	// List of IP ranges for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#permitted_ip_ranges PkiSecretBackendRootCert#permitted_ip_ranges}
	PermittedIpRanges *[]*string `field:"optional" json:"permittedIpRanges" yaml:"permittedIpRanges"`
	// List of URI domains for which certificates are allowed to be issued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#permitted_uri_domains PkiSecretBackendRootCert#permitted_uri_domains}
	PermittedUriDomains *[]*string `field:"optional" json:"permittedUriDomains" yaml:"permittedUriDomains"`
	// The postal code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#postal_code PkiSecretBackendRootCert#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The private key format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#private_key_format PkiSecretBackendRootCert#private_key_format}
	PrivateKeyFormat *string `field:"optional" json:"privateKeyFormat" yaml:"privateKeyFormat"`
	// The province.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#province PkiSecretBackendRootCert#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The number of bits to use in the signature algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#signature_bits PkiSecretBackendRootCert#signature_bits}
	SignatureBits *float64 `field:"optional" json:"signatureBits" yaml:"signatureBits"`
	// The street address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#street_address PkiSecretBackendRootCert#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#ttl PkiSecretBackendRootCert#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#uri_sans PkiSecretBackendRootCert#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
	// Whether to use PSS signatures when using a RSA key-type issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/pki_secret_backend_root_cert#use_pss PkiSecretBackendRootCert#use_pss}
	UsePss interface{} `field:"optional" json:"usePss" yaml:"usePss"`
}

