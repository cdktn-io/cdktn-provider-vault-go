// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionElasticsearch struct {
	// The password to be used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The URL for Elasticsearch's API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#url DatabaseSecretBackendConnection#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The username to be used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server's identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#ca_cert DatabaseSecretBackendConnection#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
	// The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server's identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#ca_path DatabaseSecretBackendConnection#ca_path}
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// The path to the certificate for the Elasticsearch client to present for communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#client_cert DatabaseSecretBackendConnection#client_cert}
	ClientCert *string `field:"optional" json:"clientCert" yaml:"clientCert"`
	// The path to the key for the Elasticsearch client to use for communication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#client_key DatabaseSecretBackendConnection#client_key}
	ClientKey *string `field:"optional" json:"clientKey" yaml:"clientKey"`
	// Whether to disable certificate verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#insecure DatabaseSecretBackendConnection#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// This, if set, is used to set the SNI host when connecting via TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#tls_server_name DatabaseSecretBackendConnection#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// Template describing how dynamic usernames are generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.7.0/docs/resources/database_secret_backend_connection#username_template DatabaseSecretBackendConnection#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

