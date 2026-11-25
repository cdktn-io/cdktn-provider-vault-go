// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transitsecretbackendimportedkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransitSecretBackendImportedKeyConfig struct {
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
	// Path where the Transit secret backend is mounted. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#mount TransitSecretBackendImportedKey#mount}
	Mount *string `field:"required" json:"mount" yaml:"mount"`
	// Name of the encryption key to create. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#name TransitSecretBackendImportedKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the type of key to create (e.g. `aes128-gcm96`, `aes256-gcm96`, `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072`, `rsa-4096`, `hmac`, `aes128-cmac`, `aes192-cmac`, `aes256-cmac`). Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#type TransitSecretBackendImportedKey#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// If set, enables taking backup of named key in plaintext format. Once enabled, cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#allow_plaintext_backup TransitSecretBackendImportedKey#allow_plaintext_backup}
	AllowPlaintextBackup interface{} `field:"optional" json:"allowPlaintextBackup" yaml:"allowPlaintextBackup"`
	// If set, the imported key can be rotated within Vault by using the rotate endpoint.
	//
	// Once the key is rotated in Vault, it no longer accepts imported key material. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#allow_rotation TransitSecretBackendImportedKey#allow_rotation}
	AllowRotation interface{} `field:"optional" json:"allowRotation" yaml:"allowRotation"`
	// The period in seconds at which this key should be rotated automatically.
	//
	// Set to 0 to disable; otherwise it must be at least one hour (3600).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#auto_rotate_period TransitSecretBackendImportedKey#auto_rotate_period}
	AutoRotatePeriod *float64 `field:"optional" json:"autoRotatePeriod" yaml:"autoRotatePeriod"`
	// A base64-encoded string containing the wrapped ephemeral AES key and the encrypted key material (write-only).
	//
	// Required with `ciphertext_wo_version`. Vault requires either this or `public_key` for the initial import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#ciphertext_wo TransitSecretBackendImportedKey#ciphertext_wo}
	CiphertextWo *string `field:"optional" json:"ciphertextWo" yaml:"ciphertextWo"`
	// Version identifier for `ciphertext_wo`. Increment it to import a new ciphertext as a new key version. Required with `ciphertext_wo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#ciphertext_wo_version TransitSecretBackendImportedKey#ciphertext_wo_version}
	CiphertextWoVersion *float64 `field:"optional" json:"ciphertextWoVersion" yaml:"ciphertextWoVersion"`
	// A base64-encoded string providing a context for key derivation.
	//
	// Required if `derived` is true. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#context TransitSecretBackendImportedKey#context}
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Specifies if the key is allowed to be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#deletion_allowed TransitSecretBackendImportedKey#deletion_allowed}
	DeletionAllowed interface{} `field:"optional" json:"deletionAllowed" yaml:"deletionAllowed"`
	// Specifies if key derivation is to be used. Changing this field replaces this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#derived TransitSecretBackendImportedKey#derived}
	Derived interface{} `field:"optional" json:"derived" yaml:"derived"`
	// Enables keys to be exportable. Once enabled, cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#exportable TransitSecretBackendImportedKey#exportable}
	Exportable interface{} `field:"optional" json:"exportable" yaml:"exportable"`
	// The hash function used for the RSA-OAEP step of creating the ciphertext (`SHA1`, `SHA224`, `SHA256`, `SHA384`, `SHA512`).
	//
	// Vault uses `SHA256` if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#hash_function TransitSecretBackendImportedKey#hash_function}
	HashFunction *string `field:"optional" json:"hashFunction" yaml:"hashFunction"`
	// Changes the key type of the new key version when importing key material into an existing key.
	//
	// Sent as `type` to Vault's import key version endpoint, so it has no effect on the initial import. Reported by `latest_version_type`. Requires Vault Enterprise 2.2.0 or later.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#import_version_type TransitSecretBackendImportedKey#import_version_type}
	ImportVersionType *string `field:"optional" json:"importVersionType" yaml:"importVersionType"`
	// Minimum key version to use for decryption. Vault always sets this to at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#min_decryption_version TransitSecretBackendImportedKey#min_decryption_version}
	MinDecryptionVersion *float64 `field:"optional" json:"minDecryptionVersion" yaml:"minDecryptionVersion"`
	// Minimum key version to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#min_encryption_version TransitSecretBackendImportedKey#min_encryption_version}
	MinEncryptionVersion *float64 `field:"optional" json:"minEncryptionVersion" yaml:"minEncryptionVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#namespace TransitSecretBackendImportedKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// A plaintext PEM public key to be imported.
	//
	// Vault requires either this or `ciphertext_wo` for the initial import. Changing it imports the new public key as a new key version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#public_key TransitSecretBackendImportedKey#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// Specifies the existing key version to update when importing key material to an already-imported key.
	//
	// If omitted, Vault creates a new version, except when the latest version contains only a public key and the request supplies its matching private key; in that case Vault adds the private key to the existing, latest version. Ignored when the key is first imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.14.0/docs/resources/transit_secret_backend_imported_key#version TransitSecretBackendImportedKey#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

