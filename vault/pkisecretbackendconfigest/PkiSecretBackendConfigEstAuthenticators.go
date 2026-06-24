// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigest


type PkiSecretBackendConfigEstAuthenticators struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_est#cert PkiSecretBackendConfigEst#cert}.
	Cert *map[string]*string `field:"optional" json:"cert" yaml:"cert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/pki_secret_backend_config_est#userpass PkiSecretBackendConfigEst#userpass}.
	Userpass *map[string]*string `field:"optional" json:"userpass" yaml:"userpass"`
}

