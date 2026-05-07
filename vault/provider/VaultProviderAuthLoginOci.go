// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginOci struct {
	// Authentication type to use when getting OCI credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#auth_type VaultProvider#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace. Conflicts with use_root_namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Authenticate to the root Vault namespace. Conflicts with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#use_root_namespace VaultProvider#use_root_namespace}
	UseRootNamespace interface{} `field:"optional" json:"useRootNamespace" yaml:"useRootNamespace"`
}

