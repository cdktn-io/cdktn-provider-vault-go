// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package alicloudsecretbackendrole


type AlicloudSecretBackendRoleInlinePolicies struct {
	// JSON-encoded inline policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/alicloud_secret_backend_role#policy_document AlicloudSecretBackendRole#policy_document}
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
}

