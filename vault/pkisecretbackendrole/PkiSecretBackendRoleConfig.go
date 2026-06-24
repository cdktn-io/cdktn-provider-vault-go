// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrole

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendRoleConfig struct {
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
	// The path of the PKI secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#backend PkiSecretBackendRole#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Unique name for the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#name PkiSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Flag to allow any name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_any_name PkiSecretBackendRole#allow_any_name}
	AllowAnyName interface{} `field:"optional" json:"allowAnyName" yaml:"allowAnyName"`
	// Flag to allow certificates matching the actual domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_bare_domains PkiSecretBackendRole#allow_bare_domains}
	AllowBareDomains interface{} `field:"optional" json:"allowBareDomains" yaml:"allowBareDomains"`
	// The domains of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_domains PkiSecretBackendRole#allowed_domains}
	AllowedDomains *[]*string `field:"optional" json:"allowedDomains" yaml:"allowedDomains"`
	// Flag to indicate that `allowed_domains` specifies a template expression (e.g. {{identity.entity.aliases.<mount accessor>.name}}).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_domains_template PkiSecretBackendRole#allowed_domains_template}
	AllowedDomainsTemplate interface{} `field:"optional" json:"allowedDomainsTemplate" yaml:"allowedDomainsTemplate"`
	// Defines allowed custom SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_other_sans PkiSecretBackendRole#allowed_other_sans}
	AllowedOtherSans *[]*string `field:"optional" json:"allowedOtherSans" yaml:"allowedOtherSans"`
	// Defines allowed Subject serial numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_serial_numbers PkiSecretBackendRole#allowed_serial_numbers}
	AllowedSerialNumbers *[]*string `field:"optional" json:"allowedSerialNumbers" yaml:"allowedSerialNumbers"`
	// Defines allowed URI SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_uri_sans PkiSecretBackendRole#allowed_uri_sans}
	AllowedUriSans *[]*string `field:"optional" json:"allowedUriSans" yaml:"allowedUriSans"`
	// Flag to indicate that `allowed_uri_sans` specifies a template expression (e.g. {{identity.entity.aliases.<mount accessor>.name}}).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_uri_sans_template PkiSecretBackendRole#allowed_uri_sans_template}
	AllowedUriSansTemplate interface{} `field:"optional" json:"allowedUriSansTemplate" yaml:"allowedUriSansTemplate"`
	// The allowed User ID's.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allowed_user_ids PkiSecretBackendRole#allowed_user_ids}
	AllowedUserIds *[]*string `field:"optional" json:"allowedUserIds" yaml:"allowedUserIds"`
	// Flag to allow names containing glob patterns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_glob_domains PkiSecretBackendRole#allow_glob_domains}
	AllowGlobDomains interface{} `field:"optional" json:"allowGlobDomains" yaml:"allowGlobDomains"`
	// Flag to allow IP SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_ip_sans PkiSecretBackendRole#allow_ip_sans}
	AllowIpSans interface{} `field:"optional" json:"allowIpSans" yaml:"allowIpSans"`
	// Flag to allow certificates for localhost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_localhost PkiSecretBackendRole#allow_localhost}
	AllowLocalhost interface{} `field:"optional" json:"allowLocalhost" yaml:"allowLocalhost"`
	// Flag to allow certificates matching subdomains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_subdomains PkiSecretBackendRole#allow_subdomains}
	AllowSubdomains interface{} `field:"optional" json:"allowSubdomains" yaml:"allowSubdomains"`
	// Flag to allow wildcard certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#allow_wildcard_certificates PkiSecretBackendRole#allow_wildcard_certificates}
	AllowWildcardCertificates interface{} `field:"optional" json:"allowWildcardCertificates" yaml:"allowWildcardCertificates"`
	// Flag to mark basic constraints valid when issuing non-CA certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#basic_constraints_valid_for_non_ca PkiSecretBackendRole#basic_constraints_valid_for_non_ca}
	BasicConstraintsValidForNonCa interface{} `field:"optional" json:"basicConstraintsValidForNonCa" yaml:"basicConstraintsValidForNonCa"`
	// Flag to specify certificates for client use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#client_flag PkiSecretBackendRole#client_flag}
	ClientFlag interface{} `field:"optional" json:"clientFlag" yaml:"clientFlag"`
	// Specify validations to run on the Common Name field of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#cn_validations PkiSecretBackendRole#cn_validations}
	CnValidations *[]*string `field:"optional" json:"cnValidations" yaml:"cnValidations"`
	// Flag to specify certificates for code signing use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#code_signing_flag PkiSecretBackendRole#code_signing_flag}
	CodeSigningFlag interface{} `field:"optional" json:"codeSigningFlag" yaml:"codeSigningFlag"`
	// The country of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#country PkiSecretBackendRole#country}
	Country *[]*string `field:"optional" json:"country" yaml:"country"`
	// Flag to specify certificates for email protection use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#email_protection_flag PkiSecretBackendRole#email_protection_flag}
	EmailProtectionFlag interface{} `field:"optional" json:"emailProtectionFlag" yaml:"emailProtectionFlag"`
	// Flag to allow only valid host names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#enforce_hostnames PkiSecretBackendRole#enforce_hostnames}
	EnforceHostnames interface{} `field:"optional" json:"enforceHostnames" yaml:"enforceHostnames"`
	// Specify the allowed extended key usage constraint on issued certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#ext_key_usage PkiSecretBackendRole#ext_key_usage}
	ExtKeyUsage *[]*string `field:"optional" json:"extKeyUsage" yaml:"extKeyUsage"`
	// A list of extended key usage OIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#ext_key_usage_oids PkiSecretBackendRole#ext_key_usage_oids}
	ExtKeyUsageOids *[]*string `field:"optional" json:"extKeyUsageOids" yaml:"extKeyUsageOids"`
	// Flag to generate leases with certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#generate_lease PkiSecretBackendRole#generate_lease}
	GenerateLease interface{} `field:"optional" json:"generateLease" yaml:"generateLease"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#id PkiSecretBackendRole#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the default issuer of this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#issuer_ref PkiSecretBackendRole#issuer_ref}
	IssuerRef *string `field:"optional" json:"issuerRef" yaml:"issuerRef"`
	// The number of bits of generated keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#key_bits PkiSecretBackendRole#key_bits}
	KeyBits *float64 `field:"optional" json:"keyBits" yaml:"keyBits"`
	// The generated key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#key_type PkiSecretBackendRole#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Specify the allowed key usage constraint on issued certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#key_usage PkiSecretBackendRole#key_usage}
	KeyUsage *[]*string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The locality of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#locality PkiSecretBackendRole#locality}
	Locality *[]*string `field:"optional" json:"locality" yaml:"locality"`
	// The maximum TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#max_ttl PkiSecretBackendRole#max_ttl}
	MaxTtl *string `field:"optional" json:"maxTtl" yaml:"maxTtl"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#namespace PkiSecretBackendRole#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Flag to not store certificates in the storage backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#no_store PkiSecretBackendRole#no_store}
	NoStore interface{} `field:"optional" json:"noStore" yaml:"noStore"`
	// Allows metadata to be stored keyed on the certificate's serial number.
	//
	// The field is independent of no_store, allowing metadata storage regardless of whether certificates are stored. If true, metadata is not stored and an error is returned if the metadata field is specified on issuance APIs
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#no_store_metadata PkiSecretBackendRole#no_store_metadata}
	NoStoreMetadata interface{} `field:"optional" json:"noStoreMetadata" yaml:"noStoreMetadata"`
	// Set the Not After field of the certificate with specified date value.
	//
	// The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#not_after PkiSecretBackendRole#not_after}
	NotAfter *string `field:"optional" json:"notAfter" yaml:"notAfter"`
	// Specifies the duration by which to backdate the NotBefore property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#not_before_duration PkiSecretBackendRole#not_before_duration}
	NotBeforeDuration *string `field:"optional" json:"notBeforeDuration" yaml:"notBeforeDuration"`
	// The organization of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#organization PkiSecretBackendRole#organization}
	Organization *[]*string `field:"optional" json:"organization" yaml:"organization"`
	// The organization unit of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#ou PkiSecretBackendRole#ou}
	Ou *[]*string `field:"optional" json:"ou" yaml:"ou"`
	// policy_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#policy_identifier PkiSecretBackendRole#policy_identifier}
	PolicyIdentifier interface{} `field:"optional" json:"policyIdentifier" yaml:"policyIdentifier"`
	// Specify the list of allowed policies OIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#policy_identifiers PkiSecretBackendRole#policy_identifiers}
	PolicyIdentifiers *[]*string `field:"optional" json:"policyIdentifiers" yaml:"policyIdentifiers"`
	// The postal code of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#postal_code PkiSecretBackendRole#postal_code}
	PostalCode *[]*string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The province of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#province PkiSecretBackendRole#province}
	Province *[]*string `field:"optional" json:"province" yaml:"province"`
	// Flag to force CN usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#require_cn PkiSecretBackendRole#require_cn}
	RequireCn interface{} `field:"optional" json:"requireCn" yaml:"requireCn"`
	// Specifies the source of the subject serial number.
	//
	// Valid values are json-csr (default) or json. When set to json-csr, the subject serial number is taken from the serial_number parameter and falls back to the serial number in the CSR. When set to json, the subject serial number is taken from the serial_number parameter but will ignore any value in the CSR. For backwards compatibility an empty value for this field will default to the json-csr behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#serial_number_source PkiSecretBackendRole#serial_number_source}
	SerialNumberSource *string `field:"optional" json:"serialNumberSource" yaml:"serialNumberSource"`
	// Flag to specify certificates for server use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#server_flag PkiSecretBackendRole#server_flag}
	ServerFlag interface{} `field:"optional" json:"serverFlag" yaml:"serverFlag"`
	// The number of bits to use in the signature algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#signature_bits PkiSecretBackendRole#signature_bits}
	SignatureBits *float64 `field:"optional" json:"signatureBits" yaml:"signatureBits"`
	// The street address of generated certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#street_address PkiSecretBackendRole#street_address}
	StreetAddress *[]*string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// The TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#ttl PkiSecretBackendRole#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Flag to use the CN in the CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#use_csr_common_name PkiSecretBackendRole#use_csr_common_name}
	UseCsrCommonName interface{} `field:"optional" json:"useCsrCommonName" yaml:"useCsrCommonName"`
	// Flag to use the SANs in the CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#use_csr_sans PkiSecretBackendRole#use_csr_sans}
	UseCsrSans interface{} `field:"optional" json:"useCsrSans" yaml:"useCsrSans"`
	// Specifies whether or not to use PSS signatures over PKCS#1v1.5 signatures when a RSA-type issuer is used. Ignored for ECDSA/Ed25519 issuers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_role#use_pss PkiSecretBackendRole#use_pss}
	UsePss interface{} `field:"optional" json:"usePss" yaml:"usePss"`
}

