// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alicloudsecretbackendrole


type AlicloudSecretBackendRoleRemotePolicies struct {
	// Name of the remote policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/alicloud_secret_backend_role#name AlicloudSecretBackendRole#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the remote policy (System or Custom).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/alicloud_secret_backend_role#type AlicloudSecretBackendRole#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

