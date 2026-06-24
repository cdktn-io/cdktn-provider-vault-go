// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys


type ManagedKeysGcp struct {
	// The signature algorithm to be used with the key.
	//
	// Supported values: ec_sign_p256_sha256, ec_sign_p384_sha384, rsa_sign_pss_2048_sha256, rsa_sign_pss_3072_sha256, rsa_sign_pss_4096_sha256, rsa_sign_pss_4096_sha512, rsa_sign_pkcs1_2048_sha256, rsa_sign_pkcs1_3072_sha256, rsa_sign_pkcs1_4096_sha256, rsa_sign_pkcs1_4096_sha512
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#algorithm ManagedKeys#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// The GCP service account credentials JSON to use for authenticating to GCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#credentials ManagedKeys#credentials}
	Credentials *string `field:"required" json:"credentials" yaml:"credentials"`
	// The name of the GCP Cloud KMS key.
	//
	// If no existing key exists and allow_generate_key is true, Vault will generate a key with this name
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#crypto_key ManagedKeys#crypto_key}
	CryptoKey *string `field:"required" json:"cryptoKey" yaml:"cryptoKey"`
	// The name of the key ring in GCP Cloud KMS. This needs to be created prior to key creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#key_ring ManagedKeys#key_ring}
	KeyRing *string `field:"required" json:"keyRing" yaml:"keyRing"`
	// A unique lowercase name that serves as identifying the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#name ManagedKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The GCP project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#project ManagedKeys#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The GCP region where the key ring was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#region ManagedKeys#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// If no existing key can be found in the referenced backend, instructs Vault to generate a key within the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#allow_generate_key ManagedKeys#allow_generate_key}
	AllowGenerateKey interface{} `field:"optional" json:"allowGenerateKey" yaml:"allowGenerateKey"`
	// Controls the ability for Vault to replace through generation or importing a key into the configured backend even if a key is present, if set to false those operations are forbidden if a key exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#allow_replace_key ManagedKeys#allow_replace_key}
	AllowReplaceKey interface{} `field:"optional" json:"allowReplaceKey" yaml:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the configured backend, if 'false', those operations will be forbidden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#allow_store_key ManagedKeys#allow_store_key}
	AllowStoreKey interface{} `field:"optional" json:"allowStoreKey" yaml:"allowStoreKey"`
	// Allow usage from any mount point within the namespace if 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#any_mount ManagedKeys#any_mount}
	AnyMount interface{} `field:"optional" json:"anyMount" yaml:"anyMount"`
	// The version of the key to use. (Default: 1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#crypto_key_version ManagedKeys#crypto_key_version}
	CryptoKeyVersion *string `field:"optional" json:"cryptoKeyVersion" yaml:"cryptoKeyVersion"`
	// A list of the allowed usages of this key.
	//
	// Valid values are encrypt, decrypt, sign, verify, wrap, unwrap, mac, and generate_random. Default values are sign and verify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#usages ManagedKeys#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
}

