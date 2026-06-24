// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oauthresourceserverconfigprofile


type OauthResourceServerConfigProfilePublicKeys struct {
	// The key ID (kid) for this public key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/oauth_resource_server_config_profile#key_id OauthResourceServerConfigProfile#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// The PEM-encoded public key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/oauth_resource_server_config_profile#pem OauthResourceServerConfigProfile#pem}
	Pem *string `field:"required" json:"pem" yaml:"pem"`
}

