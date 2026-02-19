// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys


type ManagedKeysAws struct {
	// The AWS access key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#access_key ManagedKeys#access_key}
	AccessKey *string `field:"required" json:"accessKey" yaml:"accessKey"`
	// The size in bits for an RSA key. This field is required when 'key_type' is 'RSA'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#key_bits ManagedKeys#key_bits}
	KeyBits *string `field:"required" json:"keyBits" yaml:"keyBits"`
	// The type of key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#key_type ManagedKeys#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// An identifier for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#kms_key ManagedKeys#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
	// A unique lowercase name that serves as identifying the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#name ManagedKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS secret key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#secret_key ManagedKeys#secret_key}
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
	// If no existing key can be found in the referenced backend, instructs Vault to generate a key within the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#allow_generate_key ManagedKeys#allow_generate_key}
	AllowGenerateKey interface{} `field:"optional" json:"allowGenerateKey" yaml:"allowGenerateKey"`
	// Controls the ability for Vault to replace through generation or importing a key into the configured backend even if a key is present, if set to false those operations are forbidden if a key exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#allow_replace_key ManagedKeys#allow_replace_key}
	AllowReplaceKey interface{} `field:"optional" json:"allowReplaceKey" yaml:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the configured backend, if 'false', those operations will be forbidden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#allow_store_key ManagedKeys#allow_store_key}
	AllowStoreKey interface{} `field:"optional" json:"allowStoreKey" yaml:"allowStoreKey"`
	// Allow usage from any mount point within the namespace if 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#any_mount ManagedKeys#any_mount}
	AnyMount interface{} `field:"optional" json:"anyMount" yaml:"anyMount"`
	// The curve to use for an ECDSA key. Used when key_type is 'ECDSA'. Required if 'allow_generate_key' is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#curve ManagedKeys#curve}
	Curve *string `field:"optional" json:"curve" yaml:"curve"`
	// Used to specify a custom AWS endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#endpoint ManagedKeys#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The AWS region where the keys are stored (or will be stored).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/managed_keys#region ManagedKeys#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

