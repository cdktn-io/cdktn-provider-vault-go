// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection


type DatabaseSecretBackendConnectionSnowflake struct {
	// Connection string to use to connect to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#connection_url DatabaseSecretBackendConnection#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// Maximum number of seconds a connection may be reused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#max_connection_lifetime DatabaseSecretBackendConnection#max_connection_lifetime}
	MaxConnectionLifetime *float64 `field:"optional" json:"maxConnectionLifetime" yaml:"maxConnectionLifetime"`
	// Maximum number of idle connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#max_idle_connections DatabaseSecretBackendConnection#max_idle_connections}
	MaxIdleConnections *float64 `field:"optional" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	// Maximum number of open connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#max_open_connections DatabaseSecretBackendConnection#max_open_connections}
	MaxOpenConnections *float64 `field:"optional" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	// The root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#password DatabaseSecretBackendConnection#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Write-only field for the root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#password_wo DatabaseSecretBackendConnection#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Version counter for root credential password write-only field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#password_wo_version DatabaseSecretBackendConnection#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// The private key configured for the admin user in Snowflake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#private_key_wo DatabaseSecretBackendConnection#private_key_wo}
	PrivateKeyWo *string `field:"optional" json:"privateKeyWo" yaml:"privateKeyWo"`
	// Version counter for the private key key-pair credentials write-only field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#private_key_wo_version DatabaseSecretBackendConnection#private_key_wo_version}
	PrivateKeyWoVersion *float64 `field:"optional" json:"privateKeyWoVersion" yaml:"privateKeyWoVersion"`
	// The root credential username used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#username DatabaseSecretBackendConnection#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Username generation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secret_backend_connection#username_template DatabaseSecretBackendConnection#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
}

