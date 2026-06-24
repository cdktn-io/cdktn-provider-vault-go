// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionRedis struct {
	// Specifies the host to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#host DatabaseSecretBackendConnection#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Specifies the password corresponding to the given username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the username for Vault to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The contents of a PEM-encoded CA cert file to use to verify the Redis server's identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#ca_cert DatabaseSecretBackendConnection#ca_cert}
	CaCert *string `field:"optional" json:"caCert" yaml:"caCert"`
	// Specifies whether to skip verification of the server certificate when using TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#insecure_tls DatabaseSecretBackendConnection#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// The transport port to use to connect to Redis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#port DatabaseSecretBackendConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Specifies whether to use TLS when connecting to Redis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secret_backend_connection#tls DatabaseSecretBackendConnection#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
}

