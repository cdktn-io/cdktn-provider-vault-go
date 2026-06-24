// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sshsecretbackendrole


type SshSecretBackendRoleAllowedUserKeyConfig struct {
	// List of allowed key lengths, vault-1.10 and above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_role#lengths SshSecretBackendRole#lengths}
	Lengths *[]*float64 `field:"required" json:"lengths" yaml:"lengths"`
	// Key type, choices: rsa, ecdsa, ec, dsa, ed25519, ssh-rsa, ssh-dss, ssh-ed25519, ecdsa-sha2-nistp256, ecdsa-sha2-nistp384, ecdsa-sha2-nistp521.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/ssh_secret_backend_role#type SshSecretBackendRole#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

