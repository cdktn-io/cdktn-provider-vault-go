// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package azuresecretbackendrole


type AzureSecretBackendRoleAzureRoles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/azure_secret_backend_role#scope AzureSecretBackendRole#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/azure_secret_backend_role#role_id AzureSecretBackendRole#role_id}.
	RoleId *string `field:"optional" json:"roleId" yaml:"roleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/azure_secret_backend_role#role_name AzureSecretBackendRole#role_name}.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

