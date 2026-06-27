// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package managedkeys


type ManagedKeysAzure struct {
	// The client id for credentials to query the Azure APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#client_id ManagedKeys#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret for credentials to query the Azure APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#client_secret ManagedKeys#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The Key Vault key to use for encryption and decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#key_name ManagedKeys#key_name}
	KeyName *string `field:"required" json:"keyName" yaml:"keyName"`
	// The type of key to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#key_type ManagedKeys#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// A unique lowercase name that serves as identifying the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#name ManagedKeys#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The tenant id for the Azure Active Directory organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#tenant_id ManagedKeys#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The Key Vault vault to use the encryption keys for encryption and decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#vault_name ManagedKeys#vault_name}
	VaultName *string `field:"required" json:"vaultName" yaml:"vaultName"`
	// If no existing key can be found in the referenced backend, instructs Vault to generate a key within the backend.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#allow_generate_key ManagedKeys#allow_generate_key}
	AllowGenerateKey interface{} `field:"optional" json:"allowGenerateKey" yaml:"allowGenerateKey"`
	// Controls the ability for Vault to replace through generation or importing a key into the configured backend even if a key is present, if set to false those operations are forbidden if a key exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#allow_replace_key ManagedKeys#allow_replace_key}
	AllowReplaceKey interface{} `field:"optional" json:"allowReplaceKey" yaml:"allowReplaceKey"`
	// Controls the ability for Vault to import a key to the configured backend, if 'false', those operations will be forbidden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#allow_store_key ManagedKeys#allow_store_key}
	AllowStoreKey interface{} `field:"optional" json:"allowStoreKey" yaml:"allowStoreKey"`
	// Allow usage from any mount point within the namespace if 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#any_mount ManagedKeys#any_mount}
	AnyMount interface{} `field:"optional" json:"anyMount" yaml:"anyMount"`
	// The Azure Cloud environment API endpoints to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#environment ManagedKeys#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The size in bits for an RSA key.
	//
	// This field is required when 'key_type' is 'RSA' or when 'allow_generate_key' is true
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#key_bits ManagedKeys#key_bits}
	KeyBits *string `field:"optional" json:"keyBits" yaml:"keyBits"`
	// The Azure Key Vault resource's DNS Suffix to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#resource ManagedKeys#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// A list of the allowed usages of this key.
	//
	// Valid values are encrypt, decrypt, sign, verify, wrap, unwrap, mac, and generate_random. Default values are sign and verify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/managed_keys#usages ManagedKeys#usages}
	Usages *[]*string `field:"optional" json:"usages" yaml:"usages"`
}

