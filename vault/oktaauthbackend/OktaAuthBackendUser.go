// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oktaauthbackend


type OktaAuthBackendUser struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend#groups OktaAuthBackend#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend#policies OktaAuthBackend#policies}.
	Policies *[]*string `field:"optional" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs/resources/okta_auth_backend#username OktaAuthBackend#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

