// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderClientAuth struct {
	// Path to a file containing the client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#cert_file VaultProvider#cert_file}
	CertFile *string `field:"required" json:"certFile" yaml:"certFile"`
	// Path to a file containing the private key that the certificate was issued for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#key_file VaultProvider#key_file}
	KeyFile *string `field:"required" json:"keyFile" yaml:"keyFile"`
}

