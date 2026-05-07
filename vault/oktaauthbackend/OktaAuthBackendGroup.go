// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oktaauthbackend


type OktaAuthBackendGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/okta_auth_backend#group_name OktaAuthBackend#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.9.0/docs/resources/okta_auth_backend#policies OktaAuthBackend#policies}.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
}

