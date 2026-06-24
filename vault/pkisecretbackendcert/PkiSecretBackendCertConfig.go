// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendcert

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendCertConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#backend PkiSecretBackendCert#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of the certificate to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#common_name PkiSecretBackendCert#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Name of the role to create the certificate against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#name PkiSecretBackendCert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#alt_names PkiSecretBackendCert#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// If enabled, a new certificate will be generated if the expiration is within min_seconds_remaining.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#auto_renew PkiSecretBackendCert#auto_renew}
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// A base 64 encoded value or an empty string to associate with the certificate's serial number.
	//
	// The role's no_store_metadata must be set to false, otherwise an error is returned when specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#cert_metadata PkiSecretBackendCert#cert_metadata}
	CertMetadata *string `field:"optional" json:"certMetadata" yaml:"certMetadata"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#exclude_cn_from_sans PkiSecretBackendCert#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// The format of data. Values "pkcs12_bundle" and "jks_bundle" require Vault version 2.1.0 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#format PkiSecretBackendCert#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#id PkiSecretBackendCert#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#ip_sans PkiSecretBackendCert#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// Specifies the default issuer of this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#issuer_ref PkiSecretBackendCert#issuer_ref}
	IssuerRef *string `field:"optional" json:"issuerRef" yaml:"issuerRef"`
	// Password for encrypting the Java keystore 		when format is set to "jks_bundle".
	//
	// If not provided,
	// 		defaults to "changeit". It is recommended to use the default password
	// 		and protect the file using other means or use a high-entropy password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#jks_password PkiSecretBackendCert#jks_password}
	JksPassword *string `field:"optional" json:"jksPassword" yaml:"jksPassword"`
	// The entry alias in the Java keystore (JKS) when format is set to "jks_bundle"  				and bundle contains a single PrivateKeyEntry.
	//
	// This field is case-sensitive, but relying
	// 			on case-only differences for unique aliases is not recommended. Defaults to "1".
	// 			This parameter is ignored by endpoints that return TrustedCertificateEntry values
	// 			(JKS trust stores), and entry aliases are assigned incrementing numeric strings starting at "1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#jks_private_key_alias PkiSecretBackendCert#jks_private_key_alias}
	JksPrivateKeyAlias *string `field:"optional" json:"jksPrivateKeyAlias" yaml:"jksPrivateKeyAlias"`
	// Generate a new certificate when the expiration is within this number of seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#min_seconds_remaining PkiSecretBackendCert#min_seconds_remaining}
	MinSecondsRemaining *float64 `field:"optional" json:"minSecondsRemaining" yaml:"minSecondsRemaining"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#namespace PkiSecretBackendCert#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Set the Not After field of the certificate with specified date value.
	//
	// The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ. Supports the Y10K end date for IEEE 802.1AR-2018 standard devices, 9999-12-31T23:59:59Z.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#not_after PkiSecretBackendCert#not_after}
	NotAfter *string `field:"optional" json:"notAfter" yaml:"notAfter"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#other_sans PkiSecretBackendCert#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// Encoder profile to use for PKCS#12 archives when  format is set to "pkcs12_bundle".
	//
	// Valid values are "modern2026" and
	// "modern2023". Defaults to "modern2026", which uses the newer PKCS#12
	// integrity format (PBMAC1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#pkcs12_encoder PkiSecretBackendCert#pkcs12_encoder}
	Pkcs12Encoder *string `field:"optional" json:"pkcs12Encoder" yaml:"pkcs12Encoder"`
	// Password for encrypting the PKCS#12  		archive when format is set to "pkcs12_bundle".
	//
	// If not provided,
	// 		defaults to "changeit". It is recommended to use the default password
	// 		and protect the file using other means or use a high-entropy password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#pkcs12_password PkiSecretBackendCert#pkcs12_password}
	Pkcs12Password *string `field:"optional" json:"pkcs12Password" yaml:"pkcs12Password"`
	// The private key format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#private_key_format PkiSecretBackendCert#private_key_format}
	PrivateKeyFormat *string `field:"optional" json:"privateKeyFormat" yaml:"privateKeyFormat"`
	// If true, the returned ca_chain field will not include any self-signed CA certificates.
	//
	// Useful if end-users already have the root CA in their trust store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#remove_roots_from_chain PkiSecretBackendCert#remove_roots_from_chain}
	RemoveRootsFromChain interface{} `field:"optional" json:"removeRootsFromChain" yaml:"removeRootsFromChain"`
	// Revoke the certificate upon resource destruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#revoke PkiSecretBackendCert#revoke}
	Revoke interface{} `field:"optional" json:"revoke" yaml:"revoke"`
	// Revoke the certificate with private key method upon resource destruction.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#revoke_with_key PkiSecretBackendCert#revoke_with_key}
	RevokeWithKey interface{} `field:"optional" json:"revokeWithKey" yaml:"revokeWithKey"`
	// Time to live.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#ttl PkiSecretBackendCert#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#uri_sans PkiSecretBackendCert#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
	// List of Subject User IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_cert#user_ids PkiSecretBackendCert#user_ids}
	UserIds *[]*string `field:"optional" json:"userIds" yaml:"userIds"`
}

