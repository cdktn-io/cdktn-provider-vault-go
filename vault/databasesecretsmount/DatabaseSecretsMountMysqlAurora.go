// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount


type DatabaseSecretsMountMysqlAurora struct {
	// Name of the database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#name DatabaseSecretsMount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of roles that are allowed to use this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#allowed_roles DatabaseSecretsMount#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// Specify alternative authorization type. (Only 'gcp_iam' is valid currently).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#auth_type DatabaseSecretsMount#auth_type}
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// Connection string to use to connect to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#connection_url DatabaseSecretsMount#connection_url}
	ConnectionUrl *string `field:"optional" json:"connectionUrl" yaml:"connectionUrl"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#data DatabaseSecretsMount#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#disable_automated_rotation DatabaseSecretsMount#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// Maximum number of seconds a connection may be reused.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#max_connection_lifetime DatabaseSecretsMount#max_connection_lifetime}
	MaxConnectionLifetime *float64 `field:"optional" json:"maxConnectionLifetime" yaml:"maxConnectionLifetime"`
	// Maximum number of idle connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#max_idle_connections DatabaseSecretsMount#max_idle_connections}
	MaxIdleConnections *float64 `field:"optional" json:"maxIdleConnections" yaml:"maxIdleConnections"`
	// Maximum number of open connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#max_open_connections DatabaseSecretsMount#max_open_connections}
	MaxOpenConnections *float64 `field:"optional" json:"maxOpenConnections" yaml:"maxOpenConnections"`
	// The root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#password DatabaseSecretsMount#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Optional name of the password policy to use for generated passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#password_policy DatabaseSecretsMount#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Write-only field for the root credential password used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#password_wo DatabaseSecretsMount#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Version counter for root credential password write-only field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#password_wo_version DatabaseSecretsMount#password_wo_version}
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// Specifies the name of the plugin to use for this connection.
	//
	// Must be prefixed with the name of one of the supported database engine types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#plugin_name DatabaseSecretsMount#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// Optional plugin version to use for this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#plugin_version DatabaseSecretsMount#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// A list of database statements to be executed to rotate the root user's credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#root_rotation_statements DatabaseSecretsMount#root_rotation_statements}
	RootRotationStatements *[]*string `field:"optional" json:"rootRotationStatements" yaml:"rootRotationStatements"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#rotation_period DatabaseSecretsMount#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#rotation_schedule DatabaseSecretsMount#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#rotation_window DatabaseSecretsMount#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// A JSON encoded credential for use with IAM authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#service_account_json DatabaseSecretsMount#service_account_json}
	ServiceAccountJson *string `field:"optional" json:"serviceAccountJson" yaml:"serviceAccountJson"`
	// Skip rotation of static role credentials on import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#skip_static_role_import_rotation DatabaseSecretsMount#skip_static_role_import_rotation}
	SkipStaticRoleImportRotation interface{} `field:"optional" json:"skipStaticRoleImportRotation" yaml:"skipStaticRoleImportRotation"`
	// x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#tls_ca DatabaseSecretsMount#tls_ca}
	TlsCa *string `field:"optional" json:"tlsCa" yaml:"tlsCa"`
	// x509 certificate for connecting to the database.
	//
	// This must be a PEM encoded version of the private key and the certificate combined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#tls_certificate_key DatabaseSecretsMount#tls_certificate_key}
	TlsCertificateKey *string `field:"optional" json:"tlsCertificateKey" yaml:"tlsCertificateKey"`
	// The root credential username used in the connection URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#username DatabaseSecretsMount#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Username generation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#username_template DatabaseSecretsMount#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
	// Specifies if the connection is verified during initial configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#verify_connection DatabaseSecretsMount#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}

