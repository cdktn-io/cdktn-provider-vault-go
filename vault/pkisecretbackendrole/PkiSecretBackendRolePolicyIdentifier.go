// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendrole


type PkiSecretBackendRolePolicyIdentifier struct {
	// OID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_role#oid PkiSecretBackendRole#oid}
	Oid *string `field:"required" json:"oid" yaml:"oid"`
	// Optional CPS URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_role#cps PkiSecretBackendRole#cps}
	Cps *string `field:"optional" json:"cps" yaml:"cps"`
	// Optional notice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/pki_secret_backend_role#notice PkiSecretBackendRole#notice}
	Notice *string `field:"optional" json:"notice" yaml:"notice"`
}

