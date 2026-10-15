// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package scimclient


type ScimClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.13.0/docs/resources/scim_client#delete ScimClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

