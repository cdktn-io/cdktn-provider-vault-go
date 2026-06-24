// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginUserpass struct {
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace. Conflicts with use_root_namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Login with password.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#password VaultProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Login with password from a file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#password_file VaultProvider#password_file}
	PasswordFile *string `field:"optional" json:"passwordFile" yaml:"passwordFile"`
	// Login with username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#username VaultProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Authenticate to the root Vault namespace. Conflicts with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#use_root_namespace VaultProvider#use_root_namespace}
	UseRootNamespace interface{} `field:"optional" json:"useRootNamespace" yaml:"useRootNamespace"`
}

