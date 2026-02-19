// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gcpsecretstaticaccount


type GcpSecretStaticAccountBinding struct {
	// Resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_static_account#resource GcpSecretStaticAccount#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// List of roles to apply to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/gcp_secret_static_account#roles GcpSecretStaticAccount#roles}
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
}

