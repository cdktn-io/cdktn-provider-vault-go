// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type VaultProviderConfig struct {
	// If true, adds the value of the `address` argument to the Terraform process environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#add_address_to_env VaultProvider#add_address_to_env}
	AddAddressToEnv *string `field:"optional" json:"addAddressToEnv" yaml:"addAddressToEnv"`
	// URL of the root of the target Vault server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#address VaultProvider#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#alias VaultProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// auth_login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login VaultProvider#auth_login}
	AuthLogin interface{} `field:"optional" json:"authLogin" yaml:"authLogin"`
	// auth_login_aws block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_aws VaultProvider#auth_login_aws}
	AuthLoginAws interface{} `field:"optional" json:"authLoginAws" yaml:"authLoginAws"`
	// auth_login_azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_azure VaultProvider#auth_login_azure}
	AuthLoginAzure interface{} `field:"optional" json:"authLoginAzure" yaml:"authLoginAzure"`
	// auth_login_cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_cert VaultProvider#auth_login_cert}
	AuthLoginCert interface{} `field:"optional" json:"authLoginCert" yaml:"authLoginCert"`
	// auth_login_gcp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_gcp VaultProvider#auth_login_gcp}
	AuthLoginGcp interface{} `field:"optional" json:"authLoginGcp" yaml:"authLoginGcp"`
	// auth_login_jwt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_jwt VaultProvider#auth_login_jwt}
	AuthLoginJwt interface{} `field:"optional" json:"authLoginJwt" yaml:"authLoginJwt"`
	// auth_login_kerberos block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_kerberos VaultProvider#auth_login_kerberos}
	AuthLoginKerberos interface{} `field:"optional" json:"authLoginKerberos" yaml:"authLoginKerberos"`
	// auth_login_oci block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_oci VaultProvider#auth_login_oci}
	AuthLoginOci interface{} `field:"optional" json:"authLoginOci" yaml:"authLoginOci"`
	// auth_login_oidc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_oidc VaultProvider#auth_login_oidc}
	AuthLoginOidc interface{} `field:"optional" json:"authLoginOidc" yaml:"authLoginOidc"`
	// auth_login_radius block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_radius VaultProvider#auth_login_radius}
	AuthLoginRadius interface{} `field:"optional" json:"authLoginRadius" yaml:"authLoginRadius"`
	// auth_login_token_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_token_file VaultProvider#auth_login_token_file}
	AuthLoginTokenFile interface{} `field:"optional" json:"authLoginTokenFile" yaml:"authLoginTokenFile"`
	// auth_login_userpass block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#auth_login_userpass VaultProvider#auth_login_userpass}
	AuthLoginUserpass interface{} `field:"optional" json:"authLoginUserpass" yaml:"authLoginUserpass"`
	// Path to directory containing CA certificate files to validate the server's certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#ca_cert_dir VaultProvider#ca_cert_dir}
	CaCertDir *string `field:"optional" json:"caCertDir" yaml:"caCertDir"`
	// Path to a CA certificate file to validate the server's certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#ca_cert_file VaultProvider#ca_cert_file}
	CaCertFile *string `field:"optional" json:"caCertFile" yaml:"caCertFile"`
	// client_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#client_auth VaultProvider#client_auth}
	ClientAuth interface{} `field:"optional" json:"clientAuth" yaml:"clientAuth"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#headers VaultProvider#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// Maximum TTL for secret leases requested by this provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#max_lease_ttl_seconds VaultProvider#max_lease_ttl_seconds}
	MaxLeaseTtlSeconds *float64 `field:"optional" json:"maxLeaseTtlSeconds" yaml:"maxLeaseTtlSeconds"`
	// Maximum number of retries when a 5xx error code is encountered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#max_retries VaultProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Maximum number of retries for Client Controlled Consistency related operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#max_retries_ccc VaultProvider#max_retries_ccc}
	MaxRetriesCcc *float64 `field:"optional" json:"maxRetriesCcc" yaml:"maxRetriesCcc"`
	// The namespace to use. Available only for Vault Enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#namespace VaultProvider#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the token namespace as the root namespace for all resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#set_namespace_from_token VaultProvider#set_namespace_from_token}
	SetNamespaceFromToken interface{} `field:"optional" json:"setNamespaceFromToken" yaml:"setNamespaceFromToken"`
	// Set this to true to prevent the creation of ephemeral child token used by this provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#skip_child_token VaultProvider#skip_child_token}
	SkipChildToken interface{} `field:"optional" json:"skipChildToken" yaml:"skipChildToken"`
	// Skip the dynamic fetching of the Vault server version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#skip_get_vault_version VaultProvider#skip_get_vault_version}
	SkipGetVaultVersion interface{} `field:"optional" json:"skipGetVaultVersion" yaml:"skipGetVaultVersion"`
	// Set this to true only if the target Vault server is an insecure development instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#skip_tls_verify VaultProvider#skip_tls_verify}
	SkipTlsVerify interface{} `field:"optional" json:"skipTlsVerify" yaml:"skipTlsVerify"`
	// Name to use as the SNI host when connecting via TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#tls_server_name VaultProvider#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// Token to use to authenticate to Vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#token VaultProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Token name to use for creating the Vault child token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#token_name VaultProvider#token_name}
	TokenName *string `field:"optional" json:"tokenName" yaml:"tokenName"`
	// Override the Vault server version, which is normally determined dynamically from the target Vault server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs#vault_version_override VaultProvider#vault_version_override}
	VaultVersionOverride *string `field:"optional" json:"vaultVersionOverride" yaml:"vaultVersionOverride"`
}

