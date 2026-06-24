// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sshsecretbackendca

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SshSecretBackendCaConfig struct {
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
	// The path of the SSH Secret Backend where the CA should be configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#backend SshSecretBackendCa#backend}
	Backend *string `field:"optional" json:"backend" yaml:"backend"`
	// Whether Vault should generate the signing key pair internally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#generate_signing_key SshSecretBackendCa#generate_signing_key}
	GenerateSigningKey interface{} `field:"optional" json:"generateSigningKey" yaml:"generateSigningKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#id SshSecretBackendCa#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies the desired key bits for the generated SSH CA key when `generate_signing_key` is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#key_bits SshSecretBackendCa#key_bits}
	KeyBits *float64 `field:"optional" json:"keyBits" yaml:"keyBits"`
	// Specifies the desired key type for the generated SSH CA key when `generate_signing_key` is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#key_type SshSecretBackendCa#key_type}
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// The id of the managed key to use. When using a managed key, this field or managed_key_name is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#managed_key_id SshSecretBackendCa#managed_key_id}
	ManagedKeyId *string `field:"optional" json:"managedKeyId" yaml:"managedKeyId"`
	// The name of the managed key to use. When using a managed key, this field or managed_key_id is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#managed_key_name SshSecretBackendCa#managed_key_name}
	ManagedKeyName *string `field:"optional" json:"managedKeyName" yaml:"managedKeyName"`
	// Target namespace. (requires Enterprise).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#namespace SshSecretBackendCa#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Private key part the SSH CA key pair; required if generate_signing_key is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#private_key SshSecretBackendCa#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Public key part the SSH CA key pair; required if generate_signing_key is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_ca#public_key SshSecretBackendCa#public_key}
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
}

