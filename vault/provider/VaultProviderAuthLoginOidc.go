// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginOidc struct {
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The callback address. Must be a valid URI without the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#callback_address VaultProvider#callback_address}
	CallbackAddress *string `field:"optional" json:"callbackAddress" yaml:"callbackAddress"`
	// The callback listener's address. Must be a valid URI without the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#callback_listener_address VaultProvider#callback_listener_address}
	CallbackListenerAddress *string `field:"optional" json:"callbackListenerAddress" yaml:"callbackListenerAddress"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace. Conflicts with use_root_namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Authenticate to the root Vault namespace. Conflicts with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#use_root_namespace VaultProvider#use_root_namespace}
	UseRootNamespace interface{} `field:"optional" json:"useRootNamespace" yaml:"useRootNamespace"`
}

