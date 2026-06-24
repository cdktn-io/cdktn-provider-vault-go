// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys


type ManagedKeysPkcs struct {
	// The name of the kms_library stanza to use from Vault's config to lookup the local library path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#library ManagedKeys#library}
	Library *string `field:"required" json:"library" yaml:"library"`
	// The encryption/decryption mechanism to use, specified as a hexadecimal (prefixed by 0x) string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#mechanism ManagedKeys#mechanism}
	Mechanism *string `field:"required" json:"mechanism" yaml:"mechanism"`
	// A unique lowercase name that serves as identifying the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#name ManagedKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The PIN for login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#pin ManagedKeys#pin}
	Pin *string `field:"required" json:"pin" yaml:"pin"`
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
	// Supplies the curve value when using the 'CKM_ECDSA' mechanism. Required if 'allow_generate_key' is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#curve ManagedKeys#curve}
	Curve *string `field:"optional" json:"curve" yaml:"curve"`
	// Force all operations to open up a read-write session to the HSM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#force_rw_session ManagedKeys#force_rw_session}
	ForceRwSession *string `field:"optional" json:"forceRwSession" yaml:"forceRwSession"`
	// Supplies the size in bits of the key when using 'CKM_RSA_PKCS_PSS', 'CKM_RSA_PKCS_OAEP' or 'CKM_RSA_PKCS' as a value for 'mechanism'.
	//
	// Required if 'allow_generate_key' is true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#key_bits ManagedKeys#key_bits}
	KeyBits *string `field:"optional" json:"keyBits" yaml:"keyBits"`
	// The id of a PKCS#11 key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#key_id ManagedKeys#key_id}
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// The label of the key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#key_label ManagedKeys#key_label}
	KeyLabel *string `field:"optional" json:"keyLabel" yaml:"keyLabel"`
	// The number of concurrent requests that may be in flight to the HSM at any given time.
	//
	// Default is 1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#max_parallel ManagedKeys#max_parallel}
	MaxParallel *float64 `field:"optional" json:"maxParallel" yaml:"maxParallel"`
	// The slot number to use, specified as a string in a decimal format (e.g. '2305843009213693953').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#slot ManagedKeys#slot}
	Slot *string `field:"optional" json:"slot" yaml:"slot"`
	// The slot token label to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#token_label ManagedKeys#token_label}
	TokenLabel *string `field:"optional" json:"tokenLabel" yaml:"tokenLabel"`
	// A list of the allowed usages of this key.
	//
	// Valid values are encrypt, decrypt, sign, verify, wrap, unwrap, mac, and generate_random. Default values are sign and verify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/managed_keys#usages ManagedKeys#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
}

