// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginJwt struct {
	// Name of the login role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#role VaultProvider#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// An optional token used to fetch group memberships specified by the distributed claim source in the jwt.
	//
	// This is supported only on Azure/Entra ID. Requires Vault 1.18+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#distributed_claim_access_token VaultProvider#distributed_claim_access_token}
	DistributedClaimAccessToken *string `field:"optional" json:"distributedClaimAccessToken" yaml:"distributedClaimAccessToken"`
	// A signed JSON Web Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs#jwt VaultProvider#jwt}
	Jwt *string `field:"optional" json:"jwt" yaml:"jwt"`
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

