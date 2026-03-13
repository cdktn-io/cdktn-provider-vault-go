// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderAuthLoginKerberos struct {
	// Disable the Kerberos FAST negotiation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#disable_fast_negotiation VaultProvider#disable_fast_negotiation}
	DisableFastNegotiation interface{} `field:"optional" json:"disableFastNegotiation" yaml:"disableFastNegotiation"`
	// The Kerberos keytab file containing the entry of the login entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#keytab_path VaultProvider#keytab_path}
	KeytabPath *string `field:"optional" json:"keytabPath" yaml:"keytabPath"`
	// A valid Kerberos configuration file e.g. /etc/krb5.conf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#krb5conf_path VaultProvider#krb5conf_path}
	Krb5ConfPath *string `field:"optional" json:"krb5ConfPath" yaml:"krb5ConfPath"`
	// The path where the authentication engine is mounted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#mount VaultProvider#mount}
	Mount *string `field:"optional" json:"mount" yaml:"mount"`
	// The authentication engine's namespace. Conflicts with use_root_namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The Kerberos server's authoritative authentication domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#realm VaultProvider#realm}
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
	// Strip the host from the username found in the keytab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#remove_instance_name VaultProvider#remove_instance_name}
	RemoveInstanceName interface{} `field:"optional" json:"removeInstanceName" yaml:"removeInstanceName"`
	// The service principle name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#service VaultProvider#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// Simple and Protected GSSAPI Negotiation Mechanism (SPNEGO) token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#token VaultProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// The username to login into Kerberos with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#username VaultProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Authenticate to the root Vault namespace. Conflicts with namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.8.0/docs#use_root_namespace VaultProvider#use_root_namespace}
	UseRootNamespace interface{} `field:"optional" json:"useRootNamespace" yaml:"useRootNamespace"`
}

