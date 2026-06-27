// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount


type DatabaseSecretsMountCassandra struct {
	// Name of the database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#name DatabaseSecretsMount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A list of roles that are allowed to use this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#allowed_roles DatabaseSecretsMount#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// The number of seconds to use as a connection timeout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#connect_timeout DatabaseSecretsMount#connect_timeout}
	ConnectTimeout *float64 `field:"optional" json:"connectTimeout" yaml:"connectTimeout"`
	// Cassandra consistency level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#consistency DatabaseSecretsMount#consistency}
	Consistency *string `field:"optional" json:"consistency" yaml:"consistency"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#data DatabaseSecretsMount#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#disable_automated_rotation DatabaseSecretsMount#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// Cassandra hosts to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#hosts DatabaseSecretsMount#hosts}
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Whether to skip verification of the server certificate when using TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#insecure_tls DatabaseSecretsMount#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// Cassandra local datacenter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#local_datacenter DatabaseSecretsMount#local_datacenter}
	LocalDatacenter *string `field:"optional" json:"localDatacenter" yaml:"localDatacenter"`
	// The password to use when authenticating with Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#password DatabaseSecretsMount#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Optional name of the password policy to use for generated passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#password_policy DatabaseSecretsMount#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
	// Concatenated PEM blocks containing a certificate and private key;
	//
	// a certificate, private key, and issuing CA certificate; or just a CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#pem_bundle DatabaseSecretsMount#pem_bundle}
	PemBundle *string `field:"optional" json:"pemBundle" yaml:"pemBundle"`
	// Specifies JSON containing a certificate and private key;
	//
	// a certificate, private key, and issuing CA certificate; or just a CA certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#pem_json DatabaseSecretsMount#pem_json}
	PemJson *string `field:"optional" json:"pemJson" yaml:"pemJson"`
	// Specifies the name of the plugin to use for this connection.
	//
	// Must be prefixed with the name of one of the supported database engine types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#plugin_name DatabaseSecretsMount#plugin_name}
	PluginName *string `field:"optional" json:"pluginName" yaml:"pluginName"`
	// Optional plugin version to use for this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#plugin_version DatabaseSecretsMount#plugin_version}
	PluginVersion *string `field:"optional" json:"pluginVersion" yaml:"pluginVersion"`
	// The transport port to use to connect to Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#port DatabaseSecretsMount#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The CQL protocol version to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#protocol_version DatabaseSecretsMount#protocol_version}
	ProtocolVersion *float64 `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// A list of database statements to be executed to rotate the root user's credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#root_rotation_statements DatabaseSecretsMount#root_rotation_statements}
	RootRotationStatements *[]*string `field:"optional" json:"rootRotationStatements" yaml:"rootRotationStatements"`
	// The period of time in seconds between each rotation of the root credential. Cannot be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#rotation_period DatabaseSecretsMount#rotation_period}
	RotationPeriod *float64 `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// The cron-style schedule for the root credential to be rotated on. Cannot be used with rotation_period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#rotation_schedule DatabaseSecretsMount#rotation_schedule}
	RotationSchedule *string `field:"optional" json:"rotationSchedule" yaml:"rotationSchedule"`
	// The maximum amount of time in seconds Vault is allowed to complete a rotation once a scheduled rotation is triggered.
	//
	// Can only be used with rotation_schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#rotation_window DatabaseSecretsMount#rotation_window}
	RotationWindow *float64 `field:"optional" json:"rotationWindow" yaml:"rotationWindow"`
	// Skip rotation of static role credentials on import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#skip_static_role_import_rotation DatabaseSecretsMount#skip_static_role_import_rotation}
	SkipStaticRoleImportRotation interface{} `field:"optional" json:"skipStaticRoleImportRotation" yaml:"skipStaticRoleImportRotation"`
	// Skip permissions checks when a connection to Cassandra is first created.
	//
	// These checks ensure that Vault is able to create roles, but can be resource intensive in clusters with many roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#skip_verification DatabaseSecretsMount#skip_verification}
	SkipVerification interface{} `field:"optional" json:"skipVerification" yaml:"skipVerification"`
	// Enable TCP keepalive for Cassandra connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#socket_keep_alive DatabaseSecretsMount#socket_keep_alive}
	SocketKeepAlive *string `field:"optional" json:"socketKeepAlive" yaml:"socketKeepAlive"`
	// Whether to use TLS when connecting to Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#tls DatabaseSecretsMount#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// SNI host for TLS connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#tls_server_name DatabaseSecretsMount#tls_server_name}
	TlsServerName *string `field:"optional" json:"tlsServerName" yaml:"tlsServerName"`
	// The username to use when authenticating with Cassandra.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#username DatabaseSecretsMount#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Template for dynamic Cassandra usernames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#username_template DatabaseSecretsMount#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
	// Specifies if the connection is verified during initial configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.1/docs/resources/database_secrets_mount#verify_connection DatabaseSecretsMount#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}

