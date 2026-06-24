// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendsign

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendSignConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#backend PkiSecretBackendSign#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of intermediate to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#common_name PkiSecretBackendSign#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// The CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#csr PkiSecretBackendSign#csr}
	Csr *string `field:"required" json:"csr" yaml:"csr"`
	// Name of the role to create the certificate against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#name PkiSecretBackendSign#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#alt_names PkiSecretBackendSign#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#auto_renew PkiSecretBackendSign#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// A base 64 encoded value or an empty string to associate with the certificate's serial number.
	//
	// The role's no_store_metadata must be set to false, otherwise an error is returned when specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#cert_metadata PkiSecretBackendSign#cert_metadata}
	CertMetadata *string `field:"optional" json:"certMetadata" yaml:"certMetadata"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#exclude_cn_from_sans PkiSecretBackendSign#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// The format of data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#format PkiSecretBackendSign#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#id PkiSecretBackendSign#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#ip_sans PkiSecretBackendSign#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// Specifies the default issuer of this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#issuer_ref PkiSecretBackendSign#issuer_ref}
	IssuerRef *string `field:"optional" json:"issuerRef" yaml:"issuerRef"`
	// Password for encrypting the Java keystore 		when format is set to "jks_bundle".
	//
	// If not provided,
	// 		defaults to "changeit". It is recommended to use the default password
	// 		and protect the file using other means or use a high-entropy password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#jks_password PkiSecretBackendSign#jks_password}
	JksPassword *string `field:"optional" json:"jksPassword" yaml:"jksPassword"`
	// Generate a new certificate when the expiration is within this number of seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#min_seconds_remaining PkiSecretBackendSign#min_seconds_remaining}
	MinSecondsRemaining *float64 `field:"optional" json:"minSecondsRemaining" yaml:"minSecondsRemaining"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#namespace PkiSecretBackendSign#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set the Not After field of the certificate with specified date value.
	//
	// The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#not_after PkiSecretBackendSign#not_after}
	NotAfter *string `field:"optional" json:"notAfter" yaml:"notAfter"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#other_sans PkiSecretBackendSign#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// Encoder profile to use for PKCS#12 archives when  format is set to "pkcs12_bundle".
	//
	// Valid values are "modern2026" and
	// "modern2023". Defaults to "modern2026", which uses the newer PKCS#12
	// integrity format (PBMAC1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#pkcs12_encoder PkiSecretBackendSign#pkcs12_encoder}
	Pkcs12Encoder *string `field:"optional" json:"pkcs12Encoder" yaml:"pkcs12Encoder"`
	// Password for encrypting the PKCS#12  		archive when format is set to "pkcs12_bundle".
	//
	// If not provided,
	// 		defaults to "changeit". It is recommended to use the default password
	// 		and protect the file using other means or use a high-entropy password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#pkcs12_password PkiSecretBackendSign#pkcs12_password}
	Pkcs12Password *string `field:"optional" json:"pkcs12Password" yaml:"pkcs12Password"`
	// If true, the returned ca_chain field will not include any self-signed CA certificates.
	//
	// Useful if end-users already have the root CA in their trust store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#remove_roots_from_chain PkiSecretBackendSign#remove_roots_from_chain}
	RemoveRootsFromChain interface{} `field:"optional" json:"removeRootsFromChain" yaml:"removeRootsFromChain"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#ttl PkiSecretBackendSign#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_sign#uri_sans PkiSecretBackendSign#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
}

