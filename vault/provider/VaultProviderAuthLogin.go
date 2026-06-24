// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLogin struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#path VaultProvider#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#method VaultProvider#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// The authentication engine's namespace. Conflicts with use_root_namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#parameters VaultProvider#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Authenticate to the root Vault namespace. Conflicts with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#use_root_namespace VaultProvider#use_root_namespace}
	UseRootNamespace interface{} `field:"optional" json:"useRootNamespace" yaml:"useRootNamespace"`
}

