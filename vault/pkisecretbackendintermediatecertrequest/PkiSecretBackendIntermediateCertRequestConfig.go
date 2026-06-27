// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendintermediatecertrequest

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PkiSecretBackendIntermediateCertRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#backend PkiSecretBackendIntermediateCertRequest#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// CN of intermediate to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#common_name PkiSecretBackendIntermediateCertRequest#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Type of intermediate to create. Must be either "existing", "exported", "internal" or "kms".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#type PkiSecretBackendIntermediateCertRequest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Set 'CA: true' in a Basic Constraints extension.
	//
	// Only needed as
	// a workaround in some compatibility scenarios with Active Directory Certificate Services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#add_basic_constraints PkiSecretBackendIntermediateCertRequest#add_basic_constraints}
	AddBasicConstraints interface{} `field:"optional" json:"addBasicConstraints" yaml:"addBasicConstraints"`
	// List of alternative names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#alt_names PkiSecretBackendIntermediateCertRequest#alt_names}
	AltNames *[]*string `field:"optional" json:"altNames" yaml:"altNames"`
	// The country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#country PkiSecretBackendIntermediateCertRequest#country}
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Flag to exclude CN from SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#exclude_cn_from_sans PkiSecretBackendIntermediateCertRequest#exclude_cn_from_sans}
	ExcludeCnFromSans interface{} `field:"optional" json:"excludeCnFromSans" yaml:"excludeCnFromSans"`
	// The format of data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#format PkiSecretBackendIntermediateCertRequest#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#id PkiSecretBackendIntermediateCertRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of alternative IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#ip_sans PkiSecretBackendIntermediateCertRequest#ip_sans}
	IpSans *[]*string `field:"optional" json:"ipSans" yaml:"ipSans"`
	// The number of bits to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#key_bits PkiSecretBackendIntermediateCertRequest#key_bits}
	KeyBits *float64 `field:"optional" json:"keyBits" yaml:"keyBits"`
	// When a new key is created with this request, optionally specifies the name for this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#key_name PkiSecretBackendIntermediateCertRequest#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Specifies the key to use for generating this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#key_ref PkiSecretBackendIntermediateCertRequest#key_ref}
	KeyRef *string `field:"optional" json:"keyRef" yaml:"keyRef"`
	// The desired key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#key_type PkiSecretBackendIntermediateCertRequest#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Specify the key usages to encode in the generated certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#key_usage PkiSecretBackendIntermediateCertRequest#key_usage}
	KeyUsage *[]*string `field:"optional" json:"keyUsage" yaml:"keyUsage"`
	// The locality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#locality PkiSecretBackendIntermediateCertRequest#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The ID of the previously configured managed key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#managed_key_id PkiSecretBackendIntermediateCertRequest#managed_key_id}
	ManagedKeyId *string `field:"optional" json:"managedKeyId" yaml:"managedKeyId"`
	// The name of the previously configured managed key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#managed_key_name PkiSecretBackendIntermediateCertRequest#managed_key_name}
	ManagedKeyName *string `field:"optional" json:"managedKeyName" yaml:"managedKeyName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#namespace PkiSecretBackendIntermediateCertRequest#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#organization PkiSecretBackendIntermediateCertRequest#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// List of other SANs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#other_sans PkiSecretBackendIntermediateCertRequest#other_sans}
	OtherSans *[]*string `field:"optional" json:"otherSans" yaml:"otherSans"`
	// The organization unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#ou PkiSecretBackendIntermediateCertRequest#ou}
	Ou *string `field:"optional" json:"ou" yaml:"ou"`
	// The postal code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#postal_code PkiSecretBackendIntermediateCertRequest#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The private key format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#private_key_format PkiSecretBackendIntermediateCertRequest#private_key_format}
	PrivateKeyFormat *string `field:"optional" json:"privateKeyFormat" yaml:"privateKeyFormat"`
	// The province.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#province PkiSecretBackendIntermediateCertRequest#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The requested Subject's named serial number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#serial_number PkiSecretBackendIntermediateCertRequest#serial_number}
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// The number of bits to use in the signature algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#signature_bits PkiSecretBackendIntermediateCertRequest#signature_bits}
	SignatureBits *float64 `field:"optional" json:"signatureBits" yaml:"signatureBits"`
	// The street address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#street_address PkiSecretBackendIntermediateCertRequest#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
	// List of alternative URIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_intermediate_cert_request#uri_sans PkiSecretBackendIntermediateCertRequest#uri_sans}
	UriSans *[]*string `field:"optional" json:"uriSans" yaml:"uriSans"`
}

