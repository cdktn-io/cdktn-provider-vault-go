// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package databasesecretsmount


type DatabaseSecretsMountCouchbase struct {
	// A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#hosts DatabaseSecretsMount#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// Name of the database connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#name DatabaseSecretsMount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the password corresponding to the given username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#password DatabaseSecretsMount#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// Specifies the username for Vault to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#username DatabaseSecretsMount#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// A list of roles that are allowed to use this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#allowed_roles DatabaseSecretsMount#allowed_roles}
	AllowedRoles *[]*string `field:"optional" json:"allowedRoles" yaml:"allowedRoles"`
	// Required if `tls` is `true`.
	//
	// Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#base64_pem DatabaseSecretsMount#base64_pem}
	Base64Pem *string `field:"optional" json:"base64Pem" yaml:"base64Pem"`
	// Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#bucket_name DatabaseSecretsMount#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#data DatabaseSecretsMount#data}
	Data *map[string]*string `field:"optional" json:"data" yaml:"data"`
	// Stops rotation of the root credential until set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#disable_automated_rotation DatabaseSecretsMount#disable_automated_rotation}
	DisableAutomatedRotation interface{} `field:"optional" json:"disableAutomatedRotation" yaml:"disableAutomatedRotation"`
	// Specifies whether to skip verification of the server certificate when using TLS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#insecure_tls DatabaseSecretsMount#insecure_tls}
	InsecureTls interface{} `field:"optional" json:"insecureTls" yaml:"insecureTls"`
	// Optional name of the password policy to use for generated passwords.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#password_policy DatabaseSecretsMount#password_policy}
	PasswordPolicy *string `field:"optional" json:"passwordPolicy" yaml:"passwordPolicy"`
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
	// Skip rotation of static role credentials on import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#skip_static_role_import_rotation DatabaseSecretsMount#skip_static_role_import_rotation}
	SkipStaticRoleImportRotation interface{} `field:"optional" json:"skipStaticRoleImportRotation" yaml:"skipStaticRoleImportRotation"`
	// Specifies whether to use TLS when connecting to Couchbase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#tls DatabaseSecretsMount#tls}
	Tls interface{} `field:"optional" json:"tls" yaml:"tls"`
	// Template describing how dynamic usernames are generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#username_template DatabaseSecretsMount#username_template}
	UsernameTemplate *string `field:"optional" json:"usernameTemplate" yaml:"usernameTemplate"`
	// Specifies if the connection is verified during initial configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/vault/5.10.0/docs/resources/database_secrets_mount#verify_connection DatabaseSecretsMount#verify_connection}
	VerifyConnection interface{} `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}

