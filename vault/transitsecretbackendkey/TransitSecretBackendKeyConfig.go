// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package transitsecretbackendkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TransitSecretBackendKeyConfig struct {
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
	// The Transit secret backend the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#backend TransitSecretBackendKey#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// Name of the encryption key to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#name TransitSecretBackendKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If set, enables taking backup of named key in the plaintext format. Once set, this cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#allow_plaintext_backup TransitSecretBackendKey#allow_plaintext_backup}
	AllowPlaintextBackup interface{} `field:"optional" json:"allowPlaintextBackup" yaml:"allowPlaintextBackup"`
	// Amount of seconds the key should live before being automatically rotated.
	//
	// A value of 0 disables automatic rotation for the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#auto_rotate_period TransitSecretBackendKey#auto_rotate_period}
	AutoRotatePeriod *float64 `field:"optional" json:"autoRotatePeriod" yaml:"autoRotatePeriod"`
	// Base64 encoded context for key derivation. Required if derived is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#context TransitSecretBackendKey#context}
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext.
	//
	// This requires derived to be set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#convergent_encryption TransitSecretBackendKey#convergent_encryption}
	ConvergentEncryption interface{} `field:"optional" json:"convergentEncryption" yaml:"convergentEncryption"`
	// Specifies if the key is allowed to be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#deletion_allowed TransitSecretBackendKey#deletion_allowed}
	DeletionAllowed interface{} `field:"optional" json:"deletionAllowed" yaml:"deletionAllowed"`
	// Specifies if key derivation is to be used.
	//
	// If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#derived TransitSecretBackendKey#derived}
	Derived interface{} `field:"optional" json:"derived" yaml:"derived"`
	// Enables keys to be exportable.
	//
	// This allows for all the valid keys in the key ring to be exported. Once set, this cannot be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#exportable TransitSecretBackendKey#exportable}
	Exportable interface{} `field:"optional" json:"exportable" yaml:"exportable"`
	// The elliptic curve algorithm to use for hybrid signatures. Supported key types are `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, and `ed25519`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#hybrid_key_type_ec TransitSecretBackendKey#hybrid_key_type_ec}
	HybridKeyTypeEc *string `field:"optional" json:"hybridKeyTypeEc" yaml:"hybridKeyTypeEc"`
	// The post-quantum algorithm to use for hybrid signatures. Currently, ML-DSA is the only supported key type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#hybrid_key_type_pqc TransitSecretBackendKey#hybrid_key_type_pqc}
	HybridKeyTypePqc *string `field:"optional" json:"hybridKeyTypePqc" yaml:"hybridKeyTypePqc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#id TransitSecretBackendKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The key size in bytes for algorithms that allow variable key sizes.
	//
	// Currently only applicable to HMAC; this value must be between 32 and 512.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#key_size TransitSecretBackendKey#key_size}
	KeySize *float64 `field:"optional" json:"keySize" yaml:"keySize"`
	// The UUID of the managed key to use when the key type is managed_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#managed_key_id TransitSecretBackendKey#managed_key_id}
	ManagedKeyId *string `field:"optional" json:"managedKeyId" yaml:"managedKeyId"`
	// The name of the managed key to use when the key type is managed_key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#managed_key_name TransitSecretBackendKey#managed_key_name}
	ManagedKeyName *string `field:"optional" json:"managedKeyName" yaml:"managedKeyName"`
	// Minimum key version to use for decryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#min_decryption_version TransitSecretBackendKey#min_decryption_version}
	MinDecryptionVersion *float64 `field:"optional" json:"minDecryptionVersion" yaml:"minDecryptionVersion"`
	// Minimum key version to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#min_encryption_version TransitSecretBackendKey#min_encryption_version}
	MinEncryptionVersion *float64 `field:"optional" json:"minEncryptionVersion" yaml:"minEncryptionVersion"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#namespace TransitSecretBackendKey#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The parameter set to use for ML-DSA.
	//
	// Required for ML-DSA and hybrid keys.  Valid values for ML-DSA are `44`, `65`, and `87`. Valid values for SLH-DSA are `slh-dsa-sha2-128s`, `slh-dsa-shake-128s`, `slh-dsa-sha2-128f`, `slh-dsa-shake-128`, `slh-dsa-sha2-192s`, `slh-dsa-shake-192s`, `slh-dsa-sha2-192f`, `slh-dsa-shake-192f`, `slh-dsa-sha2-256s`, `slh-dsa-shake-256s`, `slh-dsa-sha2-256f`, and `slh-dsa-shake-256f`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#parameter_set TransitSecretBackendKey#parameter_set}
	ParameterSet *string `field:"optional" json:"parameterSet" yaml:"parameterSet"`
	// Specifies the type of key to create.
	//
	// The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `hmac`, `rsa-2048`, `rsa-3072`, `rsa-4096`, `managed_key`, `aes128-cmac`, `aes192-cmac`, `aes256-cmac`, `ml-dsa`, `hybrid`, and `slh-dsa`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/transit_secret_backend_key#type TransitSecretBackendKey#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

